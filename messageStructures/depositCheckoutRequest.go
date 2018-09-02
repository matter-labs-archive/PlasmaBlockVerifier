package messageStructures

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type DepositIndexCheckoutRequest struct {
	Index    *big.Int
	Value    *big.Int
	Address  common.Address
	Operator common.Address
}
