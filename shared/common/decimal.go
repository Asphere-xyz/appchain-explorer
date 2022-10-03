package common

import (
	"github.com/ethereum/go-ethereum/common/math"
	"github.com/shopspring/decimal"
	"math/big"
)

func ToEtherDecimal(val *big.Int) decimal.Decimal {
	return decimal.NewFromBigInt(val, 0).Div(decimal.NewFromBigInt(math.BigPow(10, 18), 0))
}

func ToEther(val *big.Int) string {
	return ToEtherDecimal(val).String()
}
