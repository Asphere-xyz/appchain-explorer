package database

import (
	"context"
	"database/sql"
	"encoding/hex"
	"github.com/Ankr-network/ankr-protocol/shared"
	"github.com/Ankr-network/ankr-protocol/shared/entity"
	"github.com/Ankr-network/ankr-protocol/shared/types"
	_ "github.com/jackc/pgx/v4/stdlib"
	"strings"
	"time"
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
		count, err := s.CountTxsInBlock(ctx, int64(b.BlockNumber))
		if err != nil {
			return nil, err
		}
		b.TxsCount = uint32(count)
	}
	return result, err
}

func (s *Service) GetRecentTxs(ctx context.Context, timestamp int64, limit uint64) ([]*types.Transaction, error) {
	var txs []*entity.Transaction
	var err error
	from := time.Unix(timestamp, 0)
	//.Format("2006–01-02 15:04:03")
	txs, err = entity.TransactionsByTsLimit(ctx, s.db, from, limit)
	if err != nil {
		return nil, err
	}
	result := txsToProto(txs)
	return result, err
}

func (s *Service) getBlockWithTxsCount(ctx context.Context, block *entity.Block) (*types.BlockDetails, error) {
	result := blockDetailsToProto(block)
	count, err := s.CountTxsInBlock(ctx, int64(result.BlockNumber))
	if err != nil {
		return nil, err
	}
	result.TxsCount = uint32(count)
	return result, nil
}

func (s *Service) GetBlockByHash(ctx context.Context, blockHash string) (*types.BlockDetails, error) {
	hashBytes, err := hex.DecodeString(strings.TrimPrefix(blockHash, "0x"))
	if err != nil {
		return nil, err
	}
	block, err := entity.BlockByHash(ctx, s.db, hashBytes)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return s.getBlockWithTxsCount(ctx, block)
}

func (s *Service) GetBlockByNumber(ctx context.Context, blockNumber uint64) (*types.BlockDetails, error) {
	block, err := entity.BlockByNumber(ctx, s.db, int64(blockNumber))
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return s.getBlockWithTxsCount(ctx, block)
}

func (s *Service) CountTxsInBlock(ctx context.Context, blockNumber int64) (int, error) {
	txs, err := entity.TransactionsByBlockNumber(ctx, s.db, sql.NullInt64{Int64: blockNumber, Valid: true})
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
