// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contractABI

import (
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// PlasmaParentABI is the input ABI used to generate the binding from.
const PlasmaParentABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"buyoutProcessorContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"WithdrawCollateral\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"OutputChallangesDelay\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"plasmaErrorFound\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastValidBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"uint32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DepositWithdrawDelay\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"ExitDelay\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"blockStorage\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"operatorsBond\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes22\"}],\"name\":\"succesfulExits\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"limboExitContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"depositRecords\",\"outputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"status\",\"type\":\"uint8\"},{\"name\":\"hasCollateral\",\"type\":\"bool\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"withdrawStartedAt\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"DepositWithdrawCollateral\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes22\"}],\"name\":\"exitBuyoutOffers\",\"outputs\":[{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"accepted\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"challengesContract\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"bytes22\"}],\"name\":\"exitRecords\",\"outputs\":[{\"name\":\"transactionRef\",\"type\":\"bytes32\"},{\"name\":\"amount\",\"type\":\"uint256\"},{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"timePublished\",\"type\":\"uint64\"},{\"name\":\"blockNumber\",\"type\":\"uint32\"},{\"name\":\"transactionNumber\",\"type\":\"uint32\"},{\"name\":\"outputNumber\",\"type\":\"uint8\"},{\"name\":\"isValid\",\"type\":\"bool\"},{\"name\":\"isLimbo\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allDepositRecordsForUser\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"InputChallangesDelay\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"depositCounter\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"\",\"type\":\"address\"},{\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"allExitsForUser\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes22\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"exitQueue\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_priorityQueue\",\"type\":\"address\"},{\"name\":\"_blockStorage\",\"type\":\"address\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"constructor\"},{\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"fallback\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_lastValidBlockNumber\",\"type\":\"uint256\"}],\"name\":\"ErrorFoundEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_amount\",\"type\":\"uint256\"},{\"indexed\":true,\"name\":\"_depositIndex\",\"type\":\"uint256\"}],\"name\":\"DepositEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_depositIndex\",\"type\":\"uint256\"}],\"name\":\"DepositWithdrawStartedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_depositIndex\",\"type\":\"uint256\"}],\"name\":\"DepositWithdrawChallengedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_depositIndex\",\"type\":\"uint256\"}],\"name\":\"DepositWithdrawCompletedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_hash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"_data\",\"type\":\"bytes\"}],\"name\":\"TransactionPublished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_hash\",\"type\":\"bytes22\"}],\"name\":\"ExitRecordCreated\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_hash\",\"type\":\"bytes22\"}],\"name\":\"ExitChallenged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_index\",\"type\":\"uint64\"}],\"name\":\"TransactionIsPublished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"_priority\",\"type\":\"uint72\"},{\"indexed\":true,\"name\":\"_index\",\"type\":\"uint72\"},{\"indexed\":true,\"name\":\"_hash\",\"type\":\"bytes22\"}],\"name\":\"ExitStartedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_priority\",\"type\":\"uint72\"},{\"indexed\":true,\"name\":\"_partialHash\",\"type\":\"bytes22\"}],\"name\":\"LimboExitStartedEvent\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_partialHash\",\"type\":\"bytes22\"},{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"_buyoutAmount\",\"type\":\"uint256\"}],\"name\":\"ExitBuyoutOffered\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"_partialHash\",\"type\":\"bytes22\"},{\"indexed\":true,\"name\":\"_from\",\"type\":\"address\"}],\"name\":\"ExitBuyoutAccepted\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"_op\",\"type\":\"address\"},{\"name\":\"_status\",\"type\":\"uint256\"}],\"name\":\"setOperator\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_buyouts\",\"type\":\"address\"}],\"name\":\"allowDeposits\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_challenger\",\"type\":\"address\"}],\"name\":\"allowChallenges\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_limboExiter\",\"type\":\"address\"}],\"name\":\"allowLimboExits\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_headers\",\"type\":\"bytes\"}],\"name\":\"submitBlockHeaders\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"lastBlockNumber\",\"outputs\":[{\"name\":\"blockNumber\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"hashOfLastSubmittedBlock\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"incrementWeekOldCounter\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_plasmaBlockNumber\",\"type\":\"uint32\"},{\"name\":\"_outputNumber\",\"type\":\"uint8\"},{\"name\":\"_plasmaTransaction\",\"type\":\"bytes\"},{\"name\":\"_merkleProof\",\"type\":\"bytes\"}],\"name\":\"startExit\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_exitRecordHash\",\"type\":\"bytes22\"},{\"name\":\"_plasmaBlockNumber\",\"type\":\"uint32\"},{\"name\":\"_plasmaTransaction\",\"type\":\"bytes\"},{\"name\":\"_merkleProof\",\"type\":\"bytes\"},{\"name\":\"_inputNumber\",\"type\":\"uint8\"}],\"name\":\"challengeNormalExitByShowingExitBeingSpent\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_originalTransaction\",\"type\":\"bytes\"},{\"name\":\"_originalInputNumber\",\"type\":\"uint8\"},{\"name\":\"_exitRecordHash\",\"type\":\"bytes22\"},{\"name\":\"_plasmaBlockNumber\",\"type\":\"uint32\"},{\"name\":\"_plasmaTransaction\",\"type\":\"bytes\"},{\"name\":\"_merkleProof\",\"type\":\"bytes\"},{\"name\":\"_inputNumber\",\"type\":\"uint8\"}],\"name\":\"challengeNormalExitByShowingAnInputDoubleSpend\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_originalTransaction\",\"type\":\"bytes\"},{\"name\":\"_originalInputNumber\",\"type\":\"uint8\"},{\"name\":\"_exitRecordHash\",\"type\":\"bytes22\"},{\"name\":\"_plasmaBlockNumber\",\"type\":\"uint32\"},{\"name\":\"_plasmaTransaction\",\"type\":\"bytes\"},{\"name\":\"_merkleProof\",\"type\":\"bytes\"},{\"name\":\"_outputNumber\",\"type\":\"uint8\"}],\"name\":\"challengeNormalExitByShowingMismatchedInput\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_numOfExits\",\"type\":\"uint256\"}],\"name\":\"finalizeExits\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_plasmaTransaction\",\"type\":\"bytes\"}],\"name\":\"isWellFormedTransaction\",\"outputs\":[{\"name\":\"isWellFormed\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"deposit\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_for\",\"type\":\"address\"}],\"name\":\"depositFor\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_index\",\"type\":\"bytes22\"},{\"name\":\"_beneficiary\",\"type\":\"address\"}],\"name\":\"offerOutputBuyout\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_index\",\"type\":\"bytes22\"}],\"name\":\"acceptBuyoutOffer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_index\",\"type\":\"bytes22\"}],\"name\":\"returnExpiredBuyoutOffer\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"depositIndex\",\"type\":\"uint256\"}],\"name\":\"startDepositWithdraw\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":true,\"stateMutability\":\"payable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"depositIndex\",\"type\":\"uint256\"}],\"name\":\"finalizeDepositWithdraw\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"depositIndex\",\"type\":\"uint256\"},{\"name\":\"_plasmaBlockNumber\",\"type\":\"uint32\"},{\"name\":\"_plasmaTransaction\",\"type\":\"bytes\"},{\"name\":\"_merkleProof\",\"type\":\"bytes\"}],\"name\":\"challengeDepositWithdraw\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_plasmaBlockNumber1\",\"type\":\"uint32\"},{\"name\":\"_inputNumber1\",\"type\":\"uint8\"},{\"name\":\"_plasmaTransaction1\",\"type\":\"bytes\"},{\"name\":\"_merkleProof1\",\"type\":\"bytes\"},{\"name\":\"_plasmaBlockNumber2\",\"type\":\"uint32\"},{\"name\":\"_inputNumber2\",\"type\":\"uint8\"},{\"name\":\"_plasmaTransaction2\",\"type\":\"bytes\"},{\"name\":\"_merkleProof2\",\"type\":\"bytes\"}],\"name\":\"proveDoubleSpend\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_plasmaBlockNumber\",\"type\":\"uint32\"},{\"name\":\"_plasmaTransaction\",\"type\":\"bytes\"},{\"name\":\"_merkleProof\",\"type\":\"bytes\"},{\"name\":\"_originatingPlasmaTransaction\",\"type\":\"bytes\"},{\"name\":\"_originatingMerkleProof\",\"type\":\"bytes\"},{\"name\":\"_inputNumber\",\"type\":\"uint8\"}],\"name\":\"proveSpendAndWithdraw\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_plasmaBlockNumber\",\"type\":\"uint32\"},{\"name\":\"_plasmaTransaction\",\"type\":\"bytes\"},{\"name\":\"_merkleProof\",\"type\":\"bytes\"}],\"name\":\"proveInvalidDeposit\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_plasmaBlockNumber1\",\"type\":\"uint32\"},{\"name\":\"_plasmaTransaction1\",\"type\":\"bytes\"},{\"name\":\"_merkleProof1\",\"type\":\"bytes\"},{\"name\":\"_plasmaBlockNumber2\",\"type\":\"uint32\"},{\"name\":\"_plasmaTransaction2\",\"type\":\"bytes\"},{\"name\":\"_merkleProof2\",\"type\":\"bytes\"}],\"name\":\"proveDoubleFunding\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_plasmaBlockNumber\",\"type\":\"uint32\"},{\"name\":\"_plasmaInputNumberInTx\",\"type\":\"uint8\"},{\"name\":\"_plasmaTransaction\",\"type\":\"bytes\"},{\"name\":\"_merkleProof\",\"type\":\"bytes\"}],\"name\":\"proveReferencingInvalidBlock\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_plasmaBlockNumber\",\"type\":\"uint32\"},{\"name\":\"_plasmaTransaction\",\"type\":\"bytes\"},{\"name\":\"_merkleProof\",\"type\":\"bytes\"},{\"name\":\"_originatingPlasmaBlockNumber\",\"type\":\"uint32\"},{\"name\":\"_originatingPlasmaTransaction\",\"type\":\"bytes\"},{\"name\":\"_originatingMerkleProof\",\"type\":\"bytes\"},{\"name\":\"_inputOfInterest\",\"type\":\"uint256\"}],\"name\":\"proveBalanceOrOwnershipBreakingBetweenInputAndOutput\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// PlasmaParent is an auto generated Go binding around an Ethereum contract.
type PlasmaParent struct {
	PlasmaParentCaller     // Read-only binding to the contract
	PlasmaParentTransactor // Write-only binding to the contract
	PlasmaParentFilterer   // Log filterer for contract events
}

// PlasmaParentCaller is an auto generated read-only Go binding around an Ethereum contract.
type PlasmaParentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlasmaParentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PlasmaParentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlasmaParentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PlasmaParentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PlasmaParentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PlasmaParentSession struct {
	Contract     *PlasmaParent     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PlasmaParentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PlasmaParentCallerSession struct {
	Contract *PlasmaParentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// PlasmaParentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PlasmaParentTransactorSession struct {
	Contract     *PlasmaParentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// PlasmaParentRaw is an auto generated low-level Go binding around an Ethereum contract.
type PlasmaParentRaw struct {
	Contract *PlasmaParent // Generic contract binding to access the raw methods on
}

// PlasmaParentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PlasmaParentCallerRaw struct {
	Contract *PlasmaParentCaller // Generic read-only contract binding to access the raw methods on
}

// PlasmaParentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PlasmaParentTransactorRaw struct {
	Contract *PlasmaParentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPlasmaParent creates a new instance of PlasmaParent, bound to a specific deployed contract.
func NewPlasmaParent(address common.Address, backend bind.ContractBackend) (*PlasmaParent, error) {
	contract, err := bindPlasmaParent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PlasmaParent{PlasmaParentCaller: PlasmaParentCaller{contract: contract}, PlasmaParentTransactor: PlasmaParentTransactor{contract: contract}, PlasmaParentFilterer: PlasmaParentFilterer{contract: contract}}, nil
}

// NewPlasmaParentCaller creates a new read-only instance of PlasmaParent, bound to a specific deployed contract.
func NewPlasmaParentCaller(address common.Address, caller bind.ContractCaller) (*PlasmaParentCaller, error) {
	contract, err := bindPlasmaParent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PlasmaParentCaller{contract: contract}, nil
}

// NewPlasmaParentTransactor creates a new write-only instance of PlasmaParent, bound to a specific deployed contract.
func NewPlasmaParentTransactor(address common.Address, transactor bind.ContractTransactor) (*PlasmaParentTransactor, error) {
	contract, err := bindPlasmaParent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PlasmaParentTransactor{contract: contract}, nil
}

// NewPlasmaParentFilterer creates a new log filterer instance of PlasmaParent, bound to a specific deployed contract.
func NewPlasmaParentFilterer(address common.Address, filterer bind.ContractFilterer) (*PlasmaParentFilterer, error) {
	contract, err := bindPlasmaParent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PlasmaParentFilterer{contract: contract}, nil
}

// bindPlasmaParent binds a generic wrapper to an already deployed contract.
func bindPlasmaParent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PlasmaParentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlasmaParent *PlasmaParentRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PlasmaParent.Contract.PlasmaParentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlasmaParent *PlasmaParentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlasmaParent.Contract.PlasmaParentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlasmaParent *PlasmaParentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlasmaParent.Contract.PlasmaParentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PlasmaParent *PlasmaParentCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PlasmaParent.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PlasmaParent *PlasmaParentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlasmaParent.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PlasmaParent *PlasmaParentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PlasmaParent.Contract.contract.Transact(opts, method, params...)
}

// DepositWithdrawCollateral is a free data retrieval call binding the contract method 0x8286f55f.
//
// Solidity: function DepositWithdrawCollateral() constant returns(uint256)
func (_PlasmaParent *PlasmaParentCaller) DepositWithdrawCollateral(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PlasmaParent.contract.Call(opts, out, "DepositWithdrawCollateral")
	return *ret0, err
}

// DepositWithdrawCollateral is a free data retrieval call binding the contract method 0x8286f55f.
//
// Solidity: function DepositWithdrawCollateral() constant returns(uint256)
func (_PlasmaParent *PlasmaParentSession) DepositWithdrawCollateral() (*big.Int, error) {
	return _PlasmaParent.Contract.DepositWithdrawCollateral(&_PlasmaParent.CallOpts)
}

// DepositWithdrawCollateral is a free data retrieval call binding the contract method 0x8286f55f.
//
// Solidity: function DepositWithdrawCollateral() constant returns(uint256)
func (_PlasmaParent *PlasmaParentCallerSession) DepositWithdrawCollateral() (*big.Int, error) {
	return _PlasmaParent.Contract.DepositWithdrawCollateral(&_PlasmaParent.CallOpts)
}

// DepositWithdrawDelay is a free data retrieval call binding the contract method 0x3b593d6e.
//
// Solidity: function DepositWithdrawDelay() constant returns(uint256)
func (_PlasmaParent *PlasmaParentCaller) DepositWithdrawDelay(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PlasmaParent.contract.Call(opts, out, "DepositWithdrawDelay")
	return *ret0, err
}

// DepositWithdrawDelay is a free data retrieval call binding the contract method 0x3b593d6e.
//
// Solidity: function DepositWithdrawDelay() constant returns(uint256)
func (_PlasmaParent *PlasmaParentSession) DepositWithdrawDelay() (*big.Int, error) {
	return _PlasmaParent.Contract.DepositWithdrawDelay(&_PlasmaParent.CallOpts)
}

// DepositWithdrawDelay is a free data retrieval call binding the contract method 0x3b593d6e.
//
// Solidity: function DepositWithdrawDelay() constant returns(uint256)
func (_PlasmaParent *PlasmaParentCallerSession) DepositWithdrawDelay() (*big.Int, error) {
	return _PlasmaParent.Contract.DepositWithdrawDelay(&_PlasmaParent.CallOpts)
}

// ExitDelay is a free data retrieval call binding the contract method 0x48bd9892.
//
// Solidity: function ExitDelay() constant returns(uint256)
func (_PlasmaParent *PlasmaParentCaller) ExitDelay(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PlasmaParent.contract.Call(opts, out, "ExitDelay")
	return *ret0, err
}

// ExitDelay is a free data retrieval call binding the contract method 0x48bd9892.
//
// Solidity: function ExitDelay() constant returns(uint256)
func (_PlasmaParent *PlasmaParentSession) ExitDelay() (*big.Int, error) {
	return _PlasmaParent.Contract.ExitDelay(&_PlasmaParent.CallOpts)
}

// ExitDelay is a free data retrieval call binding the contract method 0x48bd9892.
//
// Solidity: function ExitDelay() constant returns(uint256)
func (_PlasmaParent *PlasmaParentCallerSession) ExitDelay() (*big.Int, error) {
	return _PlasmaParent.Contract.ExitDelay(&_PlasmaParent.CallOpts)
}

// InputChallangesDelay is a free data retrieval call binding the contract method 0xc732186e.
//
// Solidity: function InputChallangesDelay() constant returns(uint256)
func (_PlasmaParent *PlasmaParentCaller) InputChallangesDelay(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PlasmaParent.contract.Call(opts, out, "InputChallangesDelay")
	return *ret0, err
}

// InputChallangesDelay is a free data retrieval call binding the contract method 0xc732186e.
//
// Solidity: function InputChallangesDelay() constant returns(uint256)
func (_PlasmaParent *PlasmaParentSession) InputChallangesDelay() (*big.Int, error) {
	return _PlasmaParent.Contract.InputChallangesDelay(&_PlasmaParent.CallOpts)
}

// InputChallangesDelay is a free data retrieval call binding the contract method 0xc732186e.
//
// Solidity: function InputChallangesDelay() constant returns(uint256)
func (_PlasmaParent *PlasmaParentCallerSession) InputChallangesDelay() (*big.Int, error) {
	return _PlasmaParent.Contract.InputChallangesDelay(&_PlasmaParent.CallOpts)
}

// OutputChallangesDelay is a free data retrieval call binding the contract method 0x1cd2e903.
//
// Solidity: function OutputChallangesDelay() constant returns(uint256)
func (_PlasmaParent *PlasmaParentCaller) OutputChallangesDelay(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PlasmaParent.contract.Call(opts, out, "OutputChallangesDelay")
	return *ret0, err
}

// OutputChallangesDelay is a free data retrieval call binding the contract method 0x1cd2e903.
//
// Solidity: function OutputChallangesDelay() constant returns(uint256)
func (_PlasmaParent *PlasmaParentSession) OutputChallangesDelay() (*big.Int, error) {
	return _PlasmaParent.Contract.OutputChallangesDelay(&_PlasmaParent.CallOpts)
}

// OutputChallangesDelay is a free data retrieval call binding the contract method 0x1cd2e903.
//
// Solidity: function OutputChallangesDelay() constant returns(uint256)
func (_PlasmaParent *PlasmaParentCallerSession) OutputChallangesDelay() (*big.Int, error) {
	return _PlasmaParent.Contract.OutputChallangesDelay(&_PlasmaParent.CallOpts)
}

// WithdrawCollateral is a free data retrieval call binding the contract method 0x158f0706.
//
// Solidity: function WithdrawCollateral() constant returns(uint256)
func (_PlasmaParent *PlasmaParentCaller) WithdrawCollateral(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PlasmaParent.contract.Call(opts, out, "WithdrawCollateral")
	return *ret0, err
}

// WithdrawCollateral is a free data retrieval call binding the contract method 0x158f0706.
//
// Solidity: function WithdrawCollateral() constant returns(uint256)
func (_PlasmaParent *PlasmaParentSession) WithdrawCollateral() (*big.Int, error) {
	return _PlasmaParent.Contract.WithdrawCollateral(&_PlasmaParent.CallOpts)
}

// WithdrawCollateral is a free data retrieval call binding the contract method 0x158f0706.
//
// Solidity: function WithdrawCollateral() constant returns(uint256)
func (_PlasmaParent *PlasmaParentCallerSession) WithdrawCollateral() (*big.Int, error) {
	return _PlasmaParent.Contract.WithdrawCollateral(&_PlasmaParent.CallOpts)
}

// AllDepositRecordsForUser is a free data retrieval call binding the contract method 0xb20a8cc1.
//
// Solidity: function allDepositRecordsForUser( address,  uint256) constant returns(uint256)
func (_PlasmaParent *PlasmaParentCaller) AllDepositRecordsForUser(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PlasmaParent.contract.Call(opts, out, "allDepositRecordsForUser", arg0, arg1)
	return *ret0, err
}

// AllDepositRecordsForUser is a free data retrieval call binding the contract method 0xb20a8cc1.
//
// Solidity: function allDepositRecordsForUser( address,  uint256) constant returns(uint256)
func (_PlasmaParent *PlasmaParentSession) AllDepositRecordsForUser(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _PlasmaParent.Contract.AllDepositRecordsForUser(&_PlasmaParent.CallOpts, arg0, arg1)
}

// AllDepositRecordsForUser is a free data retrieval call binding the contract method 0xb20a8cc1.
//
// Solidity: function allDepositRecordsForUser( address,  uint256) constant returns(uint256)
func (_PlasmaParent *PlasmaParentCallerSession) AllDepositRecordsForUser(arg0 common.Address, arg1 *big.Int) (*big.Int, error) {
	return _PlasmaParent.Contract.AllDepositRecordsForUser(&_PlasmaParent.CallOpts, arg0, arg1)
}

// AllExitsForUser is a free data retrieval call binding the contract method 0xece34f68.
//
// Solidity: function allExitsForUser( address,  uint256) constant returns(bytes22)
func (_PlasmaParent *PlasmaParentCaller) AllExitsForUser(opts *bind.CallOpts, arg0 common.Address, arg1 *big.Int) ([22]byte, error) {
	var (
		ret0 = new([22]byte)
	)
	out := ret0
	err := _PlasmaParent.contract.Call(opts, out, "allExitsForUser", arg0, arg1)
	return *ret0, err
}

// AllExitsForUser is a free data retrieval call binding the contract method 0xece34f68.
//
// Solidity: function allExitsForUser( address,  uint256) constant returns(bytes22)
func (_PlasmaParent *PlasmaParentSession) AllExitsForUser(arg0 common.Address, arg1 *big.Int) ([22]byte, error) {
	return _PlasmaParent.Contract.AllExitsForUser(&_PlasmaParent.CallOpts, arg0, arg1)
}

// AllExitsForUser is a free data retrieval call binding the contract method 0xece34f68.
//
// Solidity: function allExitsForUser( address,  uint256) constant returns(bytes22)
func (_PlasmaParent *PlasmaParentCallerSession) AllExitsForUser(arg0 common.Address, arg1 *big.Int) ([22]byte, error) {
	return _PlasmaParent.Contract.AllExitsForUser(&_PlasmaParent.CallOpts, arg0, arg1)
}

// BlockStorage is a free data retrieval call binding the contract method 0x4a673e98.
//
// Solidity: function blockStorage() constant returns(address)
func (_PlasmaParent *PlasmaParentCaller) BlockStorage(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PlasmaParent.contract.Call(opts, out, "blockStorage")
	return *ret0, err
}

// BlockStorage is a free data retrieval call binding the contract method 0x4a673e98.
//
// Solidity: function blockStorage() constant returns(address)
func (_PlasmaParent *PlasmaParentSession) BlockStorage() (common.Address, error) {
	return _PlasmaParent.Contract.BlockStorage(&_PlasmaParent.CallOpts)
}

// BlockStorage is a free data retrieval call binding the contract method 0x4a673e98.
//
// Solidity: function blockStorage() constant returns(address)
func (_PlasmaParent *PlasmaParentCallerSession) BlockStorage() (common.Address, error) {
	return _PlasmaParent.Contract.BlockStorage(&_PlasmaParent.CallOpts)
}

// BuyoutProcessorContract is a free data retrieval call binding the contract method 0x05032538.
//
// Solidity: function buyoutProcessorContract() constant returns(address)
func (_PlasmaParent *PlasmaParentCaller) BuyoutProcessorContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PlasmaParent.contract.Call(opts, out, "buyoutProcessorContract")
	return *ret0, err
}

// BuyoutProcessorContract is a free data retrieval call binding the contract method 0x05032538.
//
// Solidity: function buyoutProcessorContract() constant returns(address)
func (_PlasmaParent *PlasmaParentSession) BuyoutProcessorContract() (common.Address, error) {
	return _PlasmaParent.Contract.BuyoutProcessorContract(&_PlasmaParent.CallOpts)
}

// BuyoutProcessorContract is a free data retrieval call binding the contract method 0x05032538.
//
// Solidity: function buyoutProcessorContract() constant returns(address)
func (_PlasmaParent *PlasmaParentCallerSession) BuyoutProcessorContract() (common.Address, error) {
	return _PlasmaParent.Contract.BuyoutProcessorContract(&_PlasmaParent.CallOpts)
}

// ChallengesContract is a free data retrieval call binding the contract method 0x934a891a.
//
// Solidity: function challengesContract() constant returns(address)
func (_PlasmaParent *PlasmaParentCaller) ChallengesContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PlasmaParent.contract.Call(opts, out, "challengesContract")
	return *ret0, err
}

// ChallengesContract is a free data retrieval call binding the contract method 0x934a891a.
//
// Solidity: function challengesContract() constant returns(address)
func (_PlasmaParent *PlasmaParentSession) ChallengesContract() (common.Address, error) {
	return _PlasmaParent.Contract.ChallengesContract(&_PlasmaParent.CallOpts)
}

// ChallengesContract is a free data retrieval call binding the contract method 0x934a891a.
//
// Solidity: function challengesContract() constant returns(address)
func (_PlasmaParent *PlasmaParentCallerSession) ChallengesContract() (common.Address, error) {
	return _PlasmaParent.Contract.ChallengesContract(&_PlasmaParent.CallOpts)
}

// DepositCounter is a free data retrieval call binding the contract method 0xecb3dc88.
//
// Solidity: function depositCounter() constant returns(uint256)
func (_PlasmaParent *PlasmaParentCaller) DepositCounter(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PlasmaParent.contract.Call(opts, out, "depositCounter")
	return *ret0, err
}

// DepositCounter is a free data retrieval call binding the contract method 0xecb3dc88.
//
// Solidity: function depositCounter() constant returns(uint256)
func (_PlasmaParent *PlasmaParentSession) DepositCounter() (*big.Int, error) {
	return _PlasmaParent.Contract.DepositCounter(&_PlasmaParent.CallOpts)
}

// DepositCounter is a free data retrieval call binding the contract method 0xecb3dc88.
//
// Solidity: function depositCounter() constant returns(uint256)
func (_PlasmaParent *PlasmaParentCallerSession) DepositCounter() (*big.Int, error) {
	return _PlasmaParent.Contract.DepositCounter(&_PlasmaParent.CallOpts)
}

// DepositRecords is a free data retrieval call binding the contract method 0x72605228.
//
// Solidity: function depositRecords( uint256) constant returns(from address, status uint8, hasCollateral bool, amount uint256, withdrawStartedAt uint256)
func (_PlasmaParent *PlasmaParentCaller) DepositRecords(opts *bind.CallOpts, arg0 *big.Int) (struct {
	From              common.Address
	Status            uint8
	HasCollateral     bool
	Amount            *big.Int
	WithdrawStartedAt *big.Int
}, error) {
	ret := new(struct {
		From              common.Address
		Status            uint8
		HasCollateral     bool
		Amount            *big.Int
		WithdrawStartedAt *big.Int
	})
	out := ret
	err := _PlasmaParent.contract.Call(opts, out, "depositRecords", arg0)
	return *ret, err
}

// DepositRecords is a free data retrieval call binding the contract method 0x72605228.
//
// Solidity: function depositRecords( uint256) constant returns(from address, status uint8, hasCollateral bool, amount uint256, withdrawStartedAt uint256)
func (_PlasmaParent *PlasmaParentSession) DepositRecords(arg0 *big.Int) (struct {
	From              common.Address
	Status            uint8
	HasCollateral     bool
	Amount            *big.Int
	WithdrawStartedAt *big.Int
}, error) {
	return _PlasmaParent.Contract.DepositRecords(&_PlasmaParent.CallOpts, arg0)
}

// DepositRecords is a free data retrieval call binding the contract method 0x72605228.
//
// Solidity: function depositRecords( uint256) constant returns(from address, status uint8, hasCollateral bool, amount uint256, withdrawStartedAt uint256)
func (_PlasmaParent *PlasmaParentCallerSession) DepositRecords(arg0 *big.Int) (struct {
	From              common.Address
	Status            uint8
	HasCollateral     bool
	Amount            *big.Int
	WithdrawStartedAt *big.Int
}, error) {
	return _PlasmaParent.Contract.DepositRecords(&_PlasmaParent.CallOpts, arg0)
}

// ExitBuyoutOffers is a free data retrieval call binding the contract method 0x86f9acbd.
//
// Solidity: function exitBuyoutOffers( bytes22) constant returns(amount uint256, from address, accepted bool)
func (_PlasmaParent *PlasmaParentCaller) ExitBuyoutOffers(opts *bind.CallOpts, arg0 [22]byte) (struct {
	Amount   *big.Int
	From     common.Address
	Accepted bool
}, error) {
	ret := new(struct {
		Amount   *big.Int
		From     common.Address
		Accepted bool
	})
	out := ret
	err := _PlasmaParent.contract.Call(opts, out, "exitBuyoutOffers", arg0)
	return *ret, err
}

// ExitBuyoutOffers is a free data retrieval call binding the contract method 0x86f9acbd.
//
// Solidity: function exitBuyoutOffers( bytes22) constant returns(amount uint256, from address, accepted bool)
func (_PlasmaParent *PlasmaParentSession) ExitBuyoutOffers(arg0 [22]byte) (struct {
	Amount   *big.Int
	From     common.Address
	Accepted bool
}, error) {
	return _PlasmaParent.Contract.ExitBuyoutOffers(&_PlasmaParent.CallOpts, arg0)
}

// ExitBuyoutOffers is a free data retrieval call binding the contract method 0x86f9acbd.
//
// Solidity: function exitBuyoutOffers( bytes22) constant returns(amount uint256, from address, accepted bool)
func (_PlasmaParent *PlasmaParentCallerSession) ExitBuyoutOffers(arg0 [22]byte) (struct {
	Amount   *big.Int
	From     common.Address
	Accepted bool
}, error) {
	return _PlasmaParent.Contract.ExitBuyoutOffers(&_PlasmaParent.CallOpts, arg0)
}

// ExitQueue is a free data retrieval call binding the contract method 0xffed4bf5.
//
// Solidity: function exitQueue() constant returns(address)
func (_PlasmaParent *PlasmaParentCaller) ExitQueue(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PlasmaParent.contract.Call(opts, out, "exitQueue")
	return *ret0, err
}

// ExitQueue is a free data retrieval call binding the contract method 0xffed4bf5.
//
// Solidity: function exitQueue() constant returns(address)
func (_PlasmaParent *PlasmaParentSession) ExitQueue() (common.Address, error) {
	return _PlasmaParent.Contract.ExitQueue(&_PlasmaParent.CallOpts)
}

// ExitQueue is a free data retrieval call binding the contract method 0xffed4bf5.
//
// Solidity: function exitQueue() constant returns(address)
func (_PlasmaParent *PlasmaParentCallerSession) ExitQueue() (common.Address, error) {
	return _PlasmaParent.Contract.ExitQueue(&_PlasmaParent.CallOpts)
}

// ExitRecords is a free data retrieval call binding the contract method 0x9ef062ee.
//
// Solidity: function exitRecords( bytes22) constant returns(transactionRef bytes32, amount uint256, owner address, timePublished uint64, blockNumber uint32, transactionNumber uint32, outputNumber uint8, isValid bool, isLimbo bool)
func (_PlasmaParent *PlasmaParentCaller) ExitRecords(opts *bind.CallOpts, arg0 [22]byte) (struct {
	TransactionRef    [32]byte
	Amount            *big.Int
	Owner             common.Address
	TimePublished     uint64
	BlockNumber       uint32
	TransactionNumber uint32
	OutputNumber      uint8
	IsValid           bool
	IsLimbo           bool
}, error) {
	ret := new(struct {
		TransactionRef    [32]byte
		Amount            *big.Int
		Owner             common.Address
		TimePublished     uint64
		BlockNumber       uint32
		TransactionNumber uint32
		OutputNumber      uint8
		IsValid           bool
		IsLimbo           bool
	})
	out := ret
	err := _PlasmaParent.contract.Call(opts, out, "exitRecords", arg0)
	return *ret, err
}

// ExitRecords is a free data retrieval call binding the contract method 0x9ef062ee.
//
// Solidity: function exitRecords( bytes22) constant returns(transactionRef bytes32, amount uint256, owner address, timePublished uint64, blockNumber uint32, transactionNumber uint32, outputNumber uint8, isValid bool, isLimbo bool)
func (_PlasmaParent *PlasmaParentSession) ExitRecords(arg0 [22]byte) (struct {
	TransactionRef    [32]byte
	Amount            *big.Int
	Owner             common.Address
	TimePublished     uint64
	BlockNumber       uint32
	TransactionNumber uint32
	OutputNumber      uint8
	IsValid           bool
	IsLimbo           bool
}, error) {
	return _PlasmaParent.Contract.ExitRecords(&_PlasmaParent.CallOpts, arg0)
}

// ExitRecords is a free data retrieval call binding the contract method 0x9ef062ee.
//
// Solidity: function exitRecords( bytes22) constant returns(transactionRef bytes32, amount uint256, owner address, timePublished uint64, blockNumber uint32, transactionNumber uint32, outputNumber uint8, isValid bool, isLimbo bool)
func (_PlasmaParent *PlasmaParentCallerSession) ExitRecords(arg0 [22]byte) (struct {
	TransactionRef    [32]byte
	Amount            *big.Int
	Owner             common.Address
	TimePublished     uint64
	BlockNumber       uint32
	TransactionNumber uint32
	OutputNumber      uint8
	IsValid           bool
	IsLimbo           bool
}, error) {
	return _PlasmaParent.Contract.ExitRecords(&_PlasmaParent.CallOpts, arg0)
}

// HashOfLastSubmittedBlock is a free data retrieval call binding the contract method 0x8649d43b.
//
// Solidity: function hashOfLastSubmittedBlock() constant returns(bytes32)
func (_PlasmaParent *PlasmaParentCaller) HashOfLastSubmittedBlock(opts *bind.CallOpts) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _PlasmaParent.contract.Call(opts, out, "hashOfLastSubmittedBlock")
	return *ret0, err
}

// HashOfLastSubmittedBlock is a free data retrieval call binding the contract method 0x8649d43b.
//
// Solidity: function hashOfLastSubmittedBlock() constant returns(bytes32)
func (_PlasmaParent *PlasmaParentSession) HashOfLastSubmittedBlock() ([32]byte, error) {
	return _PlasmaParent.Contract.HashOfLastSubmittedBlock(&_PlasmaParent.CallOpts)
}

// HashOfLastSubmittedBlock is a free data retrieval call binding the contract method 0x8649d43b.
//
// Solidity: function hashOfLastSubmittedBlock() constant returns(bytes32)
func (_PlasmaParent *PlasmaParentCallerSession) HashOfLastSubmittedBlock() ([32]byte, error) {
	return _PlasmaParent.Contract.HashOfLastSubmittedBlock(&_PlasmaParent.CallOpts)
}

// IsWellFormedTransaction is a free data retrieval call binding the contract method 0x9b8c1f27.
//
// Solidity: function isWellFormedTransaction(_plasmaTransaction bytes) constant returns(isWellFormed bool)
func (_PlasmaParent *PlasmaParentCaller) IsWellFormedTransaction(opts *bind.CallOpts, _plasmaTransaction []byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PlasmaParent.contract.Call(opts, out, "isWellFormedTransaction", _plasmaTransaction)
	return *ret0, err
}

// IsWellFormedTransaction is a free data retrieval call binding the contract method 0x9b8c1f27.
//
// Solidity: function isWellFormedTransaction(_plasmaTransaction bytes) constant returns(isWellFormed bool)
func (_PlasmaParent *PlasmaParentSession) IsWellFormedTransaction(_plasmaTransaction []byte) (bool, error) {
	return _PlasmaParent.Contract.IsWellFormedTransaction(&_PlasmaParent.CallOpts, _plasmaTransaction)
}

// IsWellFormedTransaction is a free data retrieval call binding the contract method 0x9b8c1f27.
//
// Solidity: function isWellFormedTransaction(_plasmaTransaction bytes) constant returns(isWellFormed bool)
func (_PlasmaParent *PlasmaParentCallerSession) IsWellFormedTransaction(_plasmaTransaction []byte) (bool, error) {
	return _PlasmaParent.Contract.IsWellFormedTransaction(&_PlasmaParent.CallOpts, _plasmaTransaction)
}

// LastBlockNumber is a free data retrieval call binding the contract method 0x2552317c.
//
// Solidity: function lastBlockNumber() constant returns(blockNumber uint256)
func (_PlasmaParent *PlasmaParentCaller) LastBlockNumber(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PlasmaParent.contract.Call(opts, out, "lastBlockNumber")
	return *ret0, err
}

// LastBlockNumber is a free data retrieval call binding the contract method 0x2552317c.
//
// Solidity: function lastBlockNumber() constant returns(blockNumber uint256)
func (_PlasmaParent *PlasmaParentSession) LastBlockNumber() (*big.Int, error) {
	return _PlasmaParent.Contract.LastBlockNumber(&_PlasmaParent.CallOpts)
}

// LastBlockNumber is a free data retrieval call binding the contract method 0x2552317c.
//
// Solidity: function lastBlockNumber() constant returns(blockNumber uint256)
func (_PlasmaParent *PlasmaParentCallerSession) LastBlockNumber() (*big.Int, error) {
	return _PlasmaParent.Contract.LastBlockNumber(&_PlasmaParent.CallOpts)
}

// LastValidBlock is a free data retrieval call binding the contract method 0x38f263f6.
//
// Solidity: function lastValidBlock() constant returns(uint32)
func (_PlasmaParent *PlasmaParentCaller) LastValidBlock(opts *bind.CallOpts) (uint32, error) {
	var (
		ret0 = new(uint32)
	)
	out := ret0
	err := _PlasmaParent.contract.Call(opts, out, "lastValidBlock")
	return *ret0, err
}

// LastValidBlock is a free data retrieval call binding the contract method 0x38f263f6.
//
// Solidity: function lastValidBlock() constant returns(uint32)
func (_PlasmaParent *PlasmaParentSession) LastValidBlock() (uint32, error) {
	return _PlasmaParent.Contract.LastValidBlock(&_PlasmaParent.CallOpts)
}

// LastValidBlock is a free data retrieval call binding the contract method 0x38f263f6.
//
// Solidity: function lastValidBlock() constant returns(uint32)
func (_PlasmaParent *PlasmaParentCallerSession) LastValidBlock() (uint32, error) {
	return _PlasmaParent.Contract.LastValidBlock(&_PlasmaParent.CallOpts)
}

// LimboExitContract is a free data retrieval call binding the contract method 0x70f4e877.
//
// Solidity: function limboExitContract() constant returns(address)
func (_PlasmaParent *PlasmaParentCaller) LimboExitContract(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PlasmaParent.contract.Call(opts, out, "limboExitContract")
	return *ret0, err
}

// LimboExitContract is a free data retrieval call binding the contract method 0x70f4e877.
//
// Solidity: function limboExitContract() constant returns(address)
func (_PlasmaParent *PlasmaParentSession) LimboExitContract() (common.Address, error) {
	return _PlasmaParent.Contract.LimboExitContract(&_PlasmaParent.CallOpts)
}

// LimboExitContract is a free data retrieval call binding the contract method 0x70f4e877.
//
// Solidity: function limboExitContract() constant returns(address)
func (_PlasmaParent *PlasmaParentCallerSession) LimboExitContract() (common.Address, error) {
	return _PlasmaParent.Contract.LimboExitContract(&_PlasmaParent.CallOpts)
}

// OperatorsBond is a free data retrieval call binding the contract method 0x5a20954e.
//
// Solidity: function operatorsBond() constant returns(uint256)
func (_PlasmaParent *PlasmaParentCaller) OperatorsBond(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PlasmaParent.contract.Call(opts, out, "operatorsBond")
	return *ret0, err
}

// OperatorsBond is a free data retrieval call binding the contract method 0x5a20954e.
//
// Solidity: function operatorsBond() constant returns(uint256)
func (_PlasmaParent *PlasmaParentSession) OperatorsBond() (*big.Int, error) {
	return _PlasmaParent.Contract.OperatorsBond(&_PlasmaParent.CallOpts)
}

// OperatorsBond is a free data retrieval call binding the contract method 0x5a20954e.
//
// Solidity: function operatorsBond() constant returns(uint256)
func (_PlasmaParent *PlasmaParentCallerSession) OperatorsBond() (*big.Int, error) {
	return _PlasmaParent.Contract.OperatorsBond(&_PlasmaParent.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PlasmaParent *PlasmaParentCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PlasmaParent.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PlasmaParent *PlasmaParentSession) Owner() (common.Address, error) {
	return _PlasmaParent.Contract.Owner(&_PlasmaParent.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PlasmaParent *PlasmaParentCallerSession) Owner() (common.Address, error) {
	return _PlasmaParent.Contract.Owner(&_PlasmaParent.CallOpts)
}

// PlasmaErrorFound is a free data retrieval call binding the contract method 0x2c4683c4.
//
// Solidity: function plasmaErrorFound() constant returns(bool)
func (_PlasmaParent *PlasmaParentCaller) PlasmaErrorFound(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PlasmaParent.contract.Call(opts, out, "plasmaErrorFound")
	return *ret0, err
}

// PlasmaErrorFound is a free data retrieval call binding the contract method 0x2c4683c4.
//
// Solidity: function plasmaErrorFound() constant returns(bool)
func (_PlasmaParent *PlasmaParentSession) PlasmaErrorFound() (bool, error) {
	return _PlasmaParent.Contract.PlasmaErrorFound(&_PlasmaParent.CallOpts)
}

// PlasmaErrorFound is a free data retrieval call binding the contract method 0x2c4683c4.
//
// Solidity: function plasmaErrorFound() constant returns(bool)
func (_PlasmaParent *PlasmaParentCallerSession) PlasmaErrorFound() (bool, error) {
	return _PlasmaParent.Contract.PlasmaErrorFound(&_PlasmaParent.CallOpts)
}

// SuccesfulExits is a free data retrieval call binding the contract method 0x6e637cd6.
//
// Solidity: function succesfulExits( bytes22) constant returns(bool)
func (_PlasmaParent *PlasmaParentCaller) SuccesfulExits(opts *bind.CallOpts, arg0 [22]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PlasmaParent.contract.Call(opts, out, "succesfulExits", arg0)
	return *ret0, err
}

// SuccesfulExits is a free data retrieval call binding the contract method 0x6e637cd6.
//
// Solidity: function succesfulExits( bytes22) constant returns(bool)
func (_PlasmaParent *PlasmaParentSession) SuccesfulExits(arg0 [22]byte) (bool, error) {
	return _PlasmaParent.Contract.SuccesfulExits(&_PlasmaParent.CallOpts, arg0)
}

// SuccesfulExits is a free data retrieval call binding the contract method 0x6e637cd6.
//
// Solidity: function succesfulExits( bytes22) constant returns(bool)
func (_PlasmaParent *PlasmaParentCallerSession) SuccesfulExits(arg0 [22]byte) (bool, error) {
	return _PlasmaParent.Contract.SuccesfulExits(&_PlasmaParent.CallOpts, arg0)
}

// AcceptBuyoutOffer is a paid mutator transaction binding the contract method 0xeed8e608.
//
// Solidity: function acceptBuyoutOffer(_index bytes22) returns(success bool)
func (_PlasmaParent *PlasmaParentTransactor) AcceptBuyoutOffer(opts *bind.TransactOpts, _index [22]byte) (*types.Transaction, error) {
	return _PlasmaParent.contract.Transact(opts, "acceptBuyoutOffer", _index)
}

// AcceptBuyoutOffer is a paid mutator transaction binding the contract method 0xeed8e608.
//
// Solidity: function acceptBuyoutOffer(_index bytes22) returns(success bool)
func (_PlasmaParent *PlasmaParentSession) AcceptBuyoutOffer(_index [22]byte) (*types.Transaction, error) {
	return _PlasmaParent.Contract.AcceptBuyoutOffer(&_PlasmaParent.TransactOpts, _index)
}

// AcceptBuyoutOffer is a paid mutator transaction binding the contract method 0xeed8e608.
//
// Solidity: function acceptBuyoutOffer(_index bytes22) returns(success bool)
func (_PlasmaParent *PlasmaParentTransactorSession) AcceptBuyoutOffer(_index [22]byte) (*types.Transaction, error) {
	return _PlasmaParent.Contract.AcceptBuyoutOffer(&_PlasmaParent.TransactOpts, _index)
}

// AllowChallenges is a paid mutator transaction binding the contract method 0x5868e153.
//
// Solidity: function allowChallenges(_challenger address) returns(success bool)
func (_PlasmaParent *PlasmaParentTransactor) AllowChallenges(opts *bind.TransactOpts, _challenger common.Address) (*types.Transaction, error) {
	return _PlasmaParent.contract.Transact(opts, "allowChallenges", _challenger)
}

// AllowChallenges is a paid mutator transaction binding the contract method 0x5868e153.
//
// Solidity: function allowChallenges(_challenger address) returns(success bool)
func (_PlasmaParent *PlasmaParentSession) AllowChallenges(_challenger common.Address) (*types.Transaction, error) {
	return _PlasmaParent.Contract.AllowChallenges(&_PlasmaParent.TransactOpts, _challenger)
}

// AllowChallenges is a paid mutator transaction binding the contract method 0x5868e153.
//
// Solidity: function allowChallenges(_challenger address) returns(success bool)
func (_PlasmaParent *PlasmaParentTransactorSession) AllowChallenges(_challenger common.Address) (*types.Transaction, error) {
	return _PlasmaParent.Contract.AllowChallenges(&_PlasmaParent.TransactOpts, _challenger)
}

// AllowDeposits is a paid mutator transaction binding the contract method 0x17600fdf.
//
// Solidity: function allowDeposits(_buyouts address) returns(success bool)
func (_PlasmaParent *PlasmaParentTransactor) AllowDeposits(opts *bind.TransactOpts, _buyouts common.Address) (*types.Transaction, error) {
	return _PlasmaParent.contract.Transact(opts, "allowDeposits", _buyouts)
}

// AllowDeposits is a paid mutator transaction binding the contract method 0x17600fdf.
//
// Solidity: function allowDeposits(_buyouts address) returns(success bool)
func (_PlasmaParent *PlasmaParentSession) AllowDeposits(_buyouts common.Address) (*types.Transaction, error) {
	return _PlasmaParent.Contract.AllowDeposits(&_PlasmaParent.TransactOpts, _buyouts)
}

// AllowDeposits is a paid mutator transaction binding the contract method 0x17600fdf.
//
// Solidity: function allowDeposits(_buyouts address) returns(success bool)
func (_PlasmaParent *PlasmaParentTransactorSession) AllowDeposits(_buyouts common.Address) (*types.Transaction, error) {
	return _PlasmaParent.Contract.AllowDeposits(&_PlasmaParent.TransactOpts, _buyouts)
}

// AllowLimboExits is a paid mutator transaction binding the contract method 0xcb78f233.
//
// Solidity: function allowLimboExits(_limboExiter address) returns(success bool)
func (_PlasmaParent *PlasmaParentTransactor) AllowLimboExits(opts *bind.TransactOpts, _limboExiter common.Address) (*types.Transaction, error) {
	return _PlasmaParent.contract.Transact(opts, "allowLimboExits", _limboExiter)
}

// AllowLimboExits is a paid mutator transaction binding the contract method 0xcb78f233.
//
// Solidity: function allowLimboExits(_limboExiter address) returns(success bool)
func (_PlasmaParent *PlasmaParentSession) AllowLimboExits(_limboExiter common.Address) (*types.Transaction, error) {
	return _PlasmaParent.Contract.AllowLimboExits(&_PlasmaParent.TransactOpts, _limboExiter)
}

// AllowLimboExits is a paid mutator transaction binding the contract method 0xcb78f233.
//
// Solidity: function allowLimboExits(_limboExiter address) returns(success bool)
func (_PlasmaParent *PlasmaParentTransactorSession) AllowLimboExits(_limboExiter common.Address) (*types.Transaction, error) {
	return _PlasmaParent.Contract.AllowLimboExits(&_PlasmaParent.TransactOpts, _limboExiter)
}

// ChallengeDepositWithdraw is a paid mutator transaction binding the contract method 0x6a7596f2.
//
// Solidity: function challengeDepositWithdraw(depositIndex uint256, _plasmaBlockNumber uint32, _plasmaTransaction bytes, _merkleProof bytes) returns(success bool)
func (_PlasmaParent *PlasmaParentTransactor) ChallengeDepositWithdraw(opts *bind.TransactOpts, depositIndex *big.Int, _plasmaBlockNumber uint32, _plasmaTransaction []byte, _merkleProof []byte) (*types.Transaction, error) {
	return _PlasmaParent.contract.Transact(opts, "challengeDepositWithdraw", depositIndex, _plasmaBlockNumber, _plasmaTransaction, _merkleProof)
}

// ChallengeDepositWithdraw is a paid mutator transaction binding the contract method 0x6a7596f2.
//
// Solidity: function challengeDepositWithdraw(depositIndex uint256, _plasmaBlockNumber uint32, _plasmaTransaction bytes, _merkleProof bytes) returns(success bool)
func (_PlasmaParent *PlasmaParentSession) ChallengeDepositWithdraw(depositIndex *big.Int, _plasmaBlockNumber uint32, _plasmaTransaction []byte, _merkleProof []byte) (*types.Transaction, error) {
	return _PlasmaParent.Contract.ChallengeDepositWithdraw(&_PlasmaParent.TransactOpts, depositIndex, _plasmaBlockNumber, _plasmaTransaction, _merkleProof)
}

// ChallengeDepositWithdraw is a paid mutator transaction binding the contract method 0x6a7596f2.
//
// Solidity: function challengeDepositWithdraw(depositIndex uint256, _plasmaBlockNumber uint32, _plasmaTransaction bytes, _merkleProof bytes) returns(success bool)
func (_PlasmaParent *PlasmaParentTransactorSession) ChallengeDepositWithdraw(depositIndex *big.Int, _plasmaBlockNumber uint32, _plasmaTransaction []byte, _merkleProof []byte) (*types.Transaction, error) {
	return _PlasmaParent.Contract.ChallengeDepositWithdraw(&_PlasmaParent.TransactOpts, depositIndex, _plasmaBlockNumber, _plasmaTransaction, _merkleProof)
}

// ChallengeNormalExitByShowingAnInputDoubleSpend is a paid mutator transaction binding the contract method 0x0e50aaa8.
//
// Solidity: function challengeNormalExitByShowingAnInputDoubleSpend(_originalTransaction bytes, _originalInputNumber uint8, _exitRecordHash bytes22, _plasmaBlockNumber uint32, _plasmaTransaction bytes, _merkleProof bytes, _inputNumber uint8) returns(success bool)
func (_PlasmaParent *PlasmaParentTransactor) ChallengeNormalExitByShowingAnInputDoubleSpend(opts *bind.TransactOpts, _originalTransaction []byte, _originalInputNumber uint8, _exitRecordHash [22]byte, _plasmaBlockNumber uint32, _plasmaTransaction []byte, _merkleProof []byte, _inputNumber uint8) (*types.Transaction, error) {
	return _PlasmaParent.contract.Transact(opts, "challengeNormalExitByShowingAnInputDoubleSpend", _originalTransaction, _originalInputNumber, _exitRecordHash, _plasmaBlockNumber, _plasmaTransaction, _merkleProof, _inputNumber)
}

// ChallengeNormalExitByShowingAnInputDoubleSpend is a paid mutator transaction binding the contract method 0x0e50aaa8.
//
// Solidity: function challengeNormalExitByShowingAnInputDoubleSpend(_originalTransaction bytes, _originalInputNumber uint8, _exitRecordHash bytes22, _plasmaBlockNumber uint32, _plasmaTransaction bytes, _merkleProof bytes, _inputNumber uint8) returns(success bool)
func (_PlasmaParent *PlasmaParentSession) ChallengeNormalExitByShowingAnInputDoubleSpend(_originalTransaction []byte, _originalInputNumber uint8, _exitRecordHash [22]byte, _plasmaBlockNumber uint32, _plasmaTransaction []byte, _merkleProof []byte, _inputNumber uint8) (*types.Transaction, error) {
	return _PlasmaParent.Contract.ChallengeNormalExitByShowingAnInputDoubleSpend(&_PlasmaParent.TransactOpts, _originalTransaction, _originalInputNumber, _exitRecordHash, _plasmaBlockNumber, _plasmaTransaction, _merkleProof, _inputNumber)
}

// ChallengeNormalExitByShowingAnInputDoubleSpend is a paid mutator transaction binding the contract method 0x0e50aaa8.
//
// Solidity: function challengeNormalExitByShowingAnInputDoubleSpend(_originalTransaction bytes, _originalInputNumber uint8, _exitRecordHash bytes22, _plasmaBlockNumber uint32, _plasmaTransaction bytes, _merkleProof bytes, _inputNumber uint8) returns(success bool)
func (_PlasmaParent *PlasmaParentTransactorSession) ChallengeNormalExitByShowingAnInputDoubleSpend(_originalTransaction []byte, _originalInputNumber uint8, _exitRecordHash [22]byte, _plasmaBlockNumber uint32, _plasmaTransaction []byte, _merkleProof []byte, _inputNumber uint8) (*types.Transaction, error) {
	return _PlasmaParent.Contract.ChallengeNormalExitByShowingAnInputDoubleSpend(&_PlasmaParent.TransactOpts, _originalTransaction, _originalInputNumber, _exitRecordHash, _plasmaBlockNumber, _plasmaTransaction, _merkleProof, _inputNumber)
}

// ChallengeNormalExitByShowingExitBeingSpent is a paid mutator transaction binding the contract method 0xa8c818ba.
//
// Solidity: function challengeNormalExitByShowingExitBeingSpent(_exitRecordHash bytes22, _plasmaBlockNumber uint32, _plasmaTransaction bytes, _merkleProof bytes, _inputNumber uint8) returns(success bool)
func (_PlasmaParent *PlasmaParentTransactor) ChallengeNormalExitByShowingExitBeingSpent(opts *bind.TransactOpts, _exitRecordHash [22]byte, _plasmaBlockNumber uint32, _plasmaTransaction []byte, _merkleProof []byte, _inputNumber uint8) (*types.Transaction, error) {
	return _PlasmaParent.contract.Transact(opts, "challengeNormalExitByShowingExitBeingSpent", _exitRecordHash, _plasmaBlockNumber, _plasmaTransaction, _merkleProof, _inputNumber)
}

// ChallengeNormalExitByShowingExitBeingSpent is a paid mutator transaction binding the contract method 0xa8c818ba.
//
// Solidity: function challengeNormalExitByShowingExitBeingSpent(_exitRecordHash bytes22, _plasmaBlockNumber uint32, _plasmaTransaction bytes, _merkleProof bytes, _inputNumber uint8) returns(success bool)
func (_PlasmaParent *PlasmaParentSession) ChallengeNormalExitByShowingExitBeingSpent(_exitRecordHash [22]byte, _plasmaBlockNumber uint32, _plasmaTransaction []byte, _merkleProof []byte, _inputNumber uint8) (*types.Transaction, error) {
	return _PlasmaParent.Contract.ChallengeNormalExitByShowingExitBeingSpent(&_PlasmaParent.TransactOpts, _exitRecordHash, _plasmaBlockNumber, _plasmaTransaction, _merkleProof, _inputNumber)
}

// ChallengeNormalExitByShowingExitBeingSpent is a paid mutator transaction binding the contract method 0xa8c818ba.
//
// Solidity: function challengeNormalExitByShowingExitBeingSpent(_exitRecordHash bytes22, _plasmaBlockNumber uint32, _plasmaTransaction bytes, _merkleProof bytes, _inputNumber uint8) returns(success bool)
func (_PlasmaParent *PlasmaParentTransactorSession) ChallengeNormalExitByShowingExitBeingSpent(_exitRecordHash [22]byte, _plasmaBlockNumber uint32, _plasmaTransaction []byte, _merkleProof []byte, _inputNumber uint8) (*types.Transaction, error) {
	return _PlasmaParent.Contract.ChallengeNormalExitByShowingExitBeingSpent(&_PlasmaParent.TransactOpts, _exitRecordHash, _plasmaBlockNumber, _plasmaTransaction, _merkleProof, _inputNumber)
}

// ChallengeNormalExitByShowingMismatchedInput is a paid mutator transaction binding the contract method 0xd0bbfd3a.
//
// Solidity: function challengeNormalExitByShowingMismatchedInput(_originalTransaction bytes, _originalInputNumber uint8, _exitRecordHash bytes22, _plasmaBlockNumber uint32, _plasmaTransaction bytes, _merkleProof bytes, _outputNumber uint8) returns(success bool)
func (_PlasmaParent *PlasmaParentTransactor) ChallengeNormalExitByShowingMismatchedInput(opts *bind.TransactOpts, _originalTransaction []byte, _originalInputNumber uint8, _exitRecordHash [22]byte, _plasmaBlockNumber uint32, _plasmaTransaction []byte, _merkleProof []byte, _outputNumber uint8) (*types.Transaction, error) {
	return _PlasmaParent.contract.Transact(opts, "challengeNormalExitByShowingMismatchedInput", _originalTransaction, _originalInputNumber, _exitRecordHash, _plasmaBlockNumber, _plasmaTransaction, _merkleProof, _outputNumber)
}

// ChallengeNormalExitByShowingMismatchedInput is a paid mutator transaction binding the contract method 0xd0bbfd3a.
//
// Solidity: function challengeNormalExitByShowingMismatchedInput(_originalTransaction bytes, _originalInputNumber uint8, _exitRecordHash bytes22, _plasmaBlockNumber uint32, _plasmaTransaction bytes, _merkleProof bytes, _outputNumber uint8) returns(success bool)
func (_PlasmaParent *PlasmaParentSession) ChallengeNormalExitByShowingMismatchedInput(_originalTransaction []byte, _originalInputNumber uint8, _exitRecordHash [22]byte, _plasmaBlockNumber uint32, _plasmaTransaction []byte, _merkleProof []byte, _outputNumber uint8) (*types.Transaction, error) {
	return _PlasmaParent.Contract.ChallengeNormalExitByShowingMismatchedInput(&_PlasmaParent.TransactOpts, _originalTransaction, _originalInputNumber, _exitRecordHash, _plasmaBlockNumber, _plasmaTransaction, _merkleProof, _outputNumber)
}

// ChallengeNormalExitByShowingMismatchedInput is a paid mutator transaction binding the contract method 0xd0bbfd3a.
//
// Solidity: function challengeNormalExitByShowingMismatchedInput(_originalTransaction bytes, _originalInputNumber uint8, _exitRecordHash bytes22, _plasmaBlockNumber uint32, _plasmaTransaction bytes, _merkleProof bytes, _outputNumber uint8) returns(success bool)
func (_PlasmaParent *PlasmaParentTransactorSession) ChallengeNormalExitByShowingMismatchedInput(_originalTransaction []byte, _originalInputNumber uint8, _exitRecordHash [22]byte, _plasmaBlockNumber uint32, _plasmaTransaction []byte, _merkleProof []byte, _outputNumber uint8) (*types.Transaction, error) {
	return _PlasmaParent.Contract.ChallengeNormalExitByShowingMismatchedInput(&_PlasmaParent.TransactOpts, _originalTransaction, _originalInputNumber, _exitRecordHash, _plasmaBlockNumber, _plasmaTransaction, _merkleProof, _outputNumber)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(success bool)
func (_PlasmaParent *PlasmaParentTransactor) Deposit(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlasmaParent.contract.Transact(opts, "deposit")
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(success bool)
func (_PlasmaParent *PlasmaParentSession) Deposit() (*types.Transaction, error) {
	return _PlasmaParent.Contract.Deposit(&_PlasmaParent.TransactOpts)
}

// Deposit is a paid mutator transaction binding the contract method 0xd0e30db0.
//
// Solidity: function deposit() returns(success bool)
func (_PlasmaParent *PlasmaParentTransactorSession) Deposit() (*types.Transaction, error) {
	return _PlasmaParent.Contract.Deposit(&_PlasmaParent.TransactOpts)
}

// DepositFor is a paid mutator transaction binding the contract method 0xaa67c919.
//
// Solidity: function depositFor(_for address) returns(success bool)
func (_PlasmaParent *PlasmaParentTransactor) DepositFor(opts *bind.TransactOpts, _for common.Address) (*types.Transaction, error) {
	return _PlasmaParent.contract.Transact(opts, "depositFor", _for)
}

// DepositFor is a paid mutator transaction binding the contract method 0xaa67c919.
//
// Solidity: function depositFor(_for address) returns(success bool)
func (_PlasmaParent *PlasmaParentSession) DepositFor(_for common.Address) (*types.Transaction, error) {
	return _PlasmaParent.Contract.DepositFor(&_PlasmaParent.TransactOpts, _for)
}

// DepositFor is a paid mutator transaction binding the contract method 0xaa67c919.
//
// Solidity: function depositFor(_for address) returns(success bool)
func (_PlasmaParent *PlasmaParentTransactorSession) DepositFor(_for common.Address) (*types.Transaction, error) {
	return _PlasmaParent.Contract.DepositFor(&_PlasmaParent.TransactOpts, _for)
}

// FinalizeDepositWithdraw is a paid mutator transaction binding the contract method 0x24a593bd.
//
// Solidity: function finalizeDepositWithdraw(depositIndex uint256) returns(success bool)
func (_PlasmaParent *PlasmaParentTransactor) FinalizeDepositWithdraw(opts *bind.TransactOpts, depositIndex *big.Int) (*types.Transaction, error) {
	return _PlasmaParent.contract.Transact(opts, "finalizeDepositWithdraw", depositIndex)
}

// FinalizeDepositWithdraw is a paid mutator transaction binding the contract method 0x24a593bd.
//
// Solidity: function finalizeDepositWithdraw(depositIndex uint256) returns(success bool)
func (_PlasmaParent *PlasmaParentSession) FinalizeDepositWithdraw(depositIndex *big.Int) (*types.Transaction, error) {
	return _PlasmaParent.Contract.FinalizeDepositWithdraw(&_PlasmaParent.TransactOpts, depositIndex)
}

// FinalizeDepositWithdraw is a paid mutator transaction binding the contract method 0x24a593bd.
//
// Solidity: function finalizeDepositWithdraw(depositIndex uint256) returns(success bool)
func (_PlasmaParent *PlasmaParentTransactorSession) FinalizeDepositWithdraw(depositIndex *big.Int) (*types.Transaction, error) {
	return _PlasmaParent.Contract.FinalizeDepositWithdraw(&_PlasmaParent.TransactOpts, depositIndex)
}

// FinalizeExits is a paid mutator transaction binding the contract method 0x7bcce117.
//
// Solidity: function finalizeExits(_numOfExits uint256) returns(success bool)
func (_PlasmaParent *PlasmaParentTransactor) FinalizeExits(opts *bind.TransactOpts, _numOfExits *big.Int) (*types.Transaction, error) {
	return _PlasmaParent.contract.Transact(opts, "finalizeExits", _numOfExits)
}

// FinalizeExits is a paid mutator transaction binding the contract method 0x7bcce117.
//
// Solidity: function finalizeExits(_numOfExits uint256) returns(success bool)
func (_PlasmaParent *PlasmaParentSession) FinalizeExits(_numOfExits *big.Int) (*types.Transaction, error) {
	return _PlasmaParent.Contract.FinalizeExits(&_PlasmaParent.TransactOpts, _numOfExits)
}

// FinalizeExits is a paid mutator transaction binding the contract method 0x7bcce117.
//
// Solidity: function finalizeExits(_numOfExits uint256) returns(success bool)
func (_PlasmaParent *PlasmaParentTransactorSession) FinalizeExits(_numOfExits *big.Int) (*types.Transaction, error) {
	return _PlasmaParent.Contract.FinalizeExits(&_PlasmaParent.TransactOpts, _numOfExits)
}

// IncrementWeekOldCounter is a paid mutator transaction binding the contract method 0x32bc8990.
//
// Solidity: function incrementWeekOldCounter() returns()
func (_PlasmaParent *PlasmaParentTransactor) IncrementWeekOldCounter(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PlasmaParent.contract.Transact(opts, "incrementWeekOldCounter")
}

// IncrementWeekOldCounter is a paid mutator transaction binding the contract method 0x32bc8990.
//
// Solidity: function incrementWeekOldCounter() returns()
func (_PlasmaParent *PlasmaParentSession) IncrementWeekOldCounter() (*types.Transaction, error) {
	return _PlasmaParent.Contract.IncrementWeekOldCounter(&_PlasmaParent.TransactOpts)
}

// IncrementWeekOldCounter is a paid mutator transaction binding the contract method 0x32bc8990.
//
// Solidity: function incrementWeekOldCounter() returns()
func (_PlasmaParent *PlasmaParentTransactorSession) IncrementWeekOldCounter() (*types.Transaction, error) {
	return _PlasmaParent.Contract.IncrementWeekOldCounter(&_PlasmaParent.TransactOpts)
}

// OfferOutputBuyout is a paid mutator transaction binding the contract method 0x4dee4a31.
//
// Solidity: function offerOutputBuyout(_index bytes22, _beneficiary address) returns(success bool)
func (_PlasmaParent *PlasmaParentTransactor) OfferOutputBuyout(opts *bind.TransactOpts, _index [22]byte, _beneficiary common.Address) (*types.Transaction, error) {
	return _PlasmaParent.contract.Transact(opts, "offerOutputBuyout", _index, _beneficiary)
}

// OfferOutputBuyout is a paid mutator transaction binding the contract method 0x4dee4a31.
//
// Solidity: function offerOutputBuyout(_index bytes22, _beneficiary address) returns(success bool)
func (_PlasmaParent *PlasmaParentSession) OfferOutputBuyout(_index [22]byte, _beneficiary common.Address) (*types.Transaction, error) {
	return _PlasmaParent.Contract.OfferOutputBuyout(&_PlasmaParent.TransactOpts, _index, _beneficiary)
}

// OfferOutputBuyout is a paid mutator transaction binding the contract method 0x4dee4a31.
//
// Solidity: function offerOutputBuyout(_index bytes22, _beneficiary address) returns(success bool)
func (_PlasmaParent *PlasmaParentTransactorSession) OfferOutputBuyout(_index [22]byte, _beneficiary common.Address) (*types.Transaction, error) {
	return _PlasmaParent.Contract.OfferOutputBuyout(&_PlasmaParent.TransactOpts, _index, _beneficiary)
}

// ProveBalanceOrOwnershipBreakingBetweenInputAndOutput is a paid mutator transaction binding the contract method 0xaee5b5de.
//
// Solidity: function proveBalanceOrOwnershipBreakingBetweenInputAndOutput(_plasmaBlockNumber uint32, _plasmaTransaction bytes, _merkleProof bytes, _originatingPlasmaBlockNumber uint32, _originatingPlasmaTransaction bytes, _originatingMerkleProof bytes, _inputOfInterest uint256) returns(success bool)
func (_PlasmaParent *PlasmaParentTransactor) ProveBalanceOrOwnershipBreakingBetweenInputAndOutput(opts *bind.TransactOpts, _plasmaBlockNumber uint32, _plasmaTransaction []byte, _merkleProof []byte, _originatingPlasmaBlockNumber uint32, _originatingPlasmaTransaction []byte, _originatingMerkleProof []byte, _inputOfInterest *big.Int) (*types.Transaction, error) {
	return _PlasmaParent.contract.Transact(opts, "proveBalanceOrOwnershipBreakingBetweenInputAndOutput", _plasmaBlockNumber, _plasmaTransaction, _merkleProof, _originatingPlasmaBlockNumber, _originatingPlasmaTransaction, _originatingMerkleProof, _inputOfInterest)
}

// ProveBalanceOrOwnershipBreakingBetweenInputAndOutput is a paid mutator transaction binding the contract method 0xaee5b5de.
//
// Solidity: function proveBalanceOrOwnershipBreakingBetweenInputAndOutput(_plasmaBlockNumber uint32, _plasmaTransaction bytes, _merkleProof bytes, _originatingPlasmaBlockNumber uint32, _originatingPlasmaTransaction bytes, _originatingMerkleProof bytes, _inputOfInterest uint256) returns(success bool)
func (_PlasmaParent *PlasmaParentSession) ProveBalanceOrOwnershipBreakingBetweenInputAndOutput(_plasmaBlockNumber uint32, _plasmaTransaction []byte, _merkleProof []byte, _originatingPlasmaBlockNumber uint32, _originatingPlasmaTransaction []byte, _originatingMerkleProof []byte, _inputOfInterest *big.Int) (*types.Transaction, error) {
	return _PlasmaParent.Contract.ProveBalanceOrOwnershipBreakingBetweenInputAndOutput(&_PlasmaParent.TransactOpts, _plasmaBlockNumber, _plasmaTransaction, _merkleProof, _originatingPlasmaBlockNumber, _originatingPlasmaTransaction, _originatingMerkleProof, _inputOfInterest)
}

// ProveBalanceOrOwnershipBreakingBetweenInputAndOutput is a paid mutator transaction binding the contract method 0xaee5b5de.
//
// Solidity: function proveBalanceOrOwnershipBreakingBetweenInputAndOutput(_plasmaBlockNumber uint32, _plasmaTransaction bytes, _merkleProof bytes, _originatingPlasmaBlockNumber uint32, _originatingPlasmaTransaction bytes, _originatingMerkleProof bytes, _inputOfInterest uint256) returns(success bool)
func (_PlasmaParent *PlasmaParentTransactorSession) ProveBalanceOrOwnershipBreakingBetweenInputAndOutput(_plasmaBlockNumber uint32, _plasmaTransaction []byte, _merkleProof []byte, _originatingPlasmaBlockNumber uint32, _originatingPlasmaTransaction []byte, _originatingMerkleProof []byte, _inputOfInterest *big.Int) (*types.Transaction, error) {
	return _PlasmaParent.Contract.ProveBalanceOrOwnershipBreakingBetweenInputAndOutput(&_PlasmaParent.TransactOpts, _plasmaBlockNumber, _plasmaTransaction, _merkleProof, _originatingPlasmaBlockNumber, _originatingPlasmaTransaction, _originatingMerkleProof, _inputOfInterest)
}

// ProveDoubleFunding is a paid mutator transaction binding the contract method 0x411be8c0.
//
// Solidity: function proveDoubleFunding(_plasmaBlockNumber1 uint32, _plasmaTransaction1 bytes, _merkleProof1 bytes, _plasmaBlockNumber2 uint32, _plasmaTransaction2 bytes, _merkleProof2 bytes) returns(success bool)
func (_PlasmaParent *PlasmaParentTransactor) ProveDoubleFunding(opts *bind.TransactOpts, _plasmaBlockNumber1 uint32, _plasmaTransaction1 []byte, _merkleProof1 []byte, _plasmaBlockNumber2 uint32, _plasmaTransaction2 []byte, _merkleProof2 []byte) (*types.Transaction, error) {
	return _PlasmaParent.contract.Transact(opts, "proveDoubleFunding", _plasmaBlockNumber1, _plasmaTransaction1, _merkleProof1, _plasmaBlockNumber2, _plasmaTransaction2, _merkleProof2)
}

// ProveDoubleFunding is a paid mutator transaction binding the contract method 0x411be8c0.
//
// Solidity: function proveDoubleFunding(_plasmaBlockNumber1 uint32, _plasmaTransaction1 bytes, _merkleProof1 bytes, _plasmaBlockNumber2 uint32, _plasmaTransaction2 bytes, _merkleProof2 bytes) returns(success bool)
func (_PlasmaParent *PlasmaParentSession) ProveDoubleFunding(_plasmaBlockNumber1 uint32, _plasmaTransaction1 []byte, _merkleProof1 []byte, _plasmaBlockNumber2 uint32, _plasmaTransaction2 []byte, _merkleProof2 []byte) (*types.Transaction, error) {
	return _PlasmaParent.Contract.ProveDoubleFunding(&_PlasmaParent.TransactOpts, _plasmaBlockNumber1, _plasmaTransaction1, _merkleProof1, _plasmaBlockNumber2, _plasmaTransaction2, _merkleProof2)
}

// ProveDoubleFunding is a paid mutator transaction binding the contract method 0x411be8c0.
//
// Solidity: function proveDoubleFunding(_plasmaBlockNumber1 uint32, _plasmaTransaction1 bytes, _merkleProof1 bytes, _plasmaBlockNumber2 uint32, _plasmaTransaction2 bytes, _merkleProof2 bytes) returns(success bool)
func (_PlasmaParent *PlasmaParentTransactorSession) ProveDoubleFunding(_plasmaBlockNumber1 uint32, _plasmaTransaction1 []byte, _merkleProof1 []byte, _plasmaBlockNumber2 uint32, _plasmaTransaction2 []byte, _merkleProof2 []byte) (*types.Transaction, error) {
	return _PlasmaParent.Contract.ProveDoubleFunding(&_PlasmaParent.TransactOpts, _plasmaBlockNumber1, _plasmaTransaction1, _merkleProof1, _plasmaBlockNumber2, _plasmaTransaction2, _merkleProof2)
}

// ProveDoubleSpend is a paid mutator transaction binding the contract method 0xf1f52e56.
//
// Solidity: function proveDoubleSpend(_plasmaBlockNumber1 uint32, _inputNumber1 uint8, _plasmaTransaction1 bytes, _merkleProof1 bytes, _plasmaBlockNumber2 uint32, _inputNumber2 uint8, _plasmaTransaction2 bytes, _merkleProof2 bytes) returns(success bool)
func (_PlasmaParent *PlasmaParentTransactor) ProveDoubleSpend(opts *bind.TransactOpts, _plasmaBlockNumber1 uint32, _inputNumber1 uint8, _plasmaTransaction1 []byte, _merkleProof1 []byte, _plasmaBlockNumber2 uint32, _inputNumber2 uint8, _plasmaTransaction2 []byte, _merkleProof2 []byte) (*types.Transaction, error) {
	return _PlasmaParent.contract.Transact(opts, "proveDoubleSpend", _plasmaBlockNumber1, _inputNumber1, _plasmaTransaction1, _merkleProof1, _plasmaBlockNumber2, _inputNumber2, _plasmaTransaction2, _merkleProof2)
}

// ProveDoubleSpend is a paid mutator transaction binding the contract method 0xf1f52e56.
//
// Solidity: function proveDoubleSpend(_plasmaBlockNumber1 uint32, _inputNumber1 uint8, _plasmaTransaction1 bytes, _merkleProof1 bytes, _plasmaBlockNumber2 uint32, _inputNumber2 uint8, _plasmaTransaction2 bytes, _merkleProof2 bytes) returns(success bool)
func (_PlasmaParent *PlasmaParentSession) ProveDoubleSpend(_plasmaBlockNumber1 uint32, _inputNumber1 uint8, _plasmaTransaction1 []byte, _merkleProof1 []byte, _plasmaBlockNumber2 uint32, _inputNumber2 uint8, _plasmaTransaction2 []byte, _merkleProof2 []byte) (*types.Transaction, error) {
	return _PlasmaParent.Contract.ProveDoubleSpend(&_PlasmaParent.TransactOpts, _plasmaBlockNumber1, _inputNumber1, _plasmaTransaction1, _merkleProof1, _plasmaBlockNumber2, _inputNumber2, _plasmaTransaction2, _merkleProof2)
}

// ProveDoubleSpend is a paid mutator transaction binding the contract method 0xf1f52e56.
//
// Solidity: function proveDoubleSpend(_plasmaBlockNumber1 uint32, _inputNumber1 uint8, _plasmaTransaction1 bytes, _merkleProof1 bytes, _plasmaBlockNumber2 uint32, _inputNumber2 uint8, _plasmaTransaction2 bytes, _merkleProof2 bytes) returns(success bool)
func (_PlasmaParent *PlasmaParentTransactorSession) ProveDoubleSpend(_plasmaBlockNumber1 uint32, _inputNumber1 uint8, _plasmaTransaction1 []byte, _merkleProof1 []byte, _plasmaBlockNumber2 uint32, _inputNumber2 uint8, _plasmaTransaction2 []byte, _merkleProof2 []byte) (*types.Transaction, error) {
	return _PlasmaParent.Contract.ProveDoubleSpend(&_PlasmaParent.TransactOpts, _plasmaBlockNumber1, _inputNumber1, _plasmaTransaction1, _merkleProof1, _plasmaBlockNumber2, _inputNumber2, _plasmaTransaction2, _merkleProof2)
}

// ProveInvalidDeposit is a paid mutator transaction binding the contract method 0x1fd35b77.
//
// Solidity: function proveInvalidDeposit(_plasmaBlockNumber uint32, _plasmaTransaction bytes, _merkleProof bytes) returns(success bool)
func (_PlasmaParent *PlasmaParentTransactor) ProveInvalidDeposit(opts *bind.TransactOpts, _plasmaBlockNumber uint32, _plasmaTransaction []byte, _merkleProof []byte) (*types.Transaction, error) {
	return _PlasmaParent.contract.Transact(opts, "proveInvalidDeposit", _plasmaBlockNumber, _plasmaTransaction, _merkleProof)
}

// ProveInvalidDeposit is a paid mutator transaction binding the contract method 0x1fd35b77.
//
// Solidity: function proveInvalidDeposit(_plasmaBlockNumber uint32, _plasmaTransaction bytes, _merkleProof bytes) returns(success bool)
func (_PlasmaParent *PlasmaParentSession) ProveInvalidDeposit(_plasmaBlockNumber uint32, _plasmaTransaction []byte, _merkleProof []byte) (*types.Transaction, error) {
	return _PlasmaParent.Contract.ProveInvalidDeposit(&_PlasmaParent.TransactOpts, _plasmaBlockNumber, _plasmaTransaction, _merkleProof)
}

// ProveInvalidDeposit is a paid mutator transaction binding the contract method 0x1fd35b77.
//
// Solidity: function proveInvalidDeposit(_plasmaBlockNumber uint32, _plasmaTransaction bytes, _merkleProof bytes) returns(success bool)
func (_PlasmaParent *PlasmaParentTransactorSession) ProveInvalidDeposit(_plasmaBlockNumber uint32, _plasmaTransaction []byte, _merkleProof []byte) (*types.Transaction, error) {
	return _PlasmaParent.Contract.ProveInvalidDeposit(&_PlasmaParent.TransactOpts, _plasmaBlockNumber, _plasmaTransaction, _merkleProof)
}

// ProveReferencingInvalidBlock is a paid mutator transaction binding the contract method 0xab9508b5.
//
// Solidity: function proveReferencingInvalidBlock(_plasmaBlockNumber uint32, _plasmaInputNumberInTx uint8, _plasmaTransaction bytes, _merkleProof bytes) returns(success bool)
func (_PlasmaParent *PlasmaParentTransactor) ProveReferencingInvalidBlock(opts *bind.TransactOpts, _plasmaBlockNumber uint32, _plasmaInputNumberInTx uint8, _plasmaTransaction []byte, _merkleProof []byte) (*types.Transaction, error) {
	return _PlasmaParent.contract.Transact(opts, "proveReferencingInvalidBlock", _plasmaBlockNumber, _plasmaInputNumberInTx, _plasmaTransaction, _merkleProof)
}

// ProveReferencingInvalidBlock is a paid mutator transaction binding the contract method 0xab9508b5.
//
// Solidity: function proveReferencingInvalidBlock(_plasmaBlockNumber uint32, _plasmaInputNumberInTx uint8, _plasmaTransaction bytes, _merkleProof bytes) returns(success bool)
func (_PlasmaParent *PlasmaParentSession) ProveReferencingInvalidBlock(_plasmaBlockNumber uint32, _plasmaInputNumberInTx uint8, _plasmaTransaction []byte, _merkleProof []byte) (*types.Transaction, error) {
	return _PlasmaParent.Contract.ProveReferencingInvalidBlock(&_PlasmaParent.TransactOpts, _plasmaBlockNumber, _plasmaInputNumberInTx, _plasmaTransaction, _merkleProof)
}

// ProveReferencingInvalidBlock is a paid mutator transaction binding the contract method 0xab9508b5.
//
// Solidity: function proveReferencingInvalidBlock(_plasmaBlockNumber uint32, _plasmaInputNumberInTx uint8, _plasmaTransaction bytes, _merkleProof bytes) returns(success bool)
func (_PlasmaParent *PlasmaParentTransactorSession) ProveReferencingInvalidBlock(_plasmaBlockNumber uint32, _plasmaInputNumberInTx uint8, _plasmaTransaction []byte, _merkleProof []byte) (*types.Transaction, error) {
	return _PlasmaParent.Contract.ProveReferencingInvalidBlock(&_PlasmaParent.TransactOpts, _plasmaBlockNumber, _plasmaInputNumberInTx, _plasmaTransaction, _merkleProof)
}

// ProveSpendAndWithdraw is a paid mutator transaction binding the contract method 0xfdb4ad8f.
//
// Solidity: function proveSpendAndWithdraw(_plasmaBlockNumber uint32, _plasmaTransaction bytes, _merkleProof bytes, _originatingPlasmaTransaction bytes, _originatingMerkleProof bytes, _inputNumber uint8) returns(success bool)
func (_PlasmaParent *PlasmaParentTransactor) ProveSpendAndWithdraw(opts *bind.TransactOpts, _plasmaBlockNumber uint32, _plasmaTransaction []byte, _merkleProof []byte, _originatingPlasmaTransaction []byte, _originatingMerkleProof []byte, _inputNumber uint8) (*types.Transaction, error) {
	return _PlasmaParent.contract.Transact(opts, "proveSpendAndWithdraw", _plasmaBlockNumber, _plasmaTransaction, _merkleProof, _originatingPlasmaTransaction, _originatingMerkleProof, _inputNumber)
}

// ProveSpendAndWithdraw is a paid mutator transaction binding the contract method 0xfdb4ad8f.
//
// Solidity: function proveSpendAndWithdraw(_plasmaBlockNumber uint32, _plasmaTransaction bytes, _merkleProof bytes, _originatingPlasmaTransaction bytes, _originatingMerkleProof bytes, _inputNumber uint8) returns(success bool)
func (_PlasmaParent *PlasmaParentSession) ProveSpendAndWithdraw(_plasmaBlockNumber uint32, _plasmaTransaction []byte, _merkleProof []byte, _originatingPlasmaTransaction []byte, _originatingMerkleProof []byte, _inputNumber uint8) (*types.Transaction, error) {
	return _PlasmaParent.Contract.ProveSpendAndWithdraw(&_PlasmaParent.TransactOpts, _plasmaBlockNumber, _plasmaTransaction, _merkleProof, _originatingPlasmaTransaction, _originatingMerkleProof, _inputNumber)
}

// ProveSpendAndWithdraw is a paid mutator transaction binding the contract method 0xfdb4ad8f.
//
// Solidity: function proveSpendAndWithdraw(_plasmaBlockNumber uint32, _plasmaTransaction bytes, _merkleProof bytes, _originatingPlasmaTransaction bytes, _originatingMerkleProof bytes, _inputNumber uint8) returns(success bool)
func (_PlasmaParent *PlasmaParentTransactorSession) ProveSpendAndWithdraw(_plasmaBlockNumber uint32, _plasmaTransaction []byte, _merkleProof []byte, _originatingPlasmaTransaction []byte, _originatingMerkleProof []byte, _inputNumber uint8) (*types.Transaction, error) {
	return _PlasmaParent.Contract.ProveSpendAndWithdraw(&_PlasmaParent.TransactOpts, _plasmaBlockNumber, _plasmaTransaction, _merkleProof, _originatingPlasmaTransaction, _originatingMerkleProof, _inputNumber)
}

// ReturnExpiredBuyoutOffer is a paid mutator transaction binding the contract method 0x72a609ff.
//
// Solidity: function returnExpiredBuyoutOffer(_index bytes22) returns(success bool)
func (_PlasmaParent *PlasmaParentTransactor) ReturnExpiredBuyoutOffer(opts *bind.TransactOpts, _index [22]byte) (*types.Transaction, error) {
	return _PlasmaParent.contract.Transact(opts, "returnExpiredBuyoutOffer", _index)
}

// ReturnExpiredBuyoutOffer is a paid mutator transaction binding the contract method 0x72a609ff.
//
// Solidity: function returnExpiredBuyoutOffer(_index bytes22) returns(success bool)
func (_PlasmaParent *PlasmaParentSession) ReturnExpiredBuyoutOffer(_index [22]byte) (*types.Transaction, error) {
	return _PlasmaParent.Contract.ReturnExpiredBuyoutOffer(&_PlasmaParent.TransactOpts, _index)
}

// ReturnExpiredBuyoutOffer is a paid mutator transaction binding the contract method 0x72a609ff.
//
// Solidity: function returnExpiredBuyoutOffer(_index bytes22) returns(success bool)
func (_PlasmaParent *PlasmaParentTransactorSession) ReturnExpiredBuyoutOffer(_index [22]byte) (*types.Transaction, error) {
	return _PlasmaParent.Contract.ReturnExpiredBuyoutOffer(&_PlasmaParent.TransactOpts, _index)
}

// SetOperator is a paid mutator transaction binding the contract method 0x77754136.
//
// Solidity: function setOperator(_op address, _status uint256) returns(success bool)
func (_PlasmaParent *PlasmaParentTransactor) SetOperator(opts *bind.TransactOpts, _op common.Address, _status *big.Int) (*types.Transaction, error) {
	return _PlasmaParent.contract.Transact(opts, "setOperator", _op, _status)
}

// SetOperator is a paid mutator transaction binding the contract method 0x77754136.
//
// Solidity: function setOperator(_op address, _status uint256) returns(success bool)
func (_PlasmaParent *PlasmaParentSession) SetOperator(_op common.Address, _status *big.Int) (*types.Transaction, error) {
	return _PlasmaParent.Contract.SetOperator(&_PlasmaParent.TransactOpts, _op, _status)
}

// SetOperator is a paid mutator transaction binding the contract method 0x77754136.
//
// Solidity: function setOperator(_op address, _status uint256) returns(success bool)
func (_PlasmaParent *PlasmaParentTransactorSession) SetOperator(_op common.Address, _status *big.Int) (*types.Transaction, error) {
	return _PlasmaParent.Contract.SetOperator(&_PlasmaParent.TransactOpts, _op, _status)
}

// StartDepositWithdraw is a paid mutator transaction binding the contract method 0x039aebae.
//
// Solidity: function startDepositWithdraw(depositIndex uint256) returns(success bool)
func (_PlasmaParent *PlasmaParentTransactor) StartDepositWithdraw(opts *bind.TransactOpts, depositIndex *big.Int) (*types.Transaction, error) {
	return _PlasmaParent.contract.Transact(opts, "startDepositWithdraw", depositIndex)
}

// StartDepositWithdraw is a paid mutator transaction binding the contract method 0x039aebae.
//
// Solidity: function startDepositWithdraw(depositIndex uint256) returns(success bool)
func (_PlasmaParent *PlasmaParentSession) StartDepositWithdraw(depositIndex *big.Int) (*types.Transaction, error) {
	return _PlasmaParent.Contract.StartDepositWithdraw(&_PlasmaParent.TransactOpts, depositIndex)
}

// StartDepositWithdraw is a paid mutator transaction binding the contract method 0x039aebae.
//
// Solidity: function startDepositWithdraw(depositIndex uint256) returns(success bool)
func (_PlasmaParent *PlasmaParentTransactorSession) StartDepositWithdraw(depositIndex *big.Int) (*types.Transaction, error) {
	return _PlasmaParent.Contract.StartDepositWithdraw(&_PlasmaParent.TransactOpts, depositIndex)
}

// StartExit is a paid mutator transaction binding the contract method 0x50790907.
//
// Solidity: function startExit(_plasmaBlockNumber uint32, _outputNumber uint8, _plasmaTransaction bytes, _merkleProof bytes) returns(success bool)
func (_PlasmaParent *PlasmaParentTransactor) StartExit(opts *bind.TransactOpts, _plasmaBlockNumber uint32, _outputNumber uint8, _plasmaTransaction []byte, _merkleProof []byte) (*types.Transaction, error) {
	return _PlasmaParent.contract.Transact(opts, "startExit", _plasmaBlockNumber, _outputNumber, _plasmaTransaction, _merkleProof)
}

// StartExit is a paid mutator transaction binding the contract method 0x50790907.
//
// Solidity: function startExit(_plasmaBlockNumber uint32, _outputNumber uint8, _plasmaTransaction bytes, _merkleProof bytes) returns(success bool)
func (_PlasmaParent *PlasmaParentSession) StartExit(_plasmaBlockNumber uint32, _outputNumber uint8, _plasmaTransaction []byte, _merkleProof []byte) (*types.Transaction, error) {
	return _PlasmaParent.Contract.StartExit(&_PlasmaParent.TransactOpts, _plasmaBlockNumber, _outputNumber, _plasmaTransaction, _merkleProof)
}

// StartExit is a paid mutator transaction binding the contract method 0x50790907.
//
// Solidity: function startExit(_plasmaBlockNumber uint32, _outputNumber uint8, _plasmaTransaction bytes, _merkleProof bytes) returns(success bool)
func (_PlasmaParent *PlasmaParentTransactorSession) StartExit(_plasmaBlockNumber uint32, _outputNumber uint8, _plasmaTransaction []byte, _merkleProof []byte) (*types.Transaction, error) {
	return _PlasmaParent.Contract.StartExit(&_PlasmaParent.TransactOpts, _plasmaBlockNumber, _outputNumber, _plasmaTransaction, _merkleProof)
}

// SubmitBlockHeaders is a paid mutator transaction binding the contract method 0x8d92ce46.
//
// Solidity: function submitBlockHeaders(_headers bytes) returns(success bool)
func (_PlasmaParent *PlasmaParentTransactor) SubmitBlockHeaders(opts *bind.TransactOpts, _headers []byte) (*types.Transaction, error) {
	return _PlasmaParent.contract.Transact(opts, "submitBlockHeaders", _headers)
}

// SubmitBlockHeaders is a paid mutator transaction binding the contract method 0x8d92ce46.
//
// Solidity: function submitBlockHeaders(_headers bytes) returns(success bool)
func (_PlasmaParent *PlasmaParentSession) SubmitBlockHeaders(_headers []byte) (*types.Transaction, error) {
	return _PlasmaParent.Contract.SubmitBlockHeaders(&_PlasmaParent.TransactOpts, _headers)
}

// SubmitBlockHeaders is a paid mutator transaction binding the contract method 0x8d92ce46.
//
// Solidity: function submitBlockHeaders(_headers bytes) returns(success bool)
func (_PlasmaParent *PlasmaParentTransactorSession) SubmitBlockHeaders(_headers []byte) (*types.Transaction, error) {
	return _PlasmaParent.Contract.SubmitBlockHeaders(&_PlasmaParent.TransactOpts, _headers)
}

// PlasmaParentDepositEventIterator is returned from FilterDepositEvent and is used to iterate over the raw logs and unpacked data for DepositEvent events raised by the PlasmaParent contract.
type PlasmaParentDepositEventIterator struct {
	Event *PlasmaParentDepositEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PlasmaParentDepositEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlasmaParentDepositEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PlasmaParentDepositEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PlasmaParentDepositEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlasmaParentDepositEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlasmaParentDepositEvent represents a DepositEvent event raised by the PlasmaParent contract.
type PlasmaParentDepositEvent struct {
	From         common.Address
	Amount       *big.Int
	DepositIndex *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDepositEvent is a free log retrieval operation binding the contract event 0xad40ae5dc69974ba932d08b0a608e89109412d41d04850f5196f144875ae2660.
//
// Solidity: e DepositEvent(_from indexed address, _amount indexed uint256, _depositIndex indexed uint256)
func (_PlasmaParent *PlasmaParentFilterer) FilterDepositEvent(opts *bind.FilterOpts, _from []common.Address, _amount []*big.Int, _depositIndex []*big.Int) (*PlasmaParentDepositEventIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}
	var _depositIndexRule []interface{}
	for _, _depositIndexItem := range _depositIndex {
		_depositIndexRule = append(_depositIndexRule, _depositIndexItem)
	}

	logs, sub, err := _PlasmaParent.contract.FilterLogs(opts, "DepositEvent", _fromRule, _amountRule, _depositIndexRule)
	if err != nil {
		return nil, err
	}
	return &PlasmaParentDepositEventIterator{contract: _PlasmaParent.contract, event: "DepositEvent", logs: logs, sub: sub}, nil
}

// WatchDepositEvent is a free log subscription operation binding the contract event 0xad40ae5dc69974ba932d08b0a608e89109412d41d04850f5196f144875ae2660.
//
// Solidity: e DepositEvent(_from indexed address, _amount indexed uint256, _depositIndex indexed uint256)
func (_PlasmaParent *PlasmaParentFilterer) WatchDepositEvent(opts *bind.WatchOpts, sink chan<- *PlasmaParentDepositEvent, _from []common.Address, _amount []*big.Int, _depositIndex []*big.Int) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _amountRule []interface{}
	for _, _amountItem := range _amount {
		_amountRule = append(_amountRule, _amountItem)
	}
	var _depositIndexRule []interface{}
	for _, _depositIndexItem := range _depositIndex {
		_depositIndexRule = append(_depositIndexRule, _depositIndexItem)
	}

	logs, sub, err := _PlasmaParent.contract.WatchLogs(opts, "DepositEvent", _fromRule, _amountRule, _depositIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlasmaParentDepositEvent)
				if err := _PlasmaParent.contract.UnpackLog(event, "DepositEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// PlasmaParentDepositWithdrawChallengedEventIterator is returned from FilterDepositWithdrawChallengedEvent and is used to iterate over the raw logs and unpacked data for DepositWithdrawChallengedEvent events raised by the PlasmaParent contract.
type PlasmaParentDepositWithdrawChallengedEventIterator struct {
	Event *PlasmaParentDepositWithdrawChallengedEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PlasmaParentDepositWithdrawChallengedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlasmaParentDepositWithdrawChallengedEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PlasmaParentDepositWithdrawChallengedEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PlasmaParentDepositWithdrawChallengedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlasmaParentDepositWithdrawChallengedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlasmaParentDepositWithdrawChallengedEvent represents a DepositWithdrawChallengedEvent event raised by the PlasmaParent contract.
type PlasmaParentDepositWithdrawChallengedEvent struct {
	DepositIndex *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDepositWithdrawChallengedEvent is a free log retrieval operation binding the contract event 0x51d45e8928e39b43363b1ea18991bc166b59295a20374e0fbaeca912b3de9f39.
//
// Solidity: e DepositWithdrawChallengedEvent(_depositIndex indexed uint256)
func (_PlasmaParent *PlasmaParentFilterer) FilterDepositWithdrawChallengedEvent(opts *bind.FilterOpts, _depositIndex []*big.Int) (*PlasmaParentDepositWithdrawChallengedEventIterator, error) {

	var _depositIndexRule []interface{}
	for _, _depositIndexItem := range _depositIndex {
		_depositIndexRule = append(_depositIndexRule, _depositIndexItem)
	}

	logs, sub, err := _PlasmaParent.contract.FilterLogs(opts, "DepositWithdrawChallengedEvent", _depositIndexRule)
	if err != nil {
		return nil, err
	}
	return &PlasmaParentDepositWithdrawChallengedEventIterator{contract: _PlasmaParent.contract, event: "DepositWithdrawChallengedEvent", logs: logs, sub: sub}, nil
}

// WatchDepositWithdrawChallengedEvent is a free log subscription operation binding the contract event 0x51d45e8928e39b43363b1ea18991bc166b59295a20374e0fbaeca912b3de9f39.
//
// Solidity: e DepositWithdrawChallengedEvent(_depositIndex indexed uint256)
func (_PlasmaParent *PlasmaParentFilterer) WatchDepositWithdrawChallengedEvent(opts *bind.WatchOpts, sink chan<- *PlasmaParentDepositWithdrawChallengedEvent, _depositIndex []*big.Int) (event.Subscription, error) {

	var _depositIndexRule []interface{}
	for _, _depositIndexItem := range _depositIndex {
		_depositIndexRule = append(_depositIndexRule, _depositIndexItem)
	}

	logs, sub, err := _PlasmaParent.contract.WatchLogs(opts, "DepositWithdrawChallengedEvent", _depositIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlasmaParentDepositWithdrawChallengedEvent)
				if err := _PlasmaParent.contract.UnpackLog(event, "DepositWithdrawChallengedEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// PlasmaParentDepositWithdrawCompletedEventIterator is returned from FilterDepositWithdrawCompletedEvent and is used to iterate over the raw logs and unpacked data for DepositWithdrawCompletedEvent events raised by the PlasmaParent contract.
type PlasmaParentDepositWithdrawCompletedEventIterator struct {
	Event *PlasmaParentDepositWithdrawCompletedEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PlasmaParentDepositWithdrawCompletedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlasmaParentDepositWithdrawCompletedEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PlasmaParentDepositWithdrawCompletedEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PlasmaParentDepositWithdrawCompletedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlasmaParentDepositWithdrawCompletedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlasmaParentDepositWithdrawCompletedEvent represents a DepositWithdrawCompletedEvent event raised by the PlasmaParent contract.
type PlasmaParentDepositWithdrawCompletedEvent struct {
	DepositIndex *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDepositWithdrawCompletedEvent is a free log retrieval operation binding the contract event 0xb7f7af1d1ba18a21232c7efd7d2279f6ed111802f0f11f40af09f8959752152c.
//
// Solidity: e DepositWithdrawCompletedEvent(_depositIndex indexed uint256)
func (_PlasmaParent *PlasmaParentFilterer) FilterDepositWithdrawCompletedEvent(opts *bind.FilterOpts, _depositIndex []*big.Int) (*PlasmaParentDepositWithdrawCompletedEventIterator, error) {

	var _depositIndexRule []interface{}
	for _, _depositIndexItem := range _depositIndex {
		_depositIndexRule = append(_depositIndexRule, _depositIndexItem)
	}

	logs, sub, err := _PlasmaParent.contract.FilterLogs(opts, "DepositWithdrawCompletedEvent", _depositIndexRule)
	if err != nil {
		return nil, err
	}
	return &PlasmaParentDepositWithdrawCompletedEventIterator{contract: _PlasmaParent.contract, event: "DepositWithdrawCompletedEvent", logs: logs, sub: sub}, nil
}

// WatchDepositWithdrawCompletedEvent is a free log subscription operation binding the contract event 0xb7f7af1d1ba18a21232c7efd7d2279f6ed111802f0f11f40af09f8959752152c.
//
// Solidity: e DepositWithdrawCompletedEvent(_depositIndex indexed uint256)
func (_PlasmaParent *PlasmaParentFilterer) WatchDepositWithdrawCompletedEvent(opts *bind.WatchOpts, sink chan<- *PlasmaParentDepositWithdrawCompletedEvent, _depositIndex []*big.Int) (event.Subscription, error) {

	var _depositIndexRule []interface{}
	for _, _depositIndexItem := range _depositIndex {
		_depositIndexRule = append(_depositIndexRule, _depositIndexItem)
	}

	logs, sub, err := _PlasmaParent.contract.WatchLogs(opts, "DepositWithdrawCompletedEvent", _depositIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlasmaParentDepositWithdrawCompletedEvent)
				if err := _PlasmaParent.contract.UnpackLog(event, "DepositWithdrawCompletedEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// PlasmaParentDepositWithdrawStartedEventIterator is returned from FilterDepositWithdrawStartedEvent and is used to iterate over the raw logs and unpacked data for DepositWithdrawStartedEvent events raised by the PlasmaParent contract.
type PlasmaParentDepositWithdrawStartedEventIterator struct {
	Event *PlasmaParentDepositWithdrawStartedEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PlasmaParentDepositWithdrawStartedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlasmaParentDepositWithdrawStartedEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PlasmaParentDepositWithdrawStartedEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PlasmaParentDepositWithdrawStartedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlasmaParentDepositWithdrawStartedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlasmaParentDepositWithdrawStartedEvent represents a DepositWithdrawStartedEvent event raised by the PlasmaParent contract.
type PlasmaParentDepositWithdrawStartedEvent struct {
	DepositIndex *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterDepositWithdrawStartedEvent is a free log retrieval operation binding the contract event 0x712ab85cc5498d9a867fa8c7d02fbefb065b07611e9c6d861b0a75c11e7f780d.
//
// Solidity: e DepositWithdrawStartedEvent(_depositIndex indexed uint256)
func (_PlasmaParent *PlasmaParentFilterer) FilterDepositWithdrawStartedEvent(opts *bind.FilterOpts, _depositIndex []*big.Int) (*PlasmaParentDepositWithdrawStartedEventIterator, error) {

	var _depositIndexRule []interface{}
	for _, _depositIndexItem := range _depositIndex {
		_depositIndexRule = append(_depositIndexRule, _depositIndexItem)
	}

	logs, sub, err := _PlasmaParent.contract.FilterLogs(opts, "DepositWithdrawStartedEvent", _depositIndexRule)
	if err != nil {
		return nil, err
	}
	return &PlasmaParentDepositWithdrawStartedEventIterator{contract: _PlasmaParent.contract, event: "DepositWithdrawStartedEvent", logs: logs, sub: sub}, nil
}

// WatchDepositWithdrawStartedEvent is a free log subscription operation binding the contract event 0x712ab85cc5498d9a867fa8c7d02fbefb065b07611e9c6d861b0a75c11e7f780d.
//
// Solidity: e DepositWithdrawStartedEvent(_depositIndex indexed uint256)
func (_PlasmaParent *PlasmaParentFilterer) WatchDepositWithdrawStartedEvent(opts *bind.WatchOpts, sink chan<- *PlasmaParentDepositWithdrawStartedEvent, _depositIndex []*big.Int) (event.Subscription, error) {

	var _depositIndexRule []interface{}
	for _, _depositIndexItem := range _depositIndex {
		_depositIndexRule = append(_depositIndexRule, _depositIndexItem)
	}

	logs, sub, err := _PlasmaParent.contract.WatchLogs(opts, "DepositWithdrawStartedEvent", _depositIndexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlasmaParentDepositWithdrawStartedEvent)
				if err := _PlasmaParent.contract.UnpackLog(event, "DepositWithdrawStartedEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// PlasmaParentErrorFoundEventIterator is returned from FilterErrorFoundEvent and is used to iterate over the raw logs and unpacked data for ErrorFoundEvent events raised by the PlasmaParent contract.
type PlasmaParentErrorFoundEventIterator struct {
	Event *PlasmaParentErrorFoundEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PlasmaParentErrorFoundEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlasmaParentErrorFoundEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PlasmaParentErrorFoundEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PlasmaParentErrorFoundEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlasmaParentErrorFoundEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlasmaParentErrorFoundEvent represents a ErrorFoundEvent event raised by the PlasmaParent contract.
type PlasmaParentErrorFoundEvent struct {
	LastValidBlockNumber *big.Int
	Raw                  types.Log // Blockchain specific contextual infos
}

// FilterErrorFoundEvent is a free log retrieval operation binding the contract event 0xbd33ae2c31ade34dbd760f286fa27138a3f259b97286a475d8bb5d9b43284948.
//
// Solidity: e ErrorFoundEvent(_lastValidBlockNumber indexed uint256)
func (_PlasmaParent *PlasmaParentFilterer) FilterErrorFoundEvent(opts *bind.FilterOpts, _lastValidBlockNumber []*big.Int) (*PlasmaParentErrorFoundEventIterator, error) {

	var _lastValidBlockNumberRule []interface{}
	for _, _lastValidBlockNumberItem := range _lastValidBlockNumber {
		_lastValidBlockNumberRule = append(_lastValidBlockNumberRule, _lastValidBlockNumberItem)
	}

	logs, sub, err := _PlasmaParent.contract.FilterLogs(opts, "ErrorFoundEvent", _lastValidBlockNumberRule)
	if err != nil {
		return nil, err
	}
	return &PlasmaParentErrorFoundEventIterator{contract: _PlasmaParent.contract, event: "ErrorFoundEvent", logs: logs, sub: sub}, nil
}

// WatchErrorFoundEvent is a free log subscription operation binding the contract event 0xbd33ae2c31ade34dbd760f286fa27138a3f259b97286a475d8bb5d9b43284948.
//
// Solidity: e ErrorFoundEvent(_lastValidBlockNumber indexed uint256)
func (_PlasmaParent *PlasmaParentFilterer) WatchErrorFoundEvent(opts *bind.WatchOpts, sink chan<- *PlasmaParentErrorFoundEvent, _lastValidBlockNumber []*big.Int) (event.Subscription, error) {

	var _lastValidBlockNumberRule []interface{}
	for _, _lastValidBlockNumberItem := range _lastValidBlockNumber {
		_lastValidBlockNumberRule = append(_lastValidBlockNumberRule, _lastValidBlockNumberItem)
	}

	logs, sub, err := _PlasmaParent.contract.WatchLogs(opts, "ErrorFoundEvent", _lastValidBlockNumberRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlasmaParentErrorFoundEvent)
				if err := _PlasmaParent.contract.UnpackLog(event, "ErrorFoundEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// PlasmaParentExitBuyoutAcceptedIterator is returned from FilterExitBuyoutAccepted and is used to iterate over the raw logs and unpacked data for ExitBuyoutAccepted events raised by the PlasmaParent contract.
type PlasmaParentExitBuyoutAcceptedIterator struct {
	Event *PlasmaParentExitBuyoutAccepted // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PlasmaParentExitBuyoutAcceptedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlasmaParentExitBuyoutAccepted)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PlasmaParentExitBuyoutAccepted)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PlasmaParentExitBuyoutAcceptedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlasmaParentExitBuyoutAcceptedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlasmaParentExitBuyoutAccepted represents a ExitBuyoutAccepted event raised by the PlasmaParent contract.
type PlasmaParentExitBuyoutAccepted struct {
	PartialHash [22]byte
	From        common.Address
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterExitBuyoutAccepted is a free log retrieval operation binding the contract event 0x8a732776b9b7dffbc4e537a4acf181dfb9914a637b58054e4667661ab82b6095.
//
// Solidity: e ExitBuyoutAccepted(_partialHash indexed bytes22, _from indexed address)
func (_PlasmaParent *PlasmaParentFilterer) FilterExitBuyoutAccepted(opts *bind.FilterOpts, _partialHash [][22]byte, _from []common.Address) (*PlasmaParentExitBuyoutAcceptedIterator, error) {

	var _partialHashRule []interface{}
	for _, _partialHashItem := range _partialHash {
		_partialHashRule = append(_partialHashRule, _partialHashItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _PlasmaParent.contract.FilterLogs(opts, "ExitBuyoutAccepted", _partialHashRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return &PlasmaParentExitBuyoutAcceptedIterator{contract: _PlasmaParent.contract, event: "ExitBuyoutAccepted", logs: logs, sub: sub}, nil
}

// WatchExitBuyoutAccepted is a free log subscription operation binding the contract event 0x8a732776b9b7dffbc4e537a4acf181dfb9914a637b58054e4667661ab82b6095.
//
// Solidity: e ExitBuyoutAccepted(_partialHash indexed bytes22, _from indexed address)
func (_PlasmaParent *PlasmaParentFilterer) WatchExitBuyoutAccepted(opts *bind.WatchOpts, sink chan<- *PlasmaParentExitBuyoutAccepted, _partialHash [][22]byte, _from []common.Address) (event.Subscription, error) {

	var _partialHashRule []interface{}
	for _, _partialHashItem := range _partialHash {
		_partialHashRule = append(_partialHashRule, _partialHashItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	logs, sub, err := _PlasmaParent.contract.WatchLogs(opts, "ExitBuyoutAccepted", _partialHashRule, _fromRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlasmaParentExitBuyoutAccepted)
				if err := _PlasmaParent.contract.UnpackLog(event, "ExitBuyoutAccepted", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// PlasmaParentExitBuyoutOfferedIterator is returned from FilterExitBuyoutOffered and is used to iterate over the raw logs and unpacked data for ExitBuyoutOffered events raised by the PlasmaParent contract.
type PlasmaParentExitBuyoutOfferedIterator struct {
	Event *PlasmaParentExitBuyoutOffered // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PlasmaParentExitBuyoutOfferedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlasmaParentExitBuyoutOffered)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PlasmaParentExitBuyoutOffered)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PlasmaParentExitBuyoutOfferedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlasmaParentExitBuyoutOfferedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlasmaParentExitBuyoutOffered represents a ExitBuyoutOffered event raised by the PlasmaParent contract.
type PlasmaParentExitBuyoutOffered struct {
	PartialHash  [22]byte
	From         common.Address
	BuyoutAmount *big.Int
	Raw          types.Log // Blockchain specific contextual infos
}

// FilterExitBuyoutOffered is a free log retrieval operation binding the contract event 0xe406e90405bc7b16c4327a40c5d5592a74ca552f6c83a568fb3e8da4849cf0fa.
//
// Solidity: e ExitBuyoutOffered(_partialHash indexed bytes22, _from indexed address, _buyoutAmount indexed uint256)
func (_PlasmaParent *PlasmaParentFilterer) FilterExitBuyoutOffered(opts *bind.FilterOpts, _partialHash [][22]byte, _from []common.Address, _buyoutAmount []*big.Int) (*PlasmaParentExitBuyoutOfferedIterator, error) {

	var _partialHashRule []interface{}
	for _, _partialHashItem := range _partialHash {
		_partialHashRule = append(_partialHashRule, _partialHashItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _buyoutAmountRule []interface{}
	for _, _buyoutAmountItem := range _buyoutAmount {
		_buyoutAmountRule = append(_buyoutAmountRule, _buyoutAmountItem)
	}

	logs, sub, err := _PlasmaParent.contract.FilterLogs(opts, "ExitBuyoutOffered", _partialHashRule, _fromRule, _buyoutAmountRule)
	if err != nil {
		return nil, err
	}
	return &PlasmaParentExitBuyoutOfferedIterator{contract: _PlasmaParent.contract, event: "ExitBuyoutOffered", logs: logs, sub: sub}, nil
}

// WatchExitBuyoutOffered is a free log subscription operation binding the contract event 0xe406e90405bc7b16c4327a40c5d5592a74ca552f6c83a568fb3e8da4849cf0fa.
//
// Solidity: e ExitBuyoutOffered(_partialHash indexed bytes22, _from indexed address, _buyoutAmount indexed uint256)
func (_PlasmaParent *PlasmaParentFilterer) WatchExitBuyoutOffered(opts *bind.WatchOpts, sink chan<- *PlasmaParentExitBuyoutOffered, _partialHash [][22]byte, _from []common.Address, _buyoutAmount []*big.Int) (event.Subscription, error) {

	var _partialHashRule []interface{}
	for _, _partialHashItem := range _partialHash {
		_partialHashRule = append(_partialHashRule, _partialHashItem)
	}
	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _buyoutAmountRule []interface{}
	for _, _buyoutAmountItem := range _buyoutAmount {
		_buyoutAmountRule = append(_buyoutAmountRule, _buyoutAmountItem)
	}

	logs, sub, err := _PlasmaParent.contract.WatchLogs(opts, "ExitBuyoutOffered", _partialHashRule, _fromRule, _buyoutAmountRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlasmaParentExitBuyoutOffered)
				if err := _PlasmaParent.contract.UnpackLog(event, "ExitBuyoutOffered", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// PlasmaParentExitChallengedIterator is returned from FilterExitChallenged and is used to iterate over the raw logs and unpacked data for ExitChallenged events raised by the PlasmaParent contract.
type PlasmaParentExitChallengedIterator struct {
	Event *PlasmaParentExitChallenged // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PlasmaParentExitChallengedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlasmaParentExitChallenged)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PlasmaParentExitChallenged)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PlasmaParentExitChallengedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlasmaParentExitChallengedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlasmaParentExitChallenged represents a ExitChallenged event raised by the PlasmaParent contract.
type PlasmaParentExitChallenged struct {
	Hash [22]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterExitChallenged is a free log retrieval operation binding the contract event 0xbddf872402da7ab20174f1569fa5bda4cd087a73b2108a5b70803be137bc75fb.
//
// Solidity: e ExitChallenged(_hash indexed bytes22)
func (_PlasmaParent *PlasmaParentFilterer) FilterExitChallenged(opts *bind.FilterOpts, _hash [][22]byte) (*PlasmaParentExitChallengedIterator, error) {

	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}

	logs, sub, err := _PlasmaParent.contract.FilterLogs(opts, "ExitChallenged", _hashRule)
	if err != nil {
		return nil, err
	}
	return &PlasmaParentExitChallengedIterator{contract: _PlasmaParent.contract, event: "ExitChallenged", logs: logs, sub: sub}, nil
}

// WatchExitChallenged is a free log subscription operation binding the contract event 0xbddf872402da7ab20174f1569fa5bda4cd087a73b2108a5b70803be137bc75fb.
//
// Solidity: e ExitChallenged(_hash indexed bytes22)
func (_PlasmaParent *PlasmaParentFilterer) WatchExitChallenged(opts *bind.WatchOpts, sink chan<- *PlasmaParentExitChallenged, _hash [][22]byte) (event.Subscription, error) {

	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}

	logs, sub, err := _PlasmaParent.contract.WatchLogs(opts, "ExitChallenged", _hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlasmaParentExitChallenged)
				if err := _PlasmaParent.contract.UnpackLog(event, "ExitChallenged", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// PlasmaParentExitRecordCreatedIterator is returned from FilterExitRecordCreated and is used to iterate over the raw logs and unpacked data for ExitRecordCreated events raised by the PlasmaParent contract.
type PlasmaParentExitRecordCreatedIterator struct {
	Event *PlasmaParentExitRecordCreated // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PlasmaParentExitRecordCreatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlasmaParentExitRecordCreated)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PlasmaParentExitRecordCreated)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PlasmaParentExitRecordCreatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlasmaParentExitRecordCreatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlasmaParentExitRecordCreated represents a ExitRecordCreated event raised by the PlasmaParent contract.
type PlasmaParentExitRecordCreated struct {
	Hash [22]byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterExitRecordCreated is a free log retrieval operation binding the contract event 0x5ead8ca9ab45ce1e07a752c71e011fa2f0d7a9d4b9cfba42c2032ae0cf6ca047.
//
// Solidity: e ExitRecordCreated(_hash indexed bytes22)
func (_PlasmaParent *PlasmaParentFilterer) FilterExitRecordCreated(opts *bind.FilterOpts, _hash [][22]byte) (*PlasmaParentExitRecordCreatedIterator, error) {

	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}

	logs, sub, err := _PlasmaParent.contract.FilterLogs(opts, "ExitRecordCreated", _hashRule)
	if err != nil {
		return nil, err
	}
	return &PlasmaParentExitRecordCreatedIterator{contract: _PlasmaParent.contract, event: "ExitRecordCreated", logs: logs, sub: sub}, nil
}

// WatchExitRecordCreated is a free log subscription operation binding the contract event 0x5ead8ca9ab45ce1e07a752c71e011fa2f0d7a9d4b9cfba42c2032ae0cf6ca047.
//
// Solidity: e ExitRecordCreated(_hash indexed bytes22)
func (_PlasmaParent *PlasmaParentFilterer) WatchExitRecordCreated(opts *bind.WatchOpts, sink chan<- *PlasmaParentExitRecordCreated, _hash [][22]byte) (event.Subscription, error) {

	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}

	logs, sub, err := _PlasmaParent.contract.WatchLogs(opts, "ExitRecordCreated", _hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlasmaParentExitRecordCreated)
				if err := _PlasmaParent.contract.UnpackLog(event, "ExitRecordCreated", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// PlasmaParentExitStartedEventIterator is returned from FilterExitStartedEvent and is used to iterate over the raw logs and unpacked data for ExitStartedEvent events raised by the PlasmaParent contract.
type PlasmaParentExitStartedEventIterator struct {
	Event *PlasmaParentExitStartedEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PlasmaParentExitStartedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlasmaParentExitStartedEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PlasmaParentExitStartedEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PlasmaParentExitStartedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlasmaParentExitStartedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlasmaParentExitStartedEvent represents a ExitStartedEvent event raised by the PlasmaParent contract.
type PlasmaParentExitStartedEvent struct {
	From     common.Address
	Priority *big.Int
	Index    *big.Int
	Hash     [22]byte
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterExitStartedEvent is a free log retrieval operation binding the contract event 0xe5b1890f8a8a29ed5112599c90e8a9417367f23307d8f8349053ea6198b86673.
//
// Solidity: e ExitStartedEvent(_from indexed address, _priority uint72, _index indexed uint72, _hash indexed bytes22)
func (_PlasmaParent *PlasmaParentFilterer) FilterExitStartedEvent(opts *bind.FilterOpts, _from []common.Address, _index []*big.Int, _hash [][22]byte) (*PlasmaParentExitStartedEventIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	var _indexRule []interface{}
	for _, _indexItem := range _index {
		_indexRule = append(_indexRule, _indexItem)
	}
	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}

	logs, sub, err := _PlasmaParent.contract.FilterLogs(opts, "ExitStartedEvent", _fromRule, _indexRule, _hashRule)
	if err != nil {
		return nil, err
	}
	return &PlasmaParentExitStartedEventIterator{contract: _PlasmaParent.contract, event: "ExitStartedEvent", logs: logs, sub: sub}, nil
}

// WatchExitStartedEvent is a free log subscription operation binding the contract event 0xe5b1890f8a8a29ed5112599c90e8a9417367f23307d8f8349053ea6198b86673.
//
// Solidity: e ExitStartedEvent(_from indexed address, _priority uint72, _index indexed uint72, _hash indexed bytes22)
func (_PlasmaParent *PlasmaParentFilterer) WatchExitStartedEvent(opts *bind.WatchOpts, sink chan<- *PlasmaParentExitStartedEvent, _from []common.Address, _index []*big.Int, _hash [][22]byte) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}

	var _indexRule []interface{}
	for _, _indexItem := range _index {
		_indexRule = append(_indexRule, _indexItem)
	}
	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}

	logs, sub, err := _PlasmaParent.contract.WatchLogs(opts, "ExitStartedEvent", _fromRule, _indexRule, _hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlasmaParentExitStartedEvent)
				if err := _PlasmaParent.contract.UnpackLog(event, "ExitStartedEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// PlasmaParentLimboExitStartedEventIterator is returned from FilterLimboExitStartedEvent and is used to iterate over the raw logs and unpacked data for LimboExitStartedEvent events raised by the PlasmaParent contract.
type PlasmaParentLimboExitStartedEventIterator struct {
	Event *PlasmaParentLimboExitStartedEvent // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PlasmaParentLimboExitStartedEventIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlasmaParentLimboExitStartedEvent)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PlasmaParentLimboExitStartedEvent)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PlasmaParentLimboExitStartedEventIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlasmaParentLimboExitStartedEventIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlasmaParentLimboExitStartedEvent represents a LimboExitStartedEvent event raised by the PlasmaParent contract.
type PlasmaParentLimboExitStartedEvent struct {
	From        common.Address
	Priority    *big.Int
	PartialHash [22]byte
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterLimboExitStartedEvent is a free log retrieval operation binding the contract event 0x8ef6ad4f1f0d802a05685ccdfb001a441ee56782e5408973eafdc332363f733a.
//
// Solidity: e LimboExitStartedEvent(_from indexed address, _priority indexed uint72, _partialHash indexed bytes22)
func (_PlasmaParent *PlasmaParentFilterer) FilterLimboExitStartedEvent(opts *bind.FilterOpts, _from []common.Address, _priority []*big.Int, _partialHash [][22]byte) (*PlasmaParentLimboExitStartedEventIterator, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _priorityRule []interface{}
	for _, _priorityItem := range _priority {
		_priorityRule = append(_priorityRule, _priorityItem)
	}
	var _partialHashRule []interface{}
	for _, _partialHashItem := range _partialHash {
		_partialHashRule = append(_partialHashRule, _partialHashItem)
	}

	logs, sub, err := _PlasmaParent.contract.FilterLogs(opts, "LimboExitStartedEvent", _fromRule, _priorityRule, _partialHashRule)
	if err != nil {
		return nil, err
	}
	return &PlasmaParentLimboExitStartedEventIterator{contract: _PlasmaParent.contract, event: "LimboExitStartedEvent", logs: logs, sub: sub}, nil
}

// WatchLimboExitStartedEvent is a free log subscription operation binding the contract event 0x8ef6ad4f1f0d802a05685ccdfb001a441ee56782e5408973eafdc332363f733a.
//
// Solidity: e LimboExitStartedEvent(_from indexed address, _priority indexed uint72, _partialHash indexed bytes22)
func (_PlasmaParent *PlasmaParentFilterer) WatchLimboExitStartedEvent(opts *bind.WatchOpts, sink chan<- *PlasmaParentLimboExitStartedEvent, _from []common.Address, _priority []*big.Int, _partialHash [][22]byte) (event.Subscription, error) {

	var _fromRule []interface{}
	for _, _fromItem := range _from {
		_fromRule = append(_fromRule, _fromItem)
	}
	var _priorityRule []interface{}
	for _, _priorityItem := range _priority {
		_priorityRule = append(_priorityRule, _priorityItem)
	}
	var _partialHashRule []interface{}
	for _, _partialHashItem := range _partialHash {
		_partialHashRule = append(_partialHashRule, _partialHashItem)
	}

	logs, sub, err := _PlasmaParent.contract.WatchLogs(opts, "LimboExitStartedEvent", _fromRule, _priorityRule, _partialHashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlasmaParentLimboExitStartedEvent)
				if err := _PlasmaParent.contract.UnpackLog(event, "LimboExitStartedEvent", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// PlasmaParentTransactionIsPublishedIterator is returned from FilterTransactionIsPublished and is used to iterate over the raw logs and unpacked data for TransactionIsPublished events raised by the PlasmaParent contract.
type PlasmaParentTransactionIsPublishedIterator struct {
	Event *PlasmaParentTransactionIsPublished // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PlasmaParentTransactionIsPublishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlasmaParentTransactionIsPublished)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PlasmaParentTransactionIsPublished)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PlasmaParentTransactionIsPublishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlasmaParentTransactionIsPublishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlasmaParentTransactionIsPublished represents a TransactionIsPublished event raised by the PlasmaParent contract.
type PlasmaParentTransactionIsPublished struct {
	Index uint64
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransactionIsPublished is a free log retrieval operation binding the contract event 0xb711080b0479517f47f00fc0db5ea448f048cfb4066f65df2f6df6cb216b4af2.
//
// Solidity: e TransactionIsPublished(_index indexed uint64)
func (_PlasmaParent *PlasmaParentFilterer) FilterTransactionIsPublished(opts *bind.FilterOpts, _index []uint64) (*PlasmaParentTransactionIsPublishedIterator, error) {

	var _indexRule []interface{}
	for _, _indexItem := range _index {
		_indexRule = append(_indexRule, _indexItem)
	}

	logs, sub, err := _PlasmaParent.contract.FilterLogs(opts, "TransactionIsPublished", _indexRule)
	if err != nil {
		return nil, err
	}
	return &PlasmaParentTransactionIsPublishedIterator{contract: _PlasmaParent.contract, event: "TransactionIsPublished", logs: logs, sub: sub}, nil
}

// WatchTransactionIsPublished is a free log subscription operation binding the contract event 0xb711080b0479517f47f00fc0db5ea448f048cfb4066f65df2f6df6cb216b4af2.
//
// Solidity: e TransactionIsPublished(_index indexed uint64)
func (_PlasmaParent *PlasmaParentFilterer) WatchTransactionIsPublished(opts *bind.WatchOpts, sink chan<- *PlasmaParentTransactionIsPublished, _index []uint64) (event.Subscription, error) {

	var _indexRule []interface{}
	for _, _indexItem := range _index {
		_indexRule = append(_indexRule, _indexItem)
	}

	logs, sub, err := _PlasmaParent.contract.WatchLogs(opts, "TransactionIsPublished", _indexRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlasmaParentTransactionIsPublished)
				if err := _PlasmaParent.contract.UnpackLog(event, "TransactionIsPublished", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}

// PlasmaParentTransactionPublishedIterator is returned from FilterTransactionPublished and is used to iterate over the raw logs and unpacked data for TransactionPublished events raised by the PlasmaParent contract.
type PlasmaParentTransactionPublishedIterator struct {
	Event *PlasmaParentTransactionPublished // Event containing the contract specifics and raw log

	contract *bind.BoundContract // Generic contract to use for unpacking event data
	event    string              // Event name to use for unpacking event data

	logs chan types.Log        // Log channel receiving the found contract events
	sub  ethereum.Subscription // Subscription for errors, completion and termination
	done bool                  // Whether the subscription completed delivering logs
	fail error                 // Occurred error to stop iteration
}

// Next advances the iterator to the subsequent event, returning whether there
// are any more events found. In case of a retrieval or parsing error, false is
// returned and Error() can be queried for the exact failure.
func (it *PlasmaParentTransactionPublishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PlasmaParentTransactionPublished)
			if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
				it.fail = err
				return false
			}
			it.Event.Raw = log
			return true

		default:
			return false
		}
	}
	// Iterator still in progress, wait for either a data or an error event
	select {
	case log := <-it.logs:
		it.Event = new(PlasmaParentTransactionPublished)
		if err := it.contract.UnpackLog(it.Event, it.event, log); err != nil {
			it.fail = err
			return false
		}
		it.Event.Raw = log
		return true

	case err := <-it.sub.Err():
		it.done = true
		it.fail = err
		return it.Next()
	}
}

// Error returns any retrieval or parsing error occurred during filtering.
func (it *PlasmaParentTransactionPublishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PlasmaParentTransactionPublishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PlasmaParentTransactionPublished represents a TransactionPublished event raised by the PlasmaParent contract.
type PlasmaParentTransactionPublished struct {
	Hash [32]byte
	Data []byte
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTransactionPublished is a free log retrieval operation binding the contract event 0x5801adacaa607b7a58d602d13f84bd912554bcbaf9ca07f03badfbd7be50a9b9.
//
// Solidity: e TransactionPublished(_hash indexed bytes32, _data bytes)
func (_PlasmaParent *PlasmaParentFilterer) FilterTransactionPublished(opts *bind.FilterOpts, _hash [][32]byte) (*PlasmaParentTransactionPublishedIterator, error) {

	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}

	logs, sub, err := _PlasmaParent.contract.FilterLogs(opts, "TransactionPublished", _hashRule)
	if err != nil {
		return nil, err
	}
	return &PlasmaParentTransactionPublishedIterator{contract: _PlasmaParent.contract, event: "TransactionPublished", logs: logs, sub: sub}, nil
}

// WatchTransactionPublished is a free log subscription operation binding the contract event 0x5801adacaa607b7a58d602d13f84bd912554bcbaf9ca07f03badfbd7be50a9b9.
//
// Solidity: e TransactionPublished(_hash indexed bytes32, _data bytes)
func (_PlasmaParent *PlasmaParentFilterer) WatchTransactionPublished(opts *bind.WatchOpts, sink chan<- *PlasmaParentTransactionPublished, _hash [][32]byte) (event.Subscription, error) {

	var _hashRule []interface{}
	for _, _hashItem := range _hash {
		_hashRule = append(_hashRule, _hashItem)
	}

	logs, sub, err := _PlasmaParent.contract.WatchLogs(opts, "TransactionPublished", _hashRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PlasmaParentTransactionPublished)
				if err := _PlasmaParent.contract.UnpackLog(event, "TransactionPublished", log); err != nil {
					return err
				}
				event.Raw = log

				select {
				case sink <- event:
				case err := <-sub.Err():
					return err
				case <-quit:
					return nil
				}
			case err := <-sub.Err():
				return err
			case <-quit:
				return nil
			}
		}
	}), nil
}
