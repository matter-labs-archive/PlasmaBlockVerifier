package messageStructures

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type WithdrawStartedInformation struct {
	IsLimbo    bool
	Index      *big.Int
	LimboIndex []byte
	From       common.Address
	Amount     *big.Int
}
