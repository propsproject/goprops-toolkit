// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package rinkeby

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

// PropsTokenRinkebyABI is the input ABI used to generate the binding from.
const PropsTokenRinkebyABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"mintingFinished\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"name\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"approve\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_signature\",\"type\":\"bytes\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"transferPreSigned\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"transferPreSignedHashing\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"settle\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"totalSupply\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"hash\",\"type\":\"bytes32\"},{\"name\":\"sig\",\"type\":\"bytes\"}],\"name\":\"recover\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transferFrom\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"blacklistUserForTransfers\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"isTransferWhitelistOnly\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"decimals\",\"outputs\":[{\"name\":\"\",\"type\":\"uint8\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"unpause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_amount\",\"type\":\"uint256\"}],\"name\":\"mint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"amountOfTokenToMint\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"decreaseApprovalPreSignedHashing\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"paused\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_signature\",\"type\":\"bytes\"},{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"approvePreSigned\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"}],\"name\":\"decreaseApproval\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"}],\"name\":\"balanceOf\",\"outputs\":[{\"name\":\"balance\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"finishMinting\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[],\"name\":\"pause\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_signature\",\"type\":\"bytes\"},{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_subtractedValue\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"decreaseApprovalPreSigned\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"owner\",\"outputs\":[{\"name\":\"\",\"type\":\"address\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"symbol\",\"outputs\":[{\"name\":\"\",\"type\":\"string\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_isWhitelistOnly\",\"type\":\"bool\"}],\"name\":\"setWhitelistedOnly\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"isUserAllowedToTransfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"increaseApprovalPreSignedHashing\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"}],\"name\":\"transfer\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_signature\",\"type\":\"bytes\"},{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"increaseApprovalPreSigned\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"transferFromPreSignedHashing\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_signature\",\"type\":\"bytes\"},{\"name\":\"_from\",\"type\":\"address\"},{\"name\":\"_to\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"transferFromPreSigned\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_addedValue\",\"type\":\"uint256\"}],\"name\":\"increaseApproval\",\"outputs\":[{\"name\":\"success\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_owner\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"}],\"name\":\"allowance\",\"outputs\":[{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"transferOwnership\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[{\"name\":\"_token\",\"type\":\"address\"},{\"name\":\"_spender\",\"type\":\"address\"},{\"name\":\"_value\",\"type\":\"uint256\"},{\"name\":\"_fee\",\"type\":\"uint256\"},{\"name\":\"_nonce\",\"type\":\"uint256\"}],\"name\":\"approvePreSignedHashing\",\"outputs\":[{\"name\":\"\",\"type\":\"bytes32\"}],\"payable\":false,\"stateMutability\":\"pure\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_user\",\"type\":\"address\"}],\"name\":\"whitelistUserForTransfers\",\"outputs\":[],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"user\",\"type\":\"address\"}],\"name\":\"UserAllowedToTransfer\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"flag\",\"type\":\"bool\"}],\"name\":\"TransferWhitelistOnly\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"fromBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"toBalance\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"TransferDetails\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"name\":\"timestamp\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"recipient\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Settlement\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"}],\"name\":\"Mint\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"MintFinished\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Pause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[],\"name\":\"Unpause\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"previousOwner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"newOwner\",\"type\":\"address\"}],\"name\":\"OwnershipTransferred\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"delegate\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"TransferPreSigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"delegate\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"amount\",\"type\":\"uint256\"},{\"indexed\":false,\"name\":\"fee\",\"type\":\"uint256\"}],\"name\":\"ApprovalPreSigned\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"owner\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"spender\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Approval\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"from\",\"type\":\"address\"},{\"indexed\":true,\"name\":\"to\",\"type\":\"address\"},{\"indexed\":false,\"name\":\"value\",\"type\":\"uint256\"}],\"name\":\"Transfer\",\"type\":\"event\"}]"

// PropsTokenRinkeby is an auto generated Go binding around an Ethereum contract.
type PropsTokenRinkeby struct {
	PropsTokenRinkebyCaller     // Read-only binding to the contract
	PropsTokenRinkebyTransactor // Write-only binding to the contract
	PropsTokenRinkebyFilterer   // Log filterer for contract events
}

// PropsTokenRinkebyCaller is an auto generated read-only Go binding around an Ethereum contract.
type PropsTokenRinkebyCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PropsTokenRinkebyTransactor is an auto generated write-only Go binding around an Ethereum contract.
type PropsTokenRinkebyTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PropsTokenRinkebyFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type PropsTokenRinkebyFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// PropsTokenRinkebySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type PropsTokenRinkebySession struct {
	Contract     *PropsTokenRinkeby // Generic contract binding to set the session for
	CallOpts     bind.CallOpts      // Call options to use throughout this session
	TransactOpts bind.TransactOpts  // Transaction auth options to use throughout this session
}

// PropsTokenRinkebyCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type PropsTokenRinkebyCallerSession struct {
	Contract *PropsTokenRinkebyCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts            // Call options to use throughout this session
}

// PropsTokenRinkebyTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type PropsTokenRinkebyTransactorSession struct {
	Contract     *PropsTokenRinkebyTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts            // Transaction auth options to use throughout this session
}

// PropsTokenRinkebyRaw is an auto generated low-level Go binding around an Ethereum contract.
type PropsTokenRinkebyRaw struct {
	Contract *PropsTokenRinkeby // Generic contract binding to access the raw methods on
}

// PropsTokenRinkebyCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type PropsTokenRinkebyCallerRaw struct {
	Contract *PropsTokenRinkebyCaller // Generic read-only contract binding to access the raw methods on
}

// PropsTokenRinkebyTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type PropsTokenRinkebyTransactorRaw struct {
	Contract *PropsTokenRinkebyTransactor // Generic write-only contract binding to access the raw methods on
}

// NewPropsTokenRinkeby creates a new instance of PropsTokenRinkeby, bound to a specific deployed contract.
func NewPropsTokenRinkeby(address common.Address, backend bind.ContractBackend) (*PropsTokenRinkeby, error) {
	contract, err := bindPropsTokenRinkeby(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &PropsTokenRinkeby{PropsTokenRinkebyCaller: PropsTokenRinkebyCaller{contract: contract}, PropsTokenRinkebyTransactor: PropsTokenRinkebyTransactor{contract: contract}, PropsTokenRinkebyFilterer: PropsTokenRinkebyFilterer{contract: contract}}, nil
}

// NewPropsTokenRinkebyCaller creates a new read-only instance of PropsTokenRinkeby, bound to a specific deployed contract.
func NewPropsTokenRinkebyCaller(address common.Address, caller bind.ContractCaller) (*PropsTokenRinkebyCaller, error) {
	contract, err := bindPropsTokenRinkeby(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &PropsTokenRinkebyCaller{contract: contract}, nil
}

// NewPropsTokenRinkebyTransactor creates a new write-only instance of PropsTokenRinkeby, bound to a specific deployed contract.
func NewPropsTokenRinkebyTransactor(address common.Address, transactor bind.ContractTransactor) (*PropsTokenRinkebyTransactor, error) {
	contract, err := bindPropsTokenRinkeby(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &PropsTokenRinkebyTransactor{contract: contract}, nil
}

// NewPropsTokenRinkebyFilterer creates a new log filterer instance of PropsTokenRinkeby, bound to a specific deployed contract.
func NewPropsTokenRinkebyFilterer(address common.Address, filterer bind.ContractFilterer) (*PropsTokenRinkebyFilterer, error) {
	contract, err := bindPropsTokenRinkeby(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &PropsTokenRinkebyFilterer{contract: contract}, nil
}

// bindPropsTokenRinkeby binds a generic wrapper to an already deployed contract.
func bindPropsTokenRinkeby(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(PropsTokenRinkebyABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PropsTokenRinkeby *PropsTokenRinkebyRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PropsTokenRinkeby.Contract.PropsTokenRinkebyCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PropsTokenRinkeby *PropsTokenRinkebyRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PropsTokenRinkeby.Contract.PropsTokenRinkebyTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PropsTokenRinkeby *PropsTokenRinkebyRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PropsTokenRinkeby.Contract.PropsTokenRinkebyTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_PropsTokenRinkeby *PropsTokenRinkebyCallerRaw) Call(opts *bind.CallOpts, result interface{}, method string, params ...interface{}) error {
	return _PropsTokenRinkeby.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_PropsTokenRinkeby *PropsTokenRinkebyTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PropsTokenRinkeby.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_PropsTokenRinkeby *PropsTokenRinkebyTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _PropsTokenRinkeby.Contract.contract.Transact(opts, method, params...)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_PropsTokenRinkeby *PropsTokenRinkebyCaller) Allowance(opts *bind.CallOpts, _owner common.Address, _spender common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PropsTokenRinkeby.contract.Call(opts, out, "allowance", _owner, _spender)
	return *ret0, err
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_PropsTokenRinkeby *PropsTokenRinkebySession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _PropsTokenRinkeby.Contract.Allowance(&_PropsTokenRinkeby.CallOpts, _owner, _spender)
}

// Allowance is a free data retrieval call binding the contract method 0xdd62ed3e.
//
// Solidity: function allowance(_owner address, _spender address) constant returns(uint256)
func (_PropsTokenRinkeby *PropsTokenRinkebyCallerSession) Allowance(_owner common.Address, _spender common.Address) (*big.Int, error) {
	return _PropsTokenRinkeby.Contract.Allowance(&_PropsTokenRinkeby.CallOpts, _owner, _spender)
}

// AmountOfTokenToMint is a free data retrieval call binding the contract method 0x5478bc6b.
//
// Solidity: function amountOfTokenToMint() constant returns(uint256)
func (_PropsTokenRinkeby *PropsTokenRinkebyCaller) AmountOfTokenToMint(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PropsTokenRinkeby.contract.Call(opts, out, "amountOfTokenToMint")
	return *ret0, err
}

// AmountOfTokenToMint is a free data retrieval call binding the contract method 0x5478bc6b.
//
// Solidity: function amountOfTokenToMint() constant returns(uint256)
func (_PropsTokenRinkeby *PropsTokenRinkebySession) AmountOfTokenToMint() (*big.Int, error) {
	return _PropsTokenRinkeby.Contract.AmountOfTokenToMint(&_PropsTokenRinkeby.CallOpts)
}

// AmountOfTokenToMint is a free data retrieval call binding the contract method 0x5478bc6b.
//
// Solidity: function amountOfTokenToMint() constant returns(uint256)
func (_PropsTokenRinkeby *PropsTokenRinkebyCallerSession) AmountOfTokenToMint() (*big.Int, error) {
	return _PropsTokenRinkeby.Contract.AmountOfTokenToMint(&_PropsTokenRinkeby.CallOpts)
}

// ApprovePreSignedHashing is a free data retrieval call binding the contract method 0xf7ac9c2e.
//
// Solidity: function approvePreSignedHashing(_token address, _spender address, _value uint256, _fee uint256, _nonce uint256) constant returns(bytes32)
func (_PropsTokenRinkeby *PropsTokenRinkebyCaller) ApprovePreSignedHashing(opts *bind.CallOpts, _token common.Address, _spender common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _PropsTokenRinkeby.contract.Call(opts, out, "approvePreSignedHashing", _token, _spender, _value, _fee, _nonce)
	return *ret0, err
}

// ApprovePreSignedHashing is a free data retrieval call binding the contract method 0xf7ac9c2e.
//
// Solidity: function approvePreSignedHashing(_token address, _spender address, _value uint256, _fee uint256, _nonce uint256) constant returns(bytes32)
func (_PropsTokenRinkeby *PropsTokenRinkebySession) ApprovePreSignedHashing(_token common.Address, _spender common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) ([32]byte, error) {
	return _PropsTokenRinkeby.Contract.ApprovePreSignedHashing(&_PropsTokenRinkeby.CallOpts, _token, _spender, _value, _fee, _nonce)
}

// ApprovePreSignedHashing is a free data retrieval call binding the contract method 0xf7ac9c2e.
//
// Solidity: function approvePreSignedHashing(_token address, _spender address, _value uint256, _fee uint256, _nonce uint256) constant returns(bytes32)
func (_PropsTokenRinkeby *PropsTokenRinkebyCallerSession) ApprovePreSignedHashing(_token common.Address, _spender common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) ([32]byte, error) {
	return _PropsTokenRinkeby.Contract.ApprovePreSignedHashing(&_PropsTokenRinkeby.CallOpts, _token, _spender, _value, _fee, _nonce)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_PropsTokenRinkeby *PropsTokenRinkebyCaller) BalanceOf(opts *bind.CallOpts, _owner common.Address) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PropsTokenRinkeby.contract.Call(opts, out, "balanceOf", _owner)
	return *ret0, err
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_PropsTokenRinkeby *PropsTokenRinkebySession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _PropsTokenRinkeby.Contract.BalanceOf(&_PropsTokenRinkeby.CallOpts, _owner)
}

// BalanceOf is a free data retrieval call binding the contract method 0x70a08231.
//
// Solidity: function balanceOf(_owner address) constant returns(balance uint256)
func (_PropsTokenRinkeby *PropsTokenRinkebyCallerSession) BalanceOf(_owner common.Address) (*big.Int, error) {
	return _PropsTokenRinkeby.Contract.BalanceOf(&_PropsTokenRinkeby.CallOpts, _owner)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_PropsTokenRinkeby *PropsTokenRinkebyCaller) Decimals(opts *bind.CallOpts) (uint8, error) {
	var (
		ret0 = new(uint8)
	)
	out := ret0
	err := _PropsTokenRinkeby.contract.Call(opts, out, "decimals")
	return *ret0, err
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_PropsTokenRinkeby *PropsTokenRinkebySession) Decimals() (uint8, error) {
	return _PropsTokenRinkeby.Contract.Decimals(&_PropsTokenRinkeby.CallOpts)
}

// Decimals is a free data retrieval call binding the contract method 0x313ce567.
//
// Solidity: function decimals() constant returns(uint8)
func (_PropsTokenRinkeby *PropsTokenRinkebyCallerSession) Decimals() (uint8, error) {
	return _PropsTokenRinkeby.Contract.Decimals(&_PropsTokenRinkeby.CallOpts)
}

// DecreaseApprovalPreSignedHashing is a free data retrieval call binding the contract method 0x59388d78.
//
// Solidity: function decreaseApprovalPreSignedHashing(_token address, _spender address, _subtractedValue uint256, _fee uint256, _nonce uint256) constant returns(bytes32)
func (_PropsTokenRinkeby *PropsTokenRinkebyCaller) DecreaseApprovalPreSignedHashing(opts *bind.CallOpts, _token common.Address, _spender common.Address, _subtractedValue *big.Int, _fee *big.Int, _nonce *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _PropsTokenRinkeby.contract.Call(opts, out, "decreaseApprovalPreSignedHashing", _token, _spender, _subtractedValue, _fee, _nonce)
	return *ret0, err
}

// DecreaseApprovalPreSignedHashing is a free data retrieval call binding the contract method 0x59388d78.
//
// Solidity: function decreaseApprovalPreSignedHashing(_token address, _spender address, _subtractedValue uint256, _fee uint256, _nonce uint256) constant returns(bytes32)
func (_PropsTokenRinkeby *PropsTokenRinkebySession) DecreaseApprovalPreSignedHashing(_token common.Address, _spender common.Address, _subtractedValue *big.Int, _fee *big.Int, _nonce *big.Int) ([32]byte, error) {
	return _PropsTokenRinkeby.Contract.DecreaseApprovalPreSignedHashing(&_PropsTokenRinkeby.CallOpts, _token, _spender, _subtractedValue, _fee, _nonce)
}

// DecreaseApprovalPreSignedHashing is a free data retrieval call binding the contract method 0x59388d78.
//
// Solidity: function decreaseApprovalPreSignedHashing(_token address, _spender address, _subtractedValue uint256, _fee uint256, _nonce uint256) constant returns(bytes32)
func (_PropsTokenRinkeby *PropsTokenRinkebyCallerSession) DecreaseApprovalPreSignedHashing(_token common.Address, _spender common.Address, _subtractedValue *big.Int, _fee *big.Int, _nonce *big.Int) ([32]byte, error) {
	return _PropsTokenRinkeby.Contract.DecreaseApprovalPreSignedHashing(&_PropsTokenRinkeby.CallOpts, _token, _spender, _subtractedValue, _fee, _nonce)
}

// IncreaseApprovalPreSignedHashing is a free data retrieval call binding the contract method 0xa45f71ff.
//
// Solidity: function increaseApprovalPreSignedHashing(_token address, _spender address, _addedValue uint256, _fee uint256, _nonce uint256) constant returns(bytes32)
func (_PropsTokenRinkeby *PropsTokenRinkebyCaller) IncreaseApprovalPreSignedHashing(opts *bind.CallOpts, _token common.Address, _spender common.Address, _addedValue *big.Int, _fee *big.Int, _nonce *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _PropsTokenRinkeby.contract.Call(opts, out, "increaseApprovalPreSignedHashing", _token, _spender, _addedValue, _fee, _nonce)
	return *ret0, err
}

// IncreaseApprovalPreSignedHashing is a free data retrieval call binding the contract method 0xa45f71ff.
//
// Solidity: function increaseApprovalPreSignedHashing(_token address, _spender address, _addedValue uint256, _fee uint256, _nonce uint256) constant returns(bytes32)
func (_PropsTokenRinkeby *PropsTokenRinkebySession) IncreaseApprovalPreSignedHashing(_token common.Address, _spender common.Address, _addedValue *big.Int, _fee *big.Int, _nonce *big.Int) ([32]byte, error) {
	return _PropsTokenRinkeby.Contract.IncreaseApprovalPreSignedHashing(&_PropsTokenRinkeby.CallOpts, _token, _spender, _addedValue, _fee, _nonce)
}

// IncreaseApprovalPreSignedHashing is a free data retrieval call binding the contract method 0xa45f71ff.
//
// Solidity: function increaseApprovalPreSignedHashing(_token address, _spender address, _addedValue uint256, _fee uint256, _nonce uint256) constant returns(bytes32)
func (_PropsTokenRinkeby *PropsTokenRinkebyCallerSession) IncreaseApprovalPreSignedHashing(_token common.Address, _spender common.Address, _addedValue *big.Int, _fee *big.Int, _nonce *big.Int) ([32]byte, error) {
	return _PropsTokenRinkeby.Contract.IncreaseApprovalPreSignedHashing(&_PropsTokenRinkeby.CallOpts, _token, _spender, _addedValue, _fee, _nonce)
}

// IsTransferWhitelistOnly is a free data retrieval call binding the contract method 0x2b2be309.
//
// Solidity: function isTransferWhitelistOnly() constant returns(bool)
func (_PropsTokenRinkeby *PropsTokenRinkebyCaller) IsTransferWhitelistOnly(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PropsTokenRinkeby.contract.Call(opts, out, "isTransferWhitelistOnly")
	return *ret0, err
}

// IsTransferWhitelistOnly is a free data retrieval call binding the contract method 0x2b2be309.
//
// Solidity: function isTransferWhitelistOnly() constant returns(bool)
func (_PropsTokenRinkeby *PropsTokenRinkebySession) IsTransferWhitelistOnly() (bool, error) {
	return _PropsTokenRinkeby.Contract.IsTransferWhitelistOnly(&_PropsTokenRinkeby.CallOpts)
}

// IsTransferWhitelistOnly is a free data retrieval call binding the contract method 0x2b2be309.
//
// Solidity: function isTransferWhitelistOnly() constant returns(bool)
func (_PropsTokenRinkeby *PropsTokenRinkebyCallerSession) IsTransferWhitelistOnly() (bool, error) {
	return _PropsTokenRinkeby.Contract.IsTransferWhitelistOnly(&_PropsTokenRinkeby.CallOpts)
}

// IsUserAllowedToTransfer is a free data retrieval call binding the contract method 0xa25b9384.
//
// Solidity: function isUserAllowedToTransfer(_user address) constant returns(bool)
func (_PropsTokenRinkeby *PropsTokenRinkebyCaller) IsUserAllowedToTransfer(opts *bind.CallOpts, _user common.Address) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PropsTokenRinkeby.contract.Call(opts, out, "isUserAllowedToTransfer", _user)
	return *ret0, err
}

// IsUserAllowedToTransfer is a free data retrieval call binding the contract method 0xa25b9384.
//
// Solidity: function isUserAllowedToTransfer(_user address) constant returns(bool)
func (_PropsTokenRinkeby *PropsTokenRinkebySession) IsUserAllowedToTransfer(_user common.Address) (bool, error) {
	return _PropsTokenRinkeby.Contract.IsUserAllowedToTransfer(&_PropsTokenRinkeby.CallOpts, _user)
}

// IsUserAllowedToTransfer is a free data retrieval call binding the contract method 0xa25b9384.
//
// Solidity: function isUserAllowedToTransfer(_user address) constant returns(bool)
func (_PropsTokenRinkeby *PropsTokenRinkebyCallerSession) IsUserAllowedToTransfer(_user common.Address) (bool, error) {
	return _PropsTokenRinkeby.Contract.IsUserAllowedToTransfer(&_PropsTokenRinkeby.CallOpts, _user)
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_PropsTokenRinkeby *PropsTokenRinkebyCaller) MintingFinished(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PropsTokenRinkeby.contract.Call(opts, out, "mintingFinished")
	return *ret0, err
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_PropsTokenRinkeby *PropsTokenRinkebySession) MintingFinished() (bool, error) {
	return _PropsTokenRinkeby.Contract.MintingFinished(&_PropsTokenRinkeby.CallOpts)
}

// MintingFinished is a free data retrieval call binding the contract method 0x05d2035b.
//
// Solidity: function mintingFinished() constant returns(bool)
func (_PropsTokenRinkeby *PropsTokenRinkebyCallerSession) MintingFinished() (bool, error) {
	return _PropsTokenRinkeby.Contract.MintingFinished(&_PropsTokenRinkeby.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_PropsTokenRinkeby *PropsTokenRinkebyCaller) Name(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _PropsTokenRinkeby.contract.Call(opts, out, "name")
	return *ret0, err
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_PropsTokenRinkeby *PropsTokenRinkebySession) Name() (string, error) {
	return _PropsTokenRinkeby.Contract.Name(&_PropsTokenRinkeby.CallOpts)
}

// Name is a free data retrieval call binding the contract method 0x06fdde03.
//
// Solidity: function name() constant returns(string)
func (_PropsTokenRinkeby *PropsTokenRinkebyCallerSession) Name() (string, error) {
	return _PropsTokenRinkeby.Contract.Name(&_PropsTokenRinkeby.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PropsTokenRinkeby *PropsTokenRinkebyCaller) Owner(opts *bind.CallOpts) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PropsTokenRinkeby.contract.Call(opts, out, "owner")
	return *ret0, err
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PropsTokenRinkeby *PropsTokenRinkebySession) Owner() (common.Address, error) {
	return _PropsTokenRinkeby.Contract.Owner(&_PropsTokenRinkeby.CallOpts)
}

// Owner is a free data retrieval call binding the contract method 0x8da5cb5b.
//
// Solidity: function owner() constant returns(address)
func (_PropsTokenRinkeby *PropsTokenRinkebyCallerSession) Owner() (common.Address, error) {
	return _PropsTokenRinkeby.Contract.Owner(&_PropsTokenRinkeby.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_PropsTokenRinkeby *PropsTokenRinkebyCaller) Paused(opts *bind.CallOpts) (bool, error) {
	var (
		ret0 = new(bool)
	)
	out := ret0
	err := _PropsTokenRinkeby.contract.Call(opts, out, "paused")
	return *ret0, err
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_PropsTokenRinkeby *PropsTokenRinkebySession) Paused() (bool, error) {
	return _PropsTokenRinkeby.Contract.Paused(&_PropsTokenRinkeby.CallOpts)
}

// Paused is a free data retrieval call binding the contract method 0x5c975abb.
//
// Solidity: function paused() constant returns(bool)
func (_PropsTokenRinkeby *PropsTokenRinkebyCallerSession) Paused() (bool, error) {
	return _PropsTokenRinkeby.Contract.Paused(&_PropsTokenRinkeby.CallOpts)
}

// Recover is a free data retrieval call binding the contract method 0x19045a25.
//
// Solidity: function recover(hash bytes32, sig bytes) constant returns(address)
func (_PropsTokenRinkeby *PropsTokenRinkebyCaller) Recover(opts *bind.CallOpts, hash [32]byte, sig []byte) (common.Address, error) {
	var (
		ret0 = new(common.Address)
	)
	out := ret0
	err := _PropsTokenRinkeby.contract.Call(opts, out, "recover", hash, sig)
	return *ret0, err
}

// Recover is a free data retrieval call binding the contract method 0x19045a25.
//
// Solidity: function recover(hash bytes32, sig bytes) constant returns(address)
func (_PropsTokenRinkeby *PropsTokenRinkebySession) Recover(hash [32]byte, sig []byte) (common.Address, error) {
	return _PropsTokenRinkeby.Contract.Recover(&_PropsTokenRinkeby.CallOpts, hash, sig)
}

// Recover is a free data retrieval call binding the contract method 0x19045a25.
//
// Solidity: function recover(hash bytes32, sig bytes) constant returns(address)
func (_PropsTokenRinkeby *PropsTokenRinkebyCallerSession) Recover(hash [32]byte, sig []byte) (common.Address, error) {
	return _PropsTokenRinkeby.Contract.Recover(&_PropsTokenRinkeby.CallOpts, hash, sig)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_PropsTokenRinkeby *PropsTokenRinkebyCaller) Symbol(opts *bind.CallOpts) (string, error) {
	var (
		ret0 = new(string)
	)
	out := ret0
	err := _PropsTokenRinkeby.contract.Call(opts, out, "symbol")
	return *ret0, err
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_PropsTokenRinkeby *PropsTokenRinkebySession) Symbol() (string, error) {
	return _PropsTokenRinkeby.Contract.Symbol(&_PropsTokenRinkeby.CallOpts)
}

// Symbol is a free data retrieval call binding the contract method 0x95d89b41.
//
// Solidity: function symbol() constant returns(string)
func (_PropsTokenRinkeby *PropsTokenRinkebyCallerSession) Symbol() (string, error) {
	return _PropsTokenRinkeby.Contract.Symbol(&_PropsTokenRinkeby.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_PropsTokenRinkeby *PropsTokenRinkebyCaller) TotalSupply(opts *bind.CallOpts) (*big.Int, error) {
	var (
		ret0 = new(*big.Int)
	)
	out := ret0
	err := _PropsTokenRinkeby.contract.Call(opts, out, "totalSupply")
	return *ret0, err
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_PropsTokenRinkeby *PropsTokenRinkebySession) TotalSupply() (*big.Int, error) {
	return _PropsTokenRinkeby.Contract.TotalSupply(&_PropsTokenRinkeby.CallOpts)
}

// TotalSupply is a free data retrieval call binding the contract method 0x18160ddd.
//
// Solidity: function totalSupply() constant returns(uint256)
func (_PropsTokenRinkeby *PropsTokenRinkebyCallerSession) TotalSupply() (*big.Int, error) {
	return _PropsTokenRinkeby.Contract.TotalSupply(&_PropsTokenRinkeby.CallOpts)
}

// TransferFromPreSignedHashing is a free data retrieval call binding the contract method 0xb7656dc5.
//
// Solidity: function transferFromPreSignedHashing(_token address, _from address, _to address, _value uint256, _fee uint256, _nonce uint256) constant returns(bytes32)
func (_PropsTokenRinkeby *PropsTokenRinkebyCaller) TransferFromPreSignedHashing(opts *bind.CallOpts, _token common.Address, _from common.Address, _to common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _PropsTokenRinkeby.contract.Call(opts, out, "transferFromPreSignedHashing", _token, _from, _to, _value, _fee, _nonce)
	return *ret0, err
}

// TransferFromPreSignedHashing is a free data retrieval call binding the contract method 0xb7656dc5.
//
// Solidity: function transferFromPreSignedHashing(_token address, _from address, _to address, _value uint256, _fee uint256, _nonce uint256) constant returns(bytes32)
func (_PropsTokenRinkeby *PropsTokenRinkebySession) TransferFromPreSignedHashing(_token common.Address, _from common.Address, _to common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) ([32]byte, error) {
	return _PropsTokenRinkeby.Contract.TransferFromPreSignedHashing(&_PropsTokenRinkeby.CallOpts, _token, _from, _to, _value, _fee, _nonce)
}

// TransferFromPreSignedHashing is a free data retrieval call binding the contract method 0xb7656dc5.
//
// Solidity: function transferFromPreSignedHashing(_token address, _from address, _to address, _value uint256, _fee uint256, _nonce uint256) constant returns(bytes32)
func (_PropsTokenRinkeby *PropsTokenRinkebyCallerSession) TransferFromPreSignedHashing(_token common.Address, _from common.Address, _to common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) ([32]byte, error) {
	return _PropsTokenRinkeby.Contract.TransferFromPreSignedHashing(&_PropsTokenRinkeby.CallOpts, _token, _from, _to, _value, _fee, _nonce)
}

// TransferPreSignedHashing is a free data retrieval call binding the contract method 0x15420b71.
//
// Solidity: function transferPreSignedHashing(_token address, _to address, _value uint256, _fee uint256, _nonce uint256) constant returns(bytes32)
func (_PropsTokenRinkeby *PropsTokenRinkebyCaller) TransferPreSignedHashing(opts *bind.CallOpts, _token common.Address, _to common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) ([32]byte, error) {
	var (
		ret0 = new([32]byte)
	)
	out := ret0
	err := _PropsTokenRinkeby.contract.Call(opts, out, "transferPreSignedHashing", _token, _to, _value, _fee, _nonce)
	return *ret0, err
}

// TransferPreSignedHashing is a free data retrieval call binding the contract method 0x15420b71.
//
// Solidity: function transferPreSignedHashing(_token address, _to address, _value uint256, _fee uint256, _nonce uint256) constant returns(bytes32)
func (_PropsTokenRinkeby *PropsTokenRinkebySession) TransferPreSignedHashing(_token common.Address, _to common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) ([32]byte, error) {
	return _PropsTokenRinkeby.Contract.TransferPreSignedHashing(&_PropsTokenRinkeby.CallOpts, _token, _to, _value, _fee, _nonce)
}

// TransferPreSignedHashing is a free data retrieval call binding the contract method 0x15420b71.
//
// Solidity: function transferPreSignedHashing(_token address, _to address, _value uint256, _fee uint256, _nonce uint256) constant returns(bytes32)
func (_PropsTokenRinkeby *PropsTokenRinkebyCallerSession) TransferPreSignedHashing(_token common.Address, _to common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) ([32]byte, error) {
	return _PropsTokenRinkeby.Contract.TransferPreSignedHashing(&_PropsTokenRinkeby.CallOpts, _token, _to, _value, _fee, _nonce)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_PropsTokenRinkeby *PropsTokenRinkebyTransactor) Approve(opts *bind.TransactOpts, _spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PropsTokenRinkeby.contract.Transact(opts, "approve", _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_PropsTokenRinkeby *PropsTokenRinkebySession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PropsTokenRinkeby.Contract.Approve(&_PropsTokenRinkeby.TransactOpts, _spender, _value)
}

// Approve is a paid mutator transaction binding the contract method 0x095ea7b3.
//
// Solidity: function approve(_spender address, _value uint256) returns(bool)
func (_PropsTokenRinkeby *PropsTokenRinkebyTransactorSession) Approve(_spender common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PropsTokenRinkeby.Contract.Approve(&_PropsTokenRinkeby.TransactOpts, _spender, _value)
}

// ApprovePreSigned is a paid mutator transaction binding the contract method 0x617b390b.
//
// Solidity: function approvePreSigned(_signature bytes, _spender address, _value uint256, _fee uint256, _nonce uint256) returns(bool)
func (_PropsTokenRinkeby *PropsTokenRinkebyTransactor) ApprovePreSigned(opts *bind.TransactOpts, _signature []byte, _spender common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _PropsTokenRinkeby.contract.Transact(opts, "approvePreSigned", _signature, _spender, _value, _fee, _nonce)
}

// ApprovePreSigned is a paid mutator transaction binding the contract method 0x617b390b.
//
// Solidity: function approvePreSigned(_signature bytes, _spender address, _value uint256, _fee uint256, _nonce uint256) returns(bool)
func (_PropsTokenRinkeby *PropsTokenRinkebySession) ApprovePreSigned(_signature []byte, _spender common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _PropsTokenRinkeby.Contract.ApprovePreSigned(&_PropsTokenRinkeby.TransactOpts, _signature, _spender, _value, _fee, _nonce)
}

// ApprovePreSigned is a paid mutator transaction binding the contract method 0x617b390b.
//
// Solidity: function approvePreSigned(_signature bytes, _spender address, _value uint256, _fee uint256, _nonce uint256) returns(bool)
func (_PropsTokenRinkeby *PropsTokenRinkebyTransactorSession) ApprovePreSigned(_signature []byte, _spender common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _PropsTokenRinkeby.Contract.ApprovePreSigned(&_PropsTokenRinkeby.TransactOpts, _signature, _spender, _value, _fee, _nonce)
}

// BlacklistUserForTransfers is a paid mutator transaction binding the contract method 0x260d387b.
//
// Solidity: function blacklistUserForTransfers(_user address) returns()
func (_PropsTokenRinkeby *PropsTokenRinkebyTransactor) BlacklistUserForTransfers(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _PropsTokenRinkeby.contract.Transact(opts, "blacklistUserForTransfers", _user)
}

// BlacklistUserForTransfers is a paid mutator transaction binding the contract method 0x260d387b.
//
// Solidity: function blacklistUserForTransfers(_user address) returns()
func (_PropsTokenRinkeby *PropsTokenRinkebySession) BlacklistUserForTransfers(_user common.Address) (*types.Transaction, error) {
	return _PropsTokenRinkeby.Contract.BlacklistUserForTransfers(&_PropsTokenRinkeby.TransactOpts, _user)
}

// BlacklistUserForTransfers is a paid mutator transaction binding the contract method 0x260d387b.
//
// Solidity: function blacklistUserForTransfers(_user address) returns()
func (_PropsTokenRinkeby *PropsTokenRinkebyTransactorSession) BlacklistUserForTransfers(_user common.Address) (*types.Transaction, error) {
	return _PropsTokenRinkeby.Contract.BlacklistUserForTransfers(&_PropsTokenRinkeby.TransactOpts, _user)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(success bool)
func (_PropsTokenRinkeby *PropsTokenRinkebyTransactor) DecreaseApproval(opts *bind.TransactOpts, _spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _PropsTokenRinkeby.contract.Transact(opts, "decreaseApproval", _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(success bool)
func (_PropsTokenRinkeby *PropsTokenRinkebySession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _PropsTokenRinkeby.Contract.DecreaseApproval(&_PropsTokenRinkeby.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApproval is a paid mutator transaction binding the contract method 0x66188463.
//
// Solidity: function decreaseApproval(_spender address, _subtractedValue uint256) returns(success bool)
func (_PropsTokenRinkeby *PropsTokenRinkebyTransactorSession) DecreaseApproval(_spender common.Address, _subtractedValue *big.Int) (*types.Transaction, error) {
	return _PropsTokenRinkeby.Contract.DecreaseApproval(&_PropsTokenRinkeby.TransactOpts, _spender, _subtractedValue)
}

// DecreaseApprovalPreSigned is a paid mutator transaction binding the contract method 0x8be52783.
//
// Solidity: function decreaseApprovalPreSigned(_signature bytes, _spender address, _subtractedValue uint256, _fee uint256, _nonce uint256) returns(bool)
func (_PropsTokenRinkeby *PropsTokenRinkebyTransactor) DecreaseApprovalPreSigned(opts *bind.TransactOpts, _signature []byte, _spender common.Address, _subtractedValue *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _PropsTokenRinkeby.contract.Transact(opts, "decreaseApprovalPreSigned", _signature, _spender, _subtractedValue, _fee, _nonce)
}

// DecreaseApprovalPreSigned is a paid mutator transaction binding the contract method 0x8be52783.
//
// Solidity: function decreaseApprovalPreSigned(_signature bytes, _spender address, _subtractedValue uint256, _fee uint256, _nonce uint256) returns(bool)
func (_PropsTokenRinkeby *PropsTokenRinkebySession) DecreaseApprovalPreSigned(_signature []byte, _spender common.Address, _subtractedValue *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _PropsTokenRinkeby.Contract.DecreaseApprovalPreSigned(&_PropsTokenRinkeby.TransactOpts, _signature, _spender, _subtractedValue, _fee, _nonce)
}

// DecreaseApprovalPreSigned is a paid mutator transaction binding the contract method 0x8be52783.
//
// Solidity: function decreaseApprovalPreSigned(_signature bytes, _spender address, _subtractedValue uint256, _fee uint256, _nonce uint256) returns(bool)
func (_PropsTokenRinkeby *PropsTokenRinkebyTransactorSession) DecreaseApprovalPreSigned(_signature []byte, _spender common.Address, _subtractedValue *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _PropsTokenRinkeby.Contract.DecreaseApprovalPreSigned(&_PropsTokenRinkeby.TransactOpts, _signature, _spender, _subtractedValue, _fee, _nonce)
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_PropsTokenRinkeby *PropsTokenRinkebyTransactor) FinishMinting(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PropsTokenRinkeby.contract.Transact(opts, "finishMinting")
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_PropsTokenRinkeby *PropsTokenRinkebySession) FinishMinting() (*types.Transaction, error) {
	return _PropsTokenRinkeby.Contract.FinishMinting(&_PropsTokenRinkeby.TransactOpts)
}

// FinishMinting is a paid mutator transaction binding the contract method 0x7d64bcb4.
//
// Solidity: function finishMinting() returns(bool)
func (_PropsTokenRinkeby *PropsTokenRinkebyTransactorSession) FinishMinting() (*types.Transaction, error) {
	return _PropsTokenRinkeby.Contract.FinishMinting(&_PropsTokenRinkeby.TransactOpts)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(success bool)
func (_PropsTokenRinkeby *PropsTokenRinkebyTransactor) IncreaseApproval(opts *bind.TransactOpts, _spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _PropsTokenRinkeby.contract.Transact(opts, "increaseApproval", _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(success bool)
func (_PropsTokenRinkeby *PropsTokenRinkebySession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _PropsTokenRinkeby.Contract.IncreaseApproval(&_PropsTokenRinkeby.TransactOpts, _spender, _addedValue)
}

// IncreaseApproval is a paid mutator transaction binding the contract method 0xd73dd623.
//
// Solidity: function increaseApproval(_spender address, _addedValue uint256) returns(success bool)
func (_PropsTokenRinkeby *PropsTokenRinkebyTransactorSession) IncreaseApproval(_spender common.Address, _addedValue *big.Int) (*types.Transaction, error) {
	return _PropsTokenRinkeby.Contract.IncreaseApproval(&_PropsTokenRinkeby.TransactOpts, _spender, _addedValue)
}

// IncreaseApprovalPreSigned is a paid mutator transaction binding the contract method 0xadb8249e.
//
// Solidity: function increaseApprovalPreSigned(_signature bytes, _spender address, _addedValue uint256, _fee uint256, _nonce uint256) returns(bool)
func (_PropsTokenRinkeby *PropsTokenRinkebyTransactor) IncreaseApprovalPreSigned(opts *bind.TransactOpts, _signature []byte, _spender common.Address, _addedValue *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _PropsTokenRinkeby.contract.Transact(opts, "increaseApprovalPreSigned", _signature, _spender, _addedValue, _fee, _nonce)
}

// IncreaseApprovalPreSigned is a paid mutator transaction binding the contract method 0xadb8249e.
//
// Solidity: function increaseApprovalPreSigned(_signature bytes, _spender address, _addedValue uint256, _fee uint256, _nonce uint256) returns(bool)
func (_PropsTokenRinkeby *PropsTokenRinkebySession) IncreaseApprovalPreSigned(_signature []byte, _spender common.Address, _addedValue *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _PropsTokenRinkeby.Contract.IncreaseApprovalPreSigned(&_PropsTokenRinkeby.TransactOpts, _signature, _spender, _addedValue, _fee, _nonce)
}

// IncreaseApprovalPreSigned is a paid mutator transaction binding the contract method 0xadb8249e.
//
// Solidity: function increaseApprovalPreSigned(_signature bytes, _spender address, _addedValue uint256, _fee uint256, _nonce uint256) returns(bool)
func (_PropsTokenRinkeby *PropsTokenRinkebyTransactorSession) IncreaseApprovalPreSigned(_signature []byte, _spender common.Address, _addedValue *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _PropsTokenRinkeby.Contract.IncreaseApprovalPreSigned(&_PropsTokenRinkeby.TransactOpts, _signature, _spender, _addedValue, _fee, _nonce)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_PropsTokenRinkeby *PropsTokenRinkebyTransactor) Mint(opts *bind.TransactOpts, _to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PropsTokenRinkeby.contract.Transact(opts, "mint", _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_PropsTokenRinkeby *PropsTokenRinkebySession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PropsTokenRinkeby.Contract.Mint(&_PropsTokenRinkeby.TransactOpts, _to, _amount)
}

// Mint is a paid mutator transaction binding the contract method 0x40c10f19.
//
// Solidity: function mint(_to address, _amount uint256) returns(bool)
func (_PropsTokenRinkeby *PropsTokenRinkebyTransactorSession) Mint(_to common.Address, _amount *big.Int) (*types.Transaction, error) {
	return _PropsTokenRinkeby.Contract.Mint(&_PropsTokenRinkeby.TransactOpts, _to, _amount)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PropsTokenRinkeby *PropsTokenRinkebyTransactor) Pause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PropsTokenRinkeby.contract.Transact(opts, "pause")
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PropsTokenRinkeby *PropsTokenRinkebySession) Pause() (*types.Transaction, error) {
	return _PropsTokenRinkeby.Contract.Pause(&_PropsTokenRinkeby.TransactOpts)
}

// Pause is a paid mutator transaction binding the contract method 0x8456cb59.
//
// Solidity: function pause() returns()
func (_PropsTokenRinkeby *PropsTokenRinkebyTransactorSession) Pause() (*types.Transaction, error) {
	return _PropsTokenRinkeby.Contract.Pause(&_PropsTokenRinkeby.TransactOpts)
}

// SetWhitelistedOnly is a paid mutator transaction binding the contract method 0xa157696b.
//
// Solidity: function setWhitelistedOnly(_isWhitelistOnly bool) returns()
func (_PropsTokenRinkeby *PropsTokenRinkebyTransactor) SetWhitelistedOnly(opts *bind.TransactOpts, _isWhitelistOnly bool) (*types.Transaction, error) {
	return _PropsTokenRinkeby.contract.Transact(opts, "setWhitelistedOnly", _isWhitelistOnly)
}

// SetWhitelistedOnly is a paid mutator transaction binding the contract method 0xa157696b.
//
// Solidity: function setWhitelistedOnly(_isWhitelistOnly bool) returns()
func (_PropsTokenRinkeby *PropsTokenRinkebySession) SetWhitelistedOnly(_isWhitelistOnly bool) (*types.Transaction, error) {
	return _PropsTokenRinkeby.Contract.SetWhitelistedOnly(&_PropsTokenRinkeby.TransactOpts, _isWhitelistOnly)
}

// SetWhitelistedOnly is a paid mutator transaction binding the contract method 0xa157696b.
//
// Solidity: function setWhitelistedOnly(_isWhitelistOnly bool) returns()
func (_PropsTokenRinkeby *PropsTokenRinkebyTransactorSession) SetWhitelistedOnly(_isWhitelistOnly bool) (*types.Transaction, error) {
	return _PropsTokenRinkeby.Contract.SetWhitelistedOnly(&_PropsTokenRinkeby.TransactOpts, _isWhitelistOnly)
}

// Settle is a paid mutator transaction binding the contract method 0x15afd409.
//
// Solidity: function settle(_to address, _value uint256) returns(bool)
func (_PropsTokenRinkeby *PropsTokenRinkebyTransactor) Settle(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PropsTokenRinkeby.contract.Transact(opts, "settle", _to, _value)
}

// Settle is a paid mutator transaction binding the contract method 0x15afd409.
//
// Solidity: function settle(_to address, _value uint256) returns(bool)
func (_PropsTokenRinkeby *PropsTokenRinkebySession) Settle(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PropsTokenRinkeby.Contract.Settle(&_PropsTokenRinkeby.TransactOpts, _to, _value)
}

// Settle is a paid mutator transaction binding the contract method 0x15afd409.
//
// Solidity: function settle(_to address, _value uint256) returns(bool)
func (_PropsTokenRinkeby *PropsTokenRinkebyTransactorSession) Settle(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PropsTokenRinkeby.Contract.Settle(&_PropsTokenRinkeby.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_PropsTokenRinkeby *PropsTokenRinkebyTransactor) Transfer(opts *bind.TransactOpts, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PropsTokenRinkeby.contract.Transact(opts, "transfer", _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_PropsTokenRinkeby *PropsTokenRinkebySession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PropsTokenRinkeby.Contract.Transfer(&_PropsTokenRinkeby.TransactOpts, _to, _value)
}

// Transfer is a paid mutator transaction binding the contract method 0xa9059cbb.
//
// Solidity: function transfer(_to address, _value uint256) returns(bool)
func (_PropsTokenRinkeby *PropsTokenRinkebyTransactorSession) Transfer(_to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PropsTokenRinkeby.Contract.Transfer(&_PropsTokenRinkeby.TransactOpts, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_PropsTokenRinkeby *PropsTokenRinkebyTransactor) TransferFrom(opts *bind.TransactOpts, _from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PropsTokenRinkeby.contract.Transact(opts, "transferFrom", _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_PropsTokenRinkeby *PropsTokenRinkebySession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PropsTokenRinkeby.Contract.TransferFrom(&_PropsTokenRinkeby.TransactOpts, _from, _to, _value)
}

// TransferFrom is a paid mutator transaction binding the contract method 0x23b872dd.
//
// Solidity: function transferFrom(_from address, _to address, _value uint256) returns(bool)
func (_PropsTokenRinkeby *PropsTokenRinkebyTransactorSession) TransferFrom(_from common.Address, _to common.Address, _value *big.Int) (*types.Transaction, error) {
	return _PropsTokenRinkeby.Contract.TransferFrom(&_PropsTokenRinkeby.TransactOpts, _from, _to, _value)
}

// TransferFromPreSigned is a paid mutator transaction binding the contract method 0xbca50515.
//
// Solidity: function transferFromPreSigned(_signature bytes, _from address, _to address, _value uint256, _fee uint256, _nonce uint256) returns(bool)
func (_PropsTokenRinkeby *PropsTokenRinkebyTransactor) TransferFromPreSigned(opts *bind.TransactOpts, _signature []byte, _from common.Address, _to common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _PropsTokenRinkeby.contract.Transact(opts, "transferFromPreSigned", _signature, _from, _to, _value, _fee, _nonce)
}

// TransferFromPreSigned is a paid mutator transaction binding the contract method 0xbca50515.
//
// Solidity: function transferFromPreSigned(_signature bytes, _from address, _to address, _value uint256, _fee uint256, _nonce uint256) returns(bool)
func (_PropsTokenRinkeby *PropsTokenRinkebySession) TransferFromPreSigned(_signature []byte, _from common.Address, _to common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _PropsTokenRinkeby.Contract.TransferFromPreSigned(&_PropsTokenRinkeby.TransactOpts, _signature, _from, _to, _value, _fee, _nonce)
}

// TransferFromPreSigned is a paid mutator transaction binding the contract method 0xbca50515.
//
// Solidity: function transferFromPreSigned(_signature bytes, _from address, _to address, _value uint256, _fee uint256, _nonce uint256) returns(bool)
func (_PropsTokenRinkeby *PropsTokenRinkebyTransactorSession) TransferFromPreSigned(_signature []byte, _from common.Address, _to common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _PropsTokenRinkeby.Contract.TransferFromPreSigned(&_PropsTokenRinkeby.TransactOpts, _signature, _from, _to, _value, _fee, _nonce)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_PropsTokenRinkeby *PropsTokenRinkebyTransactor) TransferOwnership(opts *bind.TransactOpts, newOwner common.Address) (*types.Transaction, error) {
	return _PropsTokenRinkeby.contract.Transact(opts, "transferOwnership", newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_PropsTokenRinkeby *PropsTokenRinkebySession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PropsTokenRinkeby.Contract.TransferOwnership(&_PropsTokenRinkeby.TransactOpts, newOwner)
}

// TransferOwnership is a paid mutator transaction binding the contract method 0xf2fde38b.
//
// Solidity: function transferOwnership(newOwner address) returns()
func (_PropsTokenRinkeby *PropsTokenRinkebyTransactorSession) TransferOwnership(newOwner common.Address) (*types.Transaction, error) {
	return _PropsTokenRinkeby.Contract.TransferOwnership(&_PropsTokenRinkeby.TransactOpts, newOwner)
}

// TransferPreSigned is a paid mutator transaction binding the contract method 0x1296830d.
//
// Solidity: function transferPreSigned(_signature bytes, _to address, _value uint256, _fee uint256, _nonce uint256) returns(bool)
func (_PropsTokenRinkeby *PropsTokenRinkebyTransactor) TransferPreSigned(opts *bind.TransactOpts, _signature []byte, _to common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _PropsTokenRinkeby.contract.Transact(opts, "transferPreSigned", _signature, _to, _value, _fee, _nonce)
}

// TransferPreSigned is a paid mutator transaction binding the contract method 0x1296830d.
//
// Solidity: function transferPreSigned(_signature bytes, _to address, _value uint256, _fee uint256, _nonce uint256) returns(bool)
func (_PropsTokenRinkeby *PropsTokenRinkebySession) TransferPreSigned(_signature []byte, _to common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _PropsTokenRinkeby.Contract.TransferPreSigned(&_PropsTokenRinkeby.TransactOpts, _signature, _to, _value, _fee, _nonce)
}

// TransferPreSigned is a paid mutator transaction binding the contract method 0x1296830d.
//
// Solidity: function transferPreSigned(_signature bytes, _to address, _value uint256, _fee uint256, _nonce uint256) returns(bool)
func (_PropsTokenRinkeby *PropsTokenRinkebyTransactorSession) TransferPreSigned(_signature []byte, _to common.Address, _value *big.Int, _fee *big.Int, _nonce *big.Int) (*types.Transaction, error) {
	return _PropsTokenRinkeby.Contract.TransferPreSigned(&_PropsTokenRinkeby.TransactOpts, _signature, _to, _value, _fee, _nonce)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PropsTokenRinkeby *PropsTokenRinkebyTransactor) Unpause(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _PropsTokenRinkeby.contract.Transact(opts, "unpause")
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PropsTokenRinkeby *PropsTokenRinkebySession) Unpause() (*types.Transaction, error) {
	return _PropsTokenRinkeby.Contract.Unpause(&_PropsTokenRinkeby.TransactOpts)
}

// Unpause is a paid mutator transaction binding the contract method 0x3f4ba83a.
//
// Solidity: function unpause() returns()
func (_PropsTokenRinkeby *PropsTokenRinkebyTransactorSession) Unpause() (*types.Transaction, error) {
	return _PropsTokenRinkeby.Contract.Unpause(&_PropsTokenRinkeby.TransactOpts)
}

// WhitelistUserForTransfers is a paid mutator transaction binding the contract method 0xfb37baa1.
//
// Solidity: function whitelistUserForTransfers(_user address) returns()
func (_PropsTokenRinkeby *PropsTokenRinkebyTransactor) WhitelistUserForTransfers(opts *bind.TransactOpts, _user common.Address) (*types.Transaction, error) {
	return _PropsTokenRinkeby.contract.Transact(opts, "whitelistUserForTransfers", _user)
}

// WhitelistUserForTransfers is a paid mutator transaction binding the contract method 0xfb37baa1.
//
// Solidity: function whitelistUserForTransfers(_user address) returns()
func (_PropsTokenRinkeby *PropsTokenRinkebySession) WhitelistUserForTransfers(_user common.Address) (*types.Transaction, error) {
	return _PropsTokenRinkeby.Contract.WhitelistUserForTransfers(&_PropsTokenRinkeby.TransactOpts, _user)
}

// WhitelistUserForTransfers is a paid mutator transaction binding the contract method 0xfb37baa1.
//
// Solidity: function whitelistUserForTransfers(_user address) returns()
func (_PropsTokenRinkeby *PropsTokenRinkebyTransactorSession) WhitelistUserForTransfers(_user common.Address) (*types.Transaction, error) {
	return _PropsTokenRinkeby.Contract.WhitelistUserForTransfers(&_PropsTokenRinkeby.TransactOpts, _user)
}

// PropsTokenRinkebyApprovalIterator is returned from FilterApproval and is used to iterate over the raw logs and unpacked data for Approval events raised by the PropsTokenRinkeby contract.
type PropsTokenRinkebyApprovalIterator struct {
	Event *PropsTokenRinkebyApproval // Event containing the contract specifics and raw log

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
func (it *PropsTokenRinkebyApprovalIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropsTokenRinkebyApproval)
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
		it.Event = new(PropsTokenRinkebyApproval)
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
func (it *PropsTokenRinkebyApprovalIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropsTokenRinkebyApprovalIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropsTokenRinkebyApproval represents a Approval event raised by the PropsTokenRinkeby contract.
type PropsTokenRinkebyApproval struct {
	Owner   common.Address
	Spender common.Address
	Value   *big.Int
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterApproval is a free log retrieval operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_PropsTokenRinkeby *PropsTokenRinkebyFilterer) FilterApproval(opts *bind.FilterOpts, owner []common.Address, spender []common.Address) (*PropsTokenRinkebyApprovalIterator, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PropsTokenRinkeby.contract.FilterLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return &PropsTokenRinkebyApprovalIterator{contract: _PropsTokenRinkeby.contract, event: "Approval", logs: logs, sub: sub}, nil
}

// WatchApproval is a free log subscription operation binding the contract event 0x8c5be1e5ebec7d5bd14f71427d1e84f3dd0314c0f7b2291e5b200ac8c7c3b925.
//
// Solidity: e Approval(owner indexed address, spender indexed address, value uint256)
func (_PropsTokenRinkeby *PropsTokenRinkebyFilterer) WatchApproval(opts *bind.WatchOpts, sink chan<- *PropsTokenRinkebyApproval, owner []common.Address, spender []common.Address) (event.Subscription, error) {

	var ownerRule []interface{}
	for _, ownerItem := range owner {
		ownerRule = append(ownerRule, ownerItem)
	}
	var spenderRule []interface{}
	for _, spenderItem := range spender {
		spenderRule = append(spenderRule, spenderItem)
	}

	logs, sub, err := _PropsTokenRinkeby.contract.WatchLogs(opts, "Approval", ownerRule, spenderRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropsTokenRinkebyApproval)
				if err := _PropsTokenRinkeby.contract.UnpackLog(event, "Approval", log); err != nil {
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

// PropsTokenRinkebyApprovalPreSignedIterator is returned from FilterApprovalPreSigned and is used to iterate over the raw logs and unpacked data for ApprovalPreSigned events raised by the PropsTokenRinkeby contract.
type PropsTokenRinkebyApprovalPreSignedIterator struct {
	Event *PropsTokenRinkebyApprovalPreSigned // Event containing the contract specifics and raw log

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
func (it *PropsTokenRinkebyApprovalPreSignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropsTokenRinkebyApprovalPreSigned)
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
		it.Event = new(PropsTokenRinkebyApprovalPreSigned)
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
func (it *PropsTokenRinkebyApprovalPreSignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropsTokenRinkebyApprovalPreSignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropsTokenRinkebyApprovalPreSigned represents a ApprovalPreSigned event raised by the PropsTokenRinkeby contract.
type PropsTokenRinkebyApprovalPreSigned struct {
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
func (_PropsTokenRinkeby *PropsTokenRinkebyFilterer) FilterApprovalPreSigned(opts *bind.FilterOpts, from []common.Address, to []common.Address, delegate []common.Address) (*PropsTokenRinkebyApprovalPreSignedIterator, error) {

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

	logs, sub, err := _PropsTokenRinkeby.contract.FilterLogs(opts, "ApprovalPreSigned", fromRule, toRule, delegateRule)
	if err != nil {
		return nil, err
	}
	return &PropsTokenRinkebyApprovalPreSignedIterator{contract: _PropsTokenRinkeby.contract, event: "ApprovalPreSigned", logs: logs, sub: sub}, nil
}

// WatchApprovalPreSigned is a free log subscription operation binding the contract event 0x43a220267705e74ee2ceafd46afc841850db6f85a662189a7def697bbdd90ffb.
//
// Solidity: e ApprovalPreSigned(from indexed address, to indexed address, delegate indexed address, amount uint256, fee uint256)
func (_PropsTokenRinkeby *PropsTokenRinkebyFilterer) WatchApprovalPreSigned(opts *bind.WatchOpts, sink chan<- *PropsTokenRinkebyApprovalPreSigned, from []common.Address, to []common.Address, delegate []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PropsTokenRinkeby.contract.WatchLogs(opts, "ApprovalPreSigned", fromRule, toRule, delegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropsTokenRinkebyApprovalPreSigned)
				if err := _PropsTokenRinkeby.contract.UnpackLog(event, "ApprovalPreSigned", log); err != nil {
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

// PropsTokenRinkebyMintIterator is returned from FilterMint and is used to iterate over the raw logs and unpacked data for Mint events raised by the PropsTokenRinkeby contract.
type PropsTokenRinkebyMintIterator struct {
	Event *PropsTokenRinkebyMint // Event containing the contract specifics and raw log

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
func (it *PropsTokenRinkebyMintIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropsTokenRinkebyMint)
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
		it.Event = new(PropsTokenRinkebyMint)
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
func (it *PropsTokenRinkebyMintIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropsTokenRinkebyMintIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropsTokenRinkebyMint represents a Mint event raised by the PropsTokenRinkeby contract.
type PropsTokenRinkebyMint struct {
	To     common.Address
	Amount *big.Int
	Raw    types.Log // Blockchain specific contextual infos
}

// FilterMint is a free log retrieval operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: e Mint(to indexed address, amount uint256)
func (_PropsTokenRinkeby *PropsTokenRinkebyFilterer) FilterMint(opts *bind.FilterOpts, to []common.Address) (*PropsTokenRinkebyMintIterator, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PropsTokenRinkeby.contract.FilterLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return &PropsTokenRinkebyMintIterator{contract: _PropsTokenRinkeby.contract, event: "Mint", logs: logs, sub: sub}, nil
}

// WatchMint is a free log subscription operation binding the contract event 0x0f6798a560793a54c3bcfe86a93cde1e73087d944c0ea20544137d4121396885.
//
// Solidity: e Mint(to indexed address, amount uint256)
func (_PropsTokenRinkeby *PropsTokenRinkebyFilterer) WatchMint(opts *bind.WatchOpts, sink chan<- *PropsTokenRinkebyMint, to []common.Address) (event.Subscription, error) {

	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PropsTokenRinkeby.contract.WatchLogs(opts, "Mint", toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropsTokenRinkebyMint)
				if err := _PropsTokenRinkeby.contract.UnpackLog(event, "Mint", log); err != nil {
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

// PropsTokenRinkebyMintFinishedIterator is returned from FilterMintFinished and is used to iterate over the raw logs and unpacked data for MintFinished events raised by the PropsTokenRinkeby contract.
type PropsTokenRinkebyMintFinishedIterator struct {
	Event *PropsTokenRinkebyMintFinished // Event containing the contract specifics and raw log

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
func (it *PropsTokenRinkebyMintFinishedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropsTokenRinkebyMintFinished)
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
		it.Event = new(PropsTokenRinkebyMintFinished)
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
func (it *PropsTokenRinkebyMintFinishedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropsTokenRinkebyMintFinishedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropsTokenRinkebyMintFinished represents a MintFinished event raised by the PropsTokenRinkeby contract.
type PropsTokenRinkebyMintFinished struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterMintFinished is a free log retrieval operation binding the contract event 0xae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa08.
//
// Solidity: e MintFinished()
func (_PropsTokenRinkeby *PropsTokenRinkebyFilterer) FilterMintFinished(opts *bind.FilterOpts) (*PropsTokenRinkebyMintFinishedIterator, error) {

	logs, sub, err := _PropsTokenRinkeby.contract.FilterLogs(opts, "MintFinished")
	if err != nil {
		return nil, err
	}
	return &PropsTokenRinkebyMintFinishedIterator{contract: _PropsTokenRinkeby.contract, event: "MintFinished", logs: logs, sub: sub}, nil
}

// WatchMintFinished is a free log subscription operation binding the contract event 0xae5184fba832cb2b1f702aca6117b8d265eaf03ad33eb133f19dde0f5920fa08.
//
// Solidity: e MintFinished()
func (_PropsTokenRinkeby *PropsTokenRinkebyFilterer) WatchMintFinished(opts *bind.WatchOpts, sink chan<- *PropsTokenRinkebyMintFinished) (event.Subscription, error) {

	logs, sub, err := _PropsTokenRinkeby.contract.WatchLogs(opts, "MintFinished")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropsTokenRinkebyMintFinished)
				if err := _PropsTokenRinkeby.contract.UnpackLog(event, "MintFinished", log); err != nil {
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

// PropsTokenRinkebyOwnershipTransferredIterator is returned from FilterOwnershipTransferred and is used to iterate over the raw logs and unpacked data for OwnershipTransferred events raised by the PropsTokenRinkeby contract.
type PropsTokenRinkebyOwnershipTransferredIterator struct {
	Event *PropsTokenRinkebyOwnershipTransferred // Event containing the contract specifics and raw log

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
func (it *PropsTokenRinkebyOwnershipTransferredIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropsTokenRinkebyOwnershipTransferred)
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
		it.Event = new(PropsTokenRinkebyOwnershipTransferred)
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
func (it *PropsTokenRinkebyOwnershipTransferredIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropsTokenRinkebyOwnershipTransferredIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropsTokenRinkebyOwnershipTransferred represents a OwnershipTransferred event raised by the PropsTokenRinkeby contract.
type PropsTokenRinkebyOwnershipTransferred struct {
	PreviousOwner common.Address
	NewOwner      common.Address
	Raw           types.Log // Blockchain specific contextual infos
}

// FilterOwnershipTransferred is a free log retrieval operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_PropsTokenRinkeby *PropsTokenRinkebyFilterer) FilterOwnershipTransferred(opts *bind.FilterOpts, previousOwner []common.Address, newOwner []common.Address) (*PropsTokenRinkebyOwnershipTransferredIterator, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PropsTokenRinkeby.contract.FilterLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return &PropsTokenRinkebyOwnershipTransferredIterator{contract: _PropsTokenRinkeby.contract, event: "OwnershipTransferred", logs: logs, sub: sub}, nil
}

// WatchOwnershipTransferred is a free log subscription operation binding the contract event 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0.
//
// Solidity: e OwnershipTransferred(previousOwner indexed address, newOwner indexed address)
func (_PropsTokenRinkeby *PropsTokenRinkebyFilterer) WatchOwnershipTransferred(opts *bind.WatchOpts, sink chan<- *PropsTokenRinkebyOwnershipTransferred, previousOwner []common.Address, newOwner []common.Address) (event.Subscription, error) {

	var previousOwnerRule []interface{}
	for _, previousOwnerItem := range previousOwner {
		previousOwnerRule = append(previousOwnerRule, previousOwnerItem)
	}
	var newOwnerRule []interface{}
	for _, newOwnerItem := range newOwner {
		newOwnerRule = append(newOwnerRule, newOwnerItem)
	}

	logs, sub, err := _PropsTokenRinkeby.contract.WatchLogs(opts, "OwnershipTransferred", previousOwnerRule, newOwnerRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropsTokenRinkebyOwnershipTransferred)
				if err := _PropsTokenRinkeby.contract.UnpackLog(event, "OwnershipTransferred", log); err != nil {
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

// PropsTokenRinkebyPauseIterator is returned from FilterPause and is used to iterate over the raw logs and unpacked data for Pause events raised by the PropsTokenRinkeby contract.
type PropsTokenRinkebyPauseIterator struct {
	Event *PropsTokenRinkebyPause // Event containing the contract specifics and raw log

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
func (it *PropsTokenRinkebyPauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropsTokenRinkebyPause)
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
		it.Event = new(PropsTokenRinkebyPause)
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
func (it *PropsTokenRinkebyPauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropsTokenRinkebyPauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropsTokenRinkebyPause represents a Pause event raised by the PropsTokenRinkeby contract.
type PropsTokenRinkebyPause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterPause is a free log retrieval operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_PropsTokenRinkeby *PropsTokenRinkebyFilterer) FilterPause(opts *bind.FilterOpts) (*PropsTokenRinkebyPauseIterator, error) {

	logs, sub, err := _PropsTokenRinkeby.contract.FilterLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return &PropsTokenRinkebyPauseIterator{contract: _PropsTokenRinkeby.contract, event: "Pause", logs: logs, sub: sub}, nil
}

// WatchPause is a free log subscription operation binding the contract event 0x6985a02210a168e66602d3235cb6db0e70f92b3ba4d376a33c0f3d9434bff625.
//
// Solidity: e Pause()
func (_PropsTokenRinkeby *PropsTokenRinkebyFilterer) WatchPause(opts *bind.WatchOpts, sink chan<- *PropsTokenRinkebyPause) (event.Subscription, error) {

	logs, sub, err := _PropsTokenRinkeby.contract.WatchLogs(opts, "Pause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropsTokenRinkebyPause)
				if err := _PropsTokenRinkeby.contract.UnpackLog(event, "Pause", log); err != nil {
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

// PropsTokenRinkebySettlementIterator is returned from FilterSettlement and is used to iterate over the raw logs and unpacked data for Settlement events raised by the PropsTokenRinkeby contract.
type PropsTokenRinkebySettlementIterator struct {
	Event *PropsTokenRinkebySettlement // Event containing the contract specifics and raw log

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
func (it *PropsTokenRinkebySettlementIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropsTokenRinkebySettlement)
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
		it.Event = new(PropsTokenRinkebySettlement)
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
func (it *PropsTokenRinkebySettlementIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropsTokenRinkebySettlementIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropsTokenRinkebySettlement represents a Settlement event raised by the PropsTokenRinkeby contract.
type PropsTokenRinkebySettlement struct {
	Timestamp *big.Int
	From      common.Address
	Recipient common.Address
	Amount    *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterSettlement is a free log retrieval operation binding the contract event 0x694d4af865c381294c69b304b233e65cd773b9dff4f5563d50d5424fe5a33c3e.
//
// Solidity: e Settlement(timestamp uint256, from address, recipient address, amount uint256)
func (_PropsTokenRinkeby *PropsTokenRinkebyFilterer) FilterSettlement(opts *bind.FilterOpts) (*PropsTokenRinkebySettlementIterator, error) {

	logs, sub, err := _PropsTokenRinkeby.contract.FilterLogs(opts, "Settlement")
	if err != nil {
		return nil, err
	}
	return &PropsTokenRinkebySettlementIterator{contract: _PropsTokenRinkeby.contract, event: "Settlement", logs: logs, sub: sub}, nil
}

// WatchSettlement is a free log subscription operation binding the contract event 0x694d4af865c381294c69b304b233e65cd773b9dff4f5563d50d5424fe5a33c3e.
//
// Solidity: e Settlement(timestamp uint256, from address, recipient address, amount uint256)
func (_PropsTokenRinkeby *PropsTokenRinkebyFilterer) WatchSettlement(opts *bind.WatchOpts, sink chan<- *PropsTokenRinkebySettlement) (event.Subscription, error) {

	logs, sub, err := _PropsTokenRinkeby.contract.WatchLogs(opts, "Settlement")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropsTokenRinkebySettlement)
				if err := _PropsTokenRinkeby.contract.UnpackLog(event, "Settlement", log); err != nil {
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

// PropsTokenRinkebyTransferIterator is returned from FilterTransfer and is used to iterate over the raw logs and unpacked data for Transfer events raised by the PropsTokenRinkeby contract.
type PropsTokenRinkebyTransferIterator struct {
	Event *PropsTokenRinkebyTransfer // Event containing the contract specifics and raw log

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
func (it *PropsTokenRinkebyTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropsTokenRinkebyTransfer)
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
		it.Event = new(PropsTokenRinkebyTransfer)
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
func (it *PropsTokenRinkebyTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropsTokenRinkebyTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropsTokenRinkebyTransfer represents a Transfer event raised by the PropsTokenRinkeby contract.
type PropsTokenRinkebyTransfer struct {
	From  common.Address
	To    common.Address
	Value *big.Int
	Raw   types.Log // Blockchain specific contextual infos
}

// FilterTransfer is a free log retrieval operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_PropsTokenRinkeby *PropsTokenRinkebyFilterer) FilterTransfer(opts *bind.FilterOpts, from []common.Address, to []common.Address) (*PropsTokenRinkebyTransferIterator, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PropsTokenRinkeby.contract.FilterLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return &PropsTokenRinkebyTransferIterator{contract: _PropsTokenRinkeby.contract, event: "Transfer", logs: logs, sub: sub}, nil
}

// WatchTransfer is a free log subscription operation binding the contract event 0xddf252ad1be2c89b69c2b068fc378daa952ba7f163c4a11628f55a4df523b3ef.
//
// Solidity: e Transfer(from indexed address, to indexed address, value uint256)
func (_PropsTokenRinkeby *PropsTokenRinkebyFilterer) WatchTransfer(opts *bind.WatchOpts, sink chan<- *PropsTokenRinkebyTransfer, from []common.Address, to []common.Address) (event.Subscription, error) {

	var fromRule []interface{}
	for _, fromItem := range from {
		fromRule = append(fromRule, fromItem)
	}
	var toRule []interface{}
	for _, toItem := range to {
		toRule = append(toRule, toItem)
	}

	logs, sub, err := _PropsTokenRinkeby.contract.WatchLogs(opts, "Transfer", fromRule, toRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropsTokenRinkebyTransfer)
				if err := _PropsTokenRinkeby.contract.UnpackLog(event, "Transfer", log); err != nil {
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

// PropsTokenRinkebyTransferDetailsIterator is returned from FilterTransferDetails and is used to iterate over the raw logs and unpacked data for TransferDetails events raised by the PropsTokenRinkeby contract.
type PropsTokenRinkebyTransferDetailsIterator struct {
	Event *PropsTokenRinkebyTransferDetails // Event containing the contract specifics and raw log

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
func (it *PropsTokenRinkebyTransferDetailsIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropsTokenRinkebyTransferDetails)
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
		it.Event = new(PropsTokenRinkebyTransferDetails)
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
func (it *PropsTokenRinkebyTransferDetailsIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropsTokenRinkebyTransferDetailsIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropsTokenRinkebyTransferDetails represents a TransferDetails event raised by the PropsTokenRinkeby contract.
type PropsTokenRinkebyTransferDetails struct {
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
func (_PropsTokenRinkeby *PropsTokenRinkebyFilterer) FilterTransferDetails(opts *bind.FilterOpts) (*PropsTokenRinkebyTransferDetailsIterator, error) {

	logs, sub, err := _PropsTokenRinkeby.contract.FilterLogs(opts, "TransferDetails")
	if err != nil {
		return nil, err
	}
	return &PropsTokenRinkebyTransferDetailsIterator{contract: _PropsTokenRinkeby.contract, event: "TransferDetails", logs: logs, sub: sub}, nil
}

// WatchTransferDetails is a free log subscription operation binding the contract event 0x850383053a4e0cc8b97828882fc80543c308bd075cc5bc454b28243814fa8846.
//
// Solidity: e TransferDetails(from address, fromBalance uint256, to address, toBalance uint256, amount uint256)
func (_PropsTokenRinkeby *PropsTokenRinkebyFilterer) WatchTransferDetails(opts *bind.WatchOpts, sink chan<- *PropsTokenRinkebyTransferDetails) (event.Subscription, error) {

	logs, sub, err := _PropsTokenRinkeby.contract.WatchLogs(opts, "TransferDetails")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropsTokenRinkebyTransferDetails)
				if err := _PropsTokenRinkeby.contract.UnpackLog(event, "TransferDetails", log); err != nil {
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

// PropsTokenRinkebyTransferPreSignedIterator is returned from FilterTransferPreSigned and is used to iterate over the raw logs and unpacked data for TransferPreSigned events raised by the PropsTokenRinkeby contract.
type PropsTokenRinkebyTransferPreSignedIterator struct {
	Event *PropsTokenRinkebyTransferPreSigned // Event containing the contract specifics and raw log

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
func (it *PropsTokenRinkebyTransferPreSignedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropsTokenRinkebyTransferPreSigned)
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
		it.Event = new(PropsTokenRinkebyTransferPreSigned)
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
func (it *PropsTokenRinkebyTransferPreSignedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropsTokenRinkebyTransferPreSignedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropsTokenRinkebyTransferPreSigned represents a TransferPreSigned event raised by the PropsTokenRinkeby contract.
type PropsTokenRinkebyTransferPreSigned struct {
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
func (_PropsTokenRinkeby *PropsTokenRinkebyFilterer) FilterTransferPreSigned(opts *bind.FilterOpts, from []common.Address, to []common.Address, delegate []common.Address) (*PropsTokenRinkebyTransferPreSignedIterator, error) {

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

	logs, sub, err := _PropsTokenRinkeby.contract.FilterLogs(opts, "TransferPreSigned", fromRule, toRule, delegateRule)
	if err != nil {
		return nil, err
	}
	return &PropsTokenRinkebyTransferPreSignedIterator{contract: _PropsTokenRinkeby.contract, event: "TransferPreSigned", logs: logs, sub: sub}, nil
}

// WatchTransferPreSigned is a free log subscription operation binding the contract event 0xec5a73fd1f178be20c1bca1b406cbf4b5c20d833b66e582fc122fb4baa0fc2a4.
//
// Solidity: e TransferPreSigned(from indexed address, to indexed address, delegate indexed address, amount uint256, fee uint256)
func (_PropsTokenRinkeby *PropsTokenRinkebyFilterer) WatchTransferPreSigned(opts *bind.WatchOpts, sink chan<- *PropsTokenRinkebyTransferPreSigned, from []common.Address, to []common.Address, delegate []common.Address) (event.Subscription, error) {

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

	logs, sub, err := _PropsTokenRinkeby.contract.WatchLogs(opts, "TransferPreSigned", fromRule, toRule, delegateRule)
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropsTokenRinkebyTransferPreSigned)
				if err := _PropsTokenRinkeby.contract.UnpackLog(event, "TransferPreSigned", log); err != nil {
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

// PropsTokenRinkebyTransferWhitelistOnlyIterator is returned from FilterTransferWhitelistOnly and is used to iterate over the raw logs and unpacked data for TransferWhitelistOnly events raised by the PropsTokenRinkeby contract.
type PropsTokenRinkebyTransferWhitelistOnlyIterator struct {
	Event *PropsTokenRinkebyTransferWhitelistOnly // Event containing the contract specifics and raw log

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
func (it *PropsTokenRinkebyTransferWhitelistOnlyIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropsTokenRinkebyTransferWhitelistOnly)
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
		it.Event = new(PropsTokenRinkebyTransferWhitelistOnly)
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
func (it *PropsTokenRinkebyTransferWhitelistOnlyIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropsTokenRinkebyTransferWhitelistOnlyIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropsTokenRinkebyTransferWhitelistOnly represents a TransferWhitelistOnly event raised by the PropsTokenRinkeby contract.
type PropsTokenRinkebyTransferWhitelistOnly struct {
	Flag bool
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterTransferWhitelistOnly is a free log retrieval operation binding the contract event 0x757db31e37ddce6c8bda07399e53891f513ac7cacd86ae8f781bab706c74d45f.
//
// Solidity: e TransferWhitelistOnly(flag bool)
func (_PropsTokenRinkeby *PropsTokenRinkebyFilterer) FilterTransferWhitelistOnly(opts *bind.FilterOpts) (*PropsTokenRinkebyTransferWhitelistOnlyIterator, error) {

	logs, sub, err := _PropsTokenRinkeby.contract.FilterLogs(opts, "TransferWhitelistOnly")
	if err != nil {
		return nil, err
	}
	return &PropsTokenRinkebyTransferWhitelistOnlyIterator{contract: _PropsTokenRinkeby.contract, event: "TransferWhitelistOnly", logs: logs, sub: sub}, nil
}

// WatchTransferWhitelistOnly is a free log subscription operation binding the contract event 0x757db31e37ddce6c8bda07399e53891f513ac7cacd86ae8f781bab706c74d45f.
//
// Solidity: e TransferWhitelistOnly(flag bool)
func (_PropsTokenRinkeby *PropsTokenRinkebyFilterer) WatchTransferWhitelistOnly(opts *bind.WatchOpts, sink chan<- *PropsTokenRinkebyTransferWhitelistOnly) (event.Subscription, error) {

	logs, sub, err := _PropsTokenRinkeby.contract.WatchLogs(opts, "TransferWhitelistOnly")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropsTokenRinkebyTransferWhitelistOnly)
				if err := _PropsTokenRinkeby.contract.UnpackLog(event, "TransferWhitelistOnly", log); err != nil {
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

// PropsTokenRinkebyUnpauseIterator is returned from FilterUnpause and is used to iterate over the raw logs and unpacked data for Unpause events raised by the PropsTokenRinkeby contract.
type PropsTokenRinkebyUnpauseIterator struct {
	Event *PropsTokenRinkebyUnpause // Event containing the contract specifics and raw log

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
func (it *PropsTokenRinkebyUnpauseIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropsTokenRinkebyUnpause)
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
		it.Event = new(PropsTokenRinkebyUnpause)
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
func (it *PropsTokenRinkebyUnpauseIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropsTokenRinkebyUnpauseIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropsTokenRinkebyUnpause represents a Unpause event raised by the PropsTokenRinkeby contract.
type PropsTokenRinkebyUnpause struct {
	Raw types.Log // Blockchain specific contextual infos
}

// FilterUnpause is a free log retrieval operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_PropsTokenRinkeby *PropsTokenRinkebyFilterer) FilterUnpause(opts *bind.FilterOpts) (*PropsTokenRinkebyUnpauseIterator, error) {

	logs, sub, err := _PropsTokenRinkeby.contract.FilterLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return &PropsTokenRinkebyUnpauseIterator{contract: _PropsTokenRinkeby.contract, event: "Unpause", logs: logs, sub: sub}, nil
}

// WatchUnpause is a free log subscription operation binding the contract event 0x7805862f689e2f13df9f062ff482ad3ad112aca9e0847911ed832e158c525b33.
//
// Solidity: e Unpause()
func (_PropsTokenRinkeby *PropsTokenRinkebyFilterer) WatchUnpause(opts *bind.WatchOpts, sink chan<- *PropsTokenRinkebyUnpause) (event.Subscription, error) {

	logs, sub, err := _PropsTokenRinkeby.contract.WatchLogs(opts, "Unpause")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropsTokenRinkebyUnpause)
				if err := _PropsTokenRinkeby.contract.UnpackLog(event, "Unpause", log); err != nil {
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

// PropsTokenRinkebyUserAllowedToTransferIterator is returned from FilterUserAllowedToTransfer and is used to iterate over the raw logs and unpacked data for UserAllowedToTransfer events raised by the PropsTokenRinkeby contract.
type PropsTokenRinkebyUserAllowedToTransferIterator struct {
	Event *PropsTokenRinkebyUserAllowedToTransfer // Event containing the contract specifics and raw log

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
func (it *PropsTokenRinkebyUserAllowedToTransferIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(PropsTokenRinkebyUserAllowedToTransfer)
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
		it.Event = new(PropsTokenRinkebyUserAllowedToTransfer)
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
func (it *PropsTokenRinkebyUserAllowedToTransferIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *PropsTokenRinkebyUserAllowedToTransferIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// PropsTokenRinkebyUserAllowedToTransfer represents a UserAllowedToTransfer event raised by the PropsTokenRinkeby contract.
type PropsTokenRinkebyUserAllowedToTransfer struct {
	User common.Address
	Raw  types.Log // Blockchain specific contextual infos
}

// FilterUserAllowedToTransfer is a free log retrieval operation binding the contract event 0xad88aec34f4861d5a8a34d064eba1c3ce7222a541db6a6cfbcef0e6aa0dac53c.
//
// Solidity: e UserAllowedToTransfer(user address)
func (_PropsTokenRinkeby *PropsTokenRinkebyFilterer) FilterUserAllowedToTransfer(opts *bind.FilterOpts) (*PropsTokenRinkebyUserAllowedToTransferIterator, error) {

	logs, sub, err := _PropsTokenRinkeby.contract.FilterLogs(opts, "UserAllowedToTransfer")
	if err != nil {
		return nil, err
	}
	return &PropsTokenRinkebyUserAllowedToTransferIterator{contract: _PropsTokenRinkeby.contract, event: "UserAllowedToTransfer", logs: logs, sub: sub}, nil
}

// WatchUserAllowedToTransfer is a free log subscription operation binding the contract event 0xad88aec34f4861d5a8a34d064eba1c3ce7222a541db6a6cfbcef0e6aa0dac53c.
//
// Solidity: e UserAllowedToTransfer(user address)
func (_PropsTokenRinkeby *PropsTokenRinkebyFilterer) WatchUserAllowedToTransfer(opts *bind.WatchOpts, sink chan<- *PropsTokenRinkebyUserAllowedToTransfer) (event.Subscription, error) {

	logs, sub, err := _PropsTokenRinkeby.contract.WatchLogs(opts, "UserAllowedToTransfer")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(PropsTokenRinkebyUserAllowedToTransfer)
				if err := _PropsTokenRinkeby.contract.UnpackLog(event, "UserAllowedToTransfer", log); err != nil {
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
