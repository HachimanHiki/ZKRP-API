// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package health

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

// HealthABI is the input ABI used to generate the binding from.
const HealthABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"c\",\"type\":\"string\"}],\"name\":\"updateCommitment\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"commitment\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"m\",\"type\":\"string\"}],\"name\":\"updateMerkleTreeRoot\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"merkleTreeRoot\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"c\",\"type\":\"string\"},{\"name\":\"m\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"}]"

// HealthBin is the compiled bytecode used for deploying new contracts.
var HealthBin = "0x608060405234801561001057600080fd5b506040516105c03803806105c083398101806040528101908080518201929190602001805182019291905050508160009080519060200190610053929190610072565b50806001908051906020019061006a929190610072565b505050610117565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100b357805160ff19168380011785556100e1565b828001600101855582156100e1579182015b828111156100e05782518255916020019190600101906100c5565b5b5090506100ee91906100f2565b5090565b61011491905b808211156101105760008160009055506001016100f8565b5090565b90565b61049a806101266000396000f300608060405260043610610062576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630af3bec2146100675780631303a484146100d0578063206b3acb14610160578063f716aee9146101c9575b600080fd5b34801561007357600080fd5b506100ce600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610259565b005b3480156100dc57600080fd5b506100e5610273565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561012557808201518184015260208101905061010a565b50505050905090810190601f1680156101525780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561016c57600080fd5b506101c7600480360381019080803590602001908201803590602001908080601f0160208091040260200160405190810160405280939291908181526020018383808284378201915050505050509192919290505050610311565b005b3480156101d557600080fd5b506101de61032b565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561021e578082015181840152602081019050610203565b50505050905090810190601f16801561024b5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b806000908051906020019061026f9291906103c9565b5050565b60008054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103095780601f106102de57610100808354040283529160200191610309565b820191906000526020600020905b8154815290600101906020018083116102ec57829003601f168201915b505050505081565b80600190805190602001906103279291906103c9565b5050565b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156103c15780601f10610396576101008083540402835291602001916103c1565b820191906000526020600020905b8154815290600101906020018083116103a457829003601f168201915b505050505081565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061040a57805160ff1916838001178555610438565b82800160010185558215610438579182015b8281111561043757825182559160200191906001019061041c565b5b5090506104459190610449565b5090565b61046b91905b8082111561046757600081600090555060010161044f565b5090565b905600a165627a7a723058206c7c9b51470a3bbb565509147781026585d78c946021c86f913e0db34846611e0029"

// DeployHealth deploys a new Ethereum contract, binding an instance of Health to it.
func DeployHealth(auth *bind.TransactOpts, backend bind.ContractBackend, c string, m string) (common.Address, *types.Transaction, *Health, error) {
	parsed, err := abi.JSON(strings.NewReader(HealthABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}

	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(HealthBin), backend, c, m)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Health{HealthCaller: HealthCaller{contract: contract}, HealthTransactor: HealthTransactor{contract: contract}, HealthFilterer: HealthFilterer{contract: contract}}, nil
}

// Health is an auto generated Go binding around an Ethereum contract.
type Health struct {
	HealthCaller     // Read-only binding to the contract
	HealthTransactor // Write-only binding to the contract
	HealthFilterer   // Log filterer for contract events
}

// HealthCaller is an auto generated read-only Go binding around an Ethereum contract.
type HealthCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HealthTransactor is an auto generated write-only Go binding around an Ethereum contract.
type HealthTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HealthFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type HealthFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// HealthSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type HealthSession struct {
	Contract     *Health           // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HealthCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type HealthCallerSession struct {
	Contract *HealthCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// HealthTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type HealthTransactorSession struct {
	Contract     *HealthTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// HealthRaw is an auto generated low-level Go binding around an Ethereum contract.
type HealthRaw struct {
	Contract *Health // Generic contract binding to access the raw methods on
}

// HealthCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type HealthCallerRaw struct {
	Contract *HealthCaller // Generic read-only contract binding to access the raw methods on
}

// HealthTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type HealthTransactorRaw struct {
	Contract *HealthTransactor // Generic write-only contract binding to access the raw methods on
}

// NewHealth creates a new instance of Health, bound to a specific deployed contract.
func NewHealth(address common.Address, backend bind.ContractBackend) (*Health, error) {
	contract, err := bindHealth(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Health{HealthCaller: HealthCaller{contract: contract}, HealthTransactor: HealthTransactor{contract: contract}, HealthFilterer: HealthFilterer{contract: contract}}, nil
}

// NewHealthCaller creates a new read-only instance of Health, bound to a specific deployed contract.
func NewHealthCaller(address common.Address, caller bind.ContractCaller) (*HealthCaller, error) {
	contract, err := bindHealth(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &HealthCaller{contract: contract}, nil
}

// NewHealthTransactor creates a new write-only instance of Health, bound to a specific deployed contract.
func NewHealthTransactor(address common.Address, transactor bind.ContractTransactor) (*HealthTransactor, error) {
	contract, err := bindHealth(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &HealthTransactor{contract: contract}, nil
}

// NewHealthFilterer creates a new log filterer instance of Health, bound to a specific deployed contract.
func NewHealthFilterer(address common.Address, filterer bind.ContractFilterer) (*HealthFilterer, error) {
	contract, err := bindHealth(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &HealthFilterer{contract: contract}, nil
}

// bindHealth binds a generic wrapper to an already deployed contract.
func bindHealth(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(HealthABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Health *HealthRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Health.Contract.HealthCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Health *HealthRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Health.Contract.HealthTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Health *HealthRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Health.Contract.HealthTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Health *HealthCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _Health.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Health *HealthTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Health.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Health *HealthTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Health.Contract.contract.Transact(opts, method, params...)
}

// Commitment is a free data retrieval call binding the contract method 0x1303a484.
//
// Solidity: function commitment() constant returns(string)
func (_Health *HealthCaller) Commitment(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Health.contract.Call(opts, out, "commitment")
	return *ret0, err
}

// Commitment is a free data retrieval call binding the contract method 0x1303a484.
//
// Solidity: function commitment() constant returns(string)
func (_Health *HealthSession) Commitment() (string, error) {
	return _Health.Contract.Commitment(&_Health.CallOpts)
}

// Commitment is a free data retrieval call binding the contract method 0x1303a484.
//
// Solidity: function commitment() constant returns(string)
func (_Health *HealthCallerSession) Commitment() (string, error) {
	return _Health.Contract.Commitment(&_Health.CallOpts)
}

// MerkleTreeRoot is a free data retrieval call binding the contract method 0xf716aee9.
//
// Solidity: function merkleTreeRoot() constant returns(string)
func (_Health *HealthCaller) MerkleTreeRoot(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _Health.contract.Call(opts, out, "merkleTreeRoot")
	return *ret0, err
}

// MerkleTreeRoot is a free data retrieval call binding the contract method 0xf716aee9.
//
// Solidity: function merkleTreeRoot() constant returns(string)
func (_Health *HealthSession) MerkleTreeRoot() (string, error) {
	return _Health.Contract.MerkleTreeRoot(&_Health.CallOpts)
}

// MerkleTreeRoot is a free data retrieval call binding the contract method 0xf716aee9.
//
// Solidity: function merkleTreeRoot() constant returns(string)
func (_Health *HealthCallerSession) MerkleTreeRoot() (string, error) {
	return _Health.Contract.MerkleTreeRoot(&_Health.CallOpts)
}

// UpdateCommitment is a paid mutator transaction binding the contract method 0x0af3bec2.
//
// Solidity: function updateCommitment(string c) returns()
func (_Health *HealthTransactor) UpdateCommitment(opts *bind.TransactOpts, c string) (*types.Transaction, error) {
	return _Health.contract.Transact(opts, "updateCommitment", c)
}

// UpdateCommitment is a paid mutator transaction binding the contract method 0x0af3bec2.
//
// Solidity: function updateCommitment(string c) returns()
func (_Health *HealthSession) UpdateCommitment(c string) (*types.Transaction, error) {
	return _Health.Contract.UpdateCommitment(&_Health.TransactOpts, c)
}

// UpdateCommitment is a paid mutator transaction binding the contract method 0x0af3bec2.
//
// Solidity: function updateCommitment(string c) returns()
func (_Health *HealthTransactorSession) UpdateCommitment(c string) (*types.Transaction, error) {
	return _Health.Contract.UpdateCommitment(&_Health.TransactOpts, c)
}

// UpdateMerkleTreeRoot is a paid mutator transaction binding the contract method 0x206b3acb.
//
// Solidity: function updateMerkleTreeRoot(string m) returns()
func (_Health *HealthTransactor) UpdateMerkleTreeRoot(opts *bind.TransactOpts, m string) (*types.Transaction, error) {
	return _Health.contract.Transact(opts, "updateMerkleTreeRoot", m)
}

// UpdateMerkleTreeRoot is a paid mutator transaction binding the contract method 0x206b3acb.
//
// Solidity: function updateMerkleTreeRoot(string m) returns()
func (_Health *HealthSession) UpdateMerkleTreeRoot(m string) (*types.Transaction, error) {
	return _Health.Contract.UpdateMerkleTreeRoot(&_Health.TransactOpts, m)
}

// UpdateMerkleTreeRoot is a paid mutator transaction binding the contract method 0x206b3acb.
//
// Solidity: function updateMerkleTreeRoot(string m) returns()
func (_Health *HealthTransactorSession) UpdateMerkleTreeRoot(m string) (*types.Transaction, error) {
	return _Health.Contract.UpdateMerkleTreeRoot(&_Health.TransactOpts, m)
}
