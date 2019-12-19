// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

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

// Reference imports to suppress errors if they are not otherwise used.
var (
	_ = big.NewInt
	_ = strings.NewReader
	_ = ethereum.NotFound
	_ = abi.U256
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// ClaimABI is the input ABI used to generate the binding from.
const ClaimABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"getClaim\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"removeClaim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"setClaim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"setSelfClaim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Claim is an auto generated Go binding around an Ethereum contract.
type Claim struct {
	ClaimCaller     // Read-only binding to the contract
	ClaimTransactor // Write-only binding to the contract
	ClaimFilterer   // Log filterer for contract events
}

// ClaimCaller is an auto generated read-only Go binding around an Ethereum contract.
type ClaimCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ClaimTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ClaimFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ClaimSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ClaimSession struct {
	Contract     *Claim            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClaimCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ClaimCallerSession struct {
	Contract *ClaimCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// ClaimTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ClaimTransactorSession struct {
	Contract     *ClaimTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ClaimRaw is an auto generated low-level Go binding around an Ethereum contract.
type ClaimRaw struct {
	Contract *Claim // Generic contract binding to access the raw methods on
}

// ClaimCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ClaimCallerRaw struct {
	Contract *ClaimCaller // Generic read-only contract binding to access the raw methods on
}

// ClaimTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ClaimTransactorRaw struct {
	Contract *ClaimTransactor // Generic write-only contract binding to access the raw methods on
}

// NewClaim creates a new instance of Claim, bound to a specific deployed contract.
func NewClaim(address common.Address, backend bind.ContractBackend) (*Claim, error) {
	contract, err := bindClaim(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Claim{ClaimCaller: ClaimCaller{contract: contract}, ClaimTransactor: ClaimTransactor{contract: contract}, ClaimFilterer: ClaimFilterer{contract: contract}}, nil
}

// NewClaimCaller creates a new read-only instance of Claim, bound to a specific deployed contract.
func NewClaimCaller(address common.Address, caller bind.ContractCaller) (*ClaimCaller, error) {
	contract, err := bindClaim(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ClaimCaller{contract: contract}, nil
}

// NewClaimTransactor creates a new write-only instance of Claim, bound to a specific deployed contract.
func NewClaimTransactor(address common.Address, transactor bind.ContractTransactor) (*ClaimTransactor, error) {
	contract, err := bindClaim(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ClaimTransactor{contract: contract}, nil
}

// NewClaimFilterer creates a new log filterer instance of Claim, bound to a specific deployed contract.
func NewClaimFilterer(address common.Address, filterer bind.ContractFilterer) (*ClaimFilterer, error) {
	contract, err := bindClaim(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ClaimFilterer{contract: contract}, nil
}

// bindClaim binds a generic wrapper to an already deployed contract.
func bindClaim(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ClaimABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Claim *ClaimRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Claim.Contract.ClaimCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Claim *ClaimRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Claim.Contract.ClaimTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Claim *ClaimRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Claim.Contract.ClaimTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Claim *ClaimCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Claim.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Claim *ClaimTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Claim.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Claim *ClaimTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Claim.Contract.contract.Transact(opts, method, params...)
}

// GetClaim is a free data retrieval call binding the contract method 0xe1661eff.
//
// Solidity: function getClaim(address issuer, address subject, bytes32 key) constant returns(bytes32)
func (_Claim *ClaimCaller) GetClaim(opts *bind.CallOpts, issuer common.Address, subject common.Address, key [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Claim.contract.Call(opts, out, "getClaim", issuer, subject, key)
	return *ret0, err
}

// GetClaim is a free data retrieval call binding the contract method 0xe1661eff.
//
// Solidity: function getClaim(address issuer, address subject, bytes32 key) constant returns(bytes32)
func (_Claim *ClaimSession) GetClaim(issuer common.Address, subject common.Address, key [32]byte) ([32]byte, error) {
	return _Claim.Contract.GetClaim(&_Claim.CallOpts, issuer, subject, key)
}

// GetClaim is a free data retrieval call binding the contract method 0xe1661eff.
//
// Solidity: function getClaim(address issuer, address subject, bytes32 key) constant returns(bytes32)
func (_Claim *ClaimCallerSession) GetClaim(issuer common.Address, subject common.Address, key [32]byte) ([32]byte, error) {
	return _Claim.Contract.GetClaim(&_Claim.CallOpts, issuer, subject, key)
}

// RemoveClaim is a paid mutator transaction binding the contract method 0xc7508ec7.
//
// Solidity: function removeClaim(address issuer, address subject, bytes32 key) returns()
func (_Claim *ClaimTransactor) RemoveClaim(opts *bind.TransactOpts, issuer common.Address, subject common.Address, key [32]byte) (*types.Transaction, error) {
	return _Claim.contract.Transact(opts, "removeClaim", issuer, subject, key)
}

// RemoveClaim is a paid mutator transaction binding the contract method 0xc7508ec7.
//
// Solidity: function removeClaim(address issuer, address subject, bytes32 key) returns()
func (_Claim *ClaimSession) RemoveClaim(issuer common.Address, subject common.Address, key [32]byte) (*types.Transaction, error) {
	return _Claim.Contract.RemoveClaim(&_Claim.TransactOpts, issuer, subject, key)
}

// RemoveClaim is a paid mutator transaction binding the contract method 0xc7508ec7.
//
// Solidity: function removeClaim(address issuer, address subject, bytes32 key) returns()
func (_Claim *ClaimTransactorSession) RemoveClaim(issuer common.Address, subject common.Address, key [32]byte) (*types.Transaction, error) {
	return _Claim.Contract.RemoveClaim(&_Claim.TransactOpts, issuer, subject, key)
}

// SetClaim is a paid mutator transaction binding the contract method 0x9918925d.
//
// Solidity: function setClaim(address subject, bytes32 key, bytes32 value) returns()
func (_Claim *ClaimTransactor) SetClaim(opts *bind.TransactOpts, subject common.Address, key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Claim.contract.Transact(opts, "setClaim", subject, key, value)
}

// SetClaim is a paid mutator transaction binding the contract method 0x9918925d.
//
// Solidity: function setClaim(address subject, bytes32 key, bytes32 value) returns()
func (_Claim *ClaimSession) SetClaim(subject common.Address, key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Claim.Contract.SetClaim(&_Claim.TransactOpts, subject, key, value)
}

// SetClaim is a paid mutator transaction binding the contract method 0x9918925d.
//
// Solidity: function setClaim(address subject, bytes32 key, bytes32 value) returns()
func (_Claim *ClaimTransactorSession) SetClaim(subject common.Address, key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Claim.Contract.SetClaim(&_Claim.TransactOpts, subject, key, value)
}

// SetSelfClaim is a paid mutator transaction binding the contract method 0x9155b01a.
//
// Solidity: function setSelfClaim(bytes32 key, bytes32 value) returns()
func (_Claim *ClaimTransactor) SetSelfClaim(opts *bind.TransactOpts, key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Claim.contract.Transact(opts, "setSelfClaim", key, value)
}

// SetSelfClaim is a paid mutator transaction binding the contract method 0x9155b01a.
//
// Solidity: function setSelfClaim(bytes32 key, bytes32 value) returns()
func (_Claim *ClaimSession) SetSelfClaim(key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Claim.Contract.SetSelfClaim(&_Claim.TransactOpts, key, value)
}

// SetSelfClaim is a paid mutator transaction binding the contract method 0x9155b01a.
//
// Solidity: function setSelfClaim(bytes32 key, bytes32 value) returns()
func (_Claim *ClaimTransactorSession) SetSelfClaim(key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Claim.Contract.SetSelfClaim(&_Claim.TransactOpts, key, value)
}
