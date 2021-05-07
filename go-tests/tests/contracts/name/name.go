// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Name

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
	_ = bind.Bind
	_ = common.Big1
	_ = types.BloomLookup
	_ = event.NewSubscription
)

// NameABI is the input ABI used to generate the binding from.
const NameABI = "[{\"inputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[],\"name\":\"getName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"newName\",\"type\":\"string\"}],\"name\":\"setName\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

// NameBin is the compiled bytecode used for deploying new contracts.
var NameBin = "0x608060405234801561001057600080fd5b506040518060400160405280600581526020017f48656c6c6f0000000000000000000000000000000000000000000000000000008152506000908051906020019061005c929190610062565b50610107565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100a357805160ff19168380011785556100d1565b828001600101855582156100d1579182015b828111156100d05782518255916020019190600101906100b5565b5b5090506100de91906100e2565b5090565b61010491905b808211156101005760008160009055506001016100e8565b5090565b90565b610310806101166000396000f3fe608060405234801561001057600080fd5b50600436106100365760003560e01c806317d7de7c1461003b578063c47f0027146100be575b600080fd5b610043610179565b6040518080602001828103825283818151815260200191508051906020019080838360005b83811015610083578082015181840152602081019050610068565b50505050905090810190601f1680156100b05780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b610177600480360360208110156100d457600080fd5b81019080803590602001906401000000008111156100f157600080fd5b82018360208201111561010357600080fd5b8035906020019184600183028401116401000000008311171561012557600080fd5b91908080601f016020809104026020016040519081016040528093929190818152602001838380828437600081840152601f19601f82011690508083019250505050505050919291929050505061021b565b005b606060008054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156102115780601f106101e657610100808354040283529160200191610211565b820191906000526020600020905b8154815290600101906020018083116101f457829003601f168201915b5050505050905090565b8060009080519060200190610231929190610235565b5050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061027657805160ff19168380011785556102a4565b828001600101855582156102a4579182015b828111156102a3578251825591602001919060010190610288565b5b5090506102b191906102b5565b5090565b6102d791905b808211156102d35760008160009055506001016102bb565b5090565b9056fea26469706673582212202cdf45b5ccc8f0b5de17d6bf56ed662fcd7adc0587b4c42a236394b4ae34d02164736f6c634300060b0033"

// DeployName deploys a new Ethereum contract, binding an instance of Name to it.
func DeployName(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Name, error) {
	parsed, err := abi.JSON(strings.NewReader(NameABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(NameBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Name{NameCaller: NameCaller{contract: contract}, NameTransactor: NameTransactor{contract: contract}, NameFilterer: NameFilterer{contract: contract}}, nil
}

// Name is an auto generated Go binding around an Ethereum contract.
type Name struct {
	NameCaller     // Read-only binding to the contract
	NameTransactor // Write-only binding to the contract
	NameFilterer   // Log filterer for contract events
}

// NameCaller is an auto generated read-only Go binding around an Ethereum contract.
type NameCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NameTransactor is an auto generated write-only Go binding around an Ethereum contract.
type NameTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NameFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type NameFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// NameSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type NameSession struct {
	Contract     *Name             // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NameCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type NameCallerSession struct {
	Contract *NameCaller   // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// NameTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type NameTransactorSession struct {
	Contract     *NameTransactor   // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// NameRaw is an auto generated low-level Go binding around an Ethereum contract.
type NameRaw struct {
	Contract *Name // Generic contract binding to access the raw methods on
}

// NameCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type NameCallerRaw struct {
	Contract *NameCaller // Generic read-only contract binding to access the raw methods on
}

// NameTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type NameTransactorRaw struct {
	Contract *NameTransactor // Generic write-only contract binding to access the raw methods on
}

// NewName creates a new instance of Name, bound to a specific deployed contract.
func NewName(address common.Address, backend bind.ContractBackend) (*Name, error) {
	contract, err := bindName(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Name{NameCaller: NameCaller{contract: contract}, NameTransactor: NameTransactor{contract: contract}, NameFilterer: NameFilterer{contract: contract}}, nil
}

// NewNameCaller creates a new read-only instance of Name, bound to a specific deployed contract.
func NewNameCaller(address common.Address, caller bind.ContractCaller) (*NameCaller, error) {
	contract, err := bindName(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &NameCaller{contract: contract}, nil
}

// NewNameTransactor creates a new write-only instance of Name, bound to a specific deployed contract.
func NewNameTransactor(address common.Address, transactor bind.ContractTransactor) (*NameTransactor, error) {
	contract, err := bindName(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &NameTransactor{contract: contract}, nil
}

// NewNameFilterer creates a new log filterer instance of Name, bound to a specific deployed contract.
func NewNameFilterer(address common.Address, filterer bind.ContractFilterer) (*NameFilterer, error) {
	contract, err := bindName(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &NameFilterer{contract: contract}, nil
}

// bindName binds a generic wrapper to an already deployed contract.
func bindName(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(NameABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Name *NameRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Name.Contract.NameCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Name *NameRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Name.Contract.NameTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Name *NameRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Name.Contract.NameTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Name *NameCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Name.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Name *NameTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Name.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Name *NameTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Name.Contract.contract.Transact(opts, method, params...)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_Name *NameCaller) GetName(opts *bind.CallOpts) (string, error) {
	var out []interface{}
	err := _Name.contract.Call(opts, &out, "getName")

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_Name *NameSession) GetName() (string, error) {
	return _Name.Contract.GetName(&_Name.CallOpts)
}

// GetName is a free data retrieval call binding the contract method 0x17d7de7c.
//
// Solidity: function getName() view returns(string)
func (_Name *NameCallerSession) GetName() (string, error) {
	return _Name.Contract.GetName(&_Name.CallOpts)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string newName) returns()
func (_Name *NameTransactor) SetName(opts *bind.TransactOpts, newName string) (*types.Transaction, error) {
	return _Name.contract.Transact(opts, "setName", newName)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string newName) returns()
func (_Name *NameSession) SetName(newName string) (*types.Transaction, error) {
	return _Name.Contract.SetName(&_Name.TransactOpts, newName)
}

// SetName is a paid mutator transaction binding the contract method 0xc47f0027.
//
// Solidity: function setName(string newName) returns()
func (_Name *NameTransactorSession) SetName(newName string) (*types.Transaction, error) {
	return _Name.Contract.SetName(&_Name.TransactOpts, newName)
}
