package staking

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Ankr-network/ankr-protocol/shared"
	"github.com/Ankr-network/ankr-protocol/shared/database"
	"github.com/Ankr-network/ankr-protocol/shared/staking/abigen"
	"github.com/Ankr-network/ankr-protocol/shared/types"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	types2 "github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"math/big"
	"net/http"
	"sort"
	"strings"
	"time"
)

type Service struct {
	state *database.StateDb
	// start
	config  *Config
	eth     *ethclient.Client
	staking *abigen.Staking
	// state
	chains []*types.Chain
}

func NewService(state *database.StateDb) *Service {
	return &Service{state: state}
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
	go func() {
		if err := s.backgroundWorker(); err != nil {
			log.Fatalf("failed to start background worker: %+v", err)
		}
	}()
	go s.refreshRecentDelegators()
	return nil
}

func (s *Service) refreshRecentDelegators() {
	updateDelegatorsTick := time.Tick(15 * time.Minute)
	updateChainsTick := time.Tick(15 * time.Minute)
	for {
		var err error
		select {
		case <-updateDelegatorsTick:
			//newDelegators := make(map[common.Address][]*types.Delegator)
			//for validator := range s.delegators {
			//	delegators, err := s.getDelegators(context.Background(), []common.Address{validator})
			//	if err != nil {
			//		log.Errorf("failed to refresh recent delegators: %+v", err)
			//	}
			//	newDelegators[validator] = delegators
			//}
			//s.delegators = newDelegators
		case <-updateChainsTick:
			s.chains, err = s.GetChains(context.Background())
			if err != nil {
				log.Errorf("failed to refresh chains: %+v", err)
			}
		}
	}
}

func (s *Service) backgroundWorker() error {
	ctx := context.TODO()
	for {
		hasMore, err := s.processEventLogs(ctx)
		if err != nil {
			return err
		} else if hasMore {
			continue
		}
		time.Sleep(30 * time.Second)
	}
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
	s.state.NewPipeline()
	defer func() {
		if err != nil {
			s.state.CancelPipline()
		} else {
			err = s.state.CommitPipeline(ctx, toBlock)
		}
	}()
	for _, l := range logs {
		if err = s.processEventLog(ctx, l); err != nil {
			return false, err
		}
	}
	return
}

func (s *Service) processEventLog(ctx context.Context, l types2.Log) error {
	if event, _ := s.staking.ParseValidatorAdded(l); event != nil {
		val, _, err := s.state.AddValidator(ctx, event)
		if err != nil {
			return err
		}
		log.WithFields(logValidatorFields(val)).Infof("validator added")
		return nil
	}
	if event, _ := s.staking.ParseValidatorModified(l); event != nil {
		val, _, err := s.state.ModifyValidator(ctx, event)
		if err != nil {
			return err
		}
		log.WithFields(logValidatorFields(val)).Infof("validator modified")
		return nil
	}
	if event, _ := s.staking.ParseValidatorRemoved(l); event != nil {
		val, _, err := s.state.RemoveValidator(ctx, event)
		if err != nil {
			return err
		}
		log.WithFields(logValidatorFields(val)).Infof("validator removed")
		return nil
	}
	if event, _ := s.staking.ParseValidatorSlashed(l); event != nil {
		val, err := s.state.SlashValidator(ctx, event)
		if err != nil {
			return err
		}
		log.WithFields(logSlashingFields(val)).Infof("validator slashed")
		return nil
	}
	if event, _ := s.staking.ParseValidatorJailed(l); event != nil {
		val, _, err := s.state.JailValidator(ctx, event)
		if err != nil {
			return err
		}
		log.WithFields(logValidatorFields(val)).Infof("validator slashed")
		return nil
	}
	if event, _ := s.staking.ParseValidatorReleased(l); event != nil {
		val, _, err := s.state.ReleaseValidator(ctx, event)
		if err != nil {
			return err
		}
		log.WithFields(logValidatorFields(val)).Infof("validator released")
		return nil
	}
	if event, _ := s.staking.ParseValidatorDeposited(l); event != nil {
		val, err := s.state.DepositValidator(ctx, event)
		if err != nil {
			return err
		}
		log.WithFields(logDepositFields(val)).Infof("validator deposited")
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
		return !s.config.HiddenNetworks[t.ChainId]
	}), nil
}

func (s *Service) GetChains(ctx context.Context) (result []*types.Chain, err error) {
	if s.chains != nil {
		return s.chains, nil
	}
	s.chains, err = s.getChains(ctx)
	return s.chains, err
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

func (s *Service) getDelegators(ctx context.Context, validators []common.Address) (result SortedDelegators, err error) {
	it, err := s.staking.FilterDelegated(&bind.FilterOpts{Context: ctx, Start: uint64(0)}, validators, nil)
	if err != nil {
		return nil, err
	}
	type delegator struct {
		amount *big.Int
		epoch  uint64
	}
	delegators := make(map[common.Address]delegator)
	for it.Next() {
		del := delegators[it.Event.Staker]
		if del.amount == nil {
			del.amount = big.NewInt(0)
		}
		del.amount.Add(del.amount, it.Event.Amount)
		del.epoch = it.Event.Epoch
		delegators[it.Event.Staker] = del
	}
	if it.Error() != nil {
		return nil, it.Error()
	}
	result = lo.MapToSlice(delegators, func(k common.Address, v delegator) *types.Delegator {
		return &types.Delegator{
			Address: k.Hex(),
			Amount:  decimal.NewFromBigInt(v.amount, -18).String(),
			Epoch:   v.epoch,
		}
	})
	sort.Sort(result)
	return result, nil
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

func (s *Service) GetDelegatorsByValidator(ctx context.Context, validator common.Address) (result []*types.Delegator, err error) {
	return s.getDelegators(ctx, []common.Address{validator})
}

func (s *Service) GetDelegators(ctx context.Context) (result []*types.Delegator, err error) {
	return s.getDelegators(ctx, []common.Address{})
}
