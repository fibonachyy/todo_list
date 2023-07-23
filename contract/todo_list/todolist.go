// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package todolist

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

// TodolistMetaData contains all meta data concerning the Todolist contract.
var TodolistMetaData = &bind.MetaData{
	ABI: "[{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"string\",\"name\":\"content\",\"type\":\"string\"}],\"name\":\"TaskAdded\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"internalType\":\"address\",\"name\":\"user\",\"type\":\"address\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"taskId\",\"type\":\"uint256\"}],\"name\":\"TaskCompleted\",\"type\":\"event\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"_content\",\"type\":\"string\"}],\"name\":\"addTask\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_taskId\",\"type\":\"uint256\"}],\"name\":\"getTask\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"},{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getTasksCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"_taskId\",\"type\":\"uint256\"}],\"name\":\"markTaskCompleted\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// TodolistABI is the input ABI used to generate the binding from.
// Deprecated: Use TodolistMetaData.ABI instead.
var TodolistABI = TodolistMetaData.ABI

// Todolist is an auto generated Go binding around an Ethereum contract.
type Todolist struct {
	TodolistCaller     // Read-only binding to the contract
	TodolistTransactor // Write-only binding to the contract
	TodolistFilterer   // Log filterer for contract events
}

// TodolistCaller is an auto generated read-only Go binding around an Ethereum contract.
type TodolistCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodolistTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TodolistTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodolistFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TodolistFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TodolistSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TodolistSession struct {
	Contract     *Todolist         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TodolistCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TodolistCallerSession struct {
	Contract *TodolistCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// TodolistTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TodolistTransactorSession struct {
	Contract     *TodolistTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// TodolistRaw is an auto generated low-level Go binding around an Ethereum contract.
type TodolistRaw struct {
	Contract *Todolist // Generic contract binding to access the raw methods on
}

// TodolistCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TodolistCallerRaw struct {
	Contract *TodolistCaller // Generic read-only contract binding to access the raw methods on
}

// TodolistTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TodolistTransactorRaw struct {
	Contract *TodolistTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTodolist creates a new instance of Todolist, bound to a specific deployed contract.
func NewTodolist(address common.Address, backend bind.ContractBackend) (*Todolist, error) {
	contract, err := bindTodolist(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Todolist{TodolistCaller: TodolistCaller{contract: contract}, TodolistTransactor: TodolistTransactor{contract: contract}, TodolistFilterer: TodolistFilterer{contract: contract}}, nil
}

// NewTodolistCaller creates a new read-only instance of Todolist, bound to a specific deployed contract.
func NewTodolistCaller(address common.Address, caller bind.ContractCaller) (*TodolistCaller, error) {
	contract, err := bindTodolist(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TodolistCaller{contract: contract}, nil
}

// NewTodolistTransactor creates a new write-only instance of Todolist, bound to a specific deployed contract.
func NewTodolistTransactor(address common.Address, transactor bind.ContractTransactor) (*TodolistTransactor, error) {
	contract, err := bindTodolist(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TodolistTransactor{contract: contract}, nil
}

// NewTodolistFilterer creates a new log filterer instance of Todolist, bound to a specific deployed contract.
func NewTodolistFilterer(address common.Address, filterer bind.ContractFilterer) (*TodolistFilterer, error) {
	contract, err := bindTodolist(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TodolistFilterer{contract: contract}, nil
}

// bindTodolist binds a generic wrapper to an already deployed contract.
func bindTodolist(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := TodolistMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Todolist *TodolistRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Todolist.Contract.TodolistCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Todolist *TodolistRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Todolist.Contract.TodolistTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Todolist *TodolistRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Todolist.Contract.TodolistTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Todolist *TodolistCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Todolist.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Todolist *TodolistTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Todolist.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Todolist *TodolistTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Todolist.Contract.contract.Transact(opts, method, params...)
}

// GetTask is a free data retrieval call binding the contract method 0x1d65e77e.
//
// Solidity: function getTask(uint256 _taskId) view returns(uint256, string, bool)
func (_Todolist *TodolistCaller) GetTask(opts *bind.CallOpts, _taskId *big.Int) (*big.Int, string, bool, error) {
	var out []interface{}
	err := _Todolist.contract.Call(opts, &out, "getTask", _taskId)

	if err != nil {
		return *new(*big.Int), *new(string), *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(string)).(*string)
	out2 := *abi.ConvertType(out[2], new(bool)).(*bool)

	return out0, out1, out2, err

}

// GetTask is a free data retrieval call binding the contract method 0x1d65e77e.
//
// Solidity: function getTask(uint256 _taskId) view returns(uint256, string, bool)
func (_Todolist *TodolistSession) GetTask(_taskId *big.Int) (*big.Int, string, bool, error) {
	return _Todolist.Contract.GetTask(&_Todolist.CallOpts, _taskId)
}

// GetTask is a free data retrieval call binding the contract method 0x1d65e77e.
//
// Solidity: function getTask(uint256 _taskId) view returns(uint256, string, bool)
func (_Todolist *TodolistCallerSession) GetTask(_taskId *big.Int) (*big.Int, string, bool, error) {
	return _Todolist.Contract.GetTask(&_Todolist.CallOpts, _taskId)
}

// GetTasksCount is a free data retrieval call binding the contract method 0x9f899273.
//
// Solidity: function getTasksCount() view returns(uint256)
func (_Todolist *TodolistCaller) GetTasksCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Todolist.contract.Call(opts, &out, "getTasksCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTasksCount is a free data retrieval call binding the contract method 0x9f899273.
//
// Solidity: function getTasksCount() view returns(uint256)
func (_Todolist *TodolistSession) GetTasksCount() (*big.Int, error) {
	return _Todolist.Contract.GetTasksCount(&_Todolist.CallOpts)
}

// GetTasksCount is a free data retrieval call binding the contract method 0x9f899273.
//
// Solidity: function getTasksCount() view returns(uint256)
func (_Todolist *TodolistCallerSession) GetTasksCount() (*big.Int, error) {
	return _Todolist.Contract.GetTasksCount(&_Todolist.CallOpts)
}

// AddTask is a paid mutator transaction binding the contract method 0x67238562.
//
// Solidity: function addTask(string _content) returns()
func (_Todolist *TodolistTransactor) AddTask(opts *bind.TransactOpts, _content string) (*types.Transaction, error) {
	return _Todolist.contract.Transact(opts, "addTask", _content)
}

// AddTask is a paid mutator transaction binding the contract method 0x67238562.
//
// Solidity: function addTask(string _content) returns()
func (_Todolist *TodolistSession) AddTask(_content string) (*types.Transaction, error) {
	return _Todolist.Contract.AddTask(&_Todolist.TransactOpts, _content)
}

// AddTask is a paid mutator transaction binding the contract method 0x67238562.
//
// Solidity: function addTask(string _content) returns()
func (_Todolist *TodolistTransactorSession) AddTask(_content string) (*types.Transaction, error) {
	return _Todolist.Contract.AddTask(&_Todolist.TransactOpts, _content)
}

// MarkTaskCompleted is a paid mutator transaction binding the contract method 0x000b2bda.
//
// Solidity: function markTaskCompleted(uint256 _taskId) returns()
func (_Todolist *TodolistTransactor) MarkTaskCompleted(opts *bind.TransactOpts, _taskId *big.Int) (*types.Transaction, error) {
	return _Todolist.contract.Transact(opts, "markTaskCompleted", _taskId)
}

// MarkTaskCompleted is a paid mutator transaction binding the contract method 0x000b2bda.
//
// Solidity: function markTaskCompleted(uint256 _taskId) returns()
func (_Todolist *TodolistSession) MarkTaskCompleted(_taskId *big.Int) (*types.Transaction, error) {
	return _Todolist.Contract.MarkTaskCompleted(&_Todolist.TransactOpts, _taskId)
}

// MarkTaskCompleted is a paid mutator transaction binding the contract method 0x000b2bda.
//
// Solidity: function markTaskCompleted(uint256 _taskId) returns()
func (_Todolist *TodolistTransactorSession) MarkTaskCompleted(_taskId *big.Int) (*types.Transaction, error) {
	return _Todolist.Contract.MarkTaskCompleted(&_Todolist.TransactOpts, _taskId)
}

// TodolistTaskAddedIterator is returned from FilterTaskAdded and is used to iterate over the raw logs and unpacked data for TaskAdded events raised by the Todolist contract.
type TodolistTaskAddedIterator struct {
	Event *TodolistTaskAdded // Event containing the contract specifics and raw log

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
func (it *TodolistTaskAddedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TodolistTaskAdded)
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
		it.Event = new(TodolistTaskAdded)
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
func (it *TodolistTaskAddedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TodolistTaskAddedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TodolistTaskAdded represents a TaskAdded event raised by the Todolist contract.
type TodolistTaskAdded struct {
	User    common.Address
	TaskId  *big.Int
	Content string
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterTaskAdded is a free log retrieval operation binding the contract event 0x793f49aab3b9e0d3125d745773b78ac45236f2beeed7c6296e60911e17a0de42.
//
// Solidity: event TaskAdded(address indexed user, uint256 taskId, string content)
func (_Todolist *TodolistFilterer) FilterTaskAdded(opts *bind.FilterOpts, user []common.Address) (*TodolistTaskAddedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Todolist.contract.FilterLogs(opts, "TaskAdded", userRule)
	if err != nil {
		return nil, err
	}
	return &TodolistTaskAddedIterator{contract: _Todolist.contract, event: "TaskAdded", logs: logs, sub: sub}, nil
}

// WatchTaskAdded is a free log subscription operation binding the contract event 0x793f49aab3b9e0d3125d745773b78ac45236f2beeed7c6296e60911e17a0de42.
//
// Solidity: event TaskAdded(address indexed user, uint256 taskId, string content)
func (_Todolist *TodolistFilterer) WatchTaskAdded(opts *bind.WatchOpts, sink chan<- *TodolistTaskAdded, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Todolist.contract.WatchLogs(opts, "TaskAdded", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TodolistTaskAdded)
				if err := _Todolist.contract.UnpackLog(event, "TaskAdded", log); err != nil {
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

// ParseTaskAdded is a log parse operation binding the contract event 0x793f49aab3b9e0d3125d745773b78ac45236f2beeed7c6296e60911e17a0de42.
//
// Solidity: event TaskAdded(address indexed user, uint256 taskId, string content)
func (_Todolist *TodolistFilterer) ParseTaskAdded(log types.Log) (*TodolistTaskAdded, error) {
	event := new(TodolistTaskAdded)
	if err := _Todolist.contract.UnpackLog(event, "TaskAdded", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// TodolistTaskCompletedIterator is returned from FilterTaskCompleted and is used to iterate over the raw logs and unpacked data for TaskCompleted events raised by the Todolist contract.
type TodolistTaskCompletedIterator struct {
	Event *TodolistTaskCompleted // Event containing the contract specifics and raw log

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
func (it *TodolistTaskCompletedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(TodolistTaskCompleted)
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
		it.Event = new(TodolistTaskCompleted)
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
func (it *TodolistTaskCompletedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *TodolistTaskCompletedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// TodolistTaskCompleted represents a TaskCompleted event raised by the Todolist contract.
type TodolistTaskCompleted struct {
	User   common.Address
	TaskId *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterTaskCompleted is a free log retrieval operation binding the contract event 0xf684d89bb0f42287bf32838774e4b7f8a60c11804b3c6d1791fa55f306792643.
//
// Solidity: event TaskCompleted(address indexed user, uint256 taskId)
func (_Todolist *TodolistFilterer) FilterTaskCompleted(opts *bind.FilterOpts, user []common.Address) (*TodolistTaskCompletedIterator, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Todolist.contract.FilterLogs(opts, "TaskCompleted", userRule)
	if err != nil {
		return nil, err
	}
	return &TodolistTaskCompletedIterator{contract: _Todolist.contract, event: "TaskCompleted", logs: logs, sub: sub}, nil
}

// WatchTaskCompleted is a free log subscription operation binding the contract event 0xf684d89bb0f42287bf32838774e4b7f8a60c11804b3c6d1791fa55f306792643.
//
// Solidity: event TaskCompleted(address indexed user, uint256 taskId)
func (_Todolist *TodolistFilterer) WatchTaskCompleted(opts *bind.WatchOpts, sink chan<- *TodolistTaskCompleted, user []common.Address) (event.Subscription, error) {

	var userRule []interface{}
	for _, userItem := range user {
		userRule = append(userRule, userItem)
	}

	logs, sub, err := _Todolist.contract.WatchLogs(opts, "TaskCompleted", userRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(TodolistTaskCompleted)
				if err := _Todolist.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
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

// ParseTaskCompleted is a log parse operation binding the contract event 0xf684d89bb0f42287bf32838774e4b7f8a60c11804b3c6d1791fa55f306792643.
//
// Solidity: event TaskCompleted(address indexed user, uint256 taskId)
func (_Todolist *TodolistFilterer) ParseTaskCompleted(log types.Log) (*TodolistTaskCompleted, error) {
	event := new(TodolistTaskCompleted)
	if err := _Todolist.contract.UnpackLog(event, "TaskCompleted", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
