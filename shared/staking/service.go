package staking

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Ankr-network/appchain-explorer/shared"
	common2 "github.com/Ankr-network/appchain-explorer/shared/common"
	"github.com/Ankr-network/appchain-explorer/shared/database"
	"github.com/Ankr-network/appchain-explorer/shared/staking/abigen"
	"github.com/Ankr-network/appchain-explorer/shared/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	types2 "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
	"sync"
	"time"
)

// we pay reward every day
var AnkrAprMultiplayer = big.NewFloat(36500)

const minApr = 0.0001

var minAprFloat = new(big.Float).SetFloat64(minApr)

type Service struct {
	state           *database.StateDb
	databaseService *database.Service
	// start
	config *Config
	eth    *ethclient.Client
	// abi
	staking     *abigen.Staking
	chainConfig *abigen.ChainConfig
	// state

	cfgMux            sync.Mutex
	cachedChainConfig *types.ChainConfig

	aprMux           sync.Mutex
	apr              *big.Float
	lastCheckedEpoch uint64

	chainId      *big.Int
	cachedChains []*types.Chain

	statsMux sync.Mutex
	stats    types.StakingStats
}

func NewService(state *database.StateDb, service *database.Service) *Service {
	return &Service{
		state:           state,
		apr:             new(big.Float),
		databaseService: service,
		stats: types.StakingStats{
			Apy:                "0",
			TotalIssuance:      "0",
			TotalStaked:        "0",
			MarketCap:          "0",
			TransferVolume_24H: "0",
			TotalInsurance:     "0",
		},
	}
}

func (s *Service) Start(cp shared.IConfigProvider) (err error) {
	config := &Config{}
	if err := cp.Parse(config); err != nil {
		return err
	}
	s.config = config
	s.eth, err = ethclient.Dial(config.Eth1Url)
	if err != nil {
		return errors.Wrapf(err, "failed to connect to Web3 node (%s)", config.Eth1Url)
	}
	s.staking, _ = abigen.NewStaking(config.StakingContract, s.eth)
	s.chainConfig, _ = abigen.NewChainConfig(config.ChainConfigContract, s.eth)
	newChainConfig, err := s.refreshChainConfig()
	if err != nil {
		return err
	}
	s.cachedChainConfig = newChainConfig

	s.chainId, err = s.eth.ChainID(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed to get chain id")
	}

	go func() {
		if err := s.backgroundWorker(); err != nil {
			log.Fatalf("failed to start background worker: %+v", err)
		}
	}()
	return nil
}

func (s *Service) backgroundWorker() (err error) {
	// do recovery
	hasMore := true
	for hasMore {
		hasMore, err = s.processEventLogs(context.TODO())
		if err != nil {
			return err
		}
	}

	apr, err := s.state.GetLastApy(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed to get last apy")
	}

	if apr != "" {
		s.apr.SetString(apr)
	}

	aprEpoch, err := s.state.GetLastApyEpoch(context.Background())
	if err != nil {
		return errors.Wrap(err, "failed to get last apy")
	}

	if aprEpoch != 0 {
		s.lastCheckedEpoch = aprEpoch
	}

	if err = s.updateStats(); err != nil {
		return errors.Wrap(err, "failed to update stats")
	}

	if err = s.updateAPY(); err != nil {
		return errors.Wrap(err, "failed to update apr")
	}

	// run worker
	refreshChainConfig := time.Tick(10 * time.Minute)
	updateChainsTick := time.Tick(15 * time.Minute)
	processEventLogsTick := time.Tick(10 * time.Second)
	updateAprTick := time.Tick(24 * time.Hour)
	updateStatsTick := time.Tick(time.Minute * 10)
	for {
		var err error
		select {
		case <-refreshChainConfig:
			var newChainConfig *types.ChainConfig
			newChainConfig, err = s.refreshChainConfig()
			if err != nil {
				return err
			}
			s.cfgMux.Lock()
			s.cachedChainConfig = newChainConfig
			s.cfgMux.Unlock()
		case <-updateAprTick:
			if err := s.updateAPY(); err != nil {
				return err
			}
		case <-updateStatsTick:
			if err := s.updateStats(); err != nil {
				return err
			}
		case <-updateChainsTick:
			s.cachedChains, err = s.GetChains(context.TODO())
			if err != nil {
				return err
			}
		case <-processEventLogsTick:
			_, err = s.processEventLogs(context.TODO())
			if err != nil {
				return err
			}
		}
	}
}

func (s *Service) refreshChainConfig() (chainConfig *types.ChainConfig, err error) {
	chainConfig = &types.ChainConfig{
		AverageBlockTime: s.config.BlockTime,
	}
	opts := &bind.CallOpts{}
	chainConfig.ActiveValidatorsLength, err = s.chainConfig.GetActiveValidatorsLength(opts)
	if err != nil {
		return nil, err
	}
	chainConfig.EpochBlockInterval, err = s.chainConfig.GetEpochBlockInterval(opts)
	if err != nil {
		return nil, err
	}
	chainConfig.MisdemeanorThreshold, err = s.chainConfig.GetMisdemeanorThreshold(opts)
	if err != nil {
		return nil, err
	}
	chainConfig.FelonyThreshold, err = s.chainConfig.GetFelonyThreshold(opts)
	if err != nil {
		return nil, err
	}
	chainConfig.ValidatorJailEpochLength, err = s.chainConfig.GetValidatorJailEpochLength(opts)
	if err != nil {
		return nil, err
	}
	chainConfig.UndelegatePeriod, err = s.chainConfig.GetUndelegatePeriod(opts)
	if err != nil {
		return nil, err
	}
	var bigValue *big.Int
	bigValue, err = s.chainConfig.GetMinValidatorStakeAmount(opts)
	if err != nil {
		return nil, err
	}
	chainConfig.MinValidatorStakeAmount = common2.ToEther(bigValue)
	bigValue, err = s.chainConfig.GetMinStakingAmount(opts)
	if err != nil {
		return nil, err
	}
	chainConfig.MinStakingAmount = common2.ToEther(bigValue)
	return
}

func (s *Service) processEventLogs(ctx context.Context) (hasMore bool, err error) {
	lastAffectedBlock := s.state.GetLastAffectedBlock(ctx)
	const confirmationBlocks = 10
	latestKnownBlock, err := s.eth.BlockNumber(context.TODO())
	if err != nil {
		return false, err
	}
	const processBatch = 100_000
	fromBlock := lastAffectedBlock
	toBlock := fromBlock + processBatch
	if toBlock > latestKnownBlock-confirmationBlocks {
		toBlock = latestKnownBlock - confirmationBlocks
	} else {
		hasMore = true
	}
	logs, err := s.eth.FilterLogs(context.TODO(), ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(fromBlock)),
		ToBlock:   big.NewInt(int64(toBlock)),
		Addresses: []common.Address{
			s.config.StakingContract,
		},
	})
	pip := s.state.NewPipeline()
	defer func() {
		if err == nil {
			err = pip.Commit(ctx, toBlock)
		} else {
			pip.Cancel()
		}
	}()
	for _, l := range logs {
		if err = s.processEventLog(ctx, pip, l); err != nil {
			return false, err
		}
	}
	return
}

func (s *Service) processEventLog(ctx context.Context, pip *database.Pipeline, l types2.Log) error {
	if event, _ := s.staking.ParseValidatorAdded(l); event != nil {
		val, _, err := pip.AddValidator(ctx, event)
		if err != nil {
			return err
		}
		log.WithFields(logValidatorFields(val)).Infof("validator added")
		return nil
	}
	if event, _ := s.staking.ParseValidatorModified(l); event != nil {
		val, _, err := pip.ModifyValidator(ctx, event)
		if err != nil {
			return err
		}
		log.WithFields(logValidatorFields(val)).Infof("validator modified")
		return nil
	}
	if event, _ := s.staking.ParseValidatorRemoved(l); event != nil {
		val, _, err := pip.RemoveValidator(ctx, event)
		if err != nil {
			return err
		}
		log.WithFields(logValidatorFields(val)).Infof("validator removed")
		return nil
	}
	if event, _ := s.staking.ParseValidatorSlashed(l); event != nil {
		val, err := pip.SlashValidator(ctx, event)
		if err != nil {
			return err
		}
		log.WithFields(logSlashingFields(val)).Infof("validator slashed")
		return nil
	}
	if event, _ := s.staking.ParseValidatorJailed(l); event != nil {
		val, _, err := pip.JailValidator(ctx, event)
		if err != nil {
			return err
		}
		log.WithFields(logValidatorFields(val)).Infof("validator slashed")
		return nil
	}
	if event, _ := s.staking.ParseValidatorReleased(l); event != nil {
		val, _, err := pip.ReleaseValidator(ctx, event)
		if err != nil {
			return err
		}
		log.WithFields(logValidatorFields(val)).Infof("validator released")
		return nil
	}
	if event, _ := s.staking.ParseValidatorDeposited(l); event != nil {
		val, err := pip.DepositValidator(ctx, event)
		if err != nil {
			return err
		}
		log.WithFields(logDepositFields(val)).Infof("validator deposited")
		return nil
	}
	if event, _ := s.staking.ParseDelegated(l); event != nil {
		ev, err := pip.AddDelegated(ctx, event)
		if err != nil {
			return err
		}
		log.WithFields(logValidatorEventFields(ev)).Infof("validator delegated")
		return nil
	}
	if event, _ := s.staking.ParseUndelegated(l); event != nil {
		ev, err := pip.AddUndelegated(ctx, event)
		if err != nil {
			return err
		}
		log.WithFields(logValidatorEventFields(ev)).Infof("validator undelegated")
		return nil
	}
	if event, _ := s.staking.ParseClaimed(l); event != nil {
		ev, err := pip.AddClaimed(ctx, event)
		if err != nil {
			return err
		}
		log.WithFields(logValidatorEventFields(ev)).Infof("validator claimed")
		return nil
	}
	if event, _ := s.staking.ParseRedelegated(l); event != nil {
		ev, err := pip.AddRedelegated(ctx, event)
		if err != nil {
			return err
		}
		log.WithFields(logValidatorEventFields(ev)).Infof("validator redelegated")
		return nil
	}
	log.WithField("topic", l.Topics[0]).Warnf("not supported event type")
	return nil
}

type chainListResponse struct {
	Name           string   `json:"name"`
	Chain          string   `json:"chain"`
	Rpc            []string `json:"rpc"`
	Faucets        []string `json:"faucets"`
	NativeCurrency struct {
		Name     string `json:"name"`
		Symbol   string `json:"symbol"`
		Decimals uint32 `json:"decimals"`
	}
	InfoUrl   string `json:"infoUrl"`
	ChainId   uint64 `json:"chainId"`
	Slip44    uint64 `json:"slip44"`
	Explorers []struct {
		Name     string `json:"name"`
		Url      string `json:"url"`
		Standard string `json:"standard"`
	}
}

func (s *Service) getChains(ctx context.Context) (result []*types.Chain, err error) {
	req, err := http.NewRequestWithContext(ctx, "GET", s.config.ChainListUrl, nil)
	if err != nil {
		return nil, err
	}
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return nil, err
	}
	requestBody, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var response []chainListResponse
	err = json.Unmarshal(requestBody, &response)
	if err != nil {
		return nil, err
	}
	chains := lo.Map(response, func(t chainListResponse, i int) *types.Chain {
		return &types.Chain{
			ChainId:    t.ChainId,
			ProjectUrl: t.InfoUrl,
			Name:       t.Name,
			ShortName:  t.Chain,
			Rpc:        t.Rpc,
			Faucet:     t.Faucets,
			NativeCurrency: &types.ChainNativeCurrency{
				Name:     t.NativeCurrency.Name,
				Symbol:   t.NativeCurrency.Symbol,
				Decimals: t.NativeCurrency.Decimals,
			},
			Explorer: lo.Map(t.Explorers, func(t struct {
				Name     string `json:"name"`
				Url      string `json:"url"`
				Standard string `json:"standard"`
			}, i int) *types.ChainExplorer {
				return &types.ChainExplorer{
					Name:     t.Name,
					Url:      t.Url,
					Standard: t.Standard,
				}
			}),
		}
	})
	return lo.Filter(chains, func(t *types.Chain, i int) bool {
		if len(s.config.VisibleNetworks) > 0 {
			return s.config.VisibleNetworks[t.ChainId]
		} else {
			return !s.config.HiddenNetworks[t.ChainId]
		}
	}), nil
}

func (s *Service) GetChains(ctx context.Context) (result []*types.Chain, err error) {
	if s.cachedChains != nil {
		return s.cachedChains, nil
	}
	s.cachedChains, err = s.getChains(ctx)
	return s.cachedChains, err
}

func (s *Service) GetChain(ctx context.Context, chain string) (*types.Chain, error) {
	chains, err := s.GetChains(ctx)
	if err != nil {
		return nil, err
	}
	for _, c := range chains {
		if strings.ToLower(c.ShortName) == strings.ToLower(chain) {
			return c, nil
		}
	}
	return nil, fmt.Errorf("unknown chain (%s)", chain)
}

func (s *Service) GetTotalStaked(ctx context.Context) (string, error) {
	res := big.NewInt(0)
	vals, err := s.staking.GetValidators(&bind.CallOpts{Context: ctx})
	if err != nil {
		return "", err
	}
	for _, val := range vals {
		status, err := s.staking.GetValidatorStatus(&bind.CallOpts{Context: ctx}, val)
		if err != nil {
			return "", err
		}
		res.Add(res, status.TotalDelegated)
	}
	return common2.ToEther(res), nil
}

func (s *Service) GetLatestBlock(ctx context.Context) (uint64, uint64, uint64, error) {
	latestKnownBlockNumber, err := s.eth.BlockByNumber(ctx, nil)
	if err != nil {
		return 0, 0, 0, err
	}
	latestAffectedBlockNumber := s.state.GetLastAffectedBlock(ctx)
	return latestKnownBlockNumber.NumberU64(), latestAffectedBlockNumber, latestKnownBlockNumber.Time(), nil
}

func (s *Service) GetChainConfig() *types.ChainConfig {
	s.cfgMux.Lock()
	result := *s.cachedChainConfig
	s.cfgMux.Unlock()
	return &result
}

func (s *Service) GetAPY() *types.ChainConfig {
	s.cfgMux.Lock()
	result := *s.cachedChainConfig
	s.cfgMux.Unlock()
	return &result
}

func (s *Service) updateAPY() error {
	validators, err := s.staking.GetValidators(nil)
	if err != nil {
		return errors.Wrap(err, "failed to get validators")
	}

	currentEpoch, err := s.staking.CurrentEpoch(nil)
	if err != nil {
		return errors.Wrap(err, "failed to get current epoch")
	}

	result := new(big.Float)
	// go through all epoch and try to calculate apr for validators based on reward that received within epoch.
	// within calculated apr's we choose the biggest one
	period := currentEpoch

	if s.lastCheckedEpoch != 0 {
		period = currentEpoch - s.lastCheckedEpoch
	}

	for i := uint64(1); i < period; i++ {
		epoch := currentEpoch - i
		for _, validator := range validators {
			status, err := s.staking.GetValidatorStatusAtEpoch(nil, validator, epoch)
			if err != nil {
				return errors.Wrap(err, "failed to get status at epoch")
			}
			if status.Status != 1 {
				log.Debugf("validator %s is not active", validator.Hex())
				continue
			}

			totalStaked := new(big.Float).SetInt(status.TotalDelegated)
			apr := new(big.Float).SetInt(status.TotalRewards)
			// calculate apr with assumption that we pay +- eq reward
			apr = apr.Quo(apr, totalStaked).Mul(apr, AnkrAprMultiplayer)
			// choose max apr
			if apr.Cmp(result) > 0 {
				result.Set(apr)
			}
		}

		if result.Cmp(new(big.Float)) > 0 {
			break
		}
	}

	s.lastCheckedEpoch = currentEpoch - 1
	if result.Cmp(new(big.Float)) > 0 {
		if result.Cmp(minAprFloat) < 0 {
			result.SetUint64(0)
		}

		s.aprMux.Lock()
		s.apr.Set(result)
		s.aprMux.Unlock()

		if err := s.state.SetLastApy(context.Background(), s.apr.String()); err != nil {
			log.WithError(err).Error("failed to set last apy")
		}
	}

	if err := s.state.SetLastApyEpoch(context.Background(), s.lastCheckedEpoch); err != nil {
		log.WithError(err).Error("failed to set last apy epoch")
	}

	return nil
}

func (s *Service) updateStats() (err error) {
	ctx := context.Background()
	stats := types.StakingStats{}
	if stats.ActiveUsers_7D, err = s.databaseService.EstimateActiveUsers(ctx, 7*time.Hour*24); err != nil {
		log.WithError(err).Error("failed to estimate active users")
	}
	if stats.TotalHolders, err = s.databaseService.EstimateTokenHolders(ctx); err != nil {
		log.WithError(err).Error("failed to estimate token holdres")
	}
	if stats.TotalIssuance, err = s.databaseService.GetTotalIssuance(ctx); err != nil {
		log.WithError(err).Error("failed to get total issuance")
	}
	if stats.TotalValidators, err = s.state.GetTotalValidators(ctx); err != nil {
		log.WithError(err).Error("failed to get total validators")
	}
	if stats.TotalDelegators, err = s.state.GetTotalDelegators(ctx); err != nil {
		log.WithError(err).Error("failed to get total delegators")
	}
	stats.ChainId = s.GetChainID().Uint64()
	stats.Apy = s.GetApr().String()
	if stats.TotalTxs, err = s.databaseService.EstimateTransactionCount(ctx); err != nil {
		log.WithError(err).Error("failed to estimate tx count")
	}
	if stats.TotalTransfers, err = s.databaseService.EstimateTransfersCount(ctx); err != nil {
		log.WithError(err).Error("failed to estimate transfers count")
	}
	if stats.TotalStaked, err = s.GetTotalStaked(ctx); err != nil {
		log.WithError(err).Error("failed to get total staked")
	}
	if stats.TransferVolume_24H, err = s.databaseService.GetTransferVolume(ctx, time.Hour*24); err != nil {
		log.WithError(err).Error("failed to get transfer volume")
	}
	if stats.MarketCap, err = s.databaseService.GetMarketCap(ctx); err != nil {
		log.WithError(err).Error("failed to get marketcap")
	}
	if stats.KnownBlock, s.stats.AffectedBlock, _, err = s.GetLatestBlock(ctx); err != nil {
		log.WithError(err).Error("failed to get latestBlock")
	}

	s.statsMux.Lock()
	s.stats = stats
	s.statsMux.Unlock()

	return
}

func (s *Service) GetStats() types.StakingStats {
	s.statsMux.Lock()
	result := s.stats
	s.statsMux.Unlock()
	return result
}

func (s *Service) GetChainID() *big.Int {
	return s.chainId
}

func (s *Service) GetApr() *big.Float {
	s.aprMux.Lock()
	result := new(big.Float).Set(s.apr)
	s.aprMux.Unlock()
	return result
}

func (s *Service) GetValidators(ctx context.Context) (result []*types.Validator, err error) {
	return s.state.GetValidators(ctx)
}

func (s *Service) GetValidatorSlashings(ctx context.Context, validator common.Address, offset, size uint32) (result []*types.ValidatorSlashing, err error) {
	return s.state.GetValidatorSlashings(ctx, validator, int64(offset), int64(size))
}

func (s *Service) GetValidatorHistories(ctx context.Context, validator common.Address, offset, size uint32) (result []*types.ValidatorHistory, err error) {
	return s.state.GetValidatorHistories(ctx, validator, int64(offset), int64(size))
}

func (s *Service) GetValidatorDeposits(ctx context.Context, validator common.Address, offset, size uint32) (result []*types.ValidatorDeposit, err error) {
	return s.state.GetValidatorDeposits(ctx, validator, int64(offset), int64(size))
}

func (s *Service) GetValidatorEvents(ctx context.Context, validator common.Address, offset, size uint32) (result []*types.ValidatorEvent, err error) {
	return s.state.GetValidatorEvents(ctx, validator, int64(offset), int64(size))
}

func (s *Service) GetDelegatorsByValidator(ctx context.Context, validator common.Address, offset, size uint32) (result []*types.Delegator, err error) {
	return s.state.GetDelegators(ctx, int64(offset), int64(size), validator)
}

func (s *Service) GetDelegators(ctx context.Context, offset, size uint32) (result []*types.Delegator, err error) {
	return s.state.GetDelegators(ctx, int64(offset), int64(size), common.Address{})
}
