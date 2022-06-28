package database

import (
	"github.com/Ankr-network/ankr-protocol/shared/entity"
	"github.com/Ankr-network/ankr-protocol/shared/types"
	"strconv"
)

func blocksToProto(blocks []*entity.Block) (result []*types.Block) {
	for _, b := range blocks {
		result = append(result, blockToProto(b))
	}
	return
}

func blockToProto(block *entity.Block) *types.Block {
	if block == nil {
		return nil
	}
	return &types.Block{
		BlockHash:   block.Hash,
		BlockNumber: uint64(block.Number),
		Coinbase:    block.MinerHash,
		Difficulty:  uint64(block.Difficulty.Float64),
		Nonce:       block.Nonce,
		ParentHash:  block.ParentHash,
		Timestamp:   uint64(block.Timestamp.Unix()),
		GasLimit:    uint64(block.GasLimit),
		GasUsed:     uint64(block.GasUsed),
		SizeBytes:   uint32(block.Size.Int64),
	}
}

func txToProto(tx *entity.Transaction) *types.Transaction {
	if tx == nil {
		return nil
	}
	return &types.Transaction{
		TxHash:            tx.Hash,
		Status:            types.TransactionStatus(tx.Status.Int64),
		Gas:               uint64(tx.Gas),
		CumulativeGasUsed: uint64(tx.CumulativeGasUsed.Float64),
		GasUsed:           uint64(tx.GasUsed.Float64),
		GasPrice:          uint64(tx.GasPrice),
		TxIndex:           uint64(tx.Index.Int64),
		Input:             tx.Input,
		Nonce:             uint64(tx.Nonce),
		Value:             strconv.FormatFloat(tx.Value, 'f', 18, 64),
		Error:             tx.Error.String,
		BlockHash:         tx.BlockHash,
		BlockNumber:       uint64(tx.BlockNumber.Int64),
		Sender:            tx.FromAddressHash,
		Recipient:         tx.ToAddressHash,
		Contract:          tx.CreatedContractAddressHash,
		RevertReason:      tx.RevertReason.String,
		Type:              uint32(tx.Type.Int64),
		InternalFailed:    tx.HasErrorInInternalTxs.Bool,
	}
}
