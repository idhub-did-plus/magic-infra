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

// IdentityABI is the input ABI used to generate the binding from.
const IdentityABI = "[{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"approvingAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"addressToAdd\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"addAssociatedAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"approvingAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"addressToAdd\",\"type\":\"address\"},{\"internalType\":\"uint8[2]\",\"name\":\"v\",\"type\":\"uint8[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"r\",\"type\":\"bytes32[2]\"},{\"internalType\":\"bytes32[2]\",\"name\":\"s\",\"type\":\"bytes32[2]\"},{\"internalType\":\"uint256[2]\",\"name\":\"timestamp\",\"type\":\"uint256[2]\"}],\"name\":\"addAssociatedAddressDelegated\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"providers\",\"type\":\"address[]\"}],\"name\":\"addProviders\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ein\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"providers\",\"type\":\"address[]\"}],\"name\":\"addProvidersFor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"resolvers\",\"type\":\"address[]\"}],\"name\":\"addResolvers\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ein\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"resolvers\",\"type\":\"address[]\"}],\"name\":\"addResolversFor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recoveryAddress\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"providers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"resolvers\",\"type\":\"address[]\"}],\"name\":\"createIdentity\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ein\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"recoveryAddress\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"associatedAddress\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"providers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"resolvers\",\"type\":\"address[]\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"createIdentityDelegated\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ein\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"getEIN\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"ein\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ein\",\"type\":\"uint256\"}],\"name\":\"getIdentity\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"recoveryAddress\",\"type\":\"address\"},{\"internalType\":\"address[]\",\"name\":\"associatedAddresses\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"providers\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"resolvers\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"hasIdentity\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ein\",\"type\":\"uint256\"}],\"name\":\"identityExists\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ein\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"}],\"name\":\"isAssociatedAddressFor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ein\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"provider\",\"type\":\"address\"}],\"name\":\"isProviderFor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ein\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"resolver\",\"type\":\"address\"}],\"name\":\"isResolverFor\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"_address\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"messageHash\",\"type\":\"bytes32\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"isSigned\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"removeAssociatedAddress\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"addressToRemove\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"removeAssociatedAddressDelegated\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"providers\",\"type\":\"address[]\"}],\"name\":\"removeProviders\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ein\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"providers\",\"type\":\"address[]\"}],\"name\":\"removeProvidersFor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address[]\",\"name\":\"resolvers\",\"type\":\"address[]\"}],\"name\":\"removeResolvers\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ein\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"resolvers\",\"type\":\"address[]\"}],\"name\":\"removeResolversFor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ein\",\"type\":\"uint256\"},{\"internalType\":\"address[]\",\"name\":\"firstChunk\",\"type\":\"address[]\"},{\"internalType\":\"address[]\",\"name\":\"lastChunk\",\"type\":\"address[]\"},{\"internalType\":\"bool\",\"name\":\"resetResolvers\",\"type\":\"bool\"}],\"name\":\"triggerDestruction\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ein\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"newAssociatedAddress\",\"type\":\"address\"},{\"internalType\":\"uint8\",\"name\":\"v\",\"type\":\"uint8\"},{\"internalType\":\"bytes32\",\"name\":\"r\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"s\",\"type\":\"bytes32\"},{\"internalType\":\"uint256\",\"name\":\"timestamp\",\"type\":\"uint256\"}],\"name\":\"triggerRecovery\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"newRecoveryAddress\",\"type\":\"address\"}],\"name\":\"triggerRecoveryAddressChange\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"ein\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"newRecoveryAddress\",\"type\":\"address\"}],\"name\":\"triggerRecoveryAddressChangeFor\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Identity is an auto generated Go binding around an Ethereum contract.
type Identity struct {
	IdentityCaller     // Read-only binding to the contract
	IdentityTransactor // Write-only binding to the contract
	IdentityFilterer   // Log filterer for contract events
}

// IdentityCaller is an auto generated read-only Go binding around an Ethereum contract.
type IdentityCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityTransactor is an auto generated write-only Go binding around an Ethereum contract.
type IdentityTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentityFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type IdentityFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// IdentitySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type IdentitySession struct {
	Contract     *Identity         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// IdentityCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type IdentityCallerSession struct {
	Contract *IdentityCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// IdentityTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type IdentityTransactorSession struct {
	Contract     *IdentityTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// IdentityRaw is an auto generated low-level Go binding around an Ethereum contract.
type IdentityRaw struct {
	Contract *Identity // Generic contract binding to access the raw methods on
}

// IdentityCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type IdentityCallerRaw struct {
	Contract *IdentityCaller // Generic read-only contract binding to access the raw methods on
}

// IdentityTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type IdentityTransactorRaw struct {
	Contract *IdentityTransactor // Generic write-only contract binding to access the raw methods on
}

// NewIdentity creates a new instance of Identity, bound to a specific deployed contract.
func NewIdentity(address common.Address, backend bind.ContractBackend) (*Identity, error) {
	contract, err := bindIdentity(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Identity{IdentityCaller: IdentityCaller{contract: contract}, IdentityTransactor: IdentityTransactor{contract: contract}, IdentityFilterer: IdentityFilterer{contract: contract}}, nil
}

// NewIdentityCaller creates a new read-only instance of Identity, bound to a specific deployed contract.
func NewIdentityCaller(address common.Address, caller bind.ContractCaller) (*IdentityCaller, error) {
	contract, err := bindIdentity(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityCaller{contract: contract}, nil
}

// NewIdentityTransactor creates a new write-only instance of Identity, bound to a specific deployed contract.
func NewIdentityTransactor(address common.Address, transactor bind.ContractTransactor) (*IdentityTransactor, error) {
	contract, err := bindIdentity(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &IdentityTransactor{contract: contract}, nil
}

// NewIdentityFilterer creates a new log filterer instance of Identity, bound to a specific deployed contract.
func NewIdentityFilterer(address common.Address, filterer bind.ContractFilterer) (*IdentityFilterer, error) {
	contract, err := bindIdentity(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &IdentityFilterer{contract: contract}, nil
}

// bindIdentity binds a generic wrapper to an already deployed contract.
func bindIdentity(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(IdentityABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Identity *IdentityRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Identity.Contract.IdentityCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Identity *IdentityRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Identity.Contract.IdentityTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Identity *IdentityRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Identity.Contract.IdentityTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Identity *IdentityCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Identity.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Identity *IdentityTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Identity.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Identity *IdentityTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Identity.Contract.contract.Transact(opts, method, params...)
}

// GetEIN is a free data retrieval call binding the contract method 0x05c62c2f.
//
// Solidity: function getEIN(address _address) constant returns(uint256 ein)
func (_Identity *IdentityCaller) GetEIN(opts *bind.CallOpts, _address common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Identity.contract.Call(opts, out, "getEIN", _address)
	return *ret0, err
}

// GetEIN is a free data retrieval call binding the contract method 0x05c62c2f.
//
// Solidity: function getEIN(address _address) constant returns(uint256 ein)
func (_Identity *IdentitySession) GetEIN(_address common.Address) (*big.Int, error) {
	return _Identity.Contract.GetEIN(&_Identity.CallOpts, _address)
}

// GetEIN is a free data retrieval call binding the contract method 0x05c62c2f.
//
// Solidity: function getEIN(address _address) constant returns(uint256 ein)
func (_Identity *IdentityCallerSession) GetEIN(_address common.Address) (*big.Int, error) {
	return _Identity.Contract.GetEIN(&_Identity.CallOpts, _address)
}

// GetIdentity is a free data retrieval call binding the contract method 0x85e3f058.
//
// Solidity: function getIdentity(uint256 ein) constant returns(address recoveryAddress, address[] associatedAddresses, address[] providers, address[] resolvers)
func (_Identity *IdentityCaller) GetIdentity(opts *bind.CallOpts, ein *big.Int) (struct {
	RecoveryAddress     common.Address
	AssociatedAddresses []common.Address
	Providers           []common.Address
	Resolvers           []common.Address
}, error) {
	ret := new(struct {
		RecoveryAddress     common.Address
		AssociatedAddresses []common.Address
		Providers           []common.Address
		Resolvers           []common.Address
	})
	out := ret
	err := _Identity.contract.Call(opts, out, "getIdentity", ein)
	return *ret, err
}

// GetIdentity is a free data retrieval call binding the contract method 0x85e3f058.
//
// Solidity: function getIdentity(uint256 ein) constant returns(address recoveryAddress, address[] associatedAddresses, address[] providers, address[] resolvers)
func (_Identity *IdentitySession) GetIdentity(ein *big.Int) (struct {
	RecoveryAddress     common.Address
	AssociatedAddresses []common.Address
	Providers           []common.Address
	Resolvers           []common.Address
}, error) {
	return _Identity.Contract.GetIdentity(&_Identity.CallOpts, ein)
}

// GetIdentity is a free data retrieval call binding the contract method 0x85e3f058.
//
// Solidity: function getIdentity(uint256 ein) constant returns(address recoveryAddress, address[] associatedAddresses, address[] providers, address[] resolvers)
func (_Identity *IdentityCallerSession) GetIdentity(ein *big.Int) (struct {
	RecoveryAddress     common.Address
	AssociatedAddresses []common.Address
	Providers           []common.Address
	Resolvers           []common.Address
}, error) {
	return _Identity.Contract.GetIdentity(&_Identity.CallOpts, ein)
}

// HasIdentity is a free data retrieval call binding the contract method 0x237f1a21.
//
// Solidity: function hasIdentity(address _address) constant returns(bool)
func (_Identity *IdentityCaller) HasIdentity(opts *bind.CallOpts, _address common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Identity.contract.Call(opts, out, "hasIdentity", _address)
	return *ret0, err
}

// HasIdentity is a free data retrieval call binding the contract method 0x237f1a21.
//
// Solidity: function hasIdentity(address _address) constant returns(bool)
func (_Identity *IdentitySession) HasIdentity(_address common.Address) (bool, error) {
	return _Identity.Contract.HasIdentity(&_Identity.CallOpts, _address)
}

// HasIdentity is a free data retrieval call binding the contract method 0x237f1a21.
//
// Solidity: function hasIdentity(address _address) constant returns(bool)
func (_Identity *IdentityCallerSession) HasIdentity(_address common.Address) (bool, error) {
	return _Identity.Contract.HasIdentity(&_Identity.CallOpts, _address)
}

// IdentityExists is a free data retrieval call binding the contract method 0x5b5aed3a.
//
// Solidity: function identityExists(uint256 ein) constant returns(bool)
func (_Identity *IdentityCaller) IdentityExists(opts *bind.CallOpts, ein *big.Int) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Identity.contract.Call(opts, out, "identityExists", ein)
	return *ret0, err
}

// IdentityExists is a free data retrieval call binding the contract method 0x5b5aed3a.
//
// Solidity: function identityExists(uint256 ein) constant returns(bool)
func (_Identity *IdentitySession) IdentityExists(ein *big.Int) (bool, error) {
	return _Identity.Contract.IdentityExists(&_Identity.CallOpts, ein)
}

// IdentityExists is a free data retrieval call binding the contract method 0x5b5aed3a.
//
// Solidity: function identityExists(uint256 ein) constant returns(bool)
func (_Identity *IdentityCallerSession) IdentityExists(ein *big.Int) (bool, error) {
	return _Identity.Contract.IdentityExists(&_Identity.CallOpts, ein)
}

// IsAssociatedAddressFor is a free data retrieval call binding the contract method 0xa687662d.
//
// Solidity: function isAssociatedAddressFor(uint256 ein, address _address) constant returns(bool)
func (_Identity *IdentityCaller) IsAssociatedAddressFor(opts *bind.CallOpts, ein *big.Int, _address common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Identity.contract.Call(opts, out, "isAssociatedAddressFor", ein, _address)
	return *ret0, err
}

// IsAssociatedAddressFor is a free data retrieval call binding the contract method 0xa687662d.
//
// Solidity: function isAssociatedAddressFor(uint256 ein, address _address) constant returns(bool)
func (_Identity *IdentitySession) IsAssociatedAddressFor(ein *big.Int, _address common.Address) (bool, error) {
	return _Identity.Contract.IsAssociatedAddressFor(&_Identity.CallOpts, ein, _address)
}

// IsAssociatedAddressFor is a free data retrieval call binding the contract method 0xa687662d.
//
// Solidity: function isAssociatedAddressFor(uint256 ein, address _address) constant returns(bool)
func (_Identity *IdentityCallerSession) IsAssociatedAddressFor(ein *big.Int, _address common.Address) (bool, error) {
	return _Identity.Contract.IsAssociatedAddressFor(&_Identity.CallOpts, ein, _address)
}

// IsProviderFor is a free data retrieval call binding the contract method 0x53a9698a.
//
// Solidity: function isProviderFor(uint256 ein, address provider) constant returns(bool)
func (_Identity *IdentityCaller) IsProviderFor(opts *bind.CallOpts, ein *big.Int, provider common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Identity.contract.Call(opts, out, "isProviderFor", ein, provider)
	return *ret0, err
}

// IsProviderFor is a free data retrieval call binding the contract method 0x53a9698a.
//
// Solidity: function isProviderFor(uint256 ein, address provider) constant returns(bool)
func (_Identity *IdentitySession) IsProviderFor(ein *big.Int, provider common.Address) (bool, error) {
	return _Identity.Contract.IsProviderFor(&_Identity.CallOpts, ein, provider)
}

// IsProviderFor is a free data retrieval call binding the contract method 0x53a9698a.
//
// Solidity: function isProviderFor(uint256 ein, address provider) constant returns(bool)
func (_Identity *IdentityCallerSession) IsProviderFor(ein *big.Int, provider common.Address) (bool, error) {
	return _Identity.Contract.IsProviderFor(&_Identity.CallOpts, ein, provider)
}

// IsResolverFor is a free data retrieval call binding the contract method 0xd4b1cdcc.
//
// Solidity: function isResolverFor(uint256 ein, address resolver) constant returns(bool)
func (_Identity *IdentityCaller) IsResolverFor(opts *bind.CallOpts, ein *big.Int, resolver common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Identity.contract.Call(opts, out, "isResolverFor", ein, resolver)
	return *ret0, err
}

// IsResolverFor is a free data retrieval call binding the contract method 0xd4b1cdcc.
//
// Solidity: function isResolverFor(uint256 ein, address resolver) constant returns(bool)
func (_Identity *IdentitySession) IsResolverFor(ein *big.Int, resolver common.Address) (bool, error) {
	return _Identity.Contract.IsResolverFor(&_Identity.CallOpts, ein, resolver)
}

// IsResolverFor is a free data retrieval call binding the contract method 0xd4b1cdcc.
//
// Solidity: function isResolverFor(uint256 ein, address resolver) constant returns(bool)
func (_Identity *IdentityCallerSession) IsResolverFor(ein *big.Int, resolver common.Address) (bool, error) {
	return _Identity.Contract.IsResolverFor(&_Identity.CallOpts, ein, resolver)
}

// IsSigned is a free data retrieval call binding the contract method 0x8677ebe8.
//
// Solidity: function isSigned(address _address, bytes32 messageHash, uint8 v, bytes32 r, bytes32 s) constant returns(bool)
func (_Identity *IdentityCaller) IsSigned(opts *bind.CallOpts, _address common.Address, messageHash [32]byte, v uint8, r [32]byte, s [32]byte) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _Identity.contract.Call(opts, out, "isSigned", _address, messageHash, v, r, s)
	return *ret0, err
}

// IsSigned is a free data retrieval call binding the contract method 0x8677ebe8.
//
// Solidity: function isSigned(address _address, bytes32 messageHash, uint8 v, bytes32 r, bytes32 s) constant returns(bool)
func (_Identity *IdentitySession) IsSigned(_address common.Address, messageHash [32]byte, v uint8, r [32]byte, s [32]byte) (bool, error) {
	return _Identity.Contract.IsSigned(&_Identity.CallOpts, _address, messageHash, v, r, s)
}

// IsSigned is a free data retrieval call binding the contract method 0x8677ebe8.
//
// Solidity: function isSigned(address _address, bytes32 messageHash, uint8 v, bytes32 r, bytes32 s) constant returns(bool)
func (_Identity *IdentityCallerSession) IsSigned(_address common.Address, messageHash [32]byte, v uint8, r [32]byte, s [32]byte) (bool, error) {
	return _Identity.Contract.IsSigned(&_Identity.CallOpts, _address, messageHash, v, r, s)
}

// AddAssociatedAddress is a paid mutator transaction binding the contract method 0x3aedf3c9.
//
// Solidity: function addAssociatedAddress(address approvingAddress, address addressToAdd, uint8 v, bytes32 r, bytes32 s, uint256 timestamp) returns()
func (_Identity *IdentityTransactor) AddAssociatedAddress(opts *bind.TransactOpts, approvingAddress common.Address, addressToAdd common.Address, v uint8, r [32]byte, s [32]byte, timestamp *big.Int) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "addAssociatedAddress", approvingAddress, addressToAdd, v, r, s, timestamp)
}

// AddAssociatedAddress is a paid mutator transaction binding the contract method 0x3aedf3c9.
//
// Solidity: function addAssociatedAddress(address approvingAddress, address addressToAdd, uint8 v, bytes32 r, bytes32 s, uint256 timestamp) returns()
func (_Identity *IdentitySession) AddAssociatedAddress(approvingAddress common.Address, addressToAdd common.Address, v uint8, r [32]byte, s [32]byte, timestamp *big.Int) (*types.Transaction, error) {
	return _Identity.Contract.AddAssociatedAddress(&_Identity.TransactOpts, approvingAddress, addressToAdd, v, r, s, timestamp)
}

// AddAssociatedAddress is a paid mutator transaction binding the contract method 0x3aedf3c9.
//
// Solidity: function addAssociatedAddress(address approvingAddress, address addressToAdd, uint8 v, bytes32 r, bytes32 s, uint256 timestamp) returns()
func (_Identity *IdentityTransactorSession) AddAssociatedAddress(approvingAddress common.Address, addressToAdd common.Address, v uint8, r [32]byte, s [32]byte, timestamp *big.Int) (*types.Transaction, error) {
	return _Identity.Contract.AddAssociatedAddress(&_Identity.TransactOpts, approvingAddress, addressToAdd, v, r, s, timestamp)
}

// AddAssociatedAddressDelegated is a paid mutator transaction binding the contract method 0xab5f6781.
//
// Solidity: function addAssociatedAddressDelegated(address approvingAddress, address addressToAdd, uint8[2] v, bytes32[2] r, bytes32[2] s, uint256[2] timestamp) returns()
func (_Identity *IdentityTransactor) AddAssociatedAddressDelegated(opts *bind.TransactOpts, approvingAddress common.Address, addressToAdd common.Address, v [2]uint8, r [2][32]byte, s [2][32]byte, timestamp [2]*big.Int) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "addAssociatedAddressDelegated", approvingAddress, addressToAdd, v, r, s, timestamp)
}

// AddAssociatedAddressDelegated is a paid mutator transaction binding the contract method 0xab5f6781.
//
// Solidity: function addAssociatedAddressDelegated(address approvingAddress, address addressToAdd, uint8[2] v, bytes32[2] r, bytes32[2] s, uint256[2] timestamp) returns()
func (_Identity *IdentitySession) AddAssociatedAddressDelegated(approvingAddress common.Address, addressToAdd common.Address, v [2]uint8, r [2][32]byte, s [2][32]byte, timestamp [2]*big.Int) (*types.Transaction, error) {
	return _Identity.Contract.AddAssociatedAddressDelegated(&_Identity.TransactOpts, approvingAddress, addressToAdd, v, r, s, timestamp)
}

// AddAssociatedAddressDelegated is a paid mutator transaction binding the contract method 0xab5f6781.
//
// Solidity: function addAssociatedAddressDelegated(address approvingAddress, address addressToAdd, uint8[2] v, bytes32[2] r, bytes32[2] s, uint256[2] timestamp) returns()
func (_Identity *IdentityTransactorSession) AddAssociatedAddressDelegated(approvingAddress common.Address, addressToAdd common.Address, v [2]uint8, r [2][32]byte, s [2][32]byte, timestamp [2]*big.Int) (*types.Transaction, error) {
	return _Identity.Contract.AddAssociatedAddressDelegated(&_Identity.TransactOpts, approvingAddress, addressToAdd, v, r, s, timestamp)
}

// AddProviders is a paid mutator transaction binding the contract method 0x7d079951.
//
// Solidity: function addProviders(address[] providers) returns()
func (_Identity *IdentityTransactor) AddProviders(opts *bind.TransactOpts, providers []common.Address) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "addProviders", providers)
}

// AddProviders is a paid mutator transaction binding the contract method 0x7d079951.
//
// Solidity: function addProviders(address[] providers) returns()
func (_Identity *IdentitySession) AddProviders(providers []common.Address) (*types.Transaction, error) {
	return _Identity.Contract.AddProviders(&_Identity.TransactOpts, providers)
}

// AddProviders is a paid mutator transaction binding the contract method 0x7d079951.
//
// Solidity: function addProviders(address[] providers) returns()
func (_Identity *IdentityTransactorSession) AddProviders(providers []common.Address) (*types.Transaction, error) {
	return _Identity.Contract.AddProviders(&_Identity.TransactOpts, providers)
}

// AddProvidersFor is a paid mutator transaction binding the contract method 0x960b11f5.
//
// Solidity: function addProvidersFor(uint256 ein, address[] providers) returns()
func (_Identity *IdentityTransactor) AddProvidersFor(opts *bind.TransactOpts, ein *big.Int, providers []common.Address) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "addProvidersFor", ein, providers)
}

// AddProvidersFor is a paid mutator transaction binding the contract method 0x960b11f5.
//
// Solidity: function addProvidersFor(uint256 ein, address[] providers) returns()
func (_Identity *IdentitySession) AddProvidersFor(ein *big.Int, providers []common.Address) (*types.Transaction, error) {
	return _Identity.Contract.AddProvidersFor(&_Identity.TransactOpts, ein, providers)
}

// AddProvidersFor is a paid mutator transaction binding the contract method 0x960b11f5.
//
// Solidity: function addProvidersFor(uint256 ein, address[] providers) returns()
func (_Identity *IdentityTransactorSession) AddProvidersFor(ein *big.Int, providers []common.Address) (*types.Transaction, error) {
	return _Identity.Contract.AddProvidersFor(&_Identity.TransactOpts, ein, providers)
}

// AddResolvers is a paid mutator transaction binding the contract method 0xe846fd91.
//
// Solidity: function addResolvers(address[] resolvers) returns()
func (_Identity *IdentityTransactor) AddResolvers(opts *bind.TransactOpts, resolvers []common.Address) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "addResolvers", resolvers)
}

// AddResolvers is a paid mutator transaction binding the contract method 0xe846fd91.
//
// Solidity: function addResolvers(address[] resolvers) returns()
func (_Identity *IdentitySession) AddResolvers(resolvers []common.Address) (*types.Transaction, error) {
	return _Identity.Contract.AddResolvers(&_Identity.TransactOpts, resolvers)
}

// AddResolvers is a paid mutator transaction binding the contract method 0xe846fd91.
//
// Solidity: function addResolvers(address[] resolvers) returns()
func (_Identity *IdentityTransactorSession) AddResolvers(resolvers []common.Address) (*types.Transaction, error) {
	return _Identity.Contract.AddResolvers(&_Identity.TransactOpts, resolvers)
}

// AddResolversFor is a paid mutator transaction binding the contract method 0x010887dc.
//
// Solidity: function addResolversFor(uint256 ein, address[] resolvers) returns()
func (_Identity *IdentityTransactor) AddResolversFor(opts *bind.TransactOpts, ein *big.Int, resolvers []common.Address) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "addResolversFor", ein, resolvers)
}

// AddResolversFor is a paid mutator transaction binding the contract method 0x010887dc.
//
// Solidity: function addResolversFor(uint256 ein, address[] resolvers) returns()
func (_Identity *IdentitySession) AddResolversFor(ein *big.Int, resolvers []common.Address) (*types.Transaction, error) {
	return _Identity.Contract.AddResolversFor(&_Identity.TransactOpts, ein, resolvers)
}

// AddResolversFor is a paid mutator transaction binding the contract method 0x010887dc.
//
// Solidity: function addResolversFor(uint256 ein, address[] resolvers) returns()
func (_Identity *IdentityTransactorSession) AddResolversFor(ein *big.Int, resolvers []common.Address) (*types.Transaction, error) {
	return _Identity.Contract.AddResolversFor(&_Identity.TransactOpts, ein, resolvers)
}

// CreateIdentity is a paid mutator transaction binding the contract method 0x268e8970.
//
// Solidity: function createIdentity(address recoveryAddress, address[] providers, address[] resolvers) returns(uint256 ein)
func (_Identity *IdentityTransactor) CreateIdentity(opts *bind.TransactOpts, recoveryAddress common.Address, providers []common.Address, resolvers []common.Address) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "createIdentity", recoveryAddress, providers, resolvers)
}

// CreateIdentity is a paid mutator transaction binding the contract method 0x268e8970.
//
// Solidity: function createIdentity(address recoveryAddress, address[] providers, address[] resolvers) returns(uint256 ein)
func (_Identity *IdentitySession) CreateIdentity(recoveryAddress common.Address, providers []common.Address, resolvers []common.Address) (*types.Transaction, error) {
	return _Identity.Contract.CreateIdentity(&_Identity.TransactOpts, recoveryAddress, providers, resolvers)
}

// CreateIdentity is a paid mutator transaction binding the contract method 0x268e8970.
//
// Solidity: function createIdentity(address recoveryAddress, address[] providers, address[] resolvers) returns(uint256 ein)
func (_Identity *IdentityTransactorSession) CreateIdentity(recoveryAddress common.Address, providers []common.Address, resolvers []common.Address) (*types.Transaction, error) {
	return _Identity.Contract.CreateIdentity(&_Identity.TransactOpts, recoveryAddress, providers, resolvers)
}

// CreateIdentityDelegated is a paid mutator transaction binding the contract method 0x14fb5646.
//
// Solidity: function createIdentityDelegated(address recoveryAddress, address associatedAddress, address[] providers, address[] resolvers, uint8 v, bytes32 r, bytes32 s, uint256 timestamp) returns(uint256 ein)
func (_Identity *IdentityTransactor) CreateIdentityDelegated(opts *bind.TransactOpts, recoveryAddress common.Address, associatedAddress common.Address, providers []common.Address, resolvers []common.Address, v uint8, r [32]byte, s [32]byte, timestamp *big.Int) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "createIdentityDelegated", recoveryAddress, associatedAddress, providers, resolvers, v, r, s, timestamp)
}

// CreateIdentityDelegated is a paid mutator transaction binding the contract method 0x14fb5646.
//
// Solidity: function createIdentityDelegated(address recoveryAddress, address associatedAddress, address[] providers, address[] resolvers, uint8 v, bytes32 r, bytes32 s, uint256 timestamp) returns(uint256 ein)
func (_Identity *IdentitySession) CreateIdentityDelegated(recoveryAddress common.Address, associatedAddress common.Address, providers []common.Address, resolvers []common.Address, v uint8, r [32]byte, s [32]byte, timestamp *big.Int) (*types.Transaction, error) {
	return _Identity.Contract.CreateIdentityDelegated(&_Identity.TransactOpts, recoveryAddress, associatedAddress, providers, resolvers, v, r, s, timestamp)
}

// CreateIdentityDelegated is a paid mutator transaction binding the contract method 0x14fb5646.
//
// Solidity: function createIdentityDelegated(address recoveryAddress, address associatedAddress, address[] providers, address[] resolvers, uint8 v, bytes32 r, bytes32 s, uint256 timestamp) returns(uint256 ein)
func (_Identity *IdentityTransactorSession) CreateIdentityDelegated(recoveryAddress common.Address, associatedAddress common.Address, providers []common.Address, resolvers []common.Address, v uint8, r [32]byte, s [32]byte, timestamp *big.Int) (*types.Transaction, error) {
	return _Identity.Contract.CreateIdentityDelegated(&_Identity.TransactOpts, recoveryAddress, associatedAddress, providers, resolvers, v, r, s, timestamp)
}

// RemoveAssociatedAddress is a paid mutator transaction binding the contract method 0x3b4656a1.
//
// Solidity: function removeAssociatedAddress() returns()
func (_Identity *IdentityTransactor) RemoveAssociatedAddress(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "removeAssociatedAddress")
}

// RemoveAssociatedAddress is a paid mutator transaction binding the contract method 0x3b4656a1.
//
// Solidity: function removeAssociatedAddress() returns()
func (_Identity *IdentitySession) RemoveAssociatedAddress() (*types.Transaction, error) {
	return _Identity.Contract.RemoveAssociatedAddress(&_Identity.TransactOpts)
}

// RemoveAssociatedAddress is a paid mutator transaction binding the contract method 0x3b4656a1.
//
// Solidity: function removeAssociatedAddress() returns()
func (_Identity *IdentityTransactorSession) RemoveAssociatedAddress() (*types.Transaction, error) {
	return _Identity.Contract.RemoveAssociatedAddress(&_Identity.TransactOpts)
}

// RemoveAssociatedAddressDelegated is a paid mutator transaction binding the contract method 0xdd9ad9ed.
//
// Solidity: function removeAssociatedAddressDelegated(address addressToRemove, uint8 v, bytes32 r, bytes32 s, uint256 timestamp) returns()
func (_Identity *IdentityTransactor) RemoveAssociatedAddressDelegated(opts *bind.TransactOpts, addressToRemove common.Address, v uint8, r [32]byte, s [32]byte, timestamp *big.Int) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "removeAssociatedAddressDelegated", addressToRemove, v, r, s, timestamp)
}

// RemoveAssociatedAddressDelegated is a paid mutator transaction binding the contract method 0xdd9ad9ed.
//
// Solidity: function removeAssociatedAddressDelegated(address addressToRemove, uint8 v, bytes32 r, bytes32 s, uint256 timestamp) returns()
func (_Identity *IdentitySession) RemoveAssociatedAddressDelegated(addressToRemove common.Address, v uint8, r [32]byte, s [32]byte, timestamp *big.Int) (*types.Transaction, error) {
	return _Identity.Contract.RemoveAssociatedAddressDelegated(&_Identity.TransactOpts, addressToRemove, v, r, s, timestamp)
}

// RemoveAssociatedAddressDelegated is a paid mutator transaction binding the contract method 0xdd9ad9ed.
//
// Solidity: function removeAssociatedAddressDelegated(address addressToRemove, uint8 v, bytes32 r, bytes32 s, uint256 timestamp) returns()
func (_Identity *IdentityTransactorSession) RemoveAssociatedAddressDelegated(addressToRemove common.Address, v uint8, r [32]byte, s [32]byte, timestamp *big.Int) (*types.Transaction, error) {
	return _Identity.Contract.RemoveAssociatedAddressDelegated(&_Identity.TransactOpts, addressToRemove, v, r, s, timestamp)
}

// RemoveProviders is a paid mutator transaction binding the contract method 0x232e7d83.
//
// Solidity: function removeProviders(address[] providers) returns()
func (_Identity *IdentityTransactor) RemoveProviders(opts *bind.TransactOpts, providers []common.Address) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "removeProviders", providers)
}

// RemoveProviders is a paid mutator transaction binding the contract method 0x232e7d83.
//
// Solidity: function removeProviders(address[] providers) returns()
func (_Identity *IdentitySession) RemoveProviders(providers []common.Address) (*types.Transaction, error) {
	return _Identity.Contract.RemoveProviders(&_Identity.TransactOpts, providers)
}

// RemoveProviders is a paid mutator transaction binding the contract method 0x232e7d83.
//
// Solidity: function removeProviders(address[] providers) returns()
func (_Identity *IdentityTransactorSession) RemoveProviders(providers []common.Address) (*types.Transaction, error) {
	return _Identity.Contract.RemoveProviders(&_Identity.TransactOpts, providers)
}

// RemoveProvidersFor is a paid mutator transaction binding the contract method 0x06c93c36.
//
// Solidity: function removeProvidersFor(uint256 ein, address[] providers) returns()
func (_Identity *IdentityTransactor) RemoveProvidersFor(opts *bind.TransactOpts, ein *big.Int, providers []common.Address) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "removeProvidersFor", ein, providers)
}

// RemoveProvidersFor is a paid mutator transaction binding the contract method 0x06c93c36.
//
// Solidity: function removeProvidersFor(uint256 ein, address[] providers) returns()
func (_Identity *IdentitySession) RemoveProvidersFor(ein *big.Int, providers []common.Address) (*types.Transaction, error) {
	return _Identity.Contract.RemoveProvidersFor(&_Identity.TransactOpts, ein, providers)
}

// RemoveProvidersFor is a paid mutator transaction binding the contract method 0x06c93c36.
//
// Solidity: function removeProvidersFor(uint256 ein, address[] providers) returns()
func (_Identity *IdentityTransactorSession) RemoveProvidersFor(ein *big.Int, providers []common.Address) (*types.Transaction, error) {
	return _Identity.Contract.RemoveProvidersFor(&_Identity.TransactOpts, ein, providers)
}

// RemoveResolvers is a paid mutator transaction binding the contract method 0xb8a5c2a6.
//
// Solidity: function removeResolvers(address[] resolvers) returns()
func (_Identity *IdentityTransactor) RemoveResolvers(opts *bind.TransactOpts, resolvers []common.Address) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "removeResolvers", resolvers)
}

// RemoveResolvers is a paid mutator transaction binding the contract method 0xb8a5c2a6.
//
// Solidity: function removeResolvers(address[] resolvers) returns()
func (_Identity *IdentitySession) RemoveResolvers(resolvers []common.Address) (*types.Transaction, error) {
	return _Identity.Contract.RemoveResolvers(&_Identity.TransactOpts, resolvers)
}

// RemoveResolvers is a paid mutator transaction binding the contract method 0xb8a5c2a6.
//
// Solidity: function removeResolvers(address[] resolvers) returns()
func (_Identity *IdentityTransactorSession) RemoveResolvers(resolvers []common.Address) (*types.Transaction, error) {
	return _Identity.Contract.RemoveResolvers(&_Identity.TransactOpts, resolvers)
}

// RemoveResolversFor is a paid mutator transaction binding the contract method 0x5726f660.
//
// Solidity: function removeResolversFor(uint256 ein, address[] resolvers) returns()
func (_Identity *IdentityTransactor) RemoveResolversFor(opts *bind.TransactOpts, ein *big.Int, resolvers []common.Address) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "removeResolversFor", ein, resolvers)
}

// RemoveResolversFor is a paid mutator transaction binding the contract method 0x5726f660.
//
// Solidity: function removeResolversFor(uint256 ein, address[] resolvers) returns()
func (_Identity *IdentitySession) RemoveResolversFor(ein *big.Int, resolvers []common.Address) (*types.Transaction, error) {
	return _Identity.Contract.RemoveResolversFor(&_Identity.TransactOpts, ein, resolvers)
}

// RemoveResolversFor is a paid mutator transaction binding the contract method 0x5726f660.
//
// Solidity: function removeResolversFor(uint256 ein, address[] resolvers) returns()
func (_Identity *IdentityTransactorSession) RemoveResolversFor(ein *big.Int, resolvers []common.Address) (*types.Transaction, error) {
	return _Identity.Contract.RemoveResolversFor(&_Identity.TransactOpts, ein, resolvers)
}

// TriggerDestruction is a paid mutator transaction binding the contract method 0xfa57b576.
//
// Solidity: function triggerDestruction(uint256 ein, address[] firstChunk, address[] lastChunk, bool resetResolvers) returns()
func (_Identity *IdentityTransactor) TriggerDestruction(opts *bind.TransactOpts, ein *big.Int, firstChunk []common.Address, lastChunk []common.Address, resetResolvers bool) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "triggerDestruction", ein, firstChunk, lastChunk, resetResolvers)
}

// TriggerDestruction is a paid mutator transaction binding the contract method 0xfa57b576.
//
// Solidity: function triggerDestruction(uint256 ein, address[] firstChunk, address[] lastChunk, bool resetResolvers) returns()
func (_Identity *IdentitySession) TriggerDestruction(ein *big.Int, firstChunk []common.Address, lastChunk []common.Address, resetResolvers bool) (*types.Transaction, error) {
	return _Identity.Contract.TriggerDestruction(&_Identity.TransactOpts, ein, firstChunk, lastChunk, resetResolvers)
}

// TriggerDestruction is a paid mutator transaction binding the contract method 0xfa57b576.
//
// Solidity: function triggerDestruction(uint256 ein, address[] firstChunk, address[] lastChunk, bool resetResolvers) returns()
func (_Identity *IdentityTransactorSession) TriggerDestruction(ein *big.Int, firstChunk []common.Address, lastChunk []common.Address, resetResolvers bool) (*types.Transaction, error) {
	return _Identity.Contract.TriggerDestruction(&_Identity.TransactOpts, ein, firstChunk, lastChunk, resetResolvers)
}

// TriggerRecovery is a paid mutator transaction binding the contract method 0x8e1bb633.
//
// Solidity: function triggerRecovery(uint256 ein, address newAssociatedAddress, uint8 v, bytes32 r, bytes32 s, uint256 timestamp) returns()
func (_Identity *IdentityTransactor) TriggerRecovery(opts *bind.TransactOpts, ein *big.Int, newAssociatedAddress common.Address, v uint8, r [32]byte, s [32]byte, timestamp *big.Int) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "triggerRecovery", ein, newAssociatedAddress, v, r, s, timestamp)
}

// TriggerRecovery is a paid mutator transaction binding the contract method 0x8e1bb633.
//
// Solidity: function triggerRecovery(uint256 ein, address newAssociatedAddress, uint8 v, bytes32 r, bytes32 s, uint256 timestamp) returns()
func (_Identity *IdentitySession) TriggerRecovery(ein *big.Int, newAssociatedAddress common.Address, v uint8, r [32]byte, s [32]byte, timestamp *big.Int) (*types.Transaction, error) {
	return _Identity.Contract.TriggerRecovery(&_Identity.TransactOpts, ein, newAssociatedAddress, v, r, s, timestamp)
}

// TriggerRecovery is a paid mutator transaction binding the contract method 0x8e1bb633.
//
// Solidity: function triggerRecovery(uint256 ein, address newAssociatedAddress, uint8 v, bytes32 r, bytes32 s, uint256 timestamp) returns()
func (_Identity *IdentityTransactorSession) TriggerRecovery(ein *big.Int, newAssociatedAddress common.Address, v uint8, r [32]byte, s [32]byte, timestamp *big.Int) (*types.Transaction, error) {
	return _Identity.Contract.TriggerRecovery(&_Identity.TransactOpts, ein, newAssociatedAddress, v, r, s, timestamp)
}

// TriggerRecoveryAddressChange is a paid mutator transaction binding the contract method 0x1bfe3508.
//
// Solidity: function triggerRecoveryAddressChange(address newRecoveryAddress) returns()
func (_Identity *IdentityTransactor) TriggerRecoveryAddressChange(opts *bind.TransactOpts, newRecoveryAddress common.Address) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "triggerRecoveryAddressChange", newRecoveryAddress)
}

// TriggerRecoveryAddressChange is a paid mutator transaction binding the contract method 0x1bfe3508.
//
// Solidity: function triggerRecoveryAddressChange(address newRecoveryAddress) returns()
func (_Identity *IdentitySession) TriggerRecoveryAddressChange(newRecoveryAddress common.Address) (*types.Transaction, error) {
	return _Identity.Contract.TriggerRecoveryAddressChange(&_Identity.TransactOpts, newRecoveryAddress)
}

// TriggerRecoveryAddressChange is a paid mutator transaction binding the contract method 0x1bfe3508.
//
// Solidity: function triggerRecoveryAddressChange(address newRecoveryAddress) returns()
func (_Identity *IdentityTransactorSession) TriggerRecoveryAddressChange(newRecoveryAddress common.Address) (*types.Transaction, error) {
	return _Identity.Contract.TriggerRecoveryAddressChange(&_Identity.TransactOpts, newRecoveryAddress)
}

// TriggerRecoveryAddressChangeFor is a paid mutator transaction binding the contract method 0x2501faa5.
//
// Solidity: function triggerRecoveryAddressChangeFor(uint256 ein, address newRecoveryAddress) returns()
func (_Identity *IdentityTransactor) TriggerRecoveryAddressChangeFor(opts *bind.TransactOpts, ein *big.Int, newRecoveryAddress common.Address) (*types.Transaction, error) {
	return _Identity.contract.Transact(opts, "triggerRecoveryAddressChangeFor", ein, newRecoveryAddress)
}

// TriggerRecoveryAddressChangeFor is a paid mutator transaction binding the contract method 0x2501faa5.
//
// Solidity: function triggerRecoveryAddressChangeFor(uint256 ein, address newRecoveryAddress) returns()
func (_Identity *IdentitySession) TriggerRecoveryAddressChangeFor(ein *big.Int, newRecoveryAddress common.Address) (*types.Transaction, error) {
	return _Identity.Contract.TriggerRecoveryAddressChangeFor(&_Identity.TransactOpts, ein, newRecoveryAddress)
}

// TriggerRecoveryAddressChangeFor is a paid mutator transaction binding the contract method 0x2501faa5.
//
// Solidity: function triggerRecoveryAddressChangeFor(uint256 ein, address newRecoveryAddress) returns()
func (_Identity *IdentityTransactorSession) TriggerRecoveryAddressChangeFor(ein *big.Int, newRecoveryAddress common.Address) (*types.Transaction, error) {
	return _Identity.Contract.TriggerRecoveryAddressChangeFor(&_Identity.TransactOpts, ein, newRecoveryAddress)
}
