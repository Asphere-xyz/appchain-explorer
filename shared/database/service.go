package database

import (
	"context"
	"database/sql"
	"github.com/Ankr-network/ankr-protocol/shared"
	"github.com/Ankr-network/ankr-protocol/shared/entity"
	"github.com/Ankr-network/ankr-protocol/shared/types"
	_ "github.com/jackc/pgx/v4/stdlib"
)

type Service struct {
	db *sql.DB
}

func NewService() *Service {
	return &Service{}
}

func (s *Service) Start(cp shared.IConfigProvider) (err error) {
	config := &Config{}
	if err := cp.Parse(config); err != nil {
		return err
	}
	s.db, err = sql.Open("pgx", config.PostgresUrl)
	if err != nil {
		return err
	}
	return nil
}

func (s *Service) GetRecentBlocks(ctx context.Context, fromBlock, limit uint64) ([]*types.Block, error) {
	var blocks []*entity.Block
	var err error
	if fromBlock > 0 {
		blocks, err = entity.BlocksByNumberLimit(ctx, s.db, fromBlock, limit)
	} else {
		blocks, err = entity.BlocksByLimit(ctx, s.db, limit)
	}
	if err != nil {
		return nil, err
	}
	result := blocksToProto(blocks)
	for _, b := range result {
		count, err := s.CountTxsInBlock(ctx, b)
		if err != nil {
			return nil, err
		}
		b.TxsCount = uint32(count)
	}
	return result, err
}

func (s *Service) GetRecentTxs(ctx context.Context, fromTs string, limit uint64) ([]*types.Transaction, error) {
	var txs []*entity.Transaction
	var err error
	if fromTs != "" {
		txs, err = entity.TransactionsByTsLimit(ctx, s.db, fromTs, limit)
	} else {
		txs, err = entity.TransactionsByLimit(ctx, s.db, limit)
	}
	if err != nil {
		return nil, err
	}
	result := txsToProto(txs)
	return result, err
}

func (s *Service) GetBlockByHash(ctx context.Context, blockHash []byte) (*types.Block, error) {
	block, err := entity.BlockByHash(ctx, s.db, blockHash)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return blockToProto(block), err
}

func (s *Service) GetBlockByNumber(ctx context.Context, blockNumber uint64) (*types.Block, error) {
	block, err := entity.BlockByNumber(ctx, s.db, int64(blockNumber))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return blockToProto(block), err
}

func (s *Service) CountTxsInBlock(ctx context.Context, block *types.Block) (int, error) {
	txs, err := entity.TransactionsByBlockNumber(ctx, s.db, sql.NullInt64{Int64: int64(block.BlockNumber), Valid: true})
	return len(txs), err
}

func (s *Service) GetTransactionByHash(ctx context.Context, txHash []byte) (*types.Transaction, error) {
	tx, err := entity.TransactionByHash(ctx, s.db, txHash)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return txToProto(tx), err
}

func (s *Service) GetTokenTransfers(ctx context.Context, tokenContract []byte, block uint64, limit uint64) ([]*types.TokenTransfer, error) {
	transfers, err := entity.TokenTransfersByBlockNumber(ctx, s.db, sql.NullInt64{Int64: int64(block)})
	if err == sql.ErrNoRows {
		return nil, nil
	}
	result := tokenTransfersToProto(transfers)
	return result, err
}
