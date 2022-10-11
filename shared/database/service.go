package database

import (
	"context"
	"database/sql"
	"encoding/binary"
	"encoding/hex"
	"github.com/Ankr-network/ankr-protocol/shared"
	"github.com/Ankr-network/ankr-protocol/shared/entity"
	"github.com/Ankr-network/ankr-protocol/shared/types"
	"github.com/jackc/pgx/v4"
	_ "github.com/jackc/pgx/v4/stdlib"
	"github.com/samber/lo"
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
	txs, err = entity.TransactionsByTsLimit(ctx, s.db, from, limit)
	if err != nil {
		return nil, err
	}
	result := txsToProto(txs)
	for i, tx := range txs {
		transfers, err := s.getTokenTransfersForTx(ctx, tx.Hash)
		if err != nil {
			return nil, err
		}
		tokens := make([]*entity.Token, len(transfers))
		for k, transfer := range transfers {
			token, err := entity.TokenByContractAddressHash(ctx, s.db, transfer.TokenContractAddressHash)
			if err != nil {
				return nil, err
			}
			tokens[k] = token
		}
		result[i].TokenTransfers = tokenTransfersToProto(transfers, tokens)
	}

	return result, err
}

func (s *Service) GetTransactionByHash(ctx context.Context, txHash string) (*types.TransactionDetails, error) {
	hashBytes, err := hex.DecodeString(strings.TrimPrefix(txHash, "0x"))
	if err != nil {
		return nil, err
	}
	tx, err := entity.TransactionByHash(ctx, s.db, hashBytes)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	result := txDetailsToProto(tx)
	transfers, err := s.getTokenTransfersForTx(ctx, tx.Hash)
	if err != nil {
		return nil, err
	}
	tokens := make([]*entity.Token, len(transfers))
	for k, transfer := range transfers {
		token, err := entity.TokenByContractAddressHash(ctx, s.db, transfer.TokenContractAddressHash)
		if err != nil {
			return nil, err
		}
		tokens[k] = token
	}
	result.TokenTransfers = tokenTransfersToProto(transfers, tokens)

	if len(tx.Input) == 0 {
		return result, err
	}
	//Try extract signature, methodname and args
	sig := tx.Input[0:4]
	identifier := int32(binary.LittleEndian.Uint32(sig))
	methods, err := entity.ContractMethodsByIdentifier(ctx, s.db, identifier)
	if err != nil {
		return nil, err
	}
	if len(methods) > 0 {
		methodSignature, args := argsToProto(methods[0], tx.Input)
		result.Args = args
		result.Signature = methodSignature
	}
	return result, err
}

func (s *Service) getBlockWithTxsCount(ctx context.Context, block *entity.Block) (*types.BlockDetails, error) {
	result := blockDetailsToProto(block)
	if result == nil {
		return nil, nil
	}
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

func (s *Service) getTokenTransfersForTx(ctx context.Context, hash []byte) ([]*entity.TokenTransfer, error) {
	transfers, err := entity.TokenTransfersByHash(ctx, s.db, hash)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return transfers, err
}

func (s *Service) getTokensForAddress(ctx context.Context, hash []byte) ([]*entity.AddressTokenBalance, error) {
	tokens, err := entity.AddressTokenBalancesByHash(ctx, s.db, hash)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return tokens, err
}

func (s *Service) getBlockValidatedByAddress(ctx context.Context, hash []byte) ([]*entity.Block, error) {
	tokens, err := entity.BlocksByMinerHash(ctx, s.db, hash)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	return tokens, err
}

func (s *Service) getTokenTransfersCountByAddress(ctx context.Context, hash []byte) (uint32, error) {
	count, err := entity.TokenTransferCountByAddressHash(ctx, s.db, hash)
	if err == sql.ErrNoRows {
		return 0, nil
	}
	return count, err
}

func (s *Service) GetAddress(ctx context.Context, hash string) (*types.Address, error) {
	hashBytes, err := hex.DecodeString(strings.TrimPrefix(hash, "0x"))
	if err != nil {
		return nil, err
	}
	addressInfo, err := entity.AddressByHash(ctx, s.db, hashBytes)
	if err == sql.ErrNoRows {
		return nil, nil
	}
	result := addressToProto(addressInfo)
	tokens, err := s.getTokensForAddress(ctx, hashBytes)
	if err != nil {
		return nil, err
	}
	validatedBlocks, err := s.getBlockValidatedByAddress(ctx, hashBytes)
	if err != nil {
		return nil, err
	}
	tokenTransfersCount, err := entity.TokenTransferCountByAddressHash(ctx, s.db, hashBytes)
	if err != nil {
		return nil, err
	}
	result.Tokens = uint32(len(tokens))
	result.BlocksValidated = uint32(len(validatedBlocks))
	result.TokenTransfers = tokenTransfersCount
	return result, nil
}

func (s *Service) EstimateTransactionCount(ctx context.Context) (uint64, error) {
	row := s.db.QueryRowContext(ctx, "SELECT reltuples::BIGINT AS estimate FROM pg_class WHERE relname='transactions'")
	if row.Err() == pgx.ErrNoRows {
		return 0, nil
	} else if row.Err() != nil {
		return 0, row.Err()
	}
	var res int64
	if err := row.Scan(&res); err != nil {
		return 0, err
	}
	return uint64(res), nil
}

func (s *Service) GetTransactionCountGraph(ctx context.Context, afterBlock, epochLength uint64) (map[uint64]uint64, error) {
	rows, err := s.db.QueryContext(ctx, "SELECT block_number / $1, count(*) AS count FROM transactions WHERE block_number >= $2 GROUP BY block_number / $1", epochLength, afterBlock)
	if err == pgx.ErrNoRows {
		return nil, nil
	} else if err != nil {
		return nil, err
	}
	defer func() { _ = rows.Close() }()
	type scanResult struct {
		BlockNumber int64
		Count       int64
	}
	var res []scanResult
	for rows.Next() {
		var t scanResult
		if err := rows.Scan(&t.BlockNumber, &t.Count); err != nil {
			return nil, err
		}
		res = append(res, t)
	}
	return lo.SliceToMap(res, func(t scanResult) (uint64, uint64) {
		return uint64(t.BlockNumber) * epochLength, uint64(t.Count)
	}), nil
}

func (s *Service) EstimateTransfersCount(ctx context.Context) (uint64, error) {
	return s.EstimateTransactionCount(ctx)
}

func (s *Service) EstimateTokenHolders(ctx context.Context) (uint64, error) {
	return 0, nil
}
