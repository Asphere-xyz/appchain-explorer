package gateway

import (
	"context"
	"github.com/Ankr-network/ankr-protocol/shared/types"
)

func (s *Server) GetDelegators(ctx context.Context, req *types.GetDelegatorsRequest) (*types.GetDelegatorsReply, error) {
	result, err := s.stakingService.GetDelegators(ctx)
	if err != nil {
		return nil, err
	}
	return &types.GetDelegatorsReply{Delegators: result}, nil
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
