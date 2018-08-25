// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package main

import (
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

// HelloABI is the input ABI used to generate the binding from.
const HelloABI = "[{\"name\":\"__init__\",\"outputs\":[],\"inputs\":[{\"type\":\"int128\",\"name\":\"_startval\"}],\"constant\":false,\"payable\":false,\"type\":\"constructor\"},{\"name\":\"increment\",\"outputs\":[],\"inputs\":[],\"constant\":false,\"payable\":false,\"type\":\"function\",\"gas\":35562},{\"name\":\"nb\",\"outputs\":[{\"type\":\"int128\",\"name\":\"out\"}],\"inputs\":[],\"constant\":true,\"payable\":false,\"type\":\"function\",\"gas\":513}]"

// HelloBin is the compiled bytecode used for deploying new contracts.
const HelloBin = `0x600035601c52740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a05260206101fc6101403934156100a757600080fd5b60605160206101fc60c03960c051806040519013156100c557600080fd5b80919012156100d357600080fd5b50610140516000556101e456600035601c52740100000000000000000000000000000000000000006020526f7fffffffffffffffffffffffffffffff6040527fffffffffffffffffffffffffffffffff8000000000000000000000000000000060605274012a05f1fffffffffffffffffffffffffdabf41c006080527ffffffffffffffffffffffffed5fa0e000000000000000000000000000000000060a05263d09de08a60005114156100d85734156100ac57600080fd5b606051600160005401806040519013156100c557600080fd5b80919012156100d357600080fd5b600055005b63d2c00e2460005114156100fe5734156100f157600080fd5b60005460005260206000f3005b60006000fd5b6100e06101e4036100e06000396100e06101e4036000f3`

// DeployHello deploys a new Ethereum contract, binding an instance of Hello to it.
func DeployHello(auth *bind.TransactOpts, backend bind.ContractBackend, _startval *big.Int) (common.Address, *types.Transaction, *Hello, error) {
	parsed, err := abi.JSON(strings.NewReader(HelloABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HelloBin), backend, _startval)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Hello{HelloCaller: HelloCaller{contract: contract}, HelloTransactor: HelloTransactor{contract: contract}, HelloFilterer: HelloFilterer{contract: contract}}, nil
}

// Hello is an auto generated Go binding around an Ethereum contract.
type Hello struct {
	HelloCaller     // Read-only binding to the contract
	HelloTransactor // Write-only binding to the contract
	HelloFilterer   // Log filterer for contract events
}

// HelloCaller is an auto generated read-only Go binding around an Ethereum contract.
type HelloCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HelloTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HelloFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HelloSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HelloSession struct {
	Contract     *Hello            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HelloCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HelloCallerSession struct {
	Contract *HelloCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// HelloTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HelloTransactorSession struct {
	Contract     *HelloTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HelloRaw is an auto generated low-level Go binding around an Ethereum contract.
type HelloRaw struct {
	Contract *Hello // Generic contract binding to access the raw methods on
}

// HelloCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HelloCallerRaw struct {
	Contract *HelloCaller // Generic read-only contract binding to access the raw methods on
}

// HelloTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HelloTransactorRaw struct {
	Contract *HelloTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHello creates a new instance of Hello, bound to a specific deployed contract.
func NewHello(address common.Address, backend bind.ContractBackend) (*Hello, error) {
	contract, err := bindHello(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Hello{HelloCaller: HelloCaller{contract: contract}, HelloTransactor: HelloTransactor{contract: contract}, HelloFilterer: HelloFilterer{contract: contract}}, nil
}

// NewHelloCaller creates a new read-only instance of Hello, bound to a specific deployed contract.
func NewHelloCaller(address common.Address, caller bind.ContractCaller) (*HelloCaller, error) {
	contract, err := bindHello(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HelloCaller{contract: contract}, nil
}

// NewHelloTransactor creates a new write-only instance of Hello, bound to a specific deployed contract.
func NewHelloTransactor(address common.Address, transactor bind.ContractTransactor) (*HelloTransactor, error) {
	contract, err := bindHello(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HelloTransactor{contract: contract}, nil
}

// NewHelloFilterer creates a new log filterer instance of Hello, bound to a specific deployed contract.
func NewHelloFilterer(address common.Address, filterer bind.ContractFilterer) (*HelloFilterer, error) {
	contract, err := bindHello(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HelloFilterer{contract: contract}, nil
}

// bindHello binds a generic wrapper to an already deployed contract.
func bindHello(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HelloABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hello *HelloRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Hello.Contract.HelloCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hello *HelloRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hello.Contract.HelloTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hello *HelloRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hello.Contract.HelloTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Hello *HelloCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Hello.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Hello *HelloTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hello.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Hello *HelloTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Hello.Contract.contract.Transact(opts, method, params...)
}

// Nb is a free data retrieval call binding the contract method 0xd2c00e24.
//
// Solidity: function nb() constant returns(out int128)
func (_Hello *HelloCaller) Nb(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _Hello.contract.Call(opts, out, "nb")
	return *ret0, err
}

// Nb is a free data retrieval call binding the contract method 0xd2c00e24.
//
// Solidity: function nb() constant returns(out int128)
func (_Hello *HelloSession) Nb() (*big.Int, error) {
	return _Hello.Contract.Nb(&_Hello.CallOpts)
}

// Nb is a free data retrieval call binding the contract method 0xd2c00e24.
//
// Solidity: function nb() constant returns(out int128)
func (_Hello *HelloCallerSession) Nb() (*big.Int, error) {
	return _Hello.Contract.Nb(&_Hello.CallOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_Hello *HelloTransactor) Increment(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Hello.contract.Transact(opts, "increment")
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_Hello *HelloSession) Increment() (*types.Transaction, error) {
	return _Hello.Contract.Increment(&_Hello.TransactOpts)
}

// Increment is a paid mutator transaction binding the contract method 0xd09de08a.
//
// Solidity: function increment() returns()
func (_Hello *HelloTransactorSession) Increment() (*types.Transaction, error) {
	return _Hello.Contract.Increment(&_Hello.TransactOpts)
}
