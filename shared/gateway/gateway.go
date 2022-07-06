package gateway

import (
	"context"
	"github.com/Ankr-network/ankr-protocol/shared/types"
)

func checkHasMore(limit uint32, length int) (int, bool) {
	hasMore := length > int(limit)
	if length <= int(limit) {
		return length, hasMore
	}
	return int(limit), hasMore
}

const DefaultRecentBlockLimit = 10
const MaxRecentBlockLimit = 100

func (s *Service) GetRecentBlocks(ctx context.Context, req *types.GetRecentBlocksRequest) (*types.GetRecentBlocksReply, error) {
	if req.Limit == 0 {
		req.Limit = DefaultRecentBlockLimit
	} else if req.Limit > MaxRecentBlockLimit {
		req.Limit = MaxRecentBlockLimit
	}
	blocks, err := s.databaseService.GetRecentBlocks(ctx, req.FromBlock, req.Limit)
	return &types.GetRecentBlocksReply{Blocks: blocks}, err
}

func (s *Service) GetBlockByHashOrNumber(ctx context.Context, req *types.GetBlockByHashOrNumberRequest) (*types.GetBlockByHashOrNumberReply, error) {
	var block *types.Block
	var err error
	if len(req.Hash) > 0 {
		block, err = s.databaseService.GetBlockByHash(ctx, req.Hash)
	} else {
		block, err = s.databaseService.GetBlockByNumber(ctx, req.Number)
	}
	return &types.GetBlockByHashOrNumberReply{Block: block}, err
}

func (s *Service) GetTransactionByHash(ctx context.Context, req *types.GetTransactionByHashRequest) (*types.GetTransactionByHashReply, error) {
	tx, err := s.databaseService.GetTransactionByHash(ctx, req.Hash)
	return &types.GetTransactionByHashReply{Transaction: tx}, err
}

func (s *Service) GetTokenTransfers(ctx context.Context, req *types.GetTokenTransfersRequest) (*types.GetTokenTransfersReply, error) {
	transfers, err := s.databaseService.GetTokenTransfers(ctx, req.TokenContract, req.FromBlock, req.Limit)
	return &types.GetTokenTransfersReply{
		Transfers: transfers,
	}, err
}
