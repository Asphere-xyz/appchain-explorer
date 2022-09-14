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
