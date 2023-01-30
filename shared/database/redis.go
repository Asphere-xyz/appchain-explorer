package database

import (
	"context"
	"fmt"
	"github.com/Ankr-network/ankr-protocol/shared"
	common2 "github.com/Ankr-network/ankr-protocol/shared/common"
	"github.com/Ankr-network/ankr-protocol/shared/staking/abigen"
	"github.com/Ankr-network/ankr-protocol/shared/types"
	"github.com/ethereum/go-ethereum/common"
	types2 "github.com/ethereum/go-ethereum/core/types"
	"github.com/go-redis/redis/v9"
	"github.com/pkg/errors"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"strconv"
	"sync"
)

const redisLatestAffectedBlock = "sidechain-explorer/latest-affected-block"
const redisLastApyEpoch = "sidechain-explorer/latest-apy-epoch"
const redisLastApy = "sidechain-explorer/latest-apy"
const redisValidatorKey = "sidechain-explorer/validators"
const redisValidatorHistoryKey = "sidechain-explorer/validator-history/%s"
const redisValidatorSlashingKey = "sidechain-explorer/validator-slashing/%s"
const redisValidatorDepositKey = "sidechain-explorer/validator-deposit/%s"
const redisValidatorEventKey = "sidechain-explorer/validator-event/%s"
const redisDelegatorsKey = "sidechain-explorer/delegators"
const redisDelegatorsIndexKey = "sidechain-explorer/delegators-index/%s"

type StateDb struct {
	con *redis.Client
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

func (s *StateDb) GetLastAffectedBlock(ctx context.Context) (result uint64) {
	if err := s.con.Get(ctx, redisLatestAffectedBlock).Scan(&result); err != nil && err != redis.Nil {
		log.Fatalf("failed to read latest affected block: %+v", err)
	}
	return result
}

func (s *StateDb) GetLastApyEpoch(ctx context.Context) (result uint64, err error) {
	err = s.con.Get(ctx, redisLastApyEpoch).Scan(&result)
	if err != nil && err != redis.Nil {
		return result, errors.Wrap(err, "failed to scan last apy epoch")
	}
	return result, nil
}

func (s *StateDb) SetLastApyEpoch(ctx context.Context, apyEpoch uint64) error {
	err := s.con.Set(ctx, redisLastApyEpoch, apyEpoch, 0).Err()
	return errors.Wrap(err, "failed to set last apy epoch")
}

func (s *StateDb) SetLastApy(ctx context.Context, apy string) error {
	err := s.con.Set(ctx, redisLastApy, apy, 0).Err()
	return errors.Wrap(err, "failed to set last apy epoch")
}

func (s *StateDb) GetLastApy(ctx context.Context) (result string, err error) {
	err = s.con.Get(ctx, redisLastApy).Scan(&result)
	if err != nil && err != redis.Nil {
		return result, errors.Wrap(err, "failed to scan last apy epoch")
	}
	return result, nil
}

func (s *StateDb) GetValidators(ctx context.Context) (result []*types.Validator, err error) {
	bytesSlice := map[string]string{}
	if bytesSlice, err = s.con.HGetAll(ctx, redisValidatorKey).Result(); err != nil {
		return nil, err
	}
	for _, bytes := range bytesSlice {
		val := &types.Validator{}
		unmarshallProto([]byte(bytes), val)
		result = append(result, val)
	}
	return
}

func (s *StateDb) GetTotalValidators(ctx context.Context) (uint64, error) {
	res, err := s.con.HLen(ctx, redisValidatorKey).Result()
	return uint64(res), err
}

func (s *StateDb) GetTotalDelegators(ctx context.Context) (uint64, error) {
	res, err := s.con.HLen(ctx, redisDelegatorsKey).Result()
	return uint64(res), err
}

func (s *StateDb) GetTotalValidatorDeposits(ctx context.Context, validator common.Address) (uint64, error) {
	res, err := s.con.LLen(ctx, fmt.Sprintf(redisValidatorDepositKey, validator.Hex())).Result()
	return uint64(res), err
}

func (s *StateDb) GetDelegators(ctx context.Context, offset, size int64, validator common.Address) (result []*types.Delegator, err error) {
	validatorString := ""
	if validator != (common.Address{}) {
		validatorString = validator.Hex()
	}
	lines, err := s.con.ZRangeArgs(ctx, redis.ZRangeArgs{
		Key:   fmt.Sprintf(redisDelegatorsIndexKey, validatorString),
		Start: offset,
		Stop:  offset + size,
		Rev:   true,
	}).Result()
	if err != nil {
		return nil, err
	} else if len(lines) == 0 {
		return
	}
	data, err := s.con.HMGet(ctx, redisDelegatorsKey, lines...).Result()
	if err != nil {
		return nil, err
	}
	for _, d := range data {
		val := &types.Delegator{}
		unmarshallProto(d.(string), val)
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
		unmarshallProto(bytes, val)
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
		unmarshallProto(bytes, val)
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
		unmarshallProto(bytes, val)
		result = append(result, val)
	}
	return
}

func (s *StateDb) GetValidatorEvents(ctx context.Context, validator common.Address, offset, size int64) (result []*types.ValidatorEvent, err error) {
	var bytesSlice [][]byte
	if err := s.con.LRange(ctx, fmt.Sprintf(redisValidatorEventKey, validator), offset, offset+size).ScanSlice(&bytesSlice); err != nil {
		return nil, err
	}
	for _, bytes := range bytesSlice {
		val := &types.ValidatorEvent{}
		unmarshallProto(bytes, val)
		result = append(result, val)
	}
	return
}

type Pipeline struct {
	pip   redis.Pipeliner
	state sync.Map
}

func (s *StateDb) NewPipeline() *Pipeline {
	return &Pipeline{
		pip: s.con.Pipeline(),
	}
}

func (p *Pipeline) Commit(ctx context.Context, toBlock uint64) error {
	p.pip.Set(ctx, redisLatestAffectedBlock, toBlock, redis.KeepTTL)
	_, err := p.pip.Exec(ctx)
	if err == redis.Nil {
		err = nil
	}
	return err
}

func (p *Pipeline) Cancel() {
	p.pip.Discard()
	p.pip = nil
}

func (p *Pipeline) AddValidator(ctx context.Context, entity *abigen.StakingValidatorAdded) (*types.Validator, *types.ValidatorHistory, error) {
	return p.createValidatorChanges(ctx, &types.Validator{
		Address:    entity.Validator.Hex(),
		Owner:      entity.Owner.Hex(),
		Status:     types.ValidatorStatus(entity.Status),
		Commission: uint32(entity.CommissionRate),
	}, &entity.Raw)
}

func (p *Pipeline) ModifyValidator(ctx context.Context, entity *abigen.StakingValidatorModified) (*types.Validator, *types.ValidatorHistory, error) {
	return p.createValidatorChanges(ctx, &types.Validator{
		Address:    entity.Validator.Hex(),
		Owner:      entity.Owner.Hex(),
		Status:     types.ValidatorStatus(entity.Status),
		Commission: uint32(entity.CommissionRate),
	}, &entity.Raw)
}

func (p *Pipeline) RemoveValidator(ctx context.Context, entity *abigen.StakingValidatorRemoved) (*types.Validator, *types.ValidatorHistory, error) {
	return p.createValidatorChanges(ctx, &types.Validator{
		Address: entity.Validator.Hex(),
		Status:  types.ValidatorStatus_VALIDATOR_STATUS_NOT_FOUND,
	}, &entity.Raw)
}

func (p *Pipeline) SlashValidator(ctx context.Context, entity *abigen.StakingValidatorSlashed) (*types.ValidatorSlashing, error) {
	slashing := &types.ValidatorSlashing{
		Validator:       entity.Validator.Hex(),
		Epoch:           entity.Epoch,
		Slashings:       entity.Slashes,
		TransactionHash: entity.Raw.TxHash.Hex(),
		BlockNumber:     entity.Raw.BlockNumber,
	}
	if err := p.pip.LPush(ctx, fmt.Sprintf(redisValidatorSlashingKey, entity.Validator), marshallProto(slashing)).Err(); err != nil {
		return nil, err
	}
	return slashing, nil
}

func (p *Pipeline) DepositValidator(ctx context.Context, entity *abigen.StakingValidatorDeposited) (*types.ValidatorDeposit, error) {
	slashing := &types.ValidatorDeposit{
		Validator:       entity.Validator.Hex(),
		Amount:          common2.ToEther(entity.Amount),
		Epoch:           entity.Epoch,
		TransactionHash: entity.Raw.TxHash.Hex(),
		BlockNumber:     entity.Raw.BlockNumber,
	}
	p.pip.LPush(ctx, fmt.Sprintf(redisValidatorDepositKey, entity.Validator), marshallProto(slashing))
	return slashing, nil
}

func (p *Pipeline) JailValidator(ctx context.Context, entity *abigen.StakingValidatorJailed) (*types.Validator, *types.ValidatorHistory, error) {
	return p.createValidatorChanges(ctx, &types.Validator{
		Address: entity.Validator.Hex(),
		Status:  types.ValidatorStatus_VALIDATOR_STATUS_JAIL,
	}, &entity.Raw)
}

func (p *Pipeline) ReleaseValidator(ctx context.Context, entity *abigen.StakingValidatorReleased) (*types.Validator, *types.ValidatorHistory, error) {
	return p.createValidatorChanges(ctx, &types.Validator{
		Address: entity.Validator.Hex(),
		Status:  types.ValidatorStatus_VALIDATOR_STATUS_ACTIVE,
	}, &entity.Raw)
}

func (p *Pipeline) createValidatorChanges(ctx context.Context, validatorChanges *types.Validator, raw *types2.Log) (*types.Validator, *types.ValidatorHistory, error) {
	// find existing validator
	existingValidator := &types.Validator{}
	if cachedValue, ok := p.state.Load(fmt.Sprintf("validator/%s", validatorChanges.Address)); ok {
		existingValidator = cachedValue.(*types.Validator)
	} else {
		bytes, err := p.pip.HGet(ctx, redisValidatorKey, validatorChanges.Address).Bytes()
		if err != nil && err != redis.Nil {
			return nil, nil, err
		} else if bytes != nil {
			unmarshallProto(bytes, existingValidator)
		}
	}
	// calc diff between two records
	diff, record := common2.CreateEntityDiff(existingValidator, validatorChanges)
	// save new validator entity
	p.pip.HSet(ctx, redisValidatorKey, validatorChanges.Address, marshallProto(&record))
	// save validator history
	history := &types.ValidatorHistory{
		Address:         validatorChanges.Address,
		Changes:         &diff,
		TransactionHash: raw.TxHash.Hex(),
		BlockNumber:     raw.BlockNumber,
	}
	p.pip.LPush(ctx, fmt.Sprintf(redisValidatorHistoryKey, validatorChanges.Address), marshallProto(history))
	// update cache
	p.state.Store(fmt.Sprintf("validator/%s", validatorChanges.Address), &record)
	return &record, history, nil
}

type baseDelegateMethods interface {
	GetValidator() string
	GetStaker() string
	GetAmount() string
	GetEpoch() uint64
}

func (p *Pipeline) createDelegatorChanges(ctx context.Context, changes baseDelegateMethods, isPlus bool) (*types.Delegator, error) {
	del := &types.Delegator{
		Staker: changes.GetStaker(),
		// these fields must be empty for the first init
		TotalDelegated: "0",
		Epoch:          0,
	}
	if cachedValue, ok := p.state.Load(fmt.Sprintf("delegator/%s", changes.GetStaker())); ok {
		del = cachedValue.(*types.Delegator)
	} else {
		bytes, err := p.pip.HGet(ctx, redisDelegatorsKey, changes.GetStaker()).Bytes()
		if err != nil && err != redis.Nil {
			return nil, err
		} else if bytes != nil {
			unmarshallProto(bytes, del)
		}
	}
	// update fields using delta
	if isPlus {
		del.TotalDelegated = decimal.RequireFromString(del.TotalDelegated).Add(decimal.RequireFromString(changes.GetAmount())).String()
	} else {
		del.TotalDelegated = decimal.RequireFromString(del.TotalDelegated).Sub(decimal.RequireFromString(changes.GetAmount())).String()
	}
	del.Epoch = changes.GetEpoch()
	// write new state
	p.pip.HSet(ctx, redisDelegatorsKey, changes.GetStaker(), marshallProto(del))
	p.state.Store(fmt.Sprintf("delegator/%s", changes.GetStaker()), del)
	// create search index
	amount, err := strconv.ParseFloat(changes.GetAmount(), 64)
	if err != nil {
		return nil, err
	}
	p.pip.ZAdd(ctx, fmt.Sprintf(redisDelegatorsIndexKey, changes.GetValidator()), redis.Z{Member: del.Staker, Score: amount})
	p.pip.ZAdd(ctx, fmt.Sprintf(redisDelegatorsIndexKey, ""), redis.Z{Member: del.Staker, Score: amount})
	return del, nil
}

func (p *Pipeline) AddDelegated(ctx context.Context, entity *abigen.StakingDelegated) (*types.ValidatorEvent, error) {
	bigAmount := common2.ToEtherDecimal(entity.Amount)
	eventFields := &types.ValidatorEvent_EventFields{
		Validator: entity.Validator.Hex(),
		Staker:    entity.Staker.Hex(),
		Amount:    bigAmount.String(),
		Epoch:     entity.Epoch,
	}
	event := &types.ValidatorEvent{
		Event: &types.ValidatorEvent_Delegated{
			Delegated: eventFields,
		},
		TransactionHash: entity.Raw.TxHash.Hex(),
		BlockNumber:     entity.Raw.BlockNumber,
	}
	eventFields.GetAmount()
	p.pip.LPush(ctx, fmt.Sprintf(redisValidatorEventKey, entity.Validator), marshallProto(event))
	if _, err := p.createDelegatorChanges(ctx, eventFields, true); err != nil {
		return nil, err
	}
	return event, nil
}

func (p *Pipeline) AddUndelegated(ctx context.Context, entity *abigen.StakingUndelegated) (*types.ValidatorEvent, error) {
	eventFields := &types.ValidatorEvent_EventFields{
		Validator: entity.Validator.Hex(),
		Staker:    entity.Staker.Hex(),
		Amount:    common2.ToEther(entity.Amount),
		Epoch:     entity.Epoch,
	}
	event := &types.ValidatorEvent{
		Event: &types.ValidatorEvent_Undelegated{
			Undelegated: eventFields,
		},
		TransactionHash: entity.Raw.TxHash.Hex(),
		BlockNumber:     entity.Raw.BlockNumber,
	}
	p.pip.LPush(ctx, fmt.Sprintf(redisValidatorEventKey, entity.Validator), marshallProto(event))
	if _, err := p.createDelegatorChanges(ctx, eventFields, false); err != nil {
		return nil, err
	}
	return event, nil
}

func (p *Pipeline) AddClaimed(ctx context.Context, entity *abigen.StakingClaimed) (*types.ValidatorEvent, error) {
	eventFields := &types.ValidatorEvent_EventFields{
		Validator: entity.Validator.Hex(),
		Staker:    entity.Staker.Hex(),
		Amount:    common2.ToEther(entity.Amount),
		Epoch:     entity.Epoch,
	}
	event := &types.ValidatorEvent{
		Event: &types.ValidatorEvent_Claimed{
			Claimed: eventFields,
		},
		TransactionHash: entity.Raw.TxHash.Hex(),
		BlockNumber:     entity.Raw.BlockNumber,
	}
	p.pip.LPush(ctx, fmt.Sprintf(redisValidatorEventKey, entity.Validator), marshallProto(event))
	return event, nil
}

func (p *Pipeline) AddRedelegated(ctx context.Context, entity *abigen.StakingRedelegated) (*types.ValidatorEvent, error) {
	eventFields := &types.ValidatorEvent_EventFields{
		Validator: entity.Validator.Hex(),
		Staker:    entity.Staker.Hex(),
		Amount:    common2.ToEther(entity.Amount),
		Dust:      common2.ToEther(entity.Dust),
		Epoch:     entity.Epoch,
	}
	event := &types.ValidatorEvent{
		Event: &types.ValidatorEvent_Redelegated{
			Redelegated: eventFields,
		},
		TransactionHash: entity.Raw.TxHash.Hex(),
		BlockNumber:     entity.Raw.BlockNumber,
	}
	p.pip.LPush(ctx, fmt.Sprintf(redisValidatorEventKey, entity.Validator), marshallProto(event))
	if _, err := p.createDelegatorChanges(ctx, eventFields, true); err != nil {
		return nil, err
	}
	return event, nil
}
