package gateway

import (
	"context"
	"github.com/Ankr-network/ankr-protocol/shared/types"
	"time"
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

const DefaultRecentTxsLimit = 10
const MaxRecentTxLimit = 100

func (s *Server) GetRecentBlocks(ctx context.Context, req *types.GetRecentBlocksRequest) (*types.GetRecentBlocksReply, error) {
	if req.Limit == 0 {
		req.Limit = DefaultRecentBlockLimit
	} else if req.Limit > MaxRecentBlockLimit {
		req.Limit = MaxRecentBlockLimit
	}
	blocks, err := s.databaseService.GetRecentBlocks(ctx, req.FromBlock, req.Limit)
	return &types.GetRecentBlocksReply{Blocks: blocks}, err
}

func (s *Server) GetRecentTxs(ctx context.Context, req *types.GetRecentTxsRequest) (*types.GetRecentTxsReply, error) {
	if req.Limit == 0 {
		req.Limit = DefaultRecentTxsLimit
	} else if req.Limit > MaxRecentTxLimit {
		req.Limit = MaxRecentTxLimit
	}
	if req.FromTs == 0 {
		req.FromTs = uint64(time.Now().Unix())
	}
	txs, err := s.databaseService.GetRecentTxs(ctx, int64(req.FromTs), req.Limit)
	return &types.GetRecentTxsReply{Txs: txs}, err
}

func (s *Server) GetBlockByHashOrNumber(ctx context.Context, req *types.GetBlockByHashOrNumberRequest) (*types.GetBlockByHashOrNumberReply, error) {
	var block *types.BlockDetails
	var err error
	if len(req.Hash) > 0 {
		block, err = s.databaseService.GetBlockByHash(ctx, req.Hash)
	} else {
		block, err = s.databaseService.GetBlockByNumber(ctx, req.Number)
	}
	return &types.GetBlockByHashOrNumberReply{Block: block}, err
}

func (s *Server) GetTransactionByHash(ctx context.Context, req *types.GetTransactionByHashRequest) (*types.GetTransactionByHashReply, error) {
	tx, err := s.databaseService.GetTransactionByHash(ctx, req.Hash)
	return &types.GetTransactionByHashReply{Transaction: tx}, err
}

func (s *Server) GetAddress(ctx context.Context, req *types.GetAddressRequest) (*types.GetAddressReply, error) {
	address, err := s.databaseService.GetAddress(ctx, req.Hash)
	return &types.GetAddressReply{
		Address: address,
	}, err
}

func (s *Server) GetTokenTransfersByAddress(ctx context.Context, req *types.GetTokenTransfersByAddressRequest) (*types.GetTokenTransfersByAddressReply, error) {
	//address, err := s.databaseService.GetAddress(ctx, req.Hash)
	//return &types.GetAddressReply{
	//	Address: address,
	//}, err
	return nil, nil
}
