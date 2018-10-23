package messageStructures

type WithdrawChallengeRequest struct {
	PartialHash []byte
	// UtxoBlock                 uint32
	// UtxoTransaction           uint32
	// UtxoOutput                uint8
	// SpendingBlock             uint32
	// SpendingTransaction       uint32
	// SpendingInput             uint8
	// SpendingInformationFilled bool
	UtxoIndex                []byte
	SpendingTransactionIndex []byte
}
