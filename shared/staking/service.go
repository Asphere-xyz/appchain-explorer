package staking

import (
	"context"
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
	"math/big"
	"time"
)

type Service struct {
	eth     *ethclient.Client
	staking *abigen.Staking
	// state
	delegators []*types.Delegator
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Start(cp shared.IConfigProvider) (err error) {
	config := &Config{}
	if err := cp.Parse(config); err != nil {
		return err
	}
	s.eth, err = ethclient.Dial(config.Eth1Url)
	if err != nil {
		return errors.Wrapf(err, "failed to connect to Web3 node (%s)", config.Eth1Url)
	}
	s.staking, _ = abigen.NewStaking(config.StakingContract, s.eth)
	go s.refreshRecentDelegators()
	return nil
}

func (s *Service) refreshRecentDelegators() {
	ticker := time.Tick(5 * time.Minute)
	for {
		<-ticker
		_, err := s.GetDelegators(context.TODO())
		if err != nil {
			log.Errorf("failed to refresh recent delegators: %+v", err)
		}
	}
}

func (s *Service) GetDelegators(ctx context.Context) (result []*types.Delegator, err error) {
	if s.delegators != nil {
		return s.delegators, nil
	}
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
	s.delegators = lo.MapToSlice(delegators, func(k common.Address, v *big.Int) *types.Delegator {
		return &types.Delegator{Address: k.Hex(), Amount: decimal.NewFromBigInt(v, -18).String()}
	})
	return s.delegators, nil
}
