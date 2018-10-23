package messageStructures

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type WithdrawStartedInformation struct {
	IsLimbo           bool
	BlockNumber       uint32
	TransactionNumber uint32
	OutputNumber      uint8
	From              common.Address
	Amount            *big.Int
	PartialHash       []byte
}
