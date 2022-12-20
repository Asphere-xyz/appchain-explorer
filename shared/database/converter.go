package database

import (
	"bytes"
	"fmt"
	"github.com/Ankr-network/ankr-protocol/shared/entity"
	"github.com/Ankr-network/ankr-protocol/shared/types"
	"github.com/ethereum/go-ethereum/accounts/abi"
	log "github.com/sirupsen/logrus"
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
		BlockHash:   fmt.Sprintf("0x%x", block.Hash),
		BlockNumber: uint32(block.Number),
		MinerHash:   fmt.Sprintf("0x%x", block.MinerHash),
		Timestamp:   uint64(block.Timestamp.Unix()),
		GasLimit:    uint64(block.GasLimit),
		GasUsed:     uint64(block.GasUsed),
		SizeBytes:   uint32(block.Size.Int64),
	}
}

func blockDetailsToProto(block *entity.Block) *types.BlockDetails {
	if block == nil {
		return nil
	}
	return &types.BlockDetails{
		BlockHash:       fmt.Sprintf("0x%x", block.Hash),
		BlockNumber:     uint32(block.Number),
		MinerHash:       fmt.Sprintf("0x%x", block.MinerHash),
		Timestamp:       uint64(block.Timestamp.Unix()),
		GasLimit:        uint64(block.GasLimit),
		GasUsed:         uint64(block.GasUsed),
		SizeBytes:       uint32(block.Size.Int64),
		ParentHash:      fmt.Sprintf("0x%x", block.ParentHash),
		Difficulty:      uint32(block.Difficulty.Float64),
		TotalDifficulty: uint32(block.TotalDifficulty.Float64),
		Nonce:           fmt.Sprintf("0x%x", block.Nonce),
	}
}

func txsToProto(txs []*entity.Transaction) (result []*types.Transaction) {
	for _, tx := range txs {
		result = append(result, txToProto(tx))
	}
	return
}

func txToProto(tx *entity.Transaction) *types.Transaction {
	if tx == nil {
		return nil
	}
	txType := types.TransactionType_UNKNOWN
	if tx.Input != nil {
		txType = types.TransactionType_CONTRACT_CALL
	}
	status := int64(2)                           // pending by default
	if tx.Status.Valid && tx.BlockNumber.Valid { // if confirmed - set status from table
		status = tx.Status.Int64
	}
	return &types.Transaction{
		TxHash:      fmt.Sprintf("0x%x", tx.Hash),
		Status:      types.TransactionStatus(status),
		Value:       strconv.FormatFloat(tx.Value/1e18, 'f', 18, 64),
		TxFee:       strconv.FormatFloat(tx.GasUsed.Float64*tx.GasPrice/1e18, 'f', 18, 64),
		BlockNumber: uint32(tx.BlockNumber.Int64),
		Timestamp:   uint32(tx.UpdatedAt.Unix()),
		Error:       tx.Error.String,
		Sender:      fmt.Sprintf("0x%x", tx.FromAddressHash),
		Recipient:   fmt.Sprintf("0x%x", tx.ToAddressHash),
		TxType:      txType,
	}
}

func txDetailsToProto(tx *entity.Transaction) *types.TransactionDetails {
	if tx == nil {
		return nil
	}
	methodId := ""
	if len(tx.Input) != 0 {
		methodId = fmt.Sprintf("0x%x", tx.Input[0:4])
	}
	status := int64(2)                           // pending by default
	if tx.Status.Valid && tx.BlockNumber.Valid { // if confirmed - set status from table
		status = tx.Status.Int64
	}
	return &types.TransactionDetails{
		TxHash:      fmt.Sprintf("0x%x", tx.Hash),
		Status:      types.TransactionStatus(status),
		Error:       tx.Error.String,
		BlockNumber: uint32(tx.BlockNumber.Int64),
		CreatedAt:   uint32(tx.InsertedAt.Unix()),
		ConfirmedAt: uint32(tx.UpdatedAt.Unix()),
		Sender:      fmt.Sprintf("0x%x", tx.FromAddressHash),
		Recipient:   fmt.Sprintf("0x%x", tx.ToAddressHash),
		Value:       strconv.FormatFloat(tx.Value/1e18, 'f', 18, 64),
		TxFee:       strconv.FormatFloat(tx.GasUsed.Float64*tx.GasPrice/1e18, 'f', 18, 64),
		GasPrice:    uint64(tx.GasPrice),
		GasLimit:    uint64(tx.Gas),
		GasUsed:     uint64(tx.GasUsed.Float64),
		Type:        uint32(tx.Type.Int64),
		Nonce:       uint32(tx.Nonce),
		Index:       uint32(tx.Index.Int64),
		MethodId:    methodId,
		RawInput:    fmt.Sprintf("0x%x", tx.Input),
	}
}

func tokenTransferToProto(transfer *entity.TokenTransfer, token *entity.Token) *types.TokenTransfer {
	if transfer == nil {
		return nil
	}
	symbol := ""
	if token.Symbol.Valid {
		symbol = token.Symbol.String
	}
	return &types.TokenTransfer{
		AddressFrom:   fmt.Sprintf("0x%x", transfer.FromAddressHash),
		AddressTo:     fmt.Sprintf("0x%x", transfer.ToAddressHash),
		TokenContract: fmt.Sprintf("0x%x", transfer.TokenContractAddressHash),
		Amount:        uint64(transfer.Amount.Float64),
		Symbol:        symbol,
		Decimals:      uint64(token.Decimals.Float64),
	}
}

func tokenTransfersToProto(transfers []*entity.TokenTransfer, tokens []*entity.Token) (result []*types.TokenTransfer) {
	for i, b := range transfers {
		result = append(result, tokenTransferToProto(b, tokens[i]))
	}
	return
}

func addressToProto(address *entity.Address) (result *types.Address) {
	if address == nil {
		return nil
	}
	return &types.Address{
		Hash:         fmt.Sprintf("0x%x", address.Hash),
		Balance:      address.FetchedCoinBalance.String,
		Transactions: uint32(address.TransactionsCount.Int64),
		GasUsed:      uint64(address.GasUsed.Int64),
	}
}

func argsToProto(contractMethod *entity.ContractMethod, raw []byte) (string, []*types.CallArg) {
	if contractMethod == nil || len(raw) == 0 {
		return "", nil
	}
	decodedABI, err := abi.JSON(bytes.NewReader(contractMethod.Abi))
	if err != nil {
		log.Printf("Failed to parse abi: %v", err)
		return "", nil
	}
	method, err := decodedABI.MethodById(raw)
	if err != nil {
		log.Printf("Failed to decode method: %v", err)
		return "", nil
	}
	values, err := method.Inputs.Unpack(raw)
	if err != nil {
		log.Printf("Failed to unpack values: %v", err)
		return "", nil
	}
	inputs := make([]*types.CallArg, len(method.Inputs))
	for i, arg := range method.Inputs {
		inputs[i] = &types.CallArg{
			Name: arg.Name,
			Type: arg.Type.String(),
			Data: fmt.Sprintf("%v", values[i]),
		}
	}
	return method.Sig, inputs
}
