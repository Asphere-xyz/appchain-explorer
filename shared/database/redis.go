package database

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Ankr-network/ankr-protocol/shared"
	common2 "github.com/Ankr-network/ankr-protocol/shared/common"
	"github.com/Ankr-network/ankr-protocol/shared/staking/abigen"
	"github.com/Ankr-network/ankr-protocol/shared/types"
	"github.com/ethereum/go-ethereum/common"
	types2 "github.com/ethereum/go-ethereum/core/types"
	"github.com/go-redis/redis/v9"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"math/big"
	"sync"
)

const redisLatestAffectedBlock = "sidechain-explorer/latest-affected-block"
const redisValidatorKey = "sidechain-explorer/validators"
const redisValidatorHistoryKey = "sidechain-explorer/validator-history/%s"
const redisValidatorSlashingKey = "sidechain-explorer/validator-slashing/%s"
const redisValidatorDepositKey = "sidechain-explorer/validator-deposit/%s"

type StateDb struct {
	con *redis.Client
	pip redis.Pipeliner
	// state
	validatorCache *sync.Map
}

func NewStateDb() *StateDb {
	return &StateDb{}
}

func (s *StateDb) Start(cp shared.IConfigProvider) (err error) {
	config := &Config{}
	if err := cp.Parse(config); err != nil {
		return err
	}
	s.con = redis.NewClient(&redis.Options{
		Addr:     config.RedisUrl,
		Password: config.RedisPassword,
	})
	return nil
}

func (s *StateDb) NewPipeline() {
	if s.pip != nil {
		log.Fatalf("you already have existing pipline")
	}
	s.pip = s.con.Pipeline()
	s.validatorCache = &sync.Map{}
}

func (s *StateDb) CommitPipeline(ctx context.Context, toBlock uint64) error {
	if s.pip == nil {
		log.Fatalf("there is no active pipeline")
	}
	s.pip.Set(ctx, redisLatestAffectedBlock, toBlock, redis.KeepTTL)
	_, err := s.pip.Exec(ctx)
	if err == redis.Nil {
		err = nil
	}
	s.pip = nil
	s.validatorCache = nil
	return err
}

func (s *StateDb) CancelPipline() {
	if s.pip == nil {
		log.Fatalf("there is no active pipeline")
	}
	s.pip.Discard()
	s.pip = nil
	s.validatorCache = nil
}

func (s *StateDb) GetLastAffectedBlock(ctx context.Context) (result uint64) {
	if err := s.con.Get(ctx, redisLatestAffectedBlock).Scan(&result); err != nil && err != redis.Nil {
		log.Fatalf("failed to read latest affected block: %+v", err)
	}
	return result
}

func (s *StateDb) GetDelegators() {
}

func (s *StateDb) GetValidators(ctx context.Context) (result []*types.Validator, err error) {
	bytesSlice := map[string]string{}
	if bytesSlice, err = s.con.HGetAll(ctx, redisValidatorKey).Result(); err != nil {
		return nil, err
	}
	for _, bytes := range bytesSlice {
		val := &types.Validator{}
		fromJSON([]byte(bytes), val)
		result = append(result, val)
	}
	return
}

func (s *StateDb) GetValidatorHistories(ctx context.Context, validator common.Address, offset, size int64) (result []*types.ValidatorHistory, err error) {
	var bytesSlice [][]byte
	if err := s.con.LRange(ctx, fmt.Sprintf(redisValidatorHistoryKey, validator), offset, offset+size).ScanSlice(&bytesSlice); err != nil {
		return nil, err
	}
	for _, bytes := range bytesSlice {
		val := &types.ValidatorHistory{}
		fromJSON(bytes, val)
		result = append(result, val)
	}
	return
}

func (s *StateDb) GetValidatorSlashings(ctx context.Context, validator common.Address, offset, size int64) (result []*types.ValidatorSlashing, err error) {
	var bytesSlice [][]byte
	if err := s.con.LRange(ctx, fmt.Sprintf(redisValidatorSlashingKey, validator), offset, offset+size).ScanSlice(&bytesSlice); err != nil {
		return nil, err
	}
	for _, bytes := range bytesSlice {
		val := &types.ValidatorSlashing{}
		fromJSON(bytes, val)
		result = append(result, val)
	}
	return
}

func (s *StateDb) GetValidatorDeposits(ctx context.Context, validator common.Address, offset, size int64) (result []*types.ValidatorDeposit, err error) {
	var bytesSlice [][]byte
	if err := s.con.LRange(ctx, fmt.Sprintf(redisValidatorDepositKey, validator), offset, offset+size).ScanSlice(&bytesSlice); err != nil {
		return nil, err
	}
	for _, bytes := range bytesSlice {
		val := &types.ValidatorDeposit{}
		fromJSON(bytes, val)
		result = append(result, val)
	}
	return
}

func (s *StateDb) AddValidator(ctx context.Context, entity *abigen.StakingValidatorAdded) (*types.Validator, *types.ValidatorHistory, error) {
	return s.createValidatorChanges(ctx, &types.Validator{
		Address:    entity.Validator.Hex(),
		Owner:      entity.Owner.Hex(),
		Status:     types.ValidatorStatus(entity.Status),
		Commission: uint32(entity.CommissionRate),
	}, &entity.Raw)
}

func (s *StateDb) ModifyValidator(ctx context.Context, entity *abigen.StakingValidatorModified) (*types.Validator, *types.ValidatorHistory, error) {
	return s.createValidatorChanges(ctx, &types.Validator{
		Address:    entity.Validator.Hex(),
		Owner:      entity.Owner.Hex(),
		Status:     types.ValidatorStatus(entity.Status),
		Commission: uint32(entity.CommissionRate),
	}, &entity.Raw)
}

func (s *StateDb) RemoveValidator(ctx context.Context, entity *abigen.StakingValidatorRemoved) (*types.Validator, *types.ValidatorHistory, error) {
	return s.createValidatorChanges(ctx, &types.Validator{
		Address: entity.Validator.Hex(),
		Status:  types.ValidatorStatus_VALIDATOR_STATUS_NOT_FOUND,
	}, &entity.Raw)
}

func (s *StateDb) SlashValidator(ctx context.Context, entity *abigen.StakingValidatorSlashed) (*types.ValidatorSlashing, error) {
	slashing := &types.ValidatorSlashing{
		Validator:       entity.Validator.Hex(),
		Epoch:           entity.Epoch,
		Slashings:       entity.Slashes,
		TransactionHash: entity.Raw.TxHash.Hex(),
		BlockNumber:     entity.Raw.BlockNumber,
	}
	if err := s.pip.LPush(ctx, fmt.Sprintf(redisValidatorSlashingKey, entity.Validator), toJSON(slashing)).Err(); err != nil {
		return nil, err
	}
	return slashing, nil
}

func (s *StateDb) DepositValidator(ctx context.Context, entity *abigen.StakingValidatorDeposited) (*types.ValidatorDeposit, error) {
	slashing := &types.ValidatorDeposit{
		Validator:       entity.Validator.Hex(),
		Amount:          decimal.NewFromBigInt(entity.Amount, -18).String(),
		Epoch:           entity.Epoch,
		TransactionHash: entity.Raw.TxHash.Hex(),
		BlockNumber:     entity.Raw.BlockNumber,
	}
	if err := s.pip.LPush(ctx, fmt.Sprintf(redisValidatorDepositKey, entity.Validator), toJSON(slashing)).Err(); err != nil {
		return nil, err
	}
	return slashing, nil
}

func (s *StateDb) JailValidator(ctx context.Context, entity *abigen.StakingValidatorJailed) (*types.Validator, *types.ValidatorHistory, error) {
	return s.createValidatorChanges(ctx, &types.Validator{
		Address: entity.Validator.Hex(),
		Status:  types.ValidatorStatus_VALIDATOR_STATUS_JAIL,
	}, &entity.Raw)
}

func (s *StateDb) ReleaseValidator(ctx context.Context, entity *abigen.StakingValidatorReleased) (*types.Validator, *types.ValidatorHistory, error) {
	return s.createValidatorChanges(ctx, &types.Validator{
		Address: entity.Validator.Hex(),
		Status:  types.ValidatorStatus_VALIDATOR_STATUS_ACTIVE,
	}, &entity.Raw)
}

func fromJSON[T interface{}](b []byte, t *T) {
	if err := json.Unmarshal(b, t); err != nil {
		log.Panicf("failed to unmarshasl JSON from redis: %+v", err)
	}
}

func toJSON[T interface{}](t T) []byte {
	bytes, err := json.Marshal(t)
	if err != nil {
		log.Panicf("failed to marshal JSON: %+v", err)
	}
	return bytes
}

func (s *StateDb) createValidatorChanges(ctx context.Context, validatorChanges *types.Validator, raw *types2.Log) (*types.Validator, *types.ValidatorHistory, error) {
	// find existing validator
	existingValidator := &types.Validator{}
	if cachedValue, ok := s.validatorCache.Load(validatorChanges.Address); ok {
		existingValidator = cachedValue.(*types.Validator)
	} else {
		bytes, err := s.pip.HGet(ctx, redisValidatorKey, validatorChanges.Address).Bytes()
		if err != nil && err != redis.Nil {
			return nil, nil, err
		} else if bytes != nil {
			fromJSON(bytes, existingValidator)
		}
	}
	// calc diff between two records
	diff, record := common2.CreateEntityDiff(existingValidator, validatorChanges)
	// save new validator entity
	if err := s.pip.HSet(ctx, redisValidatorKey, validatorChanges.Address, toJSON(record)).Err(); err != nil {
		return nil, nil, err
	}
	// save validator history
	history := &types.ValidatorHistory{
		Address:         validatorChanges.Address,
		Changes:         &diff,
		TransactionHash: raw.TxHash.Hex(),
		BlockNumber:     raw.BlockNumber,
	}
	if err := s.pip.LPush(ctx, fmt.Sprintf(redisValidatorHistoryKey, validatorChanges.Address), toJSON(history)).Err(); err != nil {
		return nil, nil, err
	}
	s.validatorCache.Store(validatorChanges.Address, &record)
	return &record, history, nil
}

func (s *StateDb) AddDelegator(delegator, validator common.Address, amount *big.Int) {
}
