package staking

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/Ankr-network/ankr-protocol/shared"
	"github.com/Ankr-network/ankr-protocol/shared/staking/abigen"
	"github.com/Ankr-network/ankr-protocol/shared/types"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/pkg/errors"
	"github.com/samber/lo"
	"github.com/shopspring/decimal"
	log "github.com/sirupsen/logrus"
	"io/ioutil"
	"math/big"
	"net/http"
	"strings"
	"time"
)

type Service struct {
	config  *Config
	eth     *ethclient.Client
	staking *abigen.Staking
	// state
	delegators []*types.Delegator
	chains     []*types.Chain
}

func NewService() *Service {
	return &Service{}
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
			s.delegators, err = s.getDelegators(context.Background())
			if err != nil {
				log.Errorf("failed to refresh recent delegators: %+v", err)
			}
		case <-updateChainsTick:
			s.chains, err = s.GetChains(context.Background())
			if err != nil {
				log.Errorf("failed to refresh chains: %+v", err)
			}
		}
	}
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

func (s *Service) getDelegators(ctx context.Context) (result []*types.Delegator, err error) {
	it, err := s.staking.FilterDelegated(&bind.FilterOpts{Context: ctx, Start: uint64(0)}, nil, nil)
	if err != nil {
		return nil, err
	}
	delegators := make(map[common.Address]*big.Int)
	for it.Next() {
		balance := delegators[it.Event.Staker]
		if balance == nil {
			balance = big.NewInt(0)
		}
		balance.Add(balance, it.Event.Amount)
		delegators[it.Event.Staker] = balance
	}
	if it.Error() != nil {
		return nil, it.Error()
	}
	return lo.MapToSlice(delegators, func(k common.Address, v *big.Int) *types.Delegator {
		return &types.Delegator{Address: k.Hex(), Amount: decimal.NewFromBigInt(v, -18).String()}
	}), nil
}

func (s *Service) GetDelegators(ctx context.Context) (result []*types.Delegator, err error) {
	if s.delegators != nil {
		return s.delegators, nil
	}
	s.delegators, err = s.getDelegators(ctx)
	return s.delegators, err
}
