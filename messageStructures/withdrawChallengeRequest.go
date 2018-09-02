package messageStructures

type WithdrawChallengeRequest struct {
	UtxoIndex                []byte
	SpendingTransactionIndex []byte
}
