package ethereuminteraction

import (
	"fmt"

	ABI "github.com/matterinc/PlasmaBlockVerifier/contractABI"
	"github.com/matterinc/PlasmaBlockVerifier/messageStructures"
	"github.com/matterinc/PlasmaCommons/transaction"
)

type WithdrawChallengeProcessor struct {
	plasmaParent *ABI.PlasmaParent
}

func NewWithdrawChallengeProcessor(plasmaParent *ABI.PlasmaParent) *WithdrawChallengeProcessor {
	newInstance := &WithdrawChallengeProcessor{plasmaParent}
	return newInstance
}

func (p *WithdrawChallengeProcessor) Process(challengeRequest *messageStructures.WithdrawChallengeRequest) (bool, error) {
	var utxoIndex [transaction.UTXOIndexLength]byte
	copy(challengeRequest.UtxoIndex[4:], utxoIndex[:])
	details := transaction.ParseIndexIntoUTXOdetails(utxoIndex)
	fmt.Println("Processign withdraw for " + fmt.Sprintln(details))
	return true, nil
}
