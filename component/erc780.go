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

// MainABI is the input ABI used to generate the binding from.
const MainABI = "[{\"constant\":true,\"inputs\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"getClaim\",\"outputs\":[{\"internalType\":\"bytes32\",\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"issuer\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"}],\"name\":\"removeClaim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"address\",\"name\":\"subject\",\"type\":\"address\"},{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"setClaim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"internalType\":\"bytes32\",\"name\":\"key\",\"type\":\"bytes32\"},{\"internalType\":\"bytes32\",\"name\":\"value\",\"type\":\"bytes32\"}],\"name\":\"setSelfClaim\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// Main is an auto generated Go binding around an Ethereum contract.
type Main struct {
	MainCaller     // Read-only binding to the contract
	MainTransactor // Write-only binding to the contract
	MainFilterer   // Log filterer for contract events
}

// MainCaller is an auto generated read-only Go binding around an Ethereum contract.
type MainCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainTransactor is an auto generated write-only Go binding around an Ethereum contract.
type MainTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type MainFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// MainSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type MainSession struct {
	Contract     *Main             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type MainCallerSession struct {
	Contract *MainCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// MainTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type MainTransactorSession struct {
	Contract     *MainTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// MainRaw is an auto generated low-level Go binding around an Ethereum contract.
type MainRaw struct {
	Contract *Main // Generic contract binding to access the raw methods on
}

// MainCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type MainCallerRaw struct {
	Contract *MainCaller // Generic read-only contract binding to access the raw methods on
}

// MainTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type MainTransactorRaw struct {
	Contract *MainTransactor // Generic write-only contract binding to access the raw methods on
}

// NewMain creates a new instance of Main, bound to a specific deployed contract.
func NewMain(address common.Address, backend bind.ContractBackend) (*Main, error) {
	contract, err := bindMain(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Main{MainCaller: MainCaller{contract: contract}, MainTransactor: MainTransactor{contract: contract}, MainFilterer: MainFilterer{contract: contract}}, nil
}

// NewMainCaller creates a new read-only instance of Main, bound to a specific deployed contract.
func NewMainCaller(address common.Address, caller bind.ContractCaller) (*MainCaller, error) {
	contract, err := bindMain(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &MainCaller{contract: contract}, nil
}

// NewMainTransactor creates a new write-only instance of Main, bound to a specific deployed contract.
func NewMainTransactor(address common.Address, transactor bind.ContractTransactor) (*MainTransactor, error) {
	contract, err := bindMain(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &MainTransactor{contract: contract}, nil
}

// NewMainFilterer creates a new log filterer instance of Main, bound to a specific deployed contract.
func NewMainFilterer(address common.Address, filterer bind.ContractFilterer) (*MainFilterer, error) {
	contract, err := bindMain(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &MainFilterer{contract: contract}, nil
}

// bindMain binds a generic wrapper to an already deployed contract.
func bindMain(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(MainABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Main.Contract.MainCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.MainTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Main *MainCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Main.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Main *MainTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Main.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Main *MainTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Main.Contract.contract.Transact(opts, method, params...)
}

// GetClaim is a free data retrieval call binding the contract method 0xe1661eff.
//
// Solidity: function getClaim(address issuer, address subject, bytes32 key) constant returns(bytes32)
func (_Main *MainCaller) GetClaim(opts *bind.CallOpts, issuer common.Address, subject common.Address, key [32]byte) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _Main.contract.Call(opts, out, "getClaim", issuer, subject, key)
	return *ret0, err
}

// GetClaim is a free data retrieval call binding the contract method 0xe1661eff.
//
// Solidity: function getClaim(address issuer, address subject, bytes32 key) constant returns(bytes32)
func (_Main *MainSession) GetClaim(issuer common.Address, subject common.Address, key [32]byte) ([32]byte, error) {
	return _Main.Contract.GetClaim(&_Main.CallOpts, issuer, subject, key)
}

// GetClaim is a free data retrieval call binding the contract method 0xe1661eff.
//
// Solidity: function getClaim(address issuer, address subject, bytes32 key) constant returns(bytes32)
func (_Main *MainCallerSession) GetClaim(issuer common.Address, subject common.Address, key [32]byte) ([32]byte, error) {
	return _Main.Contract.GetClaim(&_Main.CallOpts, issuer, subject, key)
}

// RemoveClaim is a paid mutator transaction binding the contract method 0xc7508ec7.
//
// Solidity: function removeClaim(address issuer, address subject, bytes32 key) returns()
func (_Main *MainTransactor) RemoveClaim(opts *bind.TransactOpts, issuer common.Address, subject common.Address, key [32]byte) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "removeClaim", issuer, subject, key)
}

// RemoveClaim is a paid mutator transaction binding the contract method 0xc7508ec7.
//
// Solidity: function removeClaim(address issuer, address subject, bytes32 key) returns()
func (_Main *MainSession) RemoveClaim(issuer common.Address, subject common.Address, key [32]byte) (*types.Transaction, error) {
	return _Main.Contract.RemoveClaim(&_Main.TransactOpts, issuer, subject, key)
}

// RemoveClaim is a paid mutator transaction binding the contract method 0xc7508ec7.
//
// Solidity: function removeClaim(address issuer, address subject, bytes32 key) returns()
func (_Main *MainTransactorSession) RemoveClaim(issuer common.Address, subject common.Address, key [32]byte) (*types.Transaction, error) {
	return _Main.Contract.RemoveClaim(&_Main.TransactOpts, issuer, subject, key)
}

// SetClaim is a paid mutator transaction binding the contract method 0x9918925d.
//
// Solidity: function setClaim(address subject, bytes32 key, bytes32 value) returns()
func (_Main *MainTransactor) SetClaim(opts *bind.TransactOpts, subject common.Address, key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setClaim", subject, key, value)
}

// SetClaim is a paid mutator transaction binding the contract method 0x9918925d.
//
// Solidity: function setClaim(address subject, bytes32 key, bytes32 value) returns()
func (_Main *MainSession) SetClaim(subject common.Address, key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Main.Contract.SetClaim(&_Main.TransactOpts, subject, key, value)
}

// SetClaim is a paid mutator transaction binding the contract method 0x9918925d.
//
// Solidity: function setClaim(address subject, bytes32 key, bytes32 value) returns()
func (_Main *MainTransactorSession) SetClaim(subject common.Address, key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Main.Contract.SetClaim(&_Main.TransactOpts, subject, key, value)
}

// SetSelfClaim is a paid mutator transaction binding the contract method 0x9155b01a.
//
// Solidity: function setSelfClaim(bytes32 key, bytes32 value) returns()
func (_Main *MainTransactor) SetSelfClaim(opts *bind.TransactOpts, key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Main.contract.Transact(opts, "setSelfClaim", key, value)
}

// SetSelfClaim is a paid mutator transaction binding the contract method 0x9155b01a.
//
// Solidity: function setSelfClaim(bytes32 key, bytes32 value) returns()
func (_Main *MainSession) SetSelfClaim(key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Main.Contract.SetSelfClaim(&_Main.TransactOpts, key, value)
}

// SetSelfClaim is a paid mutator transaction binding the contract method 0x9155b01a.
//
// Solidity: function setSelfClaim(bytes32 key, bytes32 value) returns()
func (_Main *MainTransactorSession) SetSelfClaim(key [32]byte, value [32]byte) (*types.Transaction, error) {
	return _Main.Contract.SetSelfClaim(&_Main.TransactOpts, key, value)
}
