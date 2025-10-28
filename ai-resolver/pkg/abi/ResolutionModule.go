// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abi

import (
	"errors"
	"math/big"
	"strings"

	ethereum "github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/event"
)

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = errors.New
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
	_ = abi.ConvertType
)

// ResolutionModuleMetaData contains all meta data concerning the ResolutionModule contract.
var ResolutionModuleMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_outcomeToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_bondToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_arbitrator\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"arbitrator\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bondToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"canDispute\",\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"canFinalize\",\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"dispute\",\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"bondAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"reason\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"disputeWindow\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"finalize\",\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"finalizeDisputed\",\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outcomeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"slashProposer\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"getDisputeTimeRemaining\",\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getResolutionState\",\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"\",\"type\":\"uint8\",\"internalType\":\"enumResolutionModule.ResolutionState\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"minBond\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"outcomeToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractOutcomeToken\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposeResolution\",\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outcomeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"bondAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"evidenceURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resolutions\",\"inputs\":[{\"name\":\"\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[{\"name\":\"state\",\"type\":\"uint8\",\"internalType\":\"enumResolutionModule.ResolutionState\"},{\"name\":\"proposedOutcome\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proposalTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"proposer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"proposerBond\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"disputer\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"disputerBond\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"evidenceURI\",\"type\":\"string\",\"internalType\":\"string\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"setArbitrator\",\"inputs\":[{\"name\":\"newArbitrator\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setDisputeWindow\",\"inputs\":[{\"name\":\"newWindow\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"setMinBond\",\"inputs\":[{\"name\":\"newMinBond\",\"type\":\"uint256\",\"internalType\":\"uint256\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"event\",\"name\":\"ArbitratorUpdated\",\"inputs\":[{\"name\":\"oldArbitrator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newArbitrator\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BondRefunded\",\"inputs\":[{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"BondSlashed\",\"inputs\":[{\"name\":\"slashedAddress\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"amount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"recipient\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"DisputeWindowUpdated\",\"inputs\":[{\"name\":\"oldWindow\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newWindow\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Disputed\",\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"disputer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"bond\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"reason\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"Finalized\",\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"outcomeId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"wasDisputed\",\"type\":\"bool\",\"indexed\":false,\"internalType\":\"bool\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"MinBondUpdated\",\"inputs\":[{\"name\":\"oldBond\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"newBond\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"ResolutionProposed\",\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"outcomeId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"proposer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"bond\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"evidenceURI\",\"type\":\"string\",\"indexed\":false,\"internalType\":\"string\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"DisputeWindowClosed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"DisputeWindowOpen\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InsufficientBond\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidBondAmount\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidOutcome\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidState\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MarketAlreadyResolved\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"Unauthorized\",\"inputs\":[]}]",
}

// ResolutionModuleABI is the input ABI used to generate the binding from.
// Deprecated: Use ResolutionModuleMetaData.ABI instead.
var ResolutionModuleABI = ResolutionModuleMetaData.ABI

// ResolutionModule is an auto generated Go binding around an Ethereum contract.
type ResolutionModule struct {
	ResolutionModuleCaller     // Read-only binding to the contract
	ResolutionModuleTransactor // Write-only binding to the contract
	ResolutionModuleFilterer   // Log filterer for contract events
}

// ResolutionModuleCaller is an auto generated read-only Go binding around an Ethereum contract.
type ResolutionModuleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResolutionModuleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ResolutionModuleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResolutionModuleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ResolutionModuleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ResolutionModuleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ResolutionModuleSession struct {
	Contract     *ResolutionModule // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ResolutionModuleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ResolutionModuleCallerSession struct {
	Contract *ResolutionModuleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// ResolutionModuleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ResolutionModuleTransactorSession struct {
	Contract     *ResolutionModuleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// ResolutionModuleRaw is an auto generated low-level Go binding around an Ethereum contract.
type ResolutionModuleRaw struct {
	Contract *ResolutionModule // Generic contract binding to access the raw methods on
}

// ResolutionModuleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ResolutionModuleCallerRaw struct {
	Contract *ResolutionModuleCaller // Generic read-only contract binding to access the raw methods on
}

// ResolutionModuleTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ResolutionModuleTransactorRaw struct {
	Contract *ResolutionModuleTransactor // Generic write-only contract binding to access the raw methods on
}

// NewResolutionModule creates a new instance of ResolutionModule, bound to a specific deployed contract.
func NewResolutionModule(address common.Address, backend bind.ContractBackend) (*ResolutionModule, error) {
	contract, err := bindResolutionModule(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ResolutionModule{ResolutionModuleCaller: ResolutionModuleCaller{contract: contract}, ResolutionModuleTransactor: ResolutionModuleTransactor{contract: contract}, ResolutionModuleFilterer: ResolutionModuleFilterer{contract: contract}}, nil
}

// NewResolutionModuleCaller creates a new read-only instance of ResolutionModule, bound to a specific deployed contract.
func NewResolutionModuleCaller(address common.Address, caller bind.ContractCaller) (*ResolutionModuleCaller, error) {
	contract, err := bindResolutionModule(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ResolutionModuleCaller{contract: contract}, nil
}

// NewResolutionModuleTransactor creates a new write-only instance of ResolutionModule, bound to a specific deployed contract.
func NewResolutionModuleTransactor(address common.Address, transactor bind.ContractTransactor) (*ResolutionModuleTransactor, error) {
	contract, err := bindResolutionModule(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ResolutionModuleTransactor{contract: contract}, nil
}

// NewResolutionModuleFilterer creates a new log filterer instance of ResolutionModule, bound to a specific deployed contract.
func NewResolutionModuleFilterer(address common.Address, filterer bind.ContractFilterer) (*ResolutionModuleFilterer, error) {
	contract, err := bindResolutionModule(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ResolutionModuleFilterer{contract: contract}, nil
}

// bindResolutionModule binds a generic wrapper to an already deployed contract.
func bindResolutionModule(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ResolutionModuleMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ResolutionModule *ResolutionModuleRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ResolutionModule.Contract.ResolutionModuleCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ResolutionModule *ResolutionModuleRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ResolutionModule.Contract.ResolutionModuleTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ResolutionModule *ResolutionModuleRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ResolutionModule.Contract.ResolutionModuleTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ResolutionModule *ResolutionModuleCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ResolutionModule.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ResolutionModule *ResolutionModuleTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ResolutionModule.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ResolutionModule *ResolutionModuleTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ResolutionModule.Contract.contract.Transact(opts, method, params...)
}

// Arbitrator is a free data retrieval call binding the contract method 0x6cc6cde1.
//
// Solidity: function arbitrator() view returns(address)
func (_ResolutionModule *ResolutionModuleCaller) Arbitrator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ResolutionModule.contract.Call(opts, &out, "arbitrator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Arbitrator is a free data retrieval call binding the contract method 0x6cc6cde1.
//
// Solidity: function arbitrator() view returns(address)
func (_ResolutionModule *ResolutionModuleSession) Arbitrator() (common.Address, error) {
	return _ResolutionModule.Contract.Arbitrator(&_ResolutionModule.CallOpts)
}

// Arbitrator is a free data retrieval call binding the contract method 0x6cc6cde1.
//
// Solidity: function arbitrator() view returns(address)
func (_ResolutionModule *ResolutionModuleCallerSession) Arbitrator() (common.Address, error) {
	return _ResolutionModule.Contract.Arbitrator(&_ResolutionModule.CallOpts)
}

// BondToken is a free data retrieval call binding the contract method 0xc28f4392.
//
// Solidity: function bondToken() view returns(address)
func (_ResolutionModule *ResolutionModuleCaller) BondToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ResolutionModule.contract.Call(opts, &out, "bondToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BondToken is a free data retrieval call binding the contract method 0xc28f4392.
//
// Solidity: function bondToken() view returns(address)
func (_ResolutionModule *ResolutionModuleSession) BondToken() (common.Address, error) {
	return _ResolutionModule.Contract.BondToken(&_ResolutionModule.CallOpts)
}

// BondToken is a free data retrieval call binding the contract method 0xc28f4392.
//
// Solidity: function bondToken() view returns(address)
func (_ResolutionModule *ResolutionModuleCallerSession) BondToken() (common.Address, error) {
	return _ResolutionModule.Contract.BondToken(&_ResolutionModule.CallOpts)
}

// CanDispute is a free data retrieval call binding the contract method 0xac1b2335.
//
// Solidity: function canDispute(uint256 marketId) view returns(bool)
func (_ResolutionModule *ResolutionModuleCaller) CanDispute(opts *bind.CallOpts, marketId *big.Int) (bool, error) {
	var out []interface{}
	err := _ResolutionModule.contract.Call(opts, &out, "canDispute", marketId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanDispute is a free data retrieval call binding the contract method 0xac1b2335.
//
// Solidity: function canDispute(uint256 marketId) view returns(bool)
func (_ResolutionModule *ResolutionModuleSession) CanDispute(marketId *big.Int) (bool, error) {
	return _ResolutionModule.Contract.CanDispute(&_ResolutionModule.CallOpts, marketId)
}

// CanDispute is a free data retrieval call binding the contract method 0xac1b2335.
//
// Solidity: function canDispute(uint256 marketId) view returns(bool)
func (_ResolutionModule *ResolutionModuleCallerSession) CanDispute(marketId *big.Int) (bool, error) {
	return _ResolutionModule.Contract.CanDispute(&_ResolutionModule.CallOpts, marketId)
}

// CanFinalize is a free data retrieval call binding the contract method 0xe4e2bfe4.
//
// Solidity: function canFinalize(uint256 marketId) view returns(bool)
func (_ResolutionModule *ResolutionModuleCaller) CanFinalize(opts *bind.CallOpts, marketId *big.Int) (bool, error) {
	var out []interface{}
	err := _ResolutionModule.contract.Call(opts, &out, "canFinalize", marketId)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// CanFinalize is a free data retrieval call binding the contract method 0xe4e2bfe4.
//
// Solidity: function canFinalize(uint256 marketId) view returns(bool)
func (_ResolutionModule *ResolutionModuleSession) CanFinalize(marketId *big.Int) (bool, error) {
	return _ResolutionModule.Contract.CanFinalize(&_ResolutionModule.CallOpts, marketId)
}

// CanFinalize is a free data retrieval call binding the contract method 0xe4e2bfe4.
//
// Solidity: function canFinalize(uint256 marketId) view returns(bool)
func (_ResolutionModule *ResolutionModuleCallerSession) CanFinalize(marketId *big.Int) (bool, error) {
	return _ResolutionModule.Contract.CanFinalize(&_ResolutionModule.CallOpts, marketId)
}

// DisputeWindow is a free data retrieval call binding the contract method 0x117f5f92.
//
// Solidity: function disputeWindow() view returns(uint256)
func (_ResolutionModule *ResolutionModuleCaller) DisputeWindow(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ResolutionModule.contract.Call(opts, &out, "disputeWindow")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// DisputeWindow is a free data retrieval call binding the contract method 0x117f5f92.
//
// Solidity: function disputeWindow() view returns(uint256)
func (_ResolutionModule *ResolutionModuleSession) DisputeWindow() (*big.Int, error) {
	return _ResolutionModule.Contract.DisputeWindow(&_ResolutionModule.CallOpts)
}

// DisputeWindow is a free data retrieval call binding the contract method 0x117f5f92.
//
// Solidity: function disputeWindow() view returns(uint256)
func (_ResolutionModule *ResolutionModuleCallerSession) DisputeWindow() (*big.Int, error) {
	return _ResolutionModule.Contract.DisputeWindow(&_ResolutionModule.CallOpts)
}

// GetDisputeTimeRemaining is a free data retrieval call binding the contract method 0xa8880fdc.
//
// Solidity: function getDisputeTimeRemaining(uint256 marketId) view returns(uint256)
func (_ResolutionModule *ResolutionModuleCaller) GetDisputeTimeRemaining(opts *bind.CallOpts, marketId *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _ResolutionModule.contract.Call(opts, &out, "getDisputeTimeRemaining", marketId)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetDisputeTimeRemaining is a free data retrieval call binding the contract method 0xa8880fdc.
//
// Solidity: function getDisputeTimeRemaining(uint256 marketId) view returns(uint256)
func (_ResolutionModule *ResolutionModuleSession) GetDisputeTimeRemaining(marketId *big.Int) (*big.Int, error) {
	return _ResolutionModule.Contract.GetDisputeTimeRemaining(&_ResolutionModule.CallOpts, marketId)
}

// GetDisputeTimeRemaining is a free data retrieval call binding the contract method 0xa8880fdc.
//
// Solidity: function getDisputeTimeRemaining(uint256 marketId) view returns(uint256)
func (_ResolutionModule *ResolutionModuleCallerSession) GetDisputeTimeRemaining(marketId *big.Int) (*big.Int, error) {
	return _ResolutionModule.Contract.GetDisputeTimeRemaining(&_ResolutionModule.CallOpts, marketId)
}

// GetResolutionState is a free data retrieval call binding the contract method 0x86e50433.
//
// Solidity: function getResolutionState(uint256 marketId) view returns(uint8)
func (_ResolutionModule *ResolutionModuleCaller) GetResolutionState(opts *bind.CallOpts, marketId *big.Int) (uint8, error) {
	var out []interface{}
	err := _ResolutionModule.contract.Call(opts, &out, "getResolutionState", marketId)

	if err != nil {
		return *new(uint8), err
	}

	out0 := *abi.ConvertType(out[0], new(uint8)).(*uint8)

	return out0, err

}

// GetResolutionState is a free data retrieval call binding the contract method 0x86e50433.
//
// Solidity: function getResolutionState(uint256 marketId) view returns(uint8)
func (_ResolutionModule *ResolutionModuleSession) GetResolutionState(marketId *big.Int) (uint8, error) {
	return _ResolutionModule.Contract.GetResolutionState(&_ResolutionModule.CallOpts, marketId)
}

// GetResolutionState is a free data retrieval call binding the contract method 0x86e50433.
//
// Solidity: function getResolutionState(uint256 marketId) view returns(uint8)
func (_ResolutionModule *ResolutionModuleCallerSession) GetResolutionState(marketId *big.Int) (uint8, error) {
	return _ResolutionModule.Contract.GetResolutionState(&_ResolutionModule.CallOpts, marketId)
}

// MinBond is a free data retrieval call binding the contract method 0x831518b7.
//
// Solidity: function minBond() view returns(uint256)
func (_ResolutionModule *ResolutionModuleCaller) MinBond(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ResolutionModule.contract.Call(opts, &out, "minBond")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinBond is a free data retrieval call binding the contract method 0x831518b7.
//
// Solidity: function minBond() view returns(uint256)
func (_ResolutionModule *ResolutionModuleSession) MinBond() (*big.Int, error) {
	return _ResolutionModule.Contract.MinBond(&_ResolutionModule.CallOpts)
}

// MinBond is a free data retrieval call binding the contract method 0x831518b7.
//
// Solidity: function minBond() view returns(uint256)
func (_ResolutionModule *ResolutionModuleCallerSession) MinBond() (*big.Int, error) {
	return _ResolutionModule.Contract.MinBond(&_ResolutionModule.CallOpts)
}

// OutcomeToken is a free data retrieval call binding the contract method 0xa998d6d8.
//
// Solidity: function outcomeToken() view returns(address)
func (_ResolutionModule *ResolutionModuleCaller) OutcomeToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ResolutionModule.contract.Call(opts, &out, "outcomeToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// OutcomeToken is a free data retrieval call binding the contract method 0xa998d6d8.
//
// Solidity: function outcomeToken() view returns(address)
func (_ResolutionModule *ResolutionModuleSession) OutcomeToken() (common.Address, error) {
	return _ResolutionModule.Contract.OutcomeToken(&_ResolutionModule.CallOpts)
}

// OutcomeToken is a free data retrieval call binding the contract method 0xa998d6d8.
//
// Solidity: function outcomeToken() view returns(address)
func (_ResolutionModule *ResolutionModuleCallerSession) OutcomeToken() (common.Address, error) {
	return _ResolutionModule.Contract.OutcomeToken(&_ResolutionModule.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ResolutionModule *ResolutionModuleCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _ResolutionModule.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ResolutionModule *ResolutionModuleSession) Owner() (common.Address, error) {
	return _ResolutionModule.Contract.Owner(&_ResolutionModule.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_ResolutionModule *ResolutionModuleCallerSession) Owner() (common.Address, error) {
	return _ResolutionModule.Contract.Owner(&_ResolutionModule.CallOpts)
}

// Resolutions is a free data retrieval call binding the contract method 0xa4b7f5ce.
//
// Solidity: function resolutions(uint256 ) view returns(uint8 state, uint256 proposedOutcome, uint256 proposalTime, address proposer, uint256 proposerBond, address disputer, uint256 disputerBond, string evidenceURI)
func (_ResolutionModule *ResolutionModuleCaller) Resolutions(opts *bind.CallOpts, arg0 *big.Int) (struct {
	State           uint8
	ProposedOutcome *big.Int
	ProposalTime    *big.Int
	Proposer        common.Address
	ProposerBond    *big.Int
	Disputer        common.Address
	DisputerBond    *big.Int
	EvidenceURI     string
}, error) {
	var out []interface{}
	err := _ResolutionModule.contract.Call(opts, &out, "resolutions", arg0)

	outstruct := new(struct {
		State           uint8
		ProposedOutcome *big.Int
		ProposalTime    *big.Int
		Proposer        common.Address
		ProposerBond    *big.Int
		Disputer        common.Address
		DisputerBond    *big.Int
		EvidenceURI     string
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.State = *abi.ConvertType(out[0], new(uint8)).(*uint8)
	outstruct.ProposedOutcome = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.ProposalTime = *abi.ConvertType(out[2], new(*big.Int)).(**big.Int)
	outstruct.Proposer = *abi.ConvertType(out[3], new(common.Address)).(*common.Address)
	outstruct.ProposerBond = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)
	outstruct.Disputer = *abi.ConvertType(out[5], new(common.Address)).(*common.Address)
	outstruct.DisputerBond = *abi.ConvertType(out[6], new(*big.Int)).(**big.Int)
	outstruct.EvidenceURI = *abi.ConvertType(out[7], new(string)).(*string)

	return *outstruct, err

}

// Resolutions is a free data retrieval call binding the contract method 0xa4b7f5ce.
//
// Solidity: function resolutions(uint256 ) view returns(uint8 state, uint256 proposedOutcome, uint256 proposalTime, address proposer, uint256 proposerBond, address disputer, uint256 disputerBond, string evidenceURI)
func (_ResolutionModule *ResolutionModuleSession) Resolutions(arg0 *big.Int) (struct {
	State           uint8
	ProposedOutcome *big.Int
	ProposalTime    *big.Int
	Proposer        common.Address
	ProposerBond    *big.Int
	Disputer        common.Address
	DisputerBond    *big.Int
	EvidenceURI     string
}, error) {
	return _ResolutionModule.Contract.Resolutions(&_ResolutionModule.CallOpts, arg0)
}

// Resolutions is a free data retrieval call binding the contract method 0xa4b7f5ce.
//
// Solidity: function resolutions(uint256 ) view returns(uint8 state, uint256 proposedOutcome, uint256 proposalTime, address proposer, uint256 proposerBond, address disputer, uint256 disputerBond, string evidenceURI)
func (_ResolutionModule *ResolutionModuleCallerSession) Resolutions(arg0 *big.Int) (struct {
	State           uint8
	ProposedOutcome *big.Int
	ProposalTime    *big.Int
	Proposer        common.Address
	ProposerBond    *big.Int
	Disputer        common.Address
	DisputerBond    *big.Int
	EvidenceURI     string
}, error) {
	return _ResolutionModule.Contract.Resolutions(&_ResolutionModule.CallOpts, arg0)
}

// Dispute is a paid mutator transaction binding the contract method 0xf8bcfca4.
//
// Solidity: function dispute(uint256 marketId, uint256 bondAmount, string reason) returns()
func (_ResolutionModule *ResolutionModuleTransactor) Dispute(opts *bind.TransactOpts, marketId *big.Int, bondAmount *big.Int, reason string) (*types.Transaction, error) {
	return _ResolutionModule.contract.Transact(opts, "dispute", marketId, bondAmount, reason)
}

// Dispute is a paid mutator transaction binding the contract method 0xf8bcfca4.
//
// Solidity: function dispute(uint256 marketId, uint256 bondAmount, string reason) returns()
func (_ResolutionModule *ResolutionModuleSession) Dispute(marketId *big.Int, bondAmount *big.Int, reason string) (*types.Transaction, error) {
	return _ResolutionModule.Contract.Dispute(&_ResolutionModule.TransactOpts, marketId, bondAmount, reason)
}

// Dispute is a paid mutator transaction binding the contract method 0xf8bcfca4.
//
// Solidity: function dispute(uint256 marketId, uint256 bondAmount, string reason) returns()
func (_ResolutionModule *ResolutionModuleTransactorSession) Dispute(marketId *big.Int, bondAmount *big.Int, reason string) (*types.Transaction, error) {
	return _ResolutionModule.Contract.Dispute(&_ResolutionModule.TransactOpts, marketId, bondAmount, reason)
}

// Finalize is a paid mutator transaction binding the contract method 0x05261aea.
//
// Solidity: function finalize(uint256 marketId) returns()
func (_ResolutionModule *ResolutionModuleTransactor) Finalize(opts *bind.TransactOpts, marketId *big.Int) (*types.Transaction, error) {
	return _ResolutionModule.contract.Transact(opts, "finalize", marketId)
}

// Finalize is a paid mutator transaction binding the contract method 0x05261aea.
//
// Solidity: function finalize(uint256 marketId) returns()
func (_ResolutionModule *ResolutionModuleSession) Finalize(marketId *big.Int) (*types.Transaction, error) {
	return _ResolutionModule.Contract.Finalize(&_ResolutionModule.TransactOpts, marketId)
}

// Finalize is a paid mutator transaction binding the contract method 0x05261aea.
//
// Solidity: function finalize(uint256 marketId) returns()
func (_ResolutionModule *ResolutionModuleTransactorSession) Finalize(marketId *big.Int) (*types.Transaction, error) {
	return _ResolutionModule.Contract.Finalize(&_ResolutionModule.TransactOpts, marketId)
}

// FinalizeDisputed is a paid mutator transaction binding the contract method 0xc1848b9a.
//
// Solidity: function finalizeDisputed(uint256 marketId, uint256 outcomeId, bool slashProposer) returns()
func (_ResolutionModule *ResolutionModuleTransactor) FinalizeDisputed(opts *bind.TransactOpts, marketId *big.Int, outcomeId *big.Int, slashProposer bool) (*types.Transaction, error) {
	return _ResolutionModule.contract.Transact(opts, "finalizeDisputed", marketId, outcomeId, slashProposer)
}

// FinalizeDisputed is a paid mutator transaction binding the contract method 0xc1848b9a.
//
// Solidity: function finalizeDisputed(uint256 marketId, uint256 outcomeId, bool slashProposer) returns()
func (_ResolutionModule *ResolutionModuleSession) FinalizeDisputed(marketId *big.Int, outcomeId *big.Int, slashProposer bool) (*types.Transaction, error) {
	return _ResolutionModule.Contract.FinalizeDisputed(&_ResolutionModule.TransactOpts, marketId, outcomeId, slashProposer)
}

// FinalizeDisputed is a paid mutator transaction binding the contract method 0xc1848b9a.
//
// Solidity: function finalizeDisputed(uint256 marketId, uint256 outcomeId, bool slashProposer) returns()
func (_ResolutionModule *ResolutionModuleTransactorSession) FinalizeDisputed(marketId *big.Int, outcomeId *big.Int, slashProposer bool) (*types.Transaction, error) {
	return _ResolutionModule.Contract.FinalizeDisputed(&_ResolutionModule.TransactOpts, marketId, outcomeId, slashProposer)
}

// ProposeResolution is a paid mutator transaction binding the contract method 0x262bc78c.
//
// Solidity: function proposeResolution(uint256 marketId, uint256 outcomeId, uint256 bondAmount, string evidenceURI) returns()
func (_ResolutionModule *ResolutionModuleTransactor) ProposeResolution(opts *bind.TransactOpts, marketId *big.Int, outcomeId *big.Int, bondAmount *big.Int, evidenceURI string) (*types.Transaction, error) {
	return _ResolutionModule.contract.Transact(opts, "proposeResolution", marketId, outcomeId, bondAmount, evidenceURI)
}

// ProposeResolution is a paid mutator transaction binding the contract method 0x262bc78c.
//
// Solidity: function proposeResolution(uint256 marketId, uint256 outcomeId, uint256 bondAmount, string evidenceURI) returns()
func (_ResolutionModule *ResolutionModuleSession) ProposeResolution(marketId *big.Int, outcomeId *big.Int, bondAmount *big.Int, evidenceURI string) (*types.Transaction, error) {
	return _ResolutionModule.Contract.ProposeResolution(&_ResolutionModule.TransactOpts, marketId, outcomeId, bondAmount, evidenceURI)
}

// ProposeResolution is a paid mutator transaction binding the contract method 0x262bc78c.
//
// Solidity: function proposeResolution(uint256 marketId, uint256 outcomeId, uint256 bondAmount, string evidenceURI) returns()
func (_ResolutionModule *ResolutionModuleTransactorSession) ProposeResolution(marketId *big.Int, outcomeId *big.Int, bondAmount *big.Int, evidenceURI string) (*types.Transaction, error) {
	return _ResolutionModule.Contract.ProposeResolution(&_ResolutionModule.TransactOpts, marketId, outcomeId, bondAmount, evidenceURI)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ResolutionModule *ResolutionModuleTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ResolutionModule.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ResolutionModule *ResolutionModuleSession) RenounceOwnership() (*types.Transaction, error) {
	return _ResolutionModule.Contract.RenounceOwnership(&_ResolutionModule.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_ResolutionModule *ResolutionModuleTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _ResolutionModule.Contract.RenounceOwnership(&_ResolutionModule.TransactOpts)
}

// SetArbitrator is a paid mutator transaction binding the contract method 0xb0eefabe.
//
// Solidity: function setArbitrator(address newArbitrator) returns()
func (_ResolutionModule *ResolutionModuleTransactor) SetArbitrator(opts *bind.TransactOpts, newArbitrator common.Address) (*types.Transaction, error) {
	return _ResolutionModule.contract.Transact(opts, "setArbitrator", newArbitrator)
}

// SetArbitrator is a paid mutator transaction binding the contract method 0xb0eefabe.
//
// Solidity: function setArbitrator(address newArbitrator) returns()
func (_ResolutionModule *ResolutionModuleSession) SetArbitrator(newArbitrator common.Address) (*types.Transaction, error) {
	return _ResolutionModule.Contract.SetArbitrator(&_ResolutionModule.TransactOpts, newArbitrator)
}

// SetArbitrator is a paid mutator transaction binding the contract method 0xb0eefabe.
//
// Solidity: function setArbitrator(address newArbitrator) returns()
func (_ResolutionModule *ResolutionModuleTransactorSession) SetArbitrator(newArbitrator common.Address) (*types.Transaction, error) {
	return _ResolutionModule.Contract.SetArbitrator(&_ResolutionModule.TransactOpts, newArbitrator)
}

// SetDisputeWindow is a paid mutator transaction binding the contract method 0x332226d0.
//
// Solidity: function setDisputeWindow(uint256 newWindow) returns()
func (_ResolutionModule *ResolutionModuleTransactor) SetDisputeWindow(opts *bind.TransactOpts, newWindow *big.Int) (*types.Transaction, error) {
	return _ResolutionModule.contract.Transact(opts, "setDisputeWindow", newWindow)
}

// SetDisputeWindow is a paid mutator transaction binding the contract method 0x332226d0.
//
// Solidity: function setDisputeWindow(uint256 newWindow) returns()
func (_ResolutionModule *ResolutionModuleSession) SetDisputeWindow(newWindow *big.Int) (*types.Transaction, error) {
	return _ResolutionModule.Contract.SetDisputeWindow(&_ResolutionModule.TransactOpts, newWindow)
}

// SetDisputeWindow is a paid mutator transaction binding the contract method 0x332226d0.
//
// Solidity: function setDisputeWindow(uint256 newWindow) returns()
func (_ResolutionModule *ResolutionModuleTransactorSession) SetDisputeWindow(newWindow *big.Int) (*types.Transaction, error) {
	return _ResolutionModule.Contract.SetDisputeWindow(&_ResolutionModule.TransactOpts, newWindow)
}

// SetMinBond is a paid mutator transaction binding the contract method 0x6eaae824.
//
// Solidity: function setMinBond(uint256 newMinBond) returns()
func (_ResolutionModule *ResolutionModuleTransactor) SetMinBond(opts *bind.TransactOpts, newMinBond *big.Int) (*types.Transaction, error) {
	return _ResolutionModule.contract.Transact(opts, "setMinBond", newMinBond)
}

// SetMinBond is a paid mutator transaction binding the contract method 0x6eaae824.
//
// Solidity: function setMinBond(uint256 newMinBond) returns()
func (_ResolutionModule *ResolutionModuleSession) SetMinBond(newMinBond *big.Int) (*types.Transaction, error) {
	return _ResolutionModule.Contract.SetMinBond(&_ResolutionModule.TransactOpts, newMinBond)
}

// SetMinBond is a paid mutator transaction binding the contract method 0x6eaae824.
//
// Solidity: function setMinBond(uint256 newMinBond) returns()
func (_ResolutionModule *ResolutionModuleTransactorSession) SetMinBond(newMinBond *big.Int) (*types.Transaction, error) {
	return _ResolutionModule.Contract.SetMinBond(&_ResolutionModule.TransactOpts, newMinBond)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ResolutionModule *ResolutionModuleTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _ResolutionModule.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ResolutionModule *ResolutionModuleSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ResolutionModule.Contract.TransferOwnership(&_ResolutionModule.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_ResolutionModule *ResolutionModuleTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _ResolutionModule.Contract.TransferOwnership(&_ResolutionModule.TransactOpts, newOwner)
}

// ResolutionModuleArbitratorUpdatedIterator is returned from FilterArbitratorUpdated and is used to iterate over the raw logs and unpacked data for ArbitratorUpdated events raised by the ResolutionModule contract.
type ResolutionModuleArbitratorUpdatedIterator struct {
	Event *ResolutionModuleArbitratorUpdated // Event containing the contract specifics and raw log

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
func (it *ResolutionModuleArbitratorUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolutionModuleArbitratorUpdated)
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
		it.Event = new(ResolutionModuleArbitratorUpdated)
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
func (it *ResolutionModuleArbitratorUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolutionModuleArbitratorUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolutionModuleArbitratorUpdated represents a ArbitratorUpdated event raised by the ResolutionModule contract.
type ResolutionModuleArbitratorUpdated struct {
	OldArbitrator common.Address
	NewArbitrator common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterArbitratorUpdated is a free log retrieval operation binding the contract event 0x0a4eb09e14cf0b6177b79042a04365517d9b189a50ec69eaabc40a6440103c90.
//
// Solidity: event ArbitratorUpdated(address indexed oldArbitrator, address indexed newArbitrator)
func (_ResolutionModule *ResolutionModuleFilterer) FilterArbitratorUpdated(opts *bind.FilterOpts, oldArbitrator []common.Address, newArbitrator []common.Address) (*ResolutionModuleArbitratorUpdatedIterator, error) {

	var oldArbitratorRule []interface{}
	for _, oldArbitratorItem := range oldArbitrator {
		oldArbitratorRule = append(oldArbitratorRule, oldArbitratorItem)
	}
	var newArbitratorRule []interface{}
	for _, newArbitratorItem := range newArbitrator {
		newArbitratorRule = append(newArbitratorRule, newArbitratorItem)
	}

	logs, sub, err := _ResolutionModule.contract.FilterLogs(opts, "ArbitratorUpdated", oldArbitratorRule, newArbitratorRule)
	if err != nil {
		return nil, err
	}
	return &ResolutionModuleArbitratorUpdatedIterator{contract: _ResolutionModule.contract, event: "ArbitratorUpdated", logs: logs, sub: sub}, nil
}

// WatchArbitratorUpdated is a free log subscription operation binding the contract event 0x0a4eb09e14cf0b6177b79042a04365517d9b189a50ec69eaabc40a6440103c90.
//
// Solidity: event ArbitratorUpdated(address indexed oldArbitrator, address indexed newArbitrator)
func (_ResolutionModule *ResolutionModuleFilterer) WatchArbitratorUpdated(opts *bind.WatchOpts, sink chan<- *ResolutionModuleArbitratorUpdated, oldArbitrator []common.Address, newArbitrator []common.Address) (event.Subscription, error) {

	var oldArbitratorRule []interface{}
	for _, oldArbitratorItem := range oldArbitrator {
		oldArbitratorRule = append(oldArbitratorRule, oldArbitratorItem)
	}
	var newArbitratorRule []interface{}
	for _, newArbitratorItem := range newArbitrator {
		newArbitratorRule = append(newArbitratorRule, newArbitratorItem)
	}

	logs, sub, err := _ResolutionModule.contract.WatchLogs(opts, "ArbitratorUpdated", oldArbitratorRule, newArbitratorRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolutionModuleArbitratorUpdated)
				if err := _ResolutionModule.contract.UnpackLog(event, "ArbitratorUpdated", log); err != nil {
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

// ParseArbitratorUpdated is a log parse operation binding the contract event 0x0a4eb09e14cf0b6177b79042a04365517d9b189a50ec69eaabc40a6440103c90.
//
// Solidity: event ArbitratorUpdated(address indexed oldArbitrator, address indexed newArbitrator)
func (_ResolutionModule *ResolutionModuleFilterer) ParseArbitratorUpdated(log types.Log) (*ResolutionModuleArbitratorUpdated, error) {
	event := new(ResolutionModuleArbitratorUpdated)
	if err := _ResolutionModule.contract.UnpackLog(event, "ArbitratorUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolutionModuleBondRefundedIterator is returned from FilterBondRefunded and is used to iterate over the raw logs and unpacked data for BondRefunded events raised by the ResolutionModule contract.
type ResolutionModuleBondRefundedIterator struct {
	Event *ResolutionModuleBondRefunded // Event containing the contract specifics and raw log

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
func (it *ResolutionModuleBondRefundedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolutionModuleBondRefunded)
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
		it.Event = new(ResolutionModuleBondRefunded)
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
func (it *ResolutionModuleBondRefundedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolutionModuleBondRefundedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolutionModuleBondRefunded represents a BondRefunded event raised by the ResolutionModule contract.
type ResolutionModuleBondRefunded struct {
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterBondRefunded is a free log retrieval operation binding the contract event 0x22060b72514e6ca1f5724d27d0d1668d9521aae214c78c1b9fa411e4401dddd1.
//
// Solidity: event BondRefunded(address indexed recipient, uint256 amount)
func (_ResolutionModule *ResolutionModuleFilterer) FilterBondRefunded(opts *bind.FilterOpts, recipient []common.Address) (*ResolutionModuleBondRefundedIterator, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ResolutionModule.contract.FilterLogs(opts, "BondRefunded", recipientRule)
	if err != nil {
		return nil, err
	}
	return &ResolutionModuleBondRefundedIterator{contract: _ResolutionModule.contract, event: "BondRefunded", logs: logs, sub: sub}, nil
}

// WatchBondRefunded is a free log subscription operation binding the contract event 0x22060b72514e6ca1f5724d27d0d1668d9521aae214c78c1b9fa411e4401dddd1.
//
// Solidity: event BondRefunded(address indexed recipient, uint256 amount)
func (_ResolutionModule *ResolutionModuleFilterer) WatchBondRefunded(opts *bind.WatchOpts, sink chan<- *ResolutionModuleBondRefunded, recipient []common.Address) (event.Subscription, error) {

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ResolutionModule.contract.WatchLogs(opts, "BondRefunded", recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolutionModuleBondRefunded)
				if err := _ResolutionModule.contract.UnpackLog(event, "BondRefunded", log); err != nil {
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

// ParseBondRefunded is a log parse operation binding the contract event 0x22060b72514e6ca1f5724d27d0d1668d9521aae214c78c1b9fa411e4401dddd1.
//
// Solidity: event BondRefunded(address indexed recipient, uint256 amount)
func (_ResolutionModule *ResolutionModuleFilterer) ParseBondRefunded(log types.Log) (*ResolutionModuleBondRefunded, error) {
	event := new(ResolutionModuleBondRefunded)
	if err := _ResolutionModule.contract.UnpackLog(event, "BondRefunded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolutionModuleBondSlashedIterator is returned from FilterBondSlashed and is used to iterate over the raw logs and unpacked data for BondSlashed events raised by the ResolutionModule contract.
type ResolutionModuleBondSlashedIterator struct {
	Event *ResolutionModuleBondSlashed // Event containing the contract specifics and raw log

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
func (it *ResolutionModuleBondSlashedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolutionModuleBondSlashed)
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
		it.Event = new(ResolutionModuleBondSlashed)
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
func (it *ResolutionModuleBondSlashedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolutionModuleBondSlashedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolutionModuleBondSlashed represents a BondSlashed event raised by the ResolutionModule contract.
type ResolutionModuleBondSlashed struct {
	SlashedAddress common.Address
	Amount         *big.Int
	Recipient      common.Address
	Raw            types.Log // Blockchain specific contextual infos
}

// FilterBondSlashed is a free log retrieval operation binding the contract event 0xf59795102dfe8996fe97ba1ee9bfbd46336e62c182e8d6556072c9288d3a6ae6.
//
// Solidity: event BondSlashed(address indexed slashedAddress, uint256 amount, address indexed recipient)
func (_ResolutionModule *ResolutionModuleFilterer) FilterBondSlashed(opts *bind.FilterOpts, slashedAddress []common.Address, recipient []common.Address) (*ResolutionModuleBondSlashedIterator, error) {

	var slashedAddressRule []interface{}
	for _, slashedAddressItem := range slashedAddress {
		slashedAddressRule = append(slashedAddressRule, slashedAddressItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ResolutionModule.contract.FilterLogs(opts, "BondSlashed", slashedAddressRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return &ResolutionModuleBondSlashedIterator{contract: _ResolutionModule.contract, event: "BondSlashed", logs: logs, sub: sub}, nil
}

// WatchBondSlashed is a free log subscription operation binding the contract event 0xf59795102dfe8996fe97ba1ee9bfbd46336e62c182e8d6556072c9288d3a6ae6.
//
// Solidity: event BondSlashed(address indexed slashedAddress, uint256 amount, address indexed recipient)
func (_ResolutionModule *ResolutionModuleFilterer) WatchBondSlashed(opts *bind.WatchOpts, sink chan<- *ResolutionModuleBondSlashed, slashedAddress []common.Address, recipient []common.Address) (event.Subscription, error) {

	var slashedAddressRule []interface{}
	for _, slashedAddressItem := range slashedAddress {
		slashedAddressRule = append(slashedAddressRule, slashedAddressItem)
	}

	var recipientRule []interface{}
	for _, recipientItem := range recipient {
		recipientRule = append(recipientRule, recipientItem)
	}

	logs, sub, err := _ResolutionModule.contract.WatchLogs(opts, "BondSlashed", slashedAddressRule, recipientRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolutionModuleBondSlashed)
				if err := _ResolutionModule.contract.UnpackLog(event, "BondSlashed", log); err != nil {
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

// ParseBondSlashed is a log parse operation binding the contract event 0xf59795102dfe8996fe97ba1ee9bfbd46336e62c182e8d6556072c9288d3a6ae6.
//
// Solidity: event BondSlashed(address indexed slashedAddress, uint256 amount, address indexed recipient)
func (_ResolutionModule *ResolutionModuleFilterer) ParseBondSlashed(log types.Log) (*ResolutionModuleBondSlashed, error) {
	event := new(ResolutionModuleBondSlashed)
	if err := _ResolutionModule.contract.UnpackLog(event, "BondSlashed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolutionModuleDisputeWindowUpdatedIterator is returned from FilterDisputeWindowUpdated and is used to iterate over the raw logs and unpacked data for DisputeWindowUpdated events raised by the ResolutionModule contract.
type ResolutionModuleDisputeWindowUpdatedIterator struct {
	Event *ResolutionModuleDisputeWindowUpdated // Event containing the contract specifics and raw log

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
func (it *ResolutionModuleDisputeWindowUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolutionModuleDisputeWindowUpdated)
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
		it.Event = new(ResolutionModuleDisputeWindowUpdated)
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
func (it *ResolutionModuleDisputeWindowUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolutionModuleDisputeWindowUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolutionModuleDisputeWindowUpdated represents a DisputeWindowUpdated event raised by the ResolutionModule contract.
type ResolutionModuleDisputeWindowUpdated struct {
	OldWindow *big.Int
	NewWindow *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterDisputeWindowUpdated is a free log retrieval operation binding the contract event 0x96e0f610d197acd2144db3bb338fe8cdc5544b8b21a40a14334068ac60e27172.
//
// Solidity: event DisputeWindowUpdated(uint256 oldWindow, uint256 newWindow)
func (_ResolutionModule *ResolutionModuleFilterer) FilterDisputeWindowUpdated(opts *bind.FilterOpts) (*ResolutionModuleDisputeWindowUpdatedIterator, error) {

	logs, sub, err := _ResolutionModule.contract.FilterLogs(opts, "DisputeWindowUpdated")
	if err != nil {
		return nil, err
	}
	return &ResolutionModuleDisputeWindowUpdatedIterator{contract: _ResolutionModule.contract, event: "DisputeWindowUpdated", logs: logs, sub: sub}, nil
}

// WatchDisputeWindowUpdated is a free log subscription operation binding the contract event 0x96e0f610d197acd2144db3bb338fe8cdc5544b8b21a40a14334068ac60e27172.
//
// Solidity: event DisputeWindowUpdated(uint256 oldWindow, uint256 newWindow)
func (_ResolutionModule *ResolutionModuleFilterer) WatchDisputeWindowUpdated(opts *bind.WatchOpts, sink chan<- *ResolutionModuleDisputeWindowUpdated) (event.Subscription, error) {

	logs, sub, err := _ResolutionModule.contract.WatchLogs(opts, "DisputeWindowUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolutionModuleDisputeWindowUpdated)
				if err := _ResolutionModule.contract.UnpackLog(event, "DisputeWindowUpdated", log); err != nil {
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

// ParseDisputeWindowUpdated is a log parse operation binding the contract event 0x96e0f610d197acd2144db3bb338fe8cdc5544b8b21a40a14334068ac60e27172.
//
// Solidity: event DisputeWindowUpdated(uint256 oldWindow, uint256 newWindow)
func (_ResolutionModule *ResolutionModuleFilterer) ParseDisputeWindowUpdated(log types.Log) (*ResolutionModuleDisputeWindowUpdated, error) {
	event := new(ResolutionModuleDisputeWindowUpdated)
	if err := _ResolutionModule.contract.UnpackLog(event, "DisputeWindowUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolutionModuleDisputedIterator is returned from FilterDisputed and is used to iterate over the raw logs and unpacked data for Disputed events raised by the ResolutionModule contract.
type ResolutionModuleDisputedIterator struct {
	Event *ResolutionModuleDisputed // Event containing the contract specifics and raw log

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
func (it *ResolutionModuleDisputedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolutionModuleDisputed)
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
		it.Event = new(ResolutionModuleDisputed)
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
func (it *ResolutionModuleDisputedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolutionModuleDisputedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolutionModuleDisputed represents a Disputed event raised by the ResolutionModule contract.
type ResolutionModuleDisputed struct {
	MarketId *big.Int
	Disputer common.Address
	Bond     *big.Int
	Reason   string
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterDisputed is a free log retrieval operation binding the contract event 0xee82e08028db05a07ca003413c29d0e6969a68d986ae687b57d586b39533bc2f.
//
// Solidity: event Disputed(uint256 indexed marketId, address indexed disputer, uint256 bond, string reason)
func (_ResolutionModule *ResolutionModuleFilterer) FilterDisputed(opts *bind.FilterOpts, marketId []*big.Int, disputer []common.Address) (*ResolutionModuleDisputedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var disputerRule []interface{}
	for _, disputerItem := range disputer {
		disputerRule = append(disputerRule, disputerItem)
	}

	logs, sub, err := _ResolutionModule.contract.FilterLogs(opts, "Disputed", marketIdRule, disputerRule)
	if err != nil {
		return nil, err
	}
	return &ResolutionModuleDisputedIterator{contract: _ResolutionModule.contract, event: "Disputed", logs: logs, sub: sub}, nil
}

// WatchDisputed is a free log subscription operation binding the contract event 0xee82e08028db05a07ca003413c29d0e6969a68d986ae687b57d586b39533bc2f.
//
// Solidity: event Disputed(uint256 indexed marketId, address indexed disputer, uint256 bond, string reason)
func (_ResolutionModule *ResolutionModuleFilterer) WatchDisputed(opts *bind.WatchOpts, sink chan<- *ResolutionModuleDisputed, marketId []*big.Int, disputer []common.Address) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var disputerRule []interface{}
	for _, disputerItem := range disputer {
		disputerRule = append(disputerRule, disputerItem)
	}

	logs, sub, err := _ResolutionModule.contract.WatchLogs(opts, "Disputed", marketIdRule, disputerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolutionModuleDisputed)
				if err := _ResolutionModule.contract.UnpackLog(event, "Disputed", log); err != nil {
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

// ParseDisputed is a log parse operation binding the contract event 0xee82e08028db05a07ca003413c29d0e6969a68d986ae687b57d586b39533bc2f.
//
// Solidity: event Disputed(uint256 indexed marketId, address indexed disputer, uint256 bond, string reason)
func (_ResolutionModule *ResolutionModuleFilterer) ParseDisputed(log types.Log) (*ResolutionModuleDisputed, error) {
	event := new(ResolutionModuleDisputed)
	if err := _ResolutionModule.contract.UnpackLog(event, "Disputed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolutionModuleFinalizedIterator is returned from FilterFinalized and is used to iterate over the raw logs and unpacked data for Finalized events raised by the ResolutionModule contract.
type ResolutionModuleFinalizedIterator struct {
	Event *ResolutionModuleFinalized // Event containing the contract specifics and raw log

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
func (it *ResolutionModuleFinalizedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolutionModuleFinalized)
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
		it.Event = new(ResolutionModuleFinalized)
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
func (it *ResolutionModuleFinalizedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolutionModuleFinalizedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolutionModuleFinalized represents a Finalized event raised by the ResolutionModule contract.
type ResolutionModuleFinalized struct {
	MarketId    *big.Int
	OutcomeId   *big.Int
	WasDisputed bool
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterFinalized is a free log retrieval operation binding the contract event 0x078425eac321e4770fa54ea86e364849803fe38ddda99fd34d7bacfeabcdc9a8.
//
// Solidity: event Finalized(uint256 indexed marketId, uint256 indexed outcomeId, bool wasDisputed)
func (_ResolutionModule *ResolutionModuleFilterer) FilterFinalized(opts *bind.FilterOpts, marketId []*big.Int, outcomeId []*big.Int) (*ResolutionModuleFinalizedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var outcomeIdRule []interface{}
	for _, outcomeIdItem := range outcomeId {
		outcomeIdRule = append(outcomeIdRule, outcomeIdItem)
	}

	logs, sub, err := _ResolutionModule.contract.FilterLogs(opts, "Finalized", marketIdRule, outcomeIdRule)
	if err != nil {
		return nil, err
	}
	return &ResolutionModuleFinalizedIterator{contract: _ResolutionModule.contract, event: "Finalized", logs: logs, sub: sub}, nil
}

// WatchFinalized is a free log subscription operation binding the contract event 0x078425eac321e4770fa54ea86e364849803fe38ddda99fd34d7bacfeabcdc9a8.
//
// Solidity: event Finalized(uint256 indexed marketId, uint256 indexed outcomeId, bool wasDisputed)
func (_ResolutionModule *ResolutionModuleFilterer) WatchFinalized(opts *bind.WatchOpts, sink chan<- *ResolutionModuleFinalized, marketId []*big.Int, outcomeId []*big.Int) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var outcomeIdRule []interface{}
	for _, outcomeIdItem := range outcomeId {
		outcomeIdRule = append(outcomeIdRule, outcomeIdItem)
	}

	logs, sub, err := _ResolutionModule.contract.WatchLogs(opts, "Finalized", marketIdRule, outcomeIdRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolutionModuleFinalized)
				if err := _ResolutionModule.contract.UnpackLog(event, "Finalized", log); err != nil {
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

// ParseFinalized is a log parse operation binding the contract event 0x078425eac321e4770fa54ea86e364849803fe38ddda99fd34d7bacfeabcdc9a8.
//
// Solidity: event Finalized(uint256 indexed marketId, uint256 indexed outcomeId, bool wasDisputed)
func (_ResolutionModule *ResolutionModuleFilterer) ParseFinalized(log types.Log) (*ResolutionModuleFinalized, error) {
	event := new(ResolutionModuleFinalized)
	if err := _ResolutionModule.contract.UnpackLog(event, "Finalized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolutionModuleMinBondUpdatedIterator is returned from FilterMinBondUpdated and is used to iterate over the raw logs and unpacked data for MinBondUpdated events raised by the ResolutionModule contract.
type ResolutionModuleMinBondUpdatedIterator struct {
	Event *ResolutionModuleMinBondUpdated // Event containing the contract specifics and raw log

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
func (it *ResolutionModuleMinBondUpdatedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolutionModuleMinBondUpdated)
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
		it.Event = new(ResolutionModuleMinBondUpdated)
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
func (it *ResolutionModuleMinBondUpdatedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolutionModuleMinBondUpdatedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolutionModuleMinBondUpdated represents a MinBondUpdated event raised by the ResolutionModule contract.
type ResolutionModuleMinBondUpdated struct {
	OldBond *big.Int
	NewBond *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterMinBondUpdated is a free log retrieval operation binding the contract event 0xbc501db9b822d1ea101cf77c85c0cb60f19f7bc67f95e272a27ac59b32e7db57.
//
// Solidity: event MinBondUpdated(uint256 oldBond, uint256 newBond)
func (_ResolutionModule *ResolutionModuleFilterer) FilterMinBondUpdated(opts *bind.FilterOpts) (*ResolutionModuleMinBondUpdatedIterator, error) {

	logs, sub, err := _ResolutionModule.contract.FilterLogs(opts, "MinBondUpdated")
	if err != nil {
		return nil, err
	}
	return &ResolutionModuleMinBondUpdatedIterator{contract: _ResolutionModule.contract, event: "MinBondUpdated", logs: logs, sub: sub}, nil
}

// WatchMinBondUpdated is a free log subscription operation binding the contract event 0xbc501db9b822d1ea101cf77c85c0cb60f19f7bc67f95e272a27ac59b32e7db57.
//
// Solidity: event MinBondUpdated(uint256 oldBond, uint256 newBond)
func (_ResolutionModule *ResolutionModuleFilterer) WatchMinBondUpdated(opts *bind.WatchOpts, sink chan<- *ResolutionModuleMinBondUpdated) (event.Subscription, error) {

	logs, sub, err := _ResolutionModule.contract.WatchLogs(opts, "MinBondUpdated")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolutionModuleMinBondUpdated)
				if err := _ResolutionModule.contract.UnpackLog(event, "MinBondUpdated", log); err != nil {
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

// ParseMinBondUpdated is a log parse operation binding the contract event 0xbc501db9b822d1ea101cf77c85c0cb60f19f7bc67f95e272a27ac59b32e7db57.
//
// Solidity: event MinBondUpdated(uint256 oldBond, uint256 newBond)
func (_ResolutionModule *ResolutionModuleFilterer) ParseMinBondUpdated(log types.Log) (*ResolutionModuleMinBondUpdated, error) {
	event := new(ResolutionModuleMinBondUpdated)
	if err := _ResolutionModule.contract.UnpackLog(event, "MinBondUpdated", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolutionModuleOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the ResolutionModule contract.
type ResolutionModuleOwnershipTransferredIterator struct {
	Event *ResolutionModuleOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *ResolutionModuleOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolutionModuleOwnershipTransferred)
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
		it.Event = new(ResolutionModuleOwnershipTransferred)
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
func (it *ResolutionModuleOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolutionModuleOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolutionModuleOwnershipTransferred represents a OwnershipTransferred event raised by the ResolutionModule contract.
type ResolutionModuleOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ResolutionModule *ResolutionModuleFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*ResolutionModuleOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ResolutionModule.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &ResolutionModuleOwnershipTransferredIterator{contract: _ResolutionModule.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ResolutionModule *ResolutionModuleFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *ResolutionModuleOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _ResolutionModule.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolutionModuleOwnershipTransferred)
				if err := _ResolutionModule.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// ParseOwnershipTransferred is a log parse operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_ResolutionModule *ResolutionModuleFilterer) ParseOwnershipTransferred(log types.Log) (*ResolutionModuleOwnershipTransferred, error) {
	event := new(ResolutionModuleOwnershipTransferred)
	if err := _ResolutionModule.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ResolutionModuleResolutionProposedIterator is returned from FilterResolutionProposed and is used to iterate over the raw logs and unpacked data for ResolutionProposed events raised by the ResolutionModule contract.
type ResolutionModuleResolutionProposedIterator struct {
	Event *ResolutionModuleResolutionProposed // Event containing the contract specifics and raw log

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
func (it *ResolutionModuleResolutionProposedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ResolutionModuleResolutionProposed)
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
		it.Event = new(ResolutionModuleResolutionProposed)
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
func (it *ResolutionModuleResolutionProposedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ResolutionModuleResolutionProposedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ResolutionModuleResolutionProposed represents a ResolutionProposed event raised by the ResolutionModule contract.
type ResolutionModuleResolutionProposed struct {
	MarketId    *big.Int
	OutcomeId   *big.Int
	Proposer    common.Address
	Bond        *big.Int
	EvidenceURI string
	Deadline    *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterResolutionProposed is a free log retrieval operation binding the contract event 0x7104ccc99e3411ea5f9127e33ba473aedac404e9984fc978a05efaa2da05f46b.
//
// Solidity: event ResolutionProposed(uint256 indexed marketId, uint256 indexed outcomeId, address indexed proposer, uint256 bond, string evidenceURI, uint256 deadline)
func (_ResolutionModule *ResolutionModuleFilterer) FilterResolutionProposed(opts *bind.FilterOpts, marketId []*big.Int, outcomeId []*big.Int, proposer []common.Address) (*ResolutionModuleResolutionProposedIterator, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var outcomeIdRule []interface{}
	for _, outcomeIdItem := range outcomeId {
		outcomeIdRule = append(outcomeIdRule, outcomeIdItem)
	}
	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}

	logs, sub, err := _ResolutionModule.contract.FilterLogs(opts, "ResolutionProposed", marketIdRule, outcomeIdRule, proposerRule)
	if err != nil {
		return nil, err
	}
	return &ResolutionModuleResolutionProposedIterator{contract: _ResolutionModule.contract, event: "ResolutionProposed", logs: logs, sub: sub}, nil
}

// WatchResolutionProposed is a free log subscription operation binding the contract event 0x7104ccc99e3411ea5f9127e33ba473aedac404e9984fc978a05efaa2da05f46b.
//
// Solidity: event ResolutionProposed(uint256 indexed marketId, uint256 indexed outcomeId, address indexed proposer, uint256 bond, string evidenceURI, uint256 deadline)
func (_ResolutionModule *ResolutionModuleFilterer) WatchResolutionProposed(opts *bind.WatchOpts, sink chan<- *ResolutionModuleResolutionProposed, marketId []*big.Int, outcomeId []*big.Int, proposer []common.Address) (event.Subscription, error) {

	var marketIdRule []interface{}
	for _, marketIdItem := range marketId {
		marketIdRule = append(marketIdRule, marketIdItem)
	}
	var outcomeIdRule []interface{}
	for _, outcomeIdItem := range outcomeId {
		outcomeIdRule = append(outcomeIdRule, outcomeIdItem)
	}
	var proposerRule []interface{}
	for _, proposerItem := range proposer {
		proposerRule = append(proposerRule, proposerItem)
	}

	logs, sub, err := _ResolutionModule.contract.WatchLogs(opts, "ResolutionProposed", marketIdRule, outcomeIdRule, proposerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ResolutionModuleResolutionProposed)
				if err := _ResolutionModule.contract.UnpackLog(event, "ResolutionProposed", log); err != nil {
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

// ParseResolutionProposed is a log parse operation binding the contract event 0x7104ccc99e3411ea5f9127e33ba473aedac404e9984fc978a05efaa2da05f46b.
//
// Solidity: event ResolutionProposed(uint256 indexed marketId, uint256 indexed outcomeId, address indexed proposer, uint256 bond, string evidenceURI, uint256 deadline)
func (_ResolutionModule *ResolutionModuleFilterer) ParseResolutionProposed(log types.Log) (*ResolutionModuleResolutionProposed, error) {
	event := new(ResolutionModuleResolutionProposed)
	if err := _ResolutionModule.contract.UnpackLog(event, "ResolutionProposed", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
