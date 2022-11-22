package gateway

import (
	"context"
	"github.com/Ankr-network/ankr-protocol/shared/types"
	"github.com/ethereum/go-ethereum/common"
	"time"
)

func (s *Server) GetValidators(ctx context.Context, req *types.GetValidatorsRequest) (*types.GetValidatorsReply, error) {
	result, err := s.stakingService.GetValidators(ctx)
	if err != nil {
		return nil, err
	}
	return &types.GetValidatorsReply{Validators: result}, nil
}

func (s *Server) GetValidatorHistory(ctx context.Context, req *types.GetValidatorHistoryRequest) (*types.GetValidatorHistoryReply, error) {
	if req.Size_ == 0 || req.Size_ > 1000 {
		req.Size_ = 1000
	}
	result, err := s.stakingService.GetValidatorHistories(ctx, common.HexToAddress(req.Validator), req.Offset, req.Size_)
	if err != nil {
		return nil, err
	}
	newSize, hasMore := checkHasMore(req.Size_, len(result))
	return &types.GetValidatorHistoryReply{ValidatorHistories: result[:newSize], HasMore: hasMore}, nil
}

func (s *Server) GetValidatorSlashings(ctx context.Context, req *types.GetValidatorSlashingsRequest) (*types.GetValidatorSlashingsReply, error) {
	if req.Size_ == 0 || req.Size_ > 1000 {
		req.Size_ = 1000
	}
	result, err := s.stakingService.GetValidatorSlashings(ctx, common.HexToAddress(req.Validator), req.Offset, req.Size_)
	if err != nil {
		return nil, err
	}
	newSize, hasMore := checkHasMore(req.Size_, len(result))
	return &types.GetValidatorSlashingsReply{ValidatorSlashings: result[:newSize], HasMore: hasMore}, nil
}

func (s *Server) GetValidatorDeposits(ctx context.Context, req *types.GetValidatorDepositsRequest) (*types.GetValidatorDepositsReply, error) {
	if req.Size_ == 0 || req.Size_ > 1000 {
		req.Size_ = 1000
	}
	result, err := s.stakingService.GetValidatorDeposits(ctx, common.HexToAddress(req.Validator), req.Offset, req.Size_)
	if err != nil {
		return nil, err
	}
	newSize, hasMore := checkHasMore(req.Size_, len(result))
	return &types.GetValidatorDepositsReply{ValidatorDeposits: result[:newSize], HasMore: hasMore}, nil
}

func (s *Server) GetValidatorEvents(ctx context.Context, req *types.GetValidatorEventsRequest) (*types.GetValidatorEventsReply, error) {
	if req.Size_ == 0 || req.Size_ > 1000 {
		req.Size_ = 1000
	}
	result, err := s.stakingService.GetValidatorEvents(ctx, common.HexToAddress(req.Validator), req.Offset, req.Size_)
	if err != nil {
		return nil, err
	}
	newSize, hasMore := checkHasMore(req.Size_, len(result))
	return &types.GetValidatorEventsReply{ValidatorEvents: result[:newSize], HasMore: hasMore}, nil
}

func (s *Server) GetDelegators(ctx context.Context, req *types.GetDelegatorsRequest) (*types.GetDelegatorsReply, error) {
	if req.Size_ == 0 || req.Size_ > 1000 {
		req.Size_ = 1000
	}
	var result []*types.Delegator
	var err error
	if req.Validator != "" {
		result, err = s.stakingService.GetDelegatorsByValidator(ctx, common.HexToAddress(req.Validator), req.Offset, req.Size_)
	} else {
		result, err = s.stakingService.GetDelegators(ctx, req.Offset, req.Size_)
	}
	if err != nil {
		return nil, err
	}
	newSize, hasMore := checkHasMore(req.Size_, len(result))
	return &types.GetDelegatorsReply{Delegators: result[:newSize], HasMore: hasMore}, nil
}

func (s *Server) GetChains(ctx context.Context, req *types.GetChainsRequest) (*types.GetChainsReply, error) {
	if req.Size_ == 0 || req.Size_ > 1000 {
		req.Size_ = 1000
	}
	if req.Chain != "" {
		chain, err := s.stakingService.GetChain(ctx, req.Chain)
		if err != nil {
			return nil, err
		}
		return &types.GetChainsReply{Chains: []*types.Chain{chain}, HasMore: false}, nil
	}
	result, err := s.stakingService.GetChains(ctx)
	if err != nil {
		return nil, err
	}
	newSize, hasMore := checkHasMore(req.Size_, len(result))
	return &types.GetChainsReply{Chains: result[:newSize], HasMore: hasMore}, nil
}

func (s *Server) GetStats(ctx context.Context, req *types.GetStatsRequest) (*types.GetStatsReply, error) {
	var err error
	stats := &types.StakingStats{
		TotalInsurance: "0",
	}
	if stats.ActiveUsers_7D, err = s.databaseService.EstimateActiveUsers(ctx, 7*time.Hour*24); err != nil {
		return nil, err
	}
	if stats.TotalHolders, err = s.databaseService.EstimateTokenHolders(ctx); err != nil {
		return nil, err
	}
	if stats.TotalInsurance, err = s.databaseService.GetTotalIssuance(ctx); err != nil {
		return nil, err
	}
	if stats.TotalValidators, err = s.stateDbService.GetTotalValidators(ctx); err != nil {
		return nil, err
	}
	if stats.TotalTxs, err = s.databaseService.EstimateTransactionCount(ctx); err != nil {
		return nil, err
	}
	if stats.TotalTransfers, err = s.databaseService.EstimateTransfersCount(ctx); err != nil {
		return nil, err
	}
	if stats.TotalStaked, err = s.stakingService.GetTotalStaked(ctx); err != nil {
		return nil, err
	}
	if stats.TransferVolume_24H, err = s.databaseService.GetTransferVolume(ctx, time.Hour*24); err != nil {
		return nil, err
	}
	if stats.MarketCap, err = s.stakingService.GetMarketCap(ctx); err != nil {
		return nil, err
	}
	if stats.KnownBlock, stats.AffectedBlock, _, err = s.stakingService.GetLatestBlock(ctx); err != nil {
		return nil, err
	}
	return &types.GetStatsReply{
		Stats: stats,
	}, nil
}

func (s *Server) GetTotalTxsGraph(ctx context.Context, _ *types.GetTotalTxsGraphRequest) (*types.GetTotalTxsGraphReply, error) {
	chainConfig := s.stakingService.GetChainConfig()
	latestKnownBlock, _, latestBlockTime, err := s.stakingService.GetLatestBlock(ctx)
	if err != nil {
		return nil, err
	}
	blocksMonthAgo := int64(latestKnownBlock) - 30*24*time.Hour.Milliseconds()/int64(chainConfig.AverageBlockTime)
	if blocksMonthAgo < 0 {
		blocksMonthAgo = int64(0)
	}
	res, err := s.databaseService.GetTransactionCountGraph(ctx, uint64(blocksMonthAgo), uint64(chainConfig.EpochBlockInterval))
	if err != nil {
		return nil, err
	}
	zeroBlockTime := latestBlockTime - latestKnownBlock*uint64(chainConfig.AverageBlockTime/1000)
	res2 := make(map[uint64]uint64)
	for k, v := range res {
		blockTime := zeroBlockTime + k*uint64(chainConfig.AverageBlockTime/1000)
		res2[blockTime] = v
	}
	return &types.GetTotalTxsGraphReply{Graph: res2}, nil
}
