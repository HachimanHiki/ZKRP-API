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
const HealthABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"c\",\"type\":\"string\"}],\"name\":\"updateCommitment\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"commitment\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"m\",\"type\":\"string\"}],\"name\":\"updateMerkleTreeRoot\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"renounceOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"merkleTreeRoot\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"c\",\"type\":\"string\"},{\"name\":\"m\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"}],\"name\":\"OwnershipRenounced\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"}]"

// HealthBin is the compiled bytecode used for deploying new contracts.
var HealthBin = "0x608060405234801561001057600080fd5b50604051610a10380380610a108339810180604052810190808051820192919060200180518201929190505050336000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555081600190805190602001906100939291906100b2565b5080600290805190602001906100aa9291906100b2565b505050610157565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f106100f357805160ff1916838001178555610121565b82800160010185558215610121579182015b82811115610120578251825591602001919060010190610105565b5b50905061012e9190610132565b5090565b61015491905b80821115610150576000816000905550600101610138565b5090565b90565b6108aa806101666000396000f300608060405260043610610083576000357c0100000000000000000000000000000000000000000000000000000000900463ffffffff1680630af3bec2146100885780631303a484146100f1578063206b3acb14610181578063715018a6146101ea5780638da5cb5b14610201578063f2fde38b14610258578063f716aee91461029b575b600080fd5b34801561009457600080fd5b506100ef600480360381019080803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929050505061032b565b005b3480156100fd57600080fd5b506101066103a0565b6040518080602001828103825283818151815260200191508051906020019080838360005b8381101561014657808201518184015260208101905061012b565b50505050905090810190601f1680156101735780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b34801561018d57600080fd5b506101e8600480360381019080803590602001908201803590602001908080601f016020809104026020016040519081016040528093929190818152602001838380828437820191505050505050919291929050505061043e565b005b3480156101f657600080fd5b506101ff6104b3565b005b34801561020d57600080fd5b506102166105b5565b604051808273ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff16815260200191505060405180910390f35b34801561026457600080fd5b50610299600480360381019080803573ffffffffffffffffffffffffffffffffffffffff1690602001909291905050506105da565b005b3480156102a757600080fd5b506102b0610641565b6040518080602001828103825283818151815260200191508051906020019080838360005b838110156102f05780820151818401526020810190506102d5565b50505050905090810190601f16801561031d5780820380516001836020036101000a031916815260200191505b509250505060405180910390f35b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561038657600080fd5b806001908051906020019061039c9291906107d9565b5050565b60018054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156104365780601f1061040b57610100808354040283529160200191610436565b820191906000526020600020905b81548152906001019060200180831161041957829003601f168201915b505050505081565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561049957600080fd5b80600290805190602001906104af9291906107d9565b5050565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561050e57600080fd5b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167ff8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c6482060405160405180910390a260008060006101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff160217905550565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1681565b6000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff163373ffffffffffffffffffffffffffffffffffffffff1614151561063557600080fd5b61063e816106df565b50565b60028054600181600116156101000203166002900480601f0160208091040260200160405190810160405280929190818152602001828054600181600116156101000203166002900480156106d75780601f106106ac576101008083540402835291602001916106d7565b820191906000526020600020905b8154815290600101906020018083116106ba57829003601f168201915b505050505081565b600073ffffffffffffffffffffffffffffffffffffffff168173ffffffffffffffffffffffffffffffffffffffff161415151561071b57600080fd5b8073ffffffffffffffffffffffffffffffffffffffff166000809054906101000a900473ffffffffffffffffffffffffffffffffffffffff1673ffffffffffffffffffffffffffffffffffffffff167f8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e060405160405180910390a3806000806101000a81548173ffffffffffffffffffffffffffffffffffffffff021916908373ffffffffffffffffffffffffffffffffffffffff16021790555050565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061081a57805160ff1916838001178555610848565b82800160010185558215610848579182015b8281111561084757825182559160200191906001019061082c565b5b5090506108559190610859565b5090565b61087b91905b8082111561087757600081600090555060010161085f565b5090565b905600a165627a7a72305820a0f0c02a16bcee36c86ba1b8cae83836c39dce830686245a5a92bf7dd1ab2d570029"

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

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Health *HealthCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _Health.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Health *HealthSession) Owner() (common.Address, error) {
	return _Health.Contract.Owner(&_Health.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_Health *HealthCallerSession) Owner() (common.Address, error) {
	return _Health.Contract.Owner(&_Health.CallOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Health *HealthTransactor) RenounceOwnership(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Health.contract.Transact(opts, "renounceOwnership")
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Health *HealthSession) RenounceOwnership() (*types.Transaction, error) {
	return _Health.Contract.RenounceOwnership(&_Health.TransactOpts)
}

// RenounceOwnership is a paid mutator transaction binding the contract method 0x715018a6.
//
// Solidity: function renounceOwnership() returns()
func (_Health *HealthTransactorSession) RenounceOwnership() (*types.Transaction, error) {
	return _Health.Contract.RenounceOwnership(&_Health.TransactOpts)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Health *HealthTransactor) TransferOwnership(opts *bind.TransactOpts, _newOwner common.Address) (*types.Transaction, error) {
	return _Health.contract.Transact(opts, "transferOwnership", _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Health *HealthSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Health.Contract.TransferOwnership(&_Health.TransactOpts, _newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(address _newOwner) returns()
func (_Health *HealthTransactorSession) TransferOwnership(_newOwner common.Address) (*types.Transaction, error) {
	return _Health.Contract.TransferOwnership(&_Health.TransactOpts, _newOwner)
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

// HealthOwnershipRenouncedIterator is returned from FilterOwnershipRenounced and is used to iterate over the raw logs and unpacked data for OwnershipRenounced events raised by the Health contract.
type HealthOwnershipRenouncedIterator struct {
	Event *HealthOwnershipRenounced // Event containing the contract specifics and raw log

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
func (it *HealthOwnershipRenouncedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HealthOwnershipRenounced)
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
		it.Event = new(HealthOwnershipRenounced)
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
func (it *HealthOwnershipRenouncedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HealthOwnershipRenouncedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HealthOwnershipRenounced represents a OwnershipRenounced event raised by the Health contract.
type HealthOwnershipRenounced struct {
	PreviousOwner common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipRenounced is a free log retrieval operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Health *HealthFilterer) FilterOwnershipRenounced(opts *bind.FilterOpts, previousOwner []common.Address) (*HealthOwnershipRenouncedIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Health.contract.FilterLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HealthOwnershipRenouncedIterator{contract: _Health.contract, event: "OwnershipRenounced", logs: logs, sub: sub}, nil
}

// WatchOwnershipRenounced is a free log subscription operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Health *HealthFilterer) WatchOwnershipRenounced(opts *bind.WatchOpts, sink chan<- *HealthOwnershipRenounced, previousOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}

	logs, sub, err := _Health.contract.WatchLogs(opts, "OwnershipRenounced", previousOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HealthOwnershipRenounced)
				if err := _Health.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
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

// ParseOwnershipRenounced is a log parse operation binding the contract event 0xf8df31144d9c2f0f6b59d69b8b98abd5459d07f2742c4df920b25aae33c64820.
//
// Solidity: event OwnershipRenounced(address indexed previousOwner)
func (_Health *HealthFilterer) ParseOwnershipRenounced(log types.Log) (*HealthOwnershipRenounced, error) {
	event := new(HealthOwnershipRenounced)
	if err := _Health.contract.UnpackLog(event, "OwnershipRenounced", log); err != nil {
		return nil, err
	}
	return event, nil
}

// HealthOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the Health contract.
type HealthOwnershipTransferredIterator struct {
	Event *HealthOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *HealthOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(HealthOwnershipTransferred)
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
		it.Event = new(HealthOwnershipTransferred)
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
func (it *HealthOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *HealthOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// HealthOwnershipTransferred represents a OwnershipTransferred event raised by the Health contract.
type HealthOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Health *HealthFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*HealthOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Health.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &HealthOwnershipTransferredIterator{contract: _Health.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: event OwnershipTransferred(address indexed previousOwner, address indexed newOwner)
func (_Health *HealthFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *HealthOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _Health.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(HealthOwnershipTransferred)
				if err := _Health.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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
func (_Health *HealthFilterer) ParseOwnershipTransferred(log types.Log) (*HealthOwnershipTransferred, error) {
	event := new(HealthOwnershipTransferred)
	if err := _Health.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
		return nil, err
	}
	return event, nil
}
