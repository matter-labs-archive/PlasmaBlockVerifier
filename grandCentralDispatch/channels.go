package grandCentralDispatch

import (
	messageStructures "github.com/matterinc/PlasmaBlockVerifier/messageStructures"
)

var ControlChannel = make(chan int)
var DepositsCheckoutsChannel = make(chan *messageStructures.DepositIndexCheckoutRequest)
var WithdrawStartedChannel = make(chan *messageStructures.WithdrawStartedInformation)
var WithdrawChallengesChannel = make(chan *messageStructures.WithdrawChallengeRequest)
var BlockInformationChannel = make(chan *messageStructures.BlockInformation)
