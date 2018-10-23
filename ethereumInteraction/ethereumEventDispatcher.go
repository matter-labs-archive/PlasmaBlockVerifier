package ethereuminteraction

import (
	"context"
	"log"
	"math/big"
	"time"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"

	"github.com/ethereum/go-ethereum/common"
	ethclient "github.com/ethereum/go-ethereum/ethclient"
	ABI "github.com/matterinc/PlasmaBlockVerifier/contractABI"
	messageStructures "github.com/matterinc/PlasmaBlockVerifier/messageStructures"
)

const loopDelay = 100000000 // 0.1 second

type EthereumNetworkEventDispatcher struct {
	ConnectionString     string
	Client               *ethclient.Client
	SignalChannel        chan int
	PlasmaParentContract *ABI.PlasmaParent
	BlockStorageContract *ABI.PlasmaBlockStorage
}

func NewEthereumNetworkEventDispatcher(connection, contractAddress string) *EthereumNetworkEventDispatcher {
	conn, err := ethclient.Dial(connection)
	if err != nil {
		return nil
	}
	signalChan := make(chan int)
	address := common.HexToAddress(contractAddress)
	plasma, err := ABI.NewPlasmaParent(address, conn)
	blockStorageAddress, err := plasma.PlasmaParentCaller.BlockStorage(nil)
	if err != nil {
		log.Fatalln("Can not get block storage contract address")
	}
	blockStorage, err := ABI.NewPlasmaBlockStorage(blockStorageAddress, conn)
	newInstance := &EthereumNetworkEventDispatcher{connection, conn, signalChan, plasma, blockStorage}
	return newInstance
}

func (p *EthereumNetworkEventDispatcher) Run(fromBlockNumber int64, blockProcessorLoopChannel chan<- *messageStructures.BlockInformation, withdrawProcessor *WithdrawStartedProcessor) {
	lastBlockNumber := big.NewInt(fromBlockNumber - 1)
	if fromBlockNumber <= 1 {
		lastBlockNumber.SetUint64(0)
	}
	one := big.NewInt(1)
	eventLoop := func() {
		for {
			select {
			case _ = <-p.SignalChannel:
				log.Println("Received control message")
				// log.Println(strconv.Itoa(controlMessage))
				return
			default:
				newBlockNumber := new(big.Int).Add(lastBlockNumber, one)
				_, err := p.Client.HeaderByNumber(context.TODO(), newBlockNumber)
				if err != nil {
					time.Sleep(loopDelay)
					continue
				}
				if newBlockNumber.Uint64()%10 == 0 {
					log.Println("Processing Ethereum block " + newBlockNumber.String())
				}
				blockNumberUInt64 := newBlockNumber.Uint64()

				filterOptions := &bind.FilterOpts{blockNumberUInt64, &blockNumberUInt64, context.TODO()}
				exitEventsIterator, err := p.PlasmaParentContract.PlasmaParentFilterer.FilterExitStartedEvent(filterOptions, nil, nil, nil)
				if err != nil {
					log.Println(err)
					continue
				}
				// iterate over exits first
				for exitEventsIterator.Next() {
					ev := exitEventsIterator.Event
					log.Println("Processing exit event")
					// this is a BUG in ethereum client, it unpacks bytesN type from the last byte, not the first

					exitRecordHash := ev.Raw.Topics[3][:22]
					exitRecordFixedLength := [22]byte{}
					copy(exitRecordFixedLength[:], exitRecordHash)
					fullInfo, err := p.PlasmaParentContract.PlasmaParentCaller.ExitRecords(nil, exitRecordFixedLength)
					if err != nil {
						log.Println("Can not get exit record data")
						log.Fatalln(err)
					}
					info := &messageStructures.WithdrawStartedInformation{false, fullInfo.BlockNumber, fullInfo.TransactionNumber, fullInfo.OutputNumber, ev.From, fullInfo.Amount, exitRecordHash}
					_, err = withdrawProcessor.Process(info)
					if err != nil {
						log.Println("Exit processor returned an error")
						log.Fatalln(err)
					}
				}
				blocksEventIterator, err := p.BlockStorageContract.PlasmaBlockStorageFilterer.FilterBlockHeaderSubmitted(filterOptions, nil, nil)
				if err != nil {
					log.Println(err)
					continue
				}
				for blocksEventIterator.Next() {
					ev := blocksEventIterator.Event
					log.Println("Processing Plasma block " + ev.BlockNumber.String() + " in Ethereum block " + newBlockNumber.String())
					plasmaBlockNumber := uint32(ev.BlockNumber.Uint64())
					merkleRoot := ev.MerkleRoot
					blockProcessingInformation := &messageStructures.BlockInformation{BlockNumber: plasmaBlockNumber, BlockHash: [32]byte{}, BlockMerkleRoot: merkleRoot}
					blockProcessorLoopChannel <- blockProcessingInformation
				}
				lastBlockNumber = newBlockNumber
			}
		}
	}
	go eventLoop()
}
