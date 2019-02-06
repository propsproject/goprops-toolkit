// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package propstoken

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

// PropsTokenABI is the input ABI used to generate the binding from.
const PropsTokenABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_signature\",\"type\":\"bytes\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"transferPreSigned\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"transferPreSignedHashing\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"settle\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"recover\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"from\",\"type\":\"address\"},{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"canTransferBeforeStartTime\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"decreaseApprovalPreSignedHashing\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_signature\",\"type\":\"bytes\"},{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"approvePreSigned\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"canTransfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_signature\",\"type\":\"bytes\"},{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"decreaseApprovalPreSigned\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"spender\",\"type\":\"address\"},{\"name\":\"subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseAllowance\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"increaseApprovalPreSignedHashing\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"to\",\"type\":\"address\"},{\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_signature\",\"type\":\"bytes\"},{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"increaseApprovalPreSigned\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"transferFromPreSignedHashing\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_signature\",\"type\":\"bytes\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"transferFromPreSigned\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"transfersStartTime\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"owner\",\"type\":\"address\"},{\"name\":\"spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"approvePreSignedHashing\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"delegate\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"TransferPreSigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"delegate\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"ApprovalPreSigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"fromBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"toBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferDetails\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Settlement\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"constant\":false,\"inputs\":[{\"name\":\"name\",\"type\":\"string\"},{\"name\":\"symbol\",\"type\":\"string\"},{\"name\":\"decimals\",\"type\":\"uint8\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_holder\",\"type\":\"address\"},{\"name\":\"_transfersStartTime\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"start\",\"type\":\"uint256\"},{\"name\":\"account\",\"type\":\"address\"}],\"name\":\"initialize\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]"

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
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_PropsToken *PropsTokenCaller) Allowance(opts *bind.CallOpts, owner common.Address, spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PropsToken.contract.Call(opts, out, "allowance", owner, spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_PropsToken *PropsTokenSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _PropsToken.Contract.Allowance(&_PropsToken.CallOpts, owner, spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(owner address, spender address) constant returns(uint256)
func (_PropsToken *PropsTokenCallerSession) Allowance(owner common.Address, spender common.Address) (*big.Int, error) {
	return _PropsToken.Contract.Allowance(&_PropsToken.CallOpts, owner, spender)
}

// ApprovePreSignedHashing is a free data retrieval call binding the contract method 0xf7ac9c2e.
//
// Solidity: function approvePreSignedHashing(_token address, _spender address, _value uint256, _fee uint256, _nonce uint256) constant returns(bytes32)
func (_PropsToken *PropsTokenCaller) ApprovePreSignedHashing(opts *bind.CallOpts, _token common.Address, _spender common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _PropsToken.contract.Call(opts, out, "approvePreSignedHashing", _token, _spender, _value, _fee, _nonce)
	return *ret0, err
}

// ApprovePreSignedHashing is a free data retrieval call binding the contract method 0xf7ac9c2e.
//
// Solidity: function approvePreSignedHashing(_token address, _spender address, _value uint256, _fee uint256, _nonce uint256) constant returns(bytes32)
func (_PropsToken *PropsTokenSession) ApprovePreSignedHashing(_token common.Address, _spender common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) ([32]byte, error) {
	return _PropsToken.Contract.ApprovePreSignedHashing(&_PropsToken.CallOpts, _token, _spender, _value, _fee, _nonce)
}

// ApprovePreSignedHashing is a free data retrieval call binding the contract method 0xf7ac9c2e.
//
// Solidity: function approvePreSignedHashing(_token address, _spender address, _value uint256, _fee uint256, _nonce uint256) constant returns(bytes32)
func (_PropsToken *PropsTokenCallerSession) ApprovePreSignedHashing(_token common.Address, _spender common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) ([32]byte, error) {
	return _PropsToken.Contract.ApprovePreSignedHashing(&_PropsToken.CallOpts, _token, _spender, _value, _fee, _nonce)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_PropsToken *PropsTokenCaller) BalanceOf(opts *bind.CallOpts, owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PropsToken.contract.Call(opts, out, "balanceOf", owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_PropsToken *PropsTokenSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _PropsToken.Contract.BalanceOf(&_PropsToken.CallOpts, owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(owner address) constant returns(uint256)
func (_PropsToken *PropsTokenCallerSession) BalanceOf(owner common.Address) (*big.Int, error) {
	return _PropsToken.Contract.BalanceOf(&_PropsToken.CallOpts, owner)
}

// CanTransfer is a free data retrieval call binding the contract method 0x78fc3cb3.
//
// Solidity: function canTransfer(account address) constant returns(bool)
func (_PropsToken *PropsTokenCaller) CanTransfer(opts *bind.CallOpts, account common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PropsToken.contract.Call(opts, out, "canTransfer", account)
	return *ret0, err
}

// CanTransfer is a free data retrieval call binding the contract method 0x78fc3cb3.
//
// Solidity: function canTransfer(account address) constant returns(bool)
func (_PropsToken *PropsTokenSession) CanTransfer(account common.Address) (bool, error) {
	return _PropsToken.Contract.CanTransfer(&_PropsToken.CallOpts, account)
}

// CanTransfer is a free data retrieval call binding the contract method 0x78fc3cb3.
//
// Solidity: function canTransfer(account address) constant returns(bool)
func (_PropsToken *PropsTokenCallerSession) CanTransfer(account common.Address) (bool, error) {
	return _PropsToken.Contract.CanTransfer(&_PropsToken.CallOpts, account)
}

// CanTransferBeforeStartTime is a free data retrieval call binding the contract method 0x45b6d7dc.
//
// Solidity: function canTransferBeforeStartTime() constant returns(address)
func (_PropsToken *PropsTokenCaller) CanTransferBeforeStartTime(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PropsToken.contract.Call(opts, out, "canTransferBeforeStartTime")
	return *ret0, err
}

// CanTransferBeforeStartTime is a free data retrieval call binding the contract method 0x45b6d7dc.
//
// Solidity: function canTransferBeforeStartTime() constant returns(address)
func (_PropsToken *PropsTokenSession) CanTransferBeforeStartTime() (common.Address, error) {
	return _PropsToken.Contract.CanTransferBeforeStartTime(&_PropsToken.CallOpts)
}

// CanTransferBeforeStartTime is a free data retrieval call binding the contract method 0x45b6d7dc.
//
// Solidity: function canTransferBeforeStartTime() constant returns(address)
func (_PropsToken *PropsTokenCallerSession) CanTransferBeforeStartTime() (common.Address, error) {
	return _PropsToken.Contract.CanTransferBeforeStartTime(&_PropsToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_PropsToken *PropsTokenCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _PropsToken.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_PropsToken *PropsTokenSession) Decimals() (uint8, error) {
	return _PropsToken.Contract.Decimals(&_PropsToken.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_PropsToken *PropsTokenCallerSession) Decimals() (uint8, error) {
	return _PropsToken.Contract.Decimals(&_PropsToken.CallOpts)
}

// DecreaseApprovalPreSignedHashing is a free data retrieval call binding the contract method 0x59388d78.
//
// Solidity: function decreaseApprovalPreSignedHashing(_token address, _spender address, _subtractedValue uint256, _fee uint256, _nonce uint256) constant returns(bytes32)
func (_PropsToken *PropsTokenCaller) DecreaseApprovalPreSignedHashing(opts *bind.CallOpts, _token common.Address, _spender common.Address, _subtractedValue *big.Int, _fee *big.Int, _nonce *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _PropsToken.contract.Call(opts, out, "decreaseApprovalPreSignedHashing", _token, _spender, _subtractedValue, _fee, _nonce)
	return *ret0, err
}

// DecreaseApprovalPreSignedHashing is a free data retrieval call binding the contract method 0x59388d78.
//
// Solidity: function decreaseApprovalPreSignedHashing(_token address, _spender address, _subtractedValue uint256, _fee uint256, _nonce uint256) constant returns(bytes32)
func (_PropsToken *PropsTokenSession) DecreaseApprovalPreSignedHashing(_token common.Address, _spender common.Address, _subtractedValue *big.Int, _fee *big.Int, _nonce *big.Int) ([32]byte, error) {
	return _PropsToken.Contract.DecreaseApprovalPreSignedHashing(&_PropsToken.CallOpts, _token, _spender, _subtractedValue, _fee, _nonce)
}

// DecreaseApprovalPreSignedHashing is a free data retrieval call binding the contract method 0x59388d78.
//
// Solidity: function decreaseApprovalPreSignedHashing(_token address, _spender address, _subtractedValue uint256, _fee uint256, _nonce uint256) constant returns(bytes32)
func (_PropsToken *PropsTokenCallerSession) DecreaseApprovalPreSignedHashing(_token common.Address, _spender common.Address, _subtractedValue *big.Int, _fee *big.Int, _nonce *big.Int) ([32]byte, error) {
	return _PropsToken.Contract.DecreaseApprovalPreSignedHashing(&_PropsToken.CallOpts, _token, _spender, _subtractedValue, _fee, _nonce)
}

// IncreaseApprovalPreSignedHashing is a free data retrieval call binding the contract method 0xa45f71ff.
//
// Solidity: function increaseApprovalPreSignedHashing(_token address, _spender address, _addedValue uint256, _fee uint256, _nonce uint256) constant returns(bytes32)
func (_PropsToken *PropsTokenCaller) IncreaseApprovalPreSignedHashing(opts *bind.CallOpts, _token common.Address, _spender common.Address, _addedValue *big.Int, _fee *big.Int, _nonce *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _PropsToken.contract.Call(opts, out, "increaseApprovalPreSignedHashing", _token, _spender, _addedValue, _fee, _nonce)
	return *ret0, err
}

// IncreaseApprovalPreSignedHashing is a free data retrieval call binding the contract method 0xa45f71ff.
//
// Solidity: function increaseApprovalPreSignedHashing(_token address, _spender address, _addedValue uint256, _fee uint256, _nonce uint256) constant returns(bytes32)
func (_PropsToken *PropsTokenSession) IncreaseApprovalPreSignedHashing(_token common.Address, _spender common.Address, _addedValue *big.Int, _fee *big.Int, _nonce *big.Int) ([32]byte, error) {
	return _PropsToken.Contract.IncreaseApprovalPreSignedHashing(&_PropsToken.CallOpts, _token, _spender, _addedValue, _fee, _nonce)
}

// IncreaseApprovalPreSignedHashing is a free data retrieval call binding the contract method 0xa45f71ff.
//
// Solidity: function increaseApprovalPreSignedHashing(_token address, _spender address, _addedValue uint256, _fee uint256, _nonce uint256) constant returns(bytes32)
func (_PropsToken *PropsTokenCallerSession) IncreaseApprovalPreSignedHashing(_token common.Address, _spender common.Address, _addedValue *big.Int, _fee *big.Int, _nonce *big.Int) ([32]byte, error) {
	return _PropsToken.Contract.IncreaseApprovalPreSignedHashing(&_PropsToken.CallOpts, _token, _spender, _addedValue, _fee, _nonce)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_PropsToken *PropsTokenCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _PropsToken.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_PropsToken *PropsTokenSession) Name() (string, error) {
	return _PropsToken.Contract.Name(&_PropsToken.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_PropsToken *PropsTokenCallerSession) Name() (string, error) {
	return _PropsToken.Contract.Name(&_PropsToken.CallOpts)
}

// Recover is a free data retrieval call binding the contract method 0x19045a25.
//
// Solidity: function recover(hash bytes32, sig bytes) constant returns(address)
func (_PropsToken *PropsTokenCaller) Recover(opts *bind.CallOpts, hash [32]byte, sig []byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PropsToken.contract.Call(opts, out, "recover", hash, sig)
	return *ret0, err
}

// Recover is a free data retrieval call binding the contract method 0x19045a25.
//
// Solidity: function recover(hash bytes32, sig bytes) constant returns(address)
func (_PropsToken *PropsTokenSession) Recover(hash [32]byte, sig []byte) (common.Address, error) {
	return _PropsToken.Contract.Recover(&_PropsToken.CallOpts, hash, sig)
}

// Recover is a free data retrieval call binding the contract method 0x19045a25.
//
// Solidity: function recover(hash bytes32, sig bytes) constant returns(address)
func (_PropsToken *PropsTokenCallerSession) Recover(hash [32]byte, sig []byte) (common.Address, error) {
	return _PropsToken.Contract.Recover(&_PropsToken.CallOpts, hash, sig)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_PropsToken *PropsTokenCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _PropsToken.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_PropsToken *PropsTokenSession) Symbol() (string, error) {
	return _PropsToken.Contract.Symbol(&_PropsToken.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_PropsToken *PropsTokenCallerSession) Symbol() (string, error) {
	return _PropsToken.Contract.Symbol(&_PropsToken.CallOpts)
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

// TransferFromPreSignedHashing is a free data retrieval call binding the contract method 0xb7656dc5.
//
// Solidity: function transferFromPreSignedHashing(_token address, _from address, _to address, _value uint256, _fee uint256, _nonce uint256) constant returns(bytes32)
func (_PropsToken *PropsTokenCaller) TransferFromPreSignedHashing(opts *bind.CallOpts, _token common.Address, _from common.Address, _to common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _PropsToken.contract.Call(opts, out, "transferFromPreSignedHashing", _token, _from, _to, _value, _fee, _nonce)
	return *ret0, err
}

// TransferFromPreSignedHashing is a free data retrieval call binding the contract method 0xb7656dc5.
//
// Solidity: function transferFromPreSignedHashing(_token address, _from address, _to address, _value uint256, _fee uint256, _nonce uint256) constant returns(bytes32)
func (_PropsToken *PropsTokenSession) TransferFromPreSignedHashing(_token common.Address, _from common.Address, _to common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) ([32]byte, error) {
	return _PropsToken.Contract.TransferFromPreSignedHashing(&_PropsToken.CallOpts, _token, _from, _to, _value, _fee, _nonce)
}

// TransferFromPreSignedHashing is a free data retrieval call binding the contract method 0xb7656dc5.
//
// Solidity: function transferFromPreSignedHashing(_token address, _from address, _to address, _value uint256, _fee uint256, _nonce uint256) constant returns(bytes32)
func (_PropsToken *PropsTokenCallerSession) TransferFromPreSignedHashing(_token common.Address, _from common.Address, _to common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) ([32]byte, error) {
	return _PropsToken.Contract.TransferFromPreSignedHashing(&_PropsToken.CallOpts, _token, _from, _to, _value, _fee, _nonce)
}

// TransferPreSignedHashing is a free data retrieval call binding the contract method 0x15420b71.
//
// Solidity: function transferPreSignedHashing(_token address, _to address, _value uint256, _fee uint256, _nonce uint256) constant returns(bytes32)
func (_PropsToken *PropsTokenCaller) TransferPreSignedHashing(opts *bind.CallOpts, _token common.Address, _to common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _PropsToken.contract.Call(opts, out, "transferPreSignedHashing", _token, _to, _value, _fee, _nonce)
	return *ret0, err
}

// TransferPreSignedHashing is a free data retrieval call binding the contract method 0x15420b71.
//
// Solidity: function transferPreSignedHashing(_token address, _to address, _value uint256, _fee uint256, _nonce uint256) constant returns(bytes32)
func (_PropsToken *PropsTokenSession) TransferPreSignedHashing(_token common.Address, _to common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) ([32]byte, error) {
	return _PropsToken.Contract.TransferPreSignedHashing(&_PropsToken.CallOpts, _token, _to, _value, _fee, _nonce)
}

// TransferPreSignedHashing is a free data retrieval call binding the contract method 0x15420b71.
//
// Solidity: function transferPreSignedHashing(_token address, _to address, _value uint256, _fee uint256, _nonce uint256) constant returns(bytes32)
func (_PropsToken *PropsTokenCallerSession) TransferPreSignedHashing(_token common.Address, _to common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) ([32]byte, error) {
	return _PropsToken.Contract.TransferPreSignedHashing(&_PropsToken.CallOpts, _token, _to, _value, _fee, _nonce)
}

// TransfersStartTime is a free data retrieval call binding the contract method 0xd0f22e9c.
//
// Solidity: function transfersStartTime() constant returns(uint256)
func (_PropsToken *PropsTokenCaller) TransfersStartTime(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PropsToken.contract.Call(opts, out, "transfersStartTime")
	return *ret0, err
}

// TransfersStartTime is a free data retrieval call binding the contract method 0xd0f22e9c.
//
// Solidity: function transfersStartTime() constant returns(uint256)
func (_PropsToken *PropsTokenSession) TransfersStartTime() (*big.Int, error) {
	return _PropsToken.Contract.TransfersStartTime(&_PropsToken.CallOpts)
}

// TransfersStartTime is a free data retrieval call binding the contract method 0xd0f22e9c.
//
// Solidity: function transfersStartTime() constant returns(uint256)
func (_PropsToken *PropsTokenCallerSession) TransfersStartTime() (*big.Int, error) {
	return _PropsToken.Contract.TransfersStartTime(&_PropsToken.CallOpts)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_PropsToken *PropsTokenTransactor) Approve(opts *bind.TransactOpts, spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _PropsToken.contract.Transact(opts, "approve", spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_PropsToken *PropsTokenSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _PropsToken.Contract.Approve(&_PropsToken.TransactOpts, spender, value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(spender address, value uint256) returns(bool)
func (_PropsToken *PropsTokenTransactorSession) Approve(spender common.Address, value *big.Int) (*types.Transaction, error) {
	return _PropsToken.Contract.Approve(&_PropsToken.TransactOpts, spender, value)
}

// ApprovePreSigned is a paid mutator transaction binding the contract method 0x617b390b.
//
// Solidity: function approvePreSigned(_signature bytes, _spender address, _value uint256, _fee uint256, _nonce uint256) returns(bool)
func (_PropsToken *PropsTokenTransactor) ApprovePreSigned(opts *bind.TransactOpts, _signature []byte, _spender common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _PropsToken.contract.Transact(opts, "approvePreSigned", _signature, _spender, _value, _fee, _nonce)
}

// ApprovePreSigned is a paid mutator transaction binding the contract method 0x617b390b.
//
// Solidity: function approvePreSigned(_signature bytes, _spender address, _value uint256, _fee uint256, _nonce uint256) returns(bool)
func (_PropsToken *PropsTokenSession) ApprovePreSigned(_signature []byte, _spender common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _PropsToken.Contract.ApprovePreSigned(&_PropsToken.TransactOpts, _signature, _spender, _value, _fee, _nonce)
}

// ApprovePreSigned is a paid mutator transaction binding the contract method 0x617b390b.
//
// Solidity: function approvePreSigned(_signature bytes, _spender address, _value uint256, _fee uint256, _nonce uint256) returns(bool)
func (_PropsToken *PropsTokenTransactorSession) ApprovePreSigned(_signature []byte, _spender common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _PropsToken.Contract.ApprovePreSigned(&_PropsToken.TransactOpts, _signature, _spender, _value, _fee, _nonce)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(spender address, subtractedValue uint256) returns(bool)
func (_PropsToken *PropsTokenTransactor) DecreaseAllowance(opts *bind.TransactOpts, spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _PropsToken.contract.Transact(opts, "decreaseAllowance", spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(spender address, subtractedValue uint256) returns(bool)
func (_PropsToken *PropsTokenSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _PropsToken.Contract.DecreaseAllowance(&_PropsToken.TransactOpts, spender, subtractedValue)
}

// DecreaseAllowance is a paid mutator transaction binding the contract method 0xa457c2d7.
//
// Solidity: function decreaseAllowance(spender address, subtractedValue uint256) returns(bool)
func (_PropsToken *PropsTokenTransactorSession) DecreaseAllowance(spender common.Address, subtractedValue *big.Int) (*types.Transaction, error) {
	return _PropsToken.Contract.DecreaseAllowance(&_PropsToken.TransactOpts, spender, subtractedValue)
}

// DecreaseApprovalPreSigned is a paid mutator transaction binding the contract method 0x8be52783.
//
// Solidity: function decreaseApprovalPreSigned(_signature bytes, _spender address, _subtractedValue uint256, _fee uint256, _nonce uint256) returns(bool)
func (_PropsToken *PropsTokenTransactor) DecreaseApprovalPreSigned(opts *bind.TransactOpts, _signature []byte, _spender common.Address, _subtractedValue *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _PropsToken.contract.Transact(opts, "decreaseApprovalPreSigned", _signature, _spender, _subtractedValue, _fee, _nonce)
}

// DecreaseApprovalPreSigned is a paid mutator transaction binding the contract method 0x8be52783.
//
// Solidity: function decreaseApprovalPreSigned(_signature bytes, _spender address, _subtractedValue uint256, _fee uint256, _nonce uint256) returns(bool)
func (_PropsToken *PropsTokenSession) DecreaseApprovalPreSigned(_signature []byte, _spender common.Address, _subtractedValue *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _PropsToken.Contract.DecreaseApprovalPreSigned(&_PropsToken.TransactOpts, _signature, _spender, _subtractedValue, _fee, _nonce)
}

// DecreaseApprovalPreSigned is a paid mutator transaction binding the contract method 0x8be52783.
//
// Solidity: function decreaseApprovalPreSigned(_signature bytes, _spender address, _subtractedValue uint256, _fee uint256, _nonce uint256) returns(bool)
func (_PropsToken *PropsTokenTransactorSession) DecreaseApprovalPreSigned(_signature []byte, _spender common.Address, _subtractedValue *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _PropsToken.Contract.DecreaseApprovalPreSigned(&_PropsToken.TransactOpts, _signature, _spender, _subtractedValue, _fee, _nonce)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(spender address, addedValue uint256) returns(bool)
func (_PropsToken *PropsTokenTransactor) IncreaseAllowance(opts *bind.TransactOpts, spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _PropsToken.contract.Transact(opts, "increaseAllowance", spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(spender address, addedValue uint256) returns(bool)
func (_PropsToken *PropsTokenSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _PropsToken.Contract.IncreaseAllowance(&_PropsToken.TransactOpts, spender, addedValue)
}

// IncreaseAllowance is a paid mutator transaction binding the contract method 0x39509351.
//
// Solidity: function increaseAllowance(spender address, addedValue uint256) returns(bool)
func (_PropsToken *PropsTokenTransactorSession) IncreaseAllowance(spender common.Address, addedValue *big.Int) (*types.Transaction, error) {
	return _PropsToken.Contract.IncreaseAllowance(&_PropsToken.TransactOpts, spender, addedValue)
}

// IncreaseApprovalPreSigned is a paid mutator transaction binding the contract method 0xadb8249e.
//
// Solidity: function increaseApprovalPreSigned(_signature bytes, _spender address, _addedValue uint256, _fee uint256, _nonce uint256) returns(bool)
func (_PropsToken *PropsTokenTransactor) IncreaseApprovalPreSigned(opts *bind.TransactOpts, _signature []byte, _spender common.Address, _addedValue *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _PropsToken.contract.Transact(opts, "increaseApprovalPreSigned", _signature, _spender, _addedValue, _fee, _nonce)
}

// IncreaseApprovalPreSigned is a paid mutator transaction binding the contract method 0xadb8249e.
//
// Solidity: function increaseApprovalPreSigned(_signature bytes, _spender address, _addedValue uint256, _fee uint256, _nonce uint256) returns(bool)
func (_PropsToken *PropsTokenSession) IncreaseApprovalPreSigned(_signature []byte, _spender common.Address, _addedValue *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _PropsToken.Contract.IncreaseApprovalPreSigned(&_PropsToken.TransactOpts, _signature, _spender, _addedValue, _fee, _nonce)
}

// IncreaseApprovalPreSigned is a paid mutator transaction binding the contract method 0xadb8249e.
//
// Solidity: function increaseApprovalPreSigned(_signature bytes, _spender address, _addedValue uint256, _fee uint256, _nonce uint256) returns(bool)
func (_PropsToken *PropsTokenTransactorSession) IncreaseApprovalPreSigned(_signature []byte, _spender common.Address, _addedValue *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _PropsToken.Contract.IncreaseApprovalPreSigned(&_PropsToken.TransactOpts, _signature, _spender, _addedValue, _fee, _nonce)
}

// Initialize is a paid mutator transaction binding the contract method 0xda35a26f.
//
// Solidity: function initialize(start uint256, account address) returns()
func (_PropsToken *PropsTokenTransactor) Initialize(opts *bind.TransactOpts, start *big.Int, account common.Address) (*types.Transaction, error) {
	return _PropsToken.contract.Transact(opts, "initialize", start, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xda35a26f.
//
// Solidity: function initialize(start uint256, account address) returns()
func (_PropsToken *PropsTokenSession) Initialize(start *big.Int, account common.Address) (*types.Transaction, error) {
	return _PropsToken.Contract.Initialize(&_PropsToken.TransactOpts, start, account)
}

// Initialize is a paid mutator transaction binding the contract method 0xda35a26f.
//
// Solidity: function initialize(start uint256, account address) returns()
func (_PropsToken *PropsTokenTransactorSession) Initialize(start *big.Int, account common.Address) (*types.Transaction, error) {
	return _PropsToken.Contract.Initialize(&_PropsToken.TransactOpts, start, account)
}

// Settle is a paid mutator transaction binding the contract method 0x15afd409.
//
// Solidity: function settle(to address, value uint256) returns(bool)
func (_PropsToken *PropsTokenTransactor) Settle(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PropsToken.contract.Transact(opts, "settle", to, value)
}

// Settle is a paid mutator transaction binding the contract method 0x15afd409.
//
// Solidity: function settle(to address, value uint256) returns(bool)
func (_PropsToken *PropsTokenSession) Settle(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PropsToken.Contract.Settle(&_PropsToken.TransactOpts, to, value)
}

// Settle is a paid mutator transaction binding the contract method 0x15afd409.
//
// Solidity: function settle(to address, value uint256) returns(bool)
func (_PropsToken *PropsTokenTransactorSession) Settle(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PropsToken.Contract.Settle(&_PropsToken.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_PropsToken *PropsTokenTransactor) Transfer(opts *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PropsToken.contract.Transact(opts, "transfer", to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_PropsToken *PropsTokenSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PropsToken.Contract.Transfer(&_PropsToken.TransactOpts, to, value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(to address, value uint256) returns(bool)
func (_PropsToken *PropsTokenTransactorSession) Transfer(to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PropsToken.Contract.Transfer(&_PropsToken.TransactOpts, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_PropsToken *PropsTokenTransactor) TransferFrom(opts *bind.TransactOpts, from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PropsToken.contract.Transact(opts, "transferFrom", from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_PropsToken *PropsTokenSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PropsToken.Contract.TransferFrom(&_PropsToken.TransactOpts, from, to, value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(from address, to address, value uint256) returns(bool)
func (_PropsToken *PropsTokenTransactorSession) TransferFrom(from common.Address, to common.Address, value *big.Int) (*types.Transaction, error) {
	return _PropsToken.Contract.TransferFrom(&_PropsToken.TransactOpts, from, to, value)
}

// TransferFromPreSigned is a paid mutator transaction binding the contract method 0xbca50515.
//
// Solidity: function transferFromPreSigned(_signature bytes, _from address, _to address, _value uint256, _fee uint256, _nonce uint256) returns(bool)
func (_PropsToken *PropsTokenTransactor) TransferFromPreSigned(opts *bind.TransactOpts, _signature []byte, _from common.Address, _to common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _PropsToken.contract.Transact(opts, "transferFromPreSigned", _signature, _from, _to, _value, _fee, _nonce)
}

// TransferFromPreSigned is a paid mutator transaction binding the contract method 0xbca50515.
//
// Solidity: function transferFromPreSigned(_signature bytes, _from address, _to address, _value uint256, _fee uint256, _nonce uint256) returns(bool)
func (_PropsToken *PropsTokenSession) TransferFromPreSigned(_signature []byte, _from common.Address, _to common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _PropsToken.Contract.TransferFromPreSigned(&_PropsToken.TransactOpts, _signature, _from, _to, _value, _fee, _nonce)
}

// TransferFromPreSigned is a paid mutator transaction binding the contract method 0xbca50515.
//
// Solidity: function transferFromPreSigned(_signature bytes, _from address, _to address, _value uint256, _fee uint256, _nonce uint256) returns(bool)
func (_PropsToken *PropsTokenTransactorSession) TransferFromPreSigned(_signature []byte, _from common.Address, _to common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _PropsToken.Contract.TransferFromPreSigned(&_PropsToken.TransactOpts, _signature, _from, _to, _value, _fee, _nonce)
}

// TransferPreSigned is a paid mutator transaction binding the contract method 0x1296830d.
//
// Solidity: function transferPreSigned(_signature bytes, _to address, _value uint256, _fee uint256, _nonce uint256) returns(bool)
func (_PropsToken *PropsTokenTransactor) TransferPreSigned(opts *bind.TransactOpts, _signature []byte, _to common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _PropsToken.contract.Transact(opts, "transferPreSigned", _signature, _to, _value, _fee, _nonce)
}

// TransferPreSigned is a paid mutator transaction binding the contract method 0x1296830d.
//
// Solidity: function transferPreSigned(_signature bytes, _to address, _value uint256, _fee uint256, _nonce uint256) returns(bool)
func (_PropsToken *PropsTokenSession) TransferPreSigned(_signature []byte, _to common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _PropsToken.Contract.TransferPreSigned(&_PropsToken.TransactOpts, _signature, _to, _value, _fee, _nonce)
}

// TransferPreSigned is a paid mutator transaction binding the contract method 0x1296830d.
//
// Solidity: function transferPreSigned(_signature bytes, _to address, _value uint256, _fee uint256, _nonce uint256) returns(bool)
func (_PropsToken *PropsTokenTransactorSession) TransferPreSigned(_signature []byte, _to common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _PropsToken.Contract.TransferPreSigned(&_PropsToken.TransactOpts, _signature, _to, _value, _fee, _nonce)
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

// PropsTokenApprovalPreSignedIterator is returned from FilterApprovalPreSigned and is used to iterate over the raw logs and unpacked data for ApprovalPreSigned events raised by the PropsToken contract.
type PropsTokenApprovalPreSignedIterator struct {
	Event *PropsTokenApprovalPreSigned // Event containing the contract specifics and raw log

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
func (it *PropsTokenApprovalPreSignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropsTokenApprovalPreSigned)
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
		it.Event = new(PropsTokenApprovalPreSigned)
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
func (it *PropsTokenApprovalPreSignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropsTokenApprovalPreSignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropsTokenApprovalPreSigned represents a ApprovalPreSigned event raised by the PropsToken contract.
type PropsTokenApprovalPreSigned struct {
	From     common.Address
	To       common.Address
	Delegate common.Address
	Amount   *big.Int
	Fee      *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterApprovalPreSigned is a free log retrieval operation binding the contract event 0x43a220267705e74ee2ceafd46afc841850db6f85a662189a7def697bbdd90ffb.
//
// Solidity: e ApprovalPreSigned(from indexed address, to indexed address, delegate indexed address, amount uint256, fee uint256)
func (_PropsToken *PropsTokenFilterer) FilterApprovalPreSigned(opts *bind.FilterOpts, from []common.Address, to []common.Address, delegate []common.Address) (*PropsTokenApprovalPreSignedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _PropsToken.contract.FilterLogs(opts, "ApprovalPreSigned", fromRule, toRule, delegateRule)
	if err != nil {
		return nil, err
	}
	return &PropsTokenApprovalPreSignedIterator{contract: _PropsToken.contract, event: "ApprovalPreSigned", logs: logs, sub: sub}, nil
}

// WatchApprovalPreSigned is a free log subscription operation binding the contract event 0x43a220267705e74ee2ceafd46afc841850db6f85a662189a7def697bbdd90ffb.
//
// Solidity: e ApprovalPreSigned(from indexed address, to indexed address, delegate indexed address, amount uint256, fee uint256)
func (_PropsToken *PropsTokenFilterer) WatchApprovalPreSigned(opts *bind.WatchOpts, sink chan<- *PropsTokenApprovalPreSigned, from []common.Address, to []common.Address, delegate []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _PropsToken.contract.WatchLogs(opts, "ApprovalPreSigned", fromRule, toRule, delegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropsTokenApprovalPreSigned)
				if err := _PropsToken.contract.UnpackLog(event, "ApprovalPreSigned", log); err != nil {
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

// PropsTokenSettlementIterator is returned from FilterSettlement and is used to iterate over the raw logs and unpacked data for Settlement events raised by the PropsToken contract.
type PropsTokenSettlementIterator struct {
	Event *PropsTokenSettlement // Event containing the contract specifics and raw log

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
func (it *PropsTokenSettlementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropsTokenSettlement)
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
		it.Event = new(PropsTokenSettlement)
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
func (it *PropsTokenSettlementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropsTokenSettlementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropsTokenSettlement represents a Settlement event raised by the PropsToken contract.
type PropsTokenSettlement struct {
	Timestamp *big.Int
	From      common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSettlement is a free log retrieval operation binding the contract event 0x694d4af865c381294c69b304b233e65cd773b9dff4f5563d50d5424fe5a33c3e.
//
// Solidity: e Settlement(timestamp uint256, from address, recipient address, amount uint256)
func (_PropsToken *PropsTokenFilterer) FilterSettlement(opts *bind.FilterOpts) (*PropsTokenSettlementIterator, error) {

	logs, sub, err := _PropsToken.contract.FilterLogs(opts, "Settlement")
	if err != nil {
		return nil, err
	}
	return &PropsTokenSettlementIterator{contract: _PropsToken.contract, event: "Settlement", logs: logs, sub: sub}, nil
}

// WatchSettlement is a free log subscription operation binding the contract event 0x694d4af865c381294c69b304b233e65cd773b9dff4f5563d50d5424fe5a33c3e.
//
// Solidity: e Settlement(timestamp uint256, from address, recipient address, amount uint256)
func (_PropsToken *PropsTokenFilterer) WatchSettlement(opts *bind.WatchOpts, sink chan<- *PropsTokenSettlement) (event.Subscription, error) {

	logs, sub, err := _PropsToken.contract.WatchLogs(opts, "Settlement")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropsTokenSettlement)
				if err := _PropsToken.contract.UnpackLog(event, "Settlement", log); err != nil {
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

// PropsTokenTransferDetailsIterator is returned from FilterTransferDetails and is used to iterate over the raw logs and unpacked data for TransferDetails events raised by the PropsToken contract.
type PropsTokenTransferDetailsIterator struct {
	Event *PropsTokenTransferDetails // Event containing the contract specifics and raw log

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
func (it *PropsTokenTransferDetailsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropsTokenTransferDetails)
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
		it.Event = new(PropsTokenTransferDetails)
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
func (it *PropsTokenTransferDetailsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropsTokenTransferDetailsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropsTokenTransferDetails represents a TransferDetails event raised by the PropsToken contract.
type PropsTokenTransferDetails struct {
	From        common.Address
	FromBalance *big.Int
	To          common.Address
	ToBalance   *big.Int
	Amount      *big.Int
	Raw         types.Log // Blockchain specific contextual infos
}

// FilterTransferDetails is a free log retrieval operation binding the contract event 0x850383053a4e0cc8b97828882fc80543c308bd075cc5bc454b28243814fa8846.
//
// Solidity: e TransferDetails(from address, fromBalance uint256, to address, toBalance uint256, amount uint256)
func (_PropsToken *PropsTokenFilterer) FilterTransferDetails(opts *bind.FilterOpts) (*PropsTokenTransferDetailsIterator, error) {

	logs, sub, err := _PropsToken.contract.FilterLogs(opts, "TransferDetails")
	if err != nil {
		return nil, err
	}
	return &PropsTokenTransferDetailsIterator{contract: _PropsToken.contract, event: "TransferDetails", logs: logs, sub: sub}, nil
}

// WatchTransferDetails is a free log subscription operation binding the contract event 0x850383053a4e0cc8b97828882fc80543c308bd075cc5bc454b28243814fa8846.
//
// Solidity: e TransferDetails(from address, fromBalance uint256, to address, toBalance uint256, amount uint256)
func (_PropsToken *PropsTokenFilterer) WatchTransferDetails(opts *bind.WatchOpts, sink chan<- *PropsTokenTransferDetails) (event.Subscription, error) {

	logs, sub, err := _PropsToken.contract.WatchLogs(opts, "TransferDetails")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropsTokenTransferDetails)
				if err := _PropsToken.contract.UnpackLog(event, "TransferDetails", log); err != nil {
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

// PropsTokenTransferPreSignedIterator is returned from FilterTransferPreSigned and is used to iterate over the raw logs and unpacked data for TransferPreSigned events raised by the PropsToken contract.
type PropsTokenTransferPreSignedIterator struct {
	Event *PropsTokenTransferPreSigned // Event containing the contract specifics and raw log

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
func (it *PropsTokenTransferPreSignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropsTokenTransferPreSigned)
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
		it.Event = new(PropsTokenTransferPreSigned)
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
func (it *PropsTokenTransferPreSignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropsTokenTransferPreSignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropsTokenTransferPreSigned represents a TransferPreSigned event raised by the PropsToken contract.
type PropsTokenTransferPreSigned struct {
	From     common.Address
	To       common.Address
	Delegate common.Address
	Amount   *big.Int
	Fee      *big.Int
	Raw      types.Log // Blockchain specific contextual infos
}

// FilterTransferPreSigned is a free log retrieval operation binding the contract event 0xec5a73fd1f178be20c1bca1b406cbf4b5c20d833b66e582fc122fb4baa0fc2a4.
//
// Solidity: e TransferPreSigned(from indexed address, to indexed address, delegate indexed address, amount uint256, fee uint256)
func (_PropsToken *PropsTokenFilterer) FilterTransferPreSigned(opts *bind.FilterOpts, from []common.Address, to []common.Address, delegate []common.Address) (*PropsTokenTransferPreSignedIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _PropsToken.contract.FilterLogs(opts, "TransferPreSigned", fromRule, toRule, delegateRule)
	if err != nil {
		return nil, err
	}
	return &PropsTokenTransferPreSignedIterator{contract: _PropsToken.contract, event: "TransferPreSigned", logs: logs, sub: sub}, nil
}

// WatchTransferPreSigned is a free log subscription operation binding the contract event 0xec5a73fd1f178be20c1bca1b406cbf4b5c20d833b66e582fc122fb4baa0fc2a4.
//
// Solidity: e TransferPreSigned(from indexed address, to indexed address, delegate indexed address, amount uint256, fee uint256)
func (_PropsToken *PropsTokenFilterer) WatchTransferPreSigned(opts *bind.WatchOpts, sink chan<- *PropsTokenTransferPreSigned, from []common.Address, to []common.Address, delegate []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}
	var delegateRule []interface{}
	for _, delegateItem := range delegate {
		delegateRule = append(delegateRule, delegateItem)
	}

	logs, sub, err := _PropsToken.contract.WatchLogs(opts, "TransferPreSigned", fromRule, toRule, delegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropsTokenTransferPreSigned)
				if err := _PropsToken.contract.UnpackLog(event, "TransferPreSigned", log); err != nil {
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
