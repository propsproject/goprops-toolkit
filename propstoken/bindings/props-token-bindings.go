// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package bindings

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

// PropsTokenABI is the input ABI used to generate the binding from.
const PropsTokenABI = "[{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_signature\",\"type\":\"bytes\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_gasPrice\",\"type\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"transferPreSigned\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_signature\",\"type\":\"bytes\"},{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_gasPrice\",\"type\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"approvePreSigned\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_who\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"}]"

// PropsToken is an auto generated Go binding around an Ethereum contract.
type PropsToken struct {
	PropsTokenCaller     // Read-only binding to the contract
	PropsTokenTransactor // Write-only binding to the contract
	PropsTokenFilterer   // Log filterer for contract events
}

// PropsTokenCaller is an auto generated read-only Go binding around an Ethereum contract.
type PropsTokenCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PropsTokenTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PropsTokenTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PropsTokenFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PropsTokenFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PropsTokenSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PropsTokenSession struct {
	Contract     *PropsToken       // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// PropsTokenCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PropsTokenCallerSession struct {
	Contract *PropsTokenCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts     // Call options to use throughout this session
}

// PropsTokenTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PropsTokenTransactorSession struct {
	Contract     *PropsTokenTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts     // Transaction auth options to use throughout this session
}

// PropsTokenRaw is an auto generated low-level Go binding around an Ethereum contract.
type PropsTokenRaw struct {
	Contract *PropsToken // Generic contract binding to access the raw methods on
}

// PropsTokenCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PropsTokenCallerRaw struct {
	Contract *PropsTokenCaller // Generic read-only contract binding to access the raw methods on
}

// PropsTokenTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PropsTokenTransactorRaw struct {
	Contract *PropsTokenTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPropsToken creates a new instance of PropsToken, bound to a specific deployed contract.
func NewPropsToken(address common.Address, backend bind.ContractBackend) (*PropsToken, error) {
	contract, err := bindPropsToken(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PropsToken{PropsTokenCaller: PropsTokenCaller{contract: contract}, PropsTokenTransactor: PropsTokenTransactor{contract: contract}, PropsTokenFilterer: PropsTokenFilterer{contract: contract}}, nil
}

// NewPropsTokenCaller creates a new read-only instance of PropsToken, bound to a specific deployed contract.
func NewPropsTokenCaller(address common.Address, caller bind.ContractCaller) (*PropsTokenCaller, error) {
	contract, err := bindPropsToken(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PropsTokenCaller{contract: contract}, nil
}

// NewPropsTokenTransactor creates a new write-only instance of PropsToken, bound to a specific deployed contract.
func NewPropsTokenTransactor(address common.Address, transactor bind.ContractTransactor) (*PropsTokenTransactor, error) {
	contract, err := bindPropsToken(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PropsTokenTransactor{contract: contract}, nil
}

// NewPropsTokenFilterer creates a new log filterer instance of PropsToken, bound to a specific deployed contract.
func NewPropsTokenFilterer(address common.Address, filterer bind.ContractFilterer) (*PropsTokenFilterer, error) {
	contract, err := bindPropsToken(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PropsTokenFilterer{contract: contract}, nil
}

// bindPropsToken binds a generic wrapper to an already deployed contract.
func bindPropsToken(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PropsTokenABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PropsToken *PropsTokenRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PropsToken.Contract.PropsTokenCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PropsToken *PropsTokenRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PropsToken.Contract.PropsTokenTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PropsToken *PropsTokenRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PropsToken.Contract.PropsTokenTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PropsToken *PropsTokenCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PropsToken.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PropsToken *PropsTokenTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PropsToken.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PropsToken *PropsTokenTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PropsToken.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_PropsToken *PropsTokenCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PropsToken.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_PropsToken *PropsTokenSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _PropsToken.Contract.Allowance(&_PropsToken.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_PropsToken *PropsTokenCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _PropsToken.Contract.Allowance(&_PropsToken.CallOpts, _owner, _spender)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_who address) constant returns(uint256)
func (_PropsToken *PropsTokenCaller) BalanceOf(opts *bind.CallOpts, _who common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PropsToken.contract.Call(opts, out, "balanceOf", _who)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_who address) constant returns(uint256)
func (_PropsToken *PropsTokenSession) BalanceOf(_who common.Address) (*big.Int, error) {
	return _PropsToken.Contract.BalanceOf(&_PropsToken.CallOpts, _who)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_who address) constant returns(uint256)
func (_PropsToken *PropsTokenCallerSession) BalanceOf(_who common.Address) (*big.Int, error) {
	return _PropsToken.Contract.BalanceOf(&_PropsToken.CallOpts, _who)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_PropsToken *PropsTokenCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PropsToken.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_PropsToken *PropsTokenSession) TotalSupply() (*big.Int, error) {
	return _PropsToken.Contract.TotalSupply(&_PropsToken.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_PropsToken *PropsTokenCallerSession) TotalSupply() (*big.Int, error) {
	return _PropsToken.Contract.TotalSupply(&_PropsToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_PropsToken *PropsTokenTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PropsToken.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_PropsToken *PropsTokenSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PropsToken.Contract.Approve(&_PropsToken.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_PropsToken *PropsTokenTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PropsToken.Contract.Approve(&_PropsToken.TransactOpts, _spender, _value)
}

// ApprovePreSigned is a paid mutator transaction binding the contract method 0x617b390b.
//
// Solidity: function approvePreSigned(_signature bytes, _spender address, _value uint256, _gasPrice uint256, _nonce uint256) returns(bool)
func (_PropsToken *PropsTokenTransactor) ApprovePreSigned(opts *bind.TransactOpts, _signature []byte, _spender common.Address, _value *big.Int, _gasPrice *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _PropsToken.contract.Transact(opts, "approvePreSigned", _signature, _spender, _value, _gasPrice, _nonce)
}

// ApprovePreSigned is a paid mutator transaction binding the contract method 0x617b390b.
//
// Solidity: function approvePreSigned(_signature bytes, _spender address, _value uint256, _gasPrice uint256, _nonce uint256) returns(bool)
func (_PropsToken *PropsTokenSession) ApprovePreSigned(_signature []byte, _spender common.Address, _value *big.Int, _gasPrice *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _PropsToken.Contract.ApprovePreSigned(&_PropsToken.TransactOpts, _signature, _spender, _value, _gasPrice, _nonce)
}

// ApprovePreSigned is a paid mutator transaction binding the contract method 0x617b390b.
//
// Solidity: function approvePreSigned(_signature bytes, _spender address, _value uint256, _gasPrice uint256, _nonce uint256) returns(bool)
func (_PropsToken *PropsTokenTransactorSession) ApprovePreSigned(_signature []byte, _spender common.Address, _value *big.Int, _gasPrice *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _PropsToken.Contract.ApprovePreSigned(&_PropsToken.TransactOpts, _signature, _spender, _value, _gasPrice, _nonce)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_PropsToken *PropsTokenTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PropsToken.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_PropsToken *PropsTokenSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PropsToken.Contract.Transfer(&_PropsToken.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_PropsToken *PropsTokenTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PropsToken.Contract.Transfer(&_PropsToken.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_PropsToken *PropsTokenTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PropsToken.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_PropsToken *PropsTokenSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PropsToken.Contract.TransferFrom(&_PropsToken.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_PropsToken *PropsTokenTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PropsToken.Contract.TransferFrom(&_PropsToken.TransactOpts, _from, _to, _value)
}

// TransferPreSigned is a paid mutator transaction binding the contract method 0x1296830d.
//
// Solidity: function transferPreSigned(_signature bytes, _to address, _value uint256, _gasPrice uint256, _nonce uint256) returns(bool)
func (_PropsToken *PropsTokenTransactor) TransferPreSigned(opts *bind.TransactOpts, _signature []byte, _to common.Address, _value *big.Int, _gasPrice *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _PropsToken.contract.Transact(opts, "transferPreSigned", _signature, _to, _value, _gasPrice, _nonce)
}

// TransferPreSigned is a paid mutator transaction binding the contract method 0x1296830d.
//
// Solidity: function transferPreSigned(_signature bytes, _to address, _value uint256, _gasPrice uint256, _nonce uint256) returns(bool)
func (_PropsToken *PropsTokenSession) TransferPreSigned(_signature []byte, _to common.Address, _value *big.Int, _gasPrice *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _PropsToken.Contract.TransferPreSigned(&_PropsToken.TransactOpts, _signature, _to, _value, _gasPrice, _nonce)
}

// TransferPreSigned is a paid mutator transaction binding the contract method 0x1296830d.
//
// Solidity: function transferPreSigned(_signature bytes, _to address, _value uint256, _gasPrice uint256, _nonce uint256) returns(bool)
func (_PropsToken *PropsTokenTransactorSession) TransferPreSigned(_signature []byte, _to common.Address, _value *big.Int, _gasPrice *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _PropsToken.Contract.TransferPreSigned(&_PropsToken.TransactOpts, _signature, _to, _value, _gasPrice, _nonce)
}

// PropsTokenApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the PropsToken contract.
type PropsTokenApprovalIterator struct {
	Event *PropsTokenApproval // Event containing the contract specifics and raw log

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
func (it *PropsTokenApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropsTokenApproval)
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
		it.Event = new(PropsTokenApproval)
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
func (it *PropsTokenApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropsTokenApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropsTokenApproval represents a Approval event raised by the PropsToken contract.
type PropsTokenApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_PropsToken *PropsTokenFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*PropsTokenApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PropsToken.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &PropsTokenApprovalIterator{contract: _PropsToken.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_PropsToken *PropsTokenFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PropsTokenApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PropsToken.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropsTokenApproval)
				if err := _PropsToken.contract.UnpackLog(event, "Approval", log); err != nil {
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

// PropsTokenTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the PropsToken contract.
type PropsTokenTransferIterator struct {
	Event *PropsTokenTransfer // Event containing the contract specifics and raw log

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
func (it *PropsTokenTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropsTokenTransfer)
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
		it.Event = new(PropsTokenTransfer)
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
func (it *PropsTokenTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropsTokenTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropsTokenTransfer represents a Transfer event raised by the PropsToken contract.
type PropsTokenTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_PropsToken *PropsTokenFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PropsTokenTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PropsToken.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PropsTokenTransferIterator{contract: _PropsToken.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_PropsToken *PropsTokenFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PropsTokenTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PropsToken.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropsTokenTransfer)
				if err := _PropsToken.contract.UnpackLog(event, "Transfer", log); err != nil {
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
