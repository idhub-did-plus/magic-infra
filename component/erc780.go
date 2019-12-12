// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package component

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

// ComponentABI is the input ABI used to generate the binding from.
const ComponentABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"getClaim\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"removeClaim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"setClaim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"setSelfClaim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Component is an auto generated Go binding around an Ethereum contract.
type Component struct {
	ComponentCaller     // Read-only binding to the contract
	ComponentTransactor // Write-only binding to the contract
	ComponentFilterer   // Log filterer for contract events
}

// ComponentCaller is an auto generated read-only Go binding around an Ethereum contract.
type ComponentCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComponentTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ComponentTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComponentFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ComponentFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ComponentSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ComponentSession struct {
	Contract     *Component        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ComponentCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ComponentCallerSession struct {
	Contract *ComponentCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// ComponentTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ComponentTransactorSession struct {
	Contract     *ComponentTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// ComponentRaw is an auto generated low-level Go binding around an Ethereum contract.
type ComponentRaw struct {
	Contract *Component // Generic contract binding to access the raw methods on
}

// ComponentCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ComponentCallerRaw struct {
	Contract *ComponentCaller // Generic read-only contract binding to access the raw methods on
}

// ComponentTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ComponentTransactorRaw struct {
	Contract *ComponentTransactor // Generic write-only contract binding to access the raw methods on
}

// NewComponent creates a new instance of Component, bound to a specific deployed contract.
func NewComponent(address common.Address, backend bind.ContractBackend) (*Component, error) {
	contract, err := bindComponent(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Component{ComponentCaller: ComponentCaller{contract: contract}, ComponentTransactor: ComponentTransactor{contract: contract}, ComponentFilterer: ComponentFilterer{contract: contract}}, nil
}

// NewComponentCaller creates a new read-only instance of Component, bound to a specific deployed contract.
func NewComponentCaller(address common.Address, caller bind.ContractCaller) (*ComponentCaller, error) {
	contract, err := bindComponent(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ComponentCaller{contract: contract}, nil
}

// NewComponentTransactor creates a new write-only instance of Component, bound to a specific deployed contract.
func NewComponentTransactor(address common.Address, transactor bind.ContractTransactor) (*ComponentTransactor, error) {
	contract, err := bindComponent(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ComponentTransactor{contract: contract}, nil
}

// NewComponentFilterer creates a new log filterer instance of Component, bound to a specific deployed contract.
func NewComponentFilterer(address common.Address, filterer bind.ContractFilterer) (*ComponentFilterer, error) {
	contract, err := bindComponent(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ComponentFilterer{contract: contract}, nil
}

// bindComponent binds a generic wrapper to an already deployed contract.
func bindComponent(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(ComponentABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Component *ComponentRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Component.Contract.ComponentCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Component *ComponentRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Component.Contract.ComponentTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Component *ComponentRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Component.Contract.ComponentTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Component *ComponentCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Component.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Component *ComponentTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Component.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Component *ComponentTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Component.Contract.contract.Transact(opts, method, params...)
}

// GetClaim is a free data retrieval call binding the contract method 0xe1661eff.
//
// Solidity: function getClaim(address issuer, address subject, bytes32 key) constant returns(bytes32)
func (_Component *ComponentCaller) GetClaim(opts *bind.CallOpts, issuer common.Address, subject common.Address, key [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Component.contract.Call(opts, out, "getClaim", issuer, subject, key)
	return *ret0, err
}

// GetClaim is a free data retrieval call binding the contract method 0xe1661eff.
//
// Solidity: function getClaim(address issuer, address subject, bytes32 key) constant returns(bytes32)
func (_Component *ComponentSession) GetClaim(issuer common.Address, subject common.Address, key [32]byte) ([32]byte, error) {
	return _Component.Contract.GetClaim(&_Component.CallOpts, issuer, subject, key)
}

// GetClaim is a free data retrieval call binding the contract method 0xe1661eff.
//
// Solidity: function getClaim(address issuer, address subject, bytes32 key) constant returns(bytes32)
func (_Component *ComponentCallerSession) GetClaim(issuer common.Address, subject common.Address, key [32]byte) ([32]byte, error) {
	return _Component.Contract.GetClaim(&_Component.CallOpts, issuer, subject, key)
}

// RemoveClaim is a paid mutator transaction binding the contract method 0xc7508ec7.
//
// Solidity: function removeClaim(address issuer, address subject, bytes32 key) returns()
func (_Component *ComponentTransactor) RemoveClaim(opts *bind.TransactOpts, issuer common.Address, subject common.Address, key [32]byte) (*types.Transaction, error) {
	return _Component.contract.Transact(opts, "removeClaim", issuer, subject, key)
}

// RemoveClaim is a paid mutator transaction binding the contract method 0xc7508ec7.
//
// Solidity: function removeClaim(address issuer, address subject, bytes32 key) returns()
func (_Component *ComponentSession) RemoveClaim(issuer common.Address, subject common.Address, key [32]byte) (*types.Transaction, error) {
	return _Component.Contract.RemoveClaim(&_Component.TransactOpts, issuer, subject, key)
}

// RemoveClaim is a paid mutator transaction binding the contract method 0xc7508ec7.
//
// Solidity: function removeClaim(address issuer, address subject, bytes32 key) returns()
func (_Component *ComponentTransactorSession) RemoveClaim(issuer common.Address, subject common.Address, key [32]byte) (*types.Transaction, error) {
	return _Component.Contract.RemoveClaim(&_Component.TransactOpts, issuer, subject, key)
}

// SetClaim is a paid mutator transaction binding the contract method 0x9918925d.
//
// Solidity: function setClaim(address subject, bytes32 key, bytes32 value) returns()
func (_Component *ComponentTransactor) SetClaim(opts *bind.TransactOpts, subject common.Address, key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Component.contract.Transact(opts, "setClaim", subject, key, value)
}

// SetClaim is a paid mutator transaction binding the contract method 0x9918925d.
//
// Solidity: function setClaim(address subject, bytes32 key, bytes32 value) returns()
func (_Component *ComponentSession) SetClaim(subject common.Address, key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Component.Contract.SetClaim(&_Component.TransactOpts, subject, key, value)
}

// SetClaim is a paid mutator transaction binding the contract method 0x9918925d.
//
// Solidity: function setClaim(address subject, bytes32 key, bytes32 value) returns()
func (_Component *ComponentTransactorSession) SetClaim(subject common.Address, key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Component.Contract.SetClaim(&_Component.TransactOpts, subject, key, value)
}

// SetSelfClaim is a paid mutator transaction binding the contract method 0x9155b01a.
//
// Solidity: function setSelfClaim(bytes32 key, bytes32 value) returns()
func (_Component *ComponentTransactor) SetSelfClaim(opts *bind.TransactOpts, key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Component.contract.Transact(opts, "setSelfClaim", key, value)
}

// SetSelfClaim is a paid mutator transaction binding the contract method 0x9155b01a.
//
// Solidity: function setSelfClaim(bytes32 key, bytes32 value) returns()
func (_Component *ComponentSession) SetSelfClaim(key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Component.Contract.SetSelfClaim(&_Component.TransactOpts, key, value)
}

// SetSelfClaim is a paid mutator transaction binding the contract method 0x9155b01a.
//
// Solidity: function setSelfClaim(bytes32 key, bytes32 value) returns()
func (_Component *ComponentTransactorSession) SetSelfClaim(key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Component.Contract.SetSelfClaim(&_Component.TransactOpts, key, value)
}
