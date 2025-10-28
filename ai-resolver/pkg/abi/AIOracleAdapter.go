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

// AIOracleAdapterProposedOutcome is an auto generated low-level Go binding around an user-defined struct.
type AIOracleAdapterProposedOutcome struct {
	MarketId     *big.Int
	OutcomeId    *big.Int
	CloseTime    *big.Int
	EvidenceHash [32]byte
	NotBefore    *big.Int
	Deadline     *big.Int
}

// AIOracleAdapterMetaData contains all meta data concerning the AIOracleAdapter contract.
var AIOracleAdapterMetaData = &bind.MetaData{
	ABI: "[{\"type\":\"constructor\",\"inputs\":[{\"name\":\"_resolutionModule\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_bondToken\",\"type\":\"address\",\"internalType\":\"address\"},{\"name\":\"_initialSigner\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"DOMAIN_SEPARATOR\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"EIP712_DOMAIN_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"PROPOSED_OUTCOME_TYPEHASH\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"addSigner\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"allowedSigners\",\"inputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"bondToken\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractIERC20\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"getProposalHash\",\"inputs\":[{\"name\":\"proposal\",\"type\":\"tuple\",\"internalType\":\"structAIOracleAdapter.ProposedOutcome\",\"components\":[{\"name\":\"marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outcomeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"closeTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"evidenceHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"notBefore\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"hashEvidence\",\"inputs\":[{\"name\":\"evidenceURIs\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"stateMutability\":\"pure\"},{\"type\":\"function\",\"name\":\"isSignatureUsed\",\"inputs\":[{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"owner\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"address\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"proposeAI\",\"inputs\":[{\"name\":\"proposal\",\"type\":\"tuple\",\"internalType\":\"structAIOracleAdapter.ProposedOutcome\",\"components\":[{\"name\":\"marketId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"outcomeId\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"closeTime\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"evidenceHash\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"},{\"name\":\"notBefore\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"deadline\",\"type\":\"uint256\",\"internalType\":\"uint256\"}]},{\"name\":\"signature\",\"type\":\"bytes\",\"internalType\":\"bytes\"},{\"name\":\"bondAmount\",\"type\":\"uint256\",\"internalType\":\"uint256\"},{\"name\":\"evidenceURIs\",\"type\":\"string[]\",\"internalType\":\"string[]\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"removeSigner\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"renounceOwnership\",\"inputs\":[],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"resolutionModule\",\"inputs\":[],\"outputs\":[{\"name\":\"\",\"type\":\"address\",\"internalType\":\"contractResolutionModule\"}],\"stateMutability\":\"view\"},{\"type\":\"function\",\"name\":\"transferOwnership\",\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\",\"internalType\":\"address\"}],\"outputs\":[],\"stateMutability\":\"nonpayable\"},{\"type\":\"function\",\"name\":\"usedSignatures\",\"inputs\":[{\"name\":\"\",\"type\":\"bytes32\",\"internalType\":\"bytes32\"}],\"outputs\":[{\"name\":\"\",\"type\":\"bool\",\"internalType\":\"bool\"}],\"stateMutability\":\"view\"},{\"type\":\"event\",\"name\":\"AIProposalSubmitted\",\"inputs\":[{\"name\":\"marketId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"outcomeId\",\"type\":\"uint256\",\"indexed\":true,\"internalType\":\"uint256\"},{\"name\":\"proposer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"aiSigner\",\"type\":\"address\",\"indexed\":false,\"internalType\":\"address\"},{\"name\":\"bondAmount\",\"type\":\"uint256\",\"indexed\":false,\"internalType\":\"uint256\"},{\"name\":\"signatureHash\",\"type\":\"bytes32\",\"indexed\":false,\"internalType\":\"bytes32\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"OwnershipTransferred\",\"inputs\":[{\"name\":\"previousOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"},{\"name\":\"newOwner\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SignerAdded\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"event\",\"name\":\"SignerRemoved\",\"inputs\":[{\"name\":\"signer\",\"type\":\"address\",\"indexed\":true,\"internalType\":\"address\"}],\"anonymous\":false},{\"type\":\"error\",\"name\":\"InvalidAddress\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidEvidenceHash\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSignature\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"InvalidSigner\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"MarketNotClosed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"OwnableInvalidOwner\",\"inputs\":[{\"name\":\"owner\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"OwnableUnauthorizedAccount\",\"inputs\":[{\"name\":\"account\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"ReentrancyGuardReentrantCall\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SafeERC20FailedOperation\",\"inputs\":[{\"name\":\"token\",\"type\":\"address\",\"internalType\":\"address\"}]},{\"type\":\"error\",\"name\":\"SignatureAlreadyUsed\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignatureExpired\",\"inputs\":[]},{\"type\":\"error\",\"name\":\"SignatureNotYetValid\",\"inputs\":[]}]",
}

// AIOracleAdapterABI is the input ABI used to generate the binding from.
// Deprecated: Use AIOracleAdapterMetaData.ABI instead.
var AIOracleAdapterABI = AIOracleAdapterMetaData.ABI

// AIOracleAdapter is an auto generated Go binding around an Ethereum contract.
type AIOracleAdapter struct {
	AIOracleAdapterCaller     // Read-only binding to the contract
	AIOracleAdapterTransactor // Write-only binding to the contract
	AIOracleAdapterFilterer   // Log filterer for contract events
}

// AIOracleAdapterCaller is an auto generated read-only Go binding around an Ethereum contract.
type AIOracleAdapterCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AIOracleAdapterTransactor is an auto generated write-only Go binding around an Ethereum contract.
type AIOracleAdapterTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AIOracleAdapterFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type AIOracleAdapterFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// AIOracleAdapterSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type AIOracleAdapterSession struct {
	Contract     *AIOracleAdapter  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// AIOracleAdapterCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type AIOracleAdapterCallerSession struct {
	Contract *AIOracleAdapterCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// AIOracleAdapterTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type AIOracleAdapterTransactorSession struct {
	Contract     *AIOracleAdapterTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// AIOracleAdapterRaw is an auto generated low-level Go binding around an Ethereum contract.
type AIOracleAdapterRaw struct {
	Contract *AIOracleAdapter // Generic contract binding to access the raw methods on
}

// AIOracleAdapterCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type AIOracleAdapterCallerRaw struct {
	Contract *AIOracleAdapterCaller // Generic read-only contract binding to access the raw methods on
}

// AIOracleAdapterTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type AIOracleAdapterTransactorRaw struct {
	Contract *AIOracleAdapterTransactor // Generic write-only contract binding to access the raw methods on
}

// NewAIOracleAdapter creates a new instance of AIOracleAdapter, bound to a specific deployed contract.
func NewAIOracleAdapter(address common.Address, backend bind.ContractBackend) (*AIOracleAdapter, error) {
	contract, err := bindAIOracleAdapter(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &AIOracleAdapter{AIOracleAdapterCaller: AIOracleAdapterCaller{contract: contract}, AIOracleAdapterTransactor: AIOracleAdapterTransactor{contract: contract}, AIOracleAdapterFilterer: AIOracleAdapterFilterer{contract: contract}}, nil
}

// NewAIOracleAdapterCaller creates a new read-only instance of AIOracleAdapter, bound to a specific deployed contract.
func NewAIOracleAdapterCaller(address common.Address, caller bind.ContractCaller) (*AIOracleAdapterCaller, error) {
	contract, err := bindAIOracleAdapter(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &AIOracleAdapterCaller{contract: contract}, nil
}

// NewAIOracleAdapterTransactor creates a new write-only instance of AIOracleAdapter, bound to a specific deployed contract.
func NewAIOracleAdapterTransactor(address common.Address, transactor bind.ContractTransactor) (*AIOracleAdapterTransactor, error) {
	contract, err := bindAIOracleAdapter(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &AIOracleAdapterTransactor{contract: contract}, nil
}

// NewAIOracleAdapterFilterer creates a new log filterer instance of AIOracleAdapter, bound to a specific deployed contract.
func NewAIOracleAdapterFilterer(address common.Address, filterer bind.ContractFilterer) (*AIOracleAdapterFilterer, error) {
	contract, err := bindAIOracleAdapter(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &AIOracleAdapterFilterer{contract: contract}, nil
}

// bindAIOracleAdapter binds a generic wrapper to an already deployed contract.
func bindAIOracleAdapter(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := AIOracleAdapterMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AIOracleAdapter *AIOracleAdapterRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AIOracleAdapter.Contract.AIOracleAdapterCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AIOracleAdapter *AIOracleAdapterRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AIOracleAdapter.Contract.AIOracleAdapterTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AIOracleAdapter *AIOracleAdapterRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AIOracleAdapter.Contract.AIOracleAdapterTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_AIOracleAdapter *AIOracleAdapterCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _AIOracleAdapter.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_AIOracleAdapter *AIOracleAdapterTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AIOracleAdapter.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_AIOracleAdapter *AIOracleAdapterTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _AIOracleAdapter.Contract.contract.Transact(opts, method, params...)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_AIOracleAdapter *AIOracleAdapterCaller) DOMAINSEPARATOR(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AIOracleAdapter.contract.Call(opts, &out, "DOMAIN_SEPARATOR")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_AIOracleAdapter *AIOracleAdapterSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _AIOracleAdapter.Contract.DOMAINSEPARATOR(&_AIOracleAdapter.CallOpts)
}

// DOMAINSEPARATOR is a free data retrieval call binding the contract method 0x3644e515.
//
// Solidity: function DOMAIN_SEPARATOR() view returns(bytes32)
func (_AIOracleAdapter *AIOracleAdapterCallerSession) DOMAINSEPARATOR() ([32]byte, error) {
	return _AIOracleAdapter.Contract.DOMAINSEPARATOR(&_AIOracleAdapter.CallOpts)
}

// EIP712DOMAINTYPEHASH is a free data retrieval call binding the contract method 0xc7977be7.
//
// Solidity: function EIP712_DOMAIN_TYPEHASH() view returns(bytes32)
func (_AIOracleAdapter *AIOracleAdapterCaller) EIP712DOMAINTYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AIOracleAdapter.contract.Call(opts, &out, "EIP712_DOMAIN_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// EIP712DOMAINTYPEHASH is a free data retrieval call binding the contract method 0xc7977be7.
//
// Solidity: function EIP712_DOMAIN_TYPEHASH() view returns(bytes32)
func (_AIOracleAdapter *AIOracleAdapterSession) EIP712DOMAINTYPEHASH() ([32]byte, error) {
	return _AIOracleAdapter.Contract.EIP712DOMAINTYPEHASH(&_AIOracleAdapter.CallOpts)
}

// EIP712DOMAINTYPEHASH is a free data retrieval call binding the contract method 0xc7977be7.
//
// Solidity: function EIP712_DOMAIN_TYPEHASH() view returns(bytes32)
func (_AIOracleAdapter *AIOracleAdapterCallerSession) EIP712DOMAINTYPEHASH() ([32]byte, error) {
	return _AIOracleAdapter.Contract.EIP712DOMAINTYPEHASH(&_AIOracleAdapter.CallOpts)
}

// PROPOSEDOUTCOMETYPEHASH is a free data retrieval call binding the contract method 0x5e453009.
//
// Solidity: function PROPOSED_OUTCOME_TYPEHASH() view returns(bytes32)
func (_AIOracleAdapter *AIOracleAdapterCaller) PROPOSEDOUTCOMETYPEHASH(opts *bind.CallOpts) ([32]byte, error) {
	var out []interface{}
	err := _AIOracleAdapter.contract.Call(opts, &out, "PROPOSED_OUTCOME_TYPEHASH")

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// PROPOSEDOUTCOMETYPEHASH is a free data retrieval call binding the contract method 0x5e453009.
//
// Solidity: function PROPOSED_OUTCOME_TYPEHASH() view returns(bytes32)
func (_AIOracleAdapter *AIOracleAdapterSession) PROPOSEDOUTCOMETYPEHASH() ([32]byte, error) {
	return _AIOracleAdapter.Contract.PROPOSEDOUTCOMETYPEHASH(&_AIOracleAdapter.CallOpts)
}

// PROPOSEDOUTCOMETYPEHASH is a free data retrieval call binding the contract method 0x5e453009.
//
// Solidity: function PROPOSED_OUTCOME_TYPEHASH() view returns(bytes32)
func (_AIOracleAdapter *AIOracleAdapterCallerSession) PROPOSEDOUTCOMETYPEHASH() ([32]byte, error) {
	return _AIOracleAdapter.Contract.PROPOSEDOUTCOMETYPEHASH(&_AIOracleAdapter.CallOpts)
}

// AllowedSigners is a free data retrieval call binding the contract method 0xf8b4d864.
//
// Solidity: function allowedSigners(address ) view returns(bool)
func (_AIOracleAdapter *AIOracleAdapterCaller) AllowedSigners(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _AIOracleAdapter.contract.Call(opts, &out, "allowedSigners", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// AllowedSigners is a free data retrieval call binding the contract method 0xf8b4d864.
//
// Solidity: function allowedSigners(address ) view returns(bool)
func (_AIOracleAdapter *AIOracleAdapterSession) AllowedSigners(arg0 common.Address) (bool, error) {
	return _AIOracleAdapter.Contract.AllowedSigners(&_AIOracleAdapter.CallOpts, arg0)
}

// AllowedSigners is a free data retrieval call binding the contract method 0xf8b4d864.
//
// Solidity: function allowedSigners(address ) view returns(bool)
func (_AIOracleAdapter *AIOracleAdapterCallerSession) AllowedSigners(arg0 common.Address) (bool, error) {
	return _AIOracleAdapter.Contract.AllowedSigners(&_AIOracleAdapter.CallOpts, arg0)
}

// BondToken is a free data retrieval call binding the contract method 0xc28f4392.
//
// Solidity: function bondToken() view returns(address)
func (_AIOracleAdapter *AIOracleAdapterCaller) BondToken(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AIOracleAdapter.contract.Call(opts, &out, "bondToken")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// BondToken is a free data retrieval call binding the contract method 0xc28f4392.
//
// Solidity: function bondToken() view returns(address)
func (_AIOracleAdapter *AIOracleAdapterSession) BondToken() (common.Address, error) {
	return _AIOracleAdapter.Contract.BondToken(&_AIOracleAdapter.CallOpts)
}

// BondToken is a free data retrieval call binding the contract method 0xc28f4392.
//
// Solidity: function bondToken() view returns(address)
func (_AIOracleAdapter *AIOracleAdapterCallerSession) BondToken() (common.Address, error) {
	return _AIOracleAdapter.Contract.BondToken(&_AIOracleAdapter.CallOpts)
}

// GetProposalHash is a free data retrieval call binding the contract method 0x79983ba0.
//
// Solidity: function getProposalHash((uint256,uint256,uint256,bytes32,uint256,uint256) proposal) view returns(bytes32)
func (_AIOracleAdapter *AIOracleAdapterCaller) GetProposalHash(opts *bind.CallOpts, proposal AIOracleAdapterProposedOutcome) ([32]byte, error) {
	var out []interface{}
	err := _AIOracleAdapter.contract.Call(opts, &out, "getProposalHash", proposal)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// GetProposalHash is a free data retrieval call binding the contract method 0x79983ba0.
//
// Solidity: function getProposalHash((uint256,uint256,uint256,bytes32,uint256,uint256) proposal) view returns(bytes32)
func (_AIOracleAdapter *AIOracleAdapterSession) GetProposalHash(proposal AIOracleAdapterProposedOutcome) ([32]byte, error) {
	return _AIOracleAdapter.Contract.GetProposalHash(&_AIOracleAdapter.CallOpts, proposal)
}

// GetProposalHash is a free data retrieval call binding the contract method 0x79983ba0.
//
// Solidity: function getProposalHash((uint256,uint256,uint256,bytes32,uint256,uint256) proposal) view returns(bytes32)
func (_AIOracleAdapter *AIOracleAdapterCallerSession) GetProposalHash(proposal AIOracleAdapterProposedOutcome) ([32]byte, error) {
	return _AIOracleAdapter.Contract.GetProposalHash(&_AIOracleAdapter.CallOpts, proposal)
}

// HashEvidence is a free data retrieval call binding the contract method 0x10ff81f0.
//
// Solidity: function hashEvidence(string[] evidenceURIs) pure returns(bytes32)
func (_AIOracleAdapter *AIOracleAdapterCaller) HashEvidence(opts *bind.CallOpts, evidenceURIs []string) ([32]byte, error) {
	var out []interface{}
	err := _AIOracleAdapter.contract.Call(opts, &out, "hashEvidence", evidenceURIs)

	if err != nil {
		return *new([32]byte), err
	}

	out0 := *abi.ConvertType(out[0], new([32]byte)).(*[32]byte)

	return out0, err

}

// HashEvidence is a free data retrieval call binding the contract method 0x10ff81f0.
//
// Solidity: function hashEvidence(string[] evidenceURIs) pure returns(bytes32)
func (_AIOracleAdapter *AIOracleAdapterSession) HashEvidence(evidenceURIs []string) ([32]byte, error) {
	return _AIOracleAdapter.Contract.HashEvidence(&_AIOracleAdapter.CallOpts, evidenceURIs)
}

// HashEvidence is a free data retrieval call binding the contract method 0x10ff81f0.
//
// Solidity: function hashEvidence(string[] evidenceURIs) pure returns(bytes32)
func (_AIOracleAdapter *AIOracleAdapterCallerSession) HashEvidence(evidenceURIs []string) ([32]byte, error) {
	return _AIOracleAdapter.Contract.HashEvidence(&_AIOracleAdapter.CallOpts, evidenceURIs)
}

// IsSignatureUsed is a free data retrieval call binding the contract method 0x1150f0f3.
//
// Solidity: function isSignatureUsed(bytes signature) view returns(bool)
func (_AIOracleAdapter *AIOracleAdapterCaller) IsSignatureUsed(opts *bind.CallOpts, signature []byte) (bool, error) {
	var out []interface{}
	err := _AIOracleAdapter.contract.Call(opts, &out, "isSignatureUsed", signature)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsSignatureUsed is a free data retrieval call binding the contract method 0x1150f0f3.
//
// Solidity: function isSignatureUsed(bytes signature) view returns(bool)
func (_AIOracleAdapter *AIOracleAdapterSession) IsSignatureUsed(signature []byte) (bool, error) {
	return _AIOracleAdapter.Contract.IsSignatureUsed(&_AIOracleAdapter.CallOpts, signature)
}

// IsSignatureUsed is a free data retrieval call binding the contract method 0x1150f0f3.
//
// Solidity: function isSignatureUsed(bytes signature) view returns(bool)
func (_AIOracleAdapter *AIOracleAdapterCallerSession) IsSignatureUsed(signature []byte) (bool, error) {
	return _AIOracleAdapter.Contract.IsSignatureUsed(&_AIOracleAdapter.CallOpts, signature)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AIOracleAdapter *AIOracleAdapterCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AIOracleAdapter.contract.Call(opts, &out, "owner")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AIOracleAdapter *AIOracleAdapterSession) Owner() (common.Address, error) {
	return _AIOracleAdapter.Contract.Owner(&_AIOracleAdapter.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() view returns(address)
func (_AIOracleAdapter *AIOracleAdapterCallerSession) Owner() (common.Address, error) {
	return _AIOracleAdapter.Contract.Owner(&_AIOracleAdapter.CallOpts)
}

// ResolutionModule is a free data retrieval call binding the contract method 0xfd38ca0f.
//
// Solidity: function resolutionModule() view returns(address)
func (_AIOracleAdapter *AIOracleAdapterCaller) ResolutionModule(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _AIOracleAdapter.contract.Call(opts, &out, "resolutionModule")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// ResolutionModule is a free data retrieval call binding the contract method 0xfd38ca0f.
//
// Solidity: function resolutionModule() view returns(address)
func (_AIOracleAdapter *AIOracleAdapterSession) ResolutionModule() (common.Address, error) {
	return _AIOracleAdapter.Contract.ResolutionModule(&_AIOracleAdapter.CallOpts)
}

// ResolutionModule is a free data retrieval call binding the contract method 0xfd38ca0f.
//
// Solidity: function resolutionModule() view returns(address)
func (_AIOracleAdapter *AIOracleAdapterCallerSession) ResolutionModule() (common.Address, error) {
	return _AIOracleAdapter.Contract.ResolutionModule(&_AIOracleAdapter.CallOpts)
}

// UsedSignatures is a free data retrieval call binding the contract method 0xf978fd61.
//
// Solidity: function usedSignatures(bytes32 ) view returns(bool)
func (_AIOracleAdapter *AIOracleAdapterCaller) UsedSignatures(opts *bind.CallOpts, arg0 [32]byte) (bool, error) {
	var out []interface{}
	err := _AIOracleAdapter.contract.Call(opts, &out, "usedSignatures", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// UsedSignatures is a free data retrieval call binding the contract method 0xf978fd61.
//
// Solidity: function usedSignatures(bytes32 ) view returns(bool)
func (_AIOracleAdapter *AIOracleAdapterSession) UsedSignatures(arg0 [32]byte) (bool, error) {
	return _AIOracleAdapter.Contract.UsedSignatures(&_AIOracleAdapter.CallOpts, arg0)
}

// UsedSignatures is a free data retrieval call binding the contract method 0xf978fd61.
//
// Solidity: function usedSignatures(bytes32 ) view returns(bool)
func (_AIOracleAdapter *AIOracleAdapterCallerSession) UsedSignatures(arg0 [32]byte) (bool, error) {
	return _AIOracleAdapter.Contract.UsedSignatures(&_AIOracleAdapter.CallOpts, arg0)
}

// AddSigner is a paid mutator transaction binding the contract method 0xeb12d61e.
//
// Solidity: function addSigner(address signer) returns()
func (_AIOracleAdapter *AIOracleAdapterTransactor) AddSigner(opts *bind.TransactOpts, signer common.Address) (*types.Transaction, error) {
	return _AIOracleAdapter.contract.Transact(opts, "addSigner", signer)
}

// AddSigner is a paid mutator transaction binding the contract method 0xeb12d61e.
//
// Solidity: function addSigner(address signer) returns()
func (_AIOracleAdapter *AIOracleAdapterSession) AddSigner(signer common.Address) (*types.Transaction, error) {
	return _AIOracleAdapter.Contract.AddSigner(&_AIOracleAdapter.TransactOpts, signer)
}

// AddSigner is a paid mutator transaction binding the contract method 0xeb12d61e.
//
// Solidity: function addSigner(address signer) returns()
func (_AIOracleAdapter *AIOracleAdapterTransactorSession) AddSigner(signer common.Address) (*types.Transaction, error) {
	return _AIOracleAdapter.Contract.AddSigner(&_AIOracleAdapter.TransactOpts, signer)
}

// ProposeAI is a paid mutator transaction binding the contract method 0x3791e646.
//
// Solidity: function proposeAI((uint256,uint256,uint256,bytes32,uint256,uint256) proposal, bytes signature, uint256 bondAmount, string[] evidenceURIs) returns()
func (_AIOracleAdapter *AIOracleAdapterTransactor) ProposeAI(opts *bind.TransactOpts, proposal AIOracleAdapterProposedOutcome, signature []byte, bondAmount *big.Int, evidenceURIs []string) (*types.Transaction, error) {
	return _AIOracleAdapter.contract.Transact(opts, "proposeAI", proposal, signature, bondAmount, evidenceURIs)
}

// ProposeAI is a paid mutator transaction binding the contract method 0x3791e646.
//
// Solidity: function proposeAI((uint256,uint256,uint256,bytes32,uint256,uint256) proposal, bytes signature, uint256 bondAmount, string[] evidenceURIs) returns()
func (_AIOracleAdapter *AIOracleAdapterSession) ProposeAI(proposal AIOracleAdapterProposedOutcome, signature []byte, bondAmount *big.Int, evidenceURIs []string) (*types.Transaction, error) {
	return _AIOracleAdapter.Contract.ProposeAI(&_AIOracleAdapter.TransactOpts, proposal, signature, bondAmount, evidenceURIs)
}

// ProposeAI is a paid mutator transaction binding the contract method 0x3791e646.
//
// Solidity: function proposeAI((uint256,uint256,uint256,bytes32,uint256,uint256) proposal, bytes signature, uint256 bondAmount, string[] evidenceURIs) returns()
func (_AIOracleAdapter *AIOracleAdapterTransactorSession) ProposeAI(proposal AIOracleAdapterProposedOutcome, signature []byte, bondAmount *big.Int, evidenceURIs []string) (*types.Transaction, error) {
	return _AIOracleAdapter.Contract.ProposeAI(&_AIOracleAdapter.TransactOpts, proposal, signature, bondAmount, evidenceURIs)
}

// RemoveSigner is a paid mutator transaction binding the contract method 0x0e316ab7.
//
// Solidity: function removeSigner(address signer) returns()
func (_AIOracleAdapter *AIOracleAdapterTransactor) RemoveSigner(opts *bind.TransactOpts, signer common.Address) (*types.Transaction, error) {
	return _AIOracleAdapter.contract.Transact(opts, "removeSigner", signer)
}

// RemoveSigner is a paid mutator transaction binding the contract method 0x0e316ab7.
//
// Solidity: function removeSigner(address signer) returns()
func (_AIOracleAdapter *AIOracleAdapterSession) RemoveSigner(signer common.Address) (*types.Transaction, error) {
	return _AIOracleAdapter.Contract.RemoveSigner(&_AIOracleAdapter.TransactOpts, signer)
}

// RemoveSigner is a paid mutator transaction binding the contract method 0x0e316ab7.
//
// Solidity: function removeSigner(address signer) returns()
func (_AIOracleAdapter *AIOracleAdapterTransactorSession) RemoveSigner(signer common.Address) (*types.Transaction, error) {
	return _AIOracleAdapter.Contract.RemoveSigner(&_AIOracleAdapter.TransactOpts, signer)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AIOracleAdapter *AIOracleAdapterTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _AIOracleAdapter.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AIOracleAdapter *AIOracleAdapterSession) RenounceOwnership() (*types.Transaction, error) {
	return _AIOracleAdapter.Contract.RenounceOwnership(&_AIOracleAdapter.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_AIOracleAdapter *AIOracleAdapterTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _AIOracleAdapter.Contract.RenounceOwnership(&_AIOracleAdapter.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AIOracleAdapter *AIOracleAdapterTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _AIOracleAdapter.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AIOracleAdapter *AIOracleAdapterSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AIOracleAdapter.Contract.TransferOwnership(&_AIOracleAdapter.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address newOwner) returns()
func (_AIOracleAdapter *AIOracleAdapterTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _AIOracleAdapter.Contract.TransferOwnership(&_AIOracleAdapter.TransactOpts, newOwner)
}

// AIOracleAdapterAIProposalSubmittedIterator is returned from FilterAIProposalSubmitted and is used to iterate over the raw logs and unpacked data for AIProposalSubmitted events raised by the AIOracleAdapter contract.
type AIOracleAdapterAIProposalSubmittedIterator struct {
	Event *AIOracleAdapterAIProposalSubmitted // Event containing the contract specifics and raw log

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
func (it *AIOracleAdapterAIProposalSubmittedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AIOracleAdapterAIProposalSubmitted)
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
		it.Event = new(AIOracleAdapterAIProposalSubmitted)
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
func (it *AIOracleAdapterAIProposalSubmittedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AIOracleAdapterAIProposalSubmittedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AIOracleAdapterAIProposalSubmitted represents a AIProposalSubmitted event raised by the AIOracleAdapter contract.
type AIOracleAdapterAIProposalSubmitted struct {
	MarketId      *big.Int
	OutcomeId     *big.Int
	Proposer      common.Address
	AiSigner      common.Address
	BondAmount    *big.Int
	SignatureHash [32]byte
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterAIProposalSubmitted is a free log retrieval operation binding the contract event 0x0b1a7dc1c0ac67b57ba259d0fbc6e4633797b01c17a70f944e6b615227ad0cb0.
//
// Solidity: event AIProposalSubmitted(uint256 indexed marketId, uint256 indexed outcomeId, address indexed proposer, address aiSigner, uint256 bondAmount, bytes32 signatureHash)
func (_AIOracleAdapter *AIOracleAdapterFilterer) FilterAIProposalSubmitted(opts *bind.FilterOpts, marketId []*big.Int, outcomeId []*big.Int, proposer []common.Address) (*AIOracleAdapterAIProposalSubmittedIterator, error) {

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

	logs, sub, err := _AIOracleAdapter.contract.FilterLogs(opts, "AIProposalSubmitted", marketIdRule, outcomeIdRule, proposerRule)
	if err != nil {
		return nil, err
	}
	return &AIOracleAdapterAIProposalSubmittedIterator{contract: _AIOracleAdapter.contract, event: "AIProposalSubmitted", logs: logs, sub: sub}, nil
}

// WatchAIProposalSubmitted is a free log subscription operation binding the contract event 0x0b1a7dc1c0ac67b57ba259d0fbc6e4633797b01c17a70f944e6b615227ad0cb0.
//
// Solidity: event AIProposalSubmitted(uint256 indexed marketId, uint256 indexed outcomeId, address indexed proposer, address aiSigner, uint256 bondAmount, bytes32 signatureHash)
func (_AIOracleAdapter *AIOracleAdapterFilterer) WatchAIProposalSubmitted(opts *bind.WatchOpts, sink chan<- *AIOracleAdapterAIProposalSubmitted, marketId []*big.Int, outcomeId []*big.Int, proposer []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _AIOracleAdapter.contract.WatchLogs(opts, "AIProposalSubmitted", marketIdRule, outcomeIdRule, proposerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AIOracleAdapterAIProposalSubmitted)
				if err := _AIOracleAdapter.contract.UnpackLog(event, "AIProposalSubmitted", log); err != nil {
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

// ParseAIProposalSubmitted is a log parse operation binding the contract event 0x0b1a7dc1c0ac67b57ba259d0fbc6e4633797b01c17a70f944e6b615227ad0cb0.
//
// Solidity: event AIProposalSubmitted(uint256 indexed marketId, uint256 indexed outcomeId, address indexed proposer, address aiSigner, uint256 bondAmount, bytes32 signatureHash)
func (_AIOracleAdapter *AIOracleAdapterFilterer) ParseAIProposalSubmitted(log types.Log) (*AIOracleAdapterAIProposalSubmitted, error) {
	event := new(AIOracleAdapterAIProposalSubmitted)
	if err := _AIOracleAdapter.contract.UnpackLog(event, "AIProposalSubmitted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AIOracleAdapterOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the AIOracleAdapter contract.
type AIOracleAdapterOwnershipTransferredIterator struct {
	Event *AIOracleAdapterOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *AIOracleAdapterOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AIOracleAdapterOwnershipTransferred)
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
		it.Event = new(AIOracleAdapterOwnershipTransferred)
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
func (it *AIOracleAdapterOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AIOracleAdapterOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AIOracleAdapterOwnershipTransferred represents a OwnershipTransferred event raised by the AIOracleAdapter contract.
type AIOracleAdapterOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AIOracleAdapter *AIOracleAdapterFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*AIOracleAdapterOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AIOracleAdapter.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &AIOracleAdapterOwnershipTransferredIterator{contract: _AIOracleAdapter.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_AIOracleAdapter *AIOracleAdapterFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *AIOracleAdapterOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _AIOracleAdapter.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AIOracleAdapterOwnershipTransferred)
				if err := _AIOracleAdapter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_AIOracleAdapter *AIOracleAdapterFilterer) ParseOwnershipTransferred(log types.Log) (*AIOracleAdapterOwnershipTransferred, error) {
	event := new(AIOracleAdapterOwnershipTransferred)
	if err := _AIOracleAdapter.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AIOracleAdapterSignerAddedIterator is returned from FilterSignerAdded and is used to iterate over the raw logs and unpacked data for SignerAdded events raised by the AIOracleAdapter contract.
type AIOracleAdapterSignerAddedIterator struct {
	Event *AIOracleAdapterSignerAdded // Event containing the contract specifics and raw log

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
func (it *AIOracleAdapterSignerAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AIOracleAdapterSignerAdded)
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
		it.Event = new(AIOracleAdapterSignerAdded)
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
func (it *AIOracleAdapterSignerAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AIOracleAdapterSignerAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AIOracleAdapterSignerAdded represents a SignerAdded event raised by the AIOracleAdapter contract.
type AIOracleAdapterSignerAdded struct {
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSignerAdded is a free log retrieval operation binding the contract event 0x47d1c22a25bb3a5d4e481b9b1e6944c2eade3181a0a20b495ed61d35b5323f24.
//
// Solidity: event SignerAdded(address indexed signer)
func (_AIOracleAdapter *AIOracleAdapterFilterer) FilterSignerAdded(opts *bind.FilterOpts, signer []common.Address) (*AIOracleAdapterSignerAddedIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _AIOracleAdapter.contract.FilterLogs(opts, "SignerAdded", signerRule)
	if err != nil {
		return nil, err
	}
	return &AIOracleAdapterSignerAddedIterator{contract: _AIOracleAdapter.contract, event: "SignerAdded", logs: logs, sub: sub}, nil
}

// WatchSignerAdded is a free log subscription operation binding the contract event 0x47d1c22a25bb3a5d4e481b9b1e6944c2eade3181a0a20b495ed61d35b5323f24.
//
// Solidity: event SignerAdded(address indexed signer)
func (_AIOracleAdapter *AIOracleAdapterFilterer) WatchSignerAdded(opts *bind.WatchOpts, sink chan<- *AIOracleAdapterSignerAdded, signer []common.Address) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _AIOracleAdapter.contract.WatchLogs(opts, "SignerAdded", signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AIOracleAdapterSignerAdded)
				if err := _AIOracleAdapter.contract.UnpackLog(event, "SignerAdded", log); err != nil {
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

// ParseSignerAdded is a log parse operation binding the contract event 0x47d1c22a25bb3a5d4e481b9b1e6944c2eade3181a0a20b495ed61d35b5323f24.
//
// Solidity: event SignerAdded(address indexed signer)
func (_AIOracleAdapter *AIOracleAdapterFilterer) ParseSignerAdded(log types.Log) (*AIOracleAdapterSignerAdded, error) {
	event := new(AIOracleAdapterSignerAdded)
	if err := _AIOracleAdapter.contract.UnpackLog(event, "SignerAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// AIOracleAdapterSignerRemovedIterator is returned from FilterSignerRemoved and is used to iterate over the raw logs and unpacked data for SignerRemoved events raised by the AIOracleAdapter contract.
type AIOracleAdapterSignerRemovedIterator struct {
	Event *AIOracleAdapterSignerRemoved // Event containing the contract specifics and raw log

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
func (it *AIOracleAdapterSignerRemovedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(AIOracleAdapterSignerRemoved)
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
		it.Event = new(AIOracleAdapterSignerRemoved)
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
func (it *AIOracleAdapterSignerRemovedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *AIOracleAdapterSignerRemovedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// AIOracleAdapterSignerRemoved represents a SignerRemoved event raised by the AIOracleAdapter contract.
type AIOracleAdapterSignerRemoved struct {
	Signer common.Address
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterSignerRemoved is a free log retrieval operation binding the contract event 0x3525e22824a8a7df2c9a6029941c824cf95b6447f1e13d5128fd3826d35afe8b.
//
// Solidity: event SignerRemoved(address indexed signer)
func (_AIOracleAdapter *AIOracleAdapterFilterer) FilterSignerRemoved(opts *bind.FilterOpts, signer []common.Address) (*AIOracleAdapterSignerRemovedIterator, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _AIOracleAdapter.contract.FilterLogs(opts, "SignerRemoved", signerRule)
	if err != nil {
		return nil, err
	}
	return &AIOracleAdapterSignerRemovedIterator{contract: _AIOracleAdapter.contract, event: "SignerRemoved", logs: logs, sub: sub}, nil
}

// WatchSignerRemoved is a free log subscription operation binding the contract event 0x3525e22824a8a7df2c9a6029941c824cf95b6447f1e13d5128fd3826d35afe8b.
//
// Solidity: event SignerRemoved(address indexed signer)
func (_AIOracleAdapter *AIOracleAdapterFilterer) WatchSignerRemoved(opts *bind.WatchOpts, sink chan<- *AIOracleAdapterSignerRemoved, signer []common.Address) (event.Subscription, error) {

	var signerRule []interface{}
	for _, signerItem := range signer {
		signerRule = append(signerRule, signerItem)
	}

	logs, sub, err := _AIOracleAdapter.contract.WatchLogs(opts, "SignerRemoved", signerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(AIOracleAdapterSignerRemoved)
				if err := _AIOracleAdapter.contract.UnpackLog(event, "SignerRemoved", log); err != nil {
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

// ParseSignerRemoved is a log parse operation binding the contract event 0x3525e22824a8a7df2c9a6029941c824cf95b6447f1e13d5128fd3826d35afe8b.
//
// Solidity: event SignerRemoved(address indexed signer)
func (_AIOracleAdapter *AIOracleAdapterFilterer) ParseSignerRemoved(log types.Log) (*AIOracleAdapterSignerRemoved, error) {
	event := new(AIOracleAdapterSignerRemoved)
	if err := _AIOracleAdapter.contract.UnpackLog(event, "SignerRemoved", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
