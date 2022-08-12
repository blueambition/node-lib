// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package tokenpair

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

// TokenPairABI is the input ABI used to generate the binding from.
const TokenPairABI = "[{\"inputs\":[{\"internalType\":\"address\",\"name\":\"lp\",\"type\":\"address\"},{\"internalType\":\"uint256\",\"name\":\"token0Price\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"token1Price\",\"type\":\"uint256\"}],\"name\":\"LPPrice\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"lp\",\"type\":\"address\"}],\"name\":\"getLPReserves\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"lp\",\"type\":\"address\"}],\"name\":\"getLPTokens\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"router\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenIn\",\"type\":\"address\"},{\"internalType\":\"address\",\"name\":\"tokenOut\",\"type\":\"address\"}],\"name\":\"getTokenOut\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]"

// TokenPair is an auto generated Go binding around an Ethereum contract.
type TokenPair struct {
	TokenPairCaller     // Read-only binding to the contract
	TokenPairTransactor // Write-only binding to the contract
	TokenPairFilterer   // Log filterer for contract events
}

// TokenPairCaller is an auto generated read-only Go binding around an Ethereum contract.
type TokenPairCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenPairTransactor is an auto generated write-only Go binding around an Ethereum contract.
type TokenPairTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenPairFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type TokenPairFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// TokenPairSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type TokenPairSession struct {
	Contract     *TokenPair        // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// TokenPairCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type TokenPairCallerSession struct {
	Contract *TokenPairCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts    // Call options to use throughout this session
}

// TokenPairTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type TokenPairTransactorSession struct {
	Contract     *TokenPairTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts    // Transaction auth options to use throughout this session
}

// TokenPairRaw is an auto generated low-level Go binding around an Ethereum contract.
type TokenPairRaw struct {
	Contract *TokenPair // Generic contract binding to access the raw methods on
}

// TokenPairCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type TokenPairCallerRaw struct {
	Contract *TokenPairCaller // Generic read-only contract binding to access the raw methods on
}

// TokenPairTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type TokenPairTransactorRaw struct {
	Contract *TokenPairTransactor // Generic write-only contract binding to access the raw methods on
}

// NewTokenPair creates a new instance of TokenPair, bound to a specific deployed contract.
func NewTokenPair(address common.Address, backend bind.ContractBackend) (*TokenPair, error) {
	contract, err := bindTokenPair(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &TokenPair{TokenPairCaller: TokenPairCaller{contract: contract}, TokenPairTransactor: TokenPairTransactor{contract: contract}, TokenPairFilterer: TokenPairFilterer{contract: contract}}, nil
}

// NewTokenPairCaller creates a new read-only instance of TokenPair, bound to a specific deployed contract.
func NewTokenPairCaller(address common.Address, caller bind.ContractCaller) (*TokenPairCaller, error) {
	contract, err := bindTokenPair(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &TokenPairCaller{contract: contract}, nil
}

// NewTokenPairTransactor creates a new write-only instance of TokenPair, bound to a specific deployed contract.
func NewTokenPairTransactor(address common.Address, transactor bind.ContractTransactor) (*TokenPairTransactor, error) {
	contract, err := bindTokenPair(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &TokenPairTransactor{contract: contract}, nil
}

// NewTokenPairFilterer creates a new log filterer instance of TokenPair, bound to a specific deployed contract.
func NewTokenPairFilterer(address common.Address, filterer bind.ContractFilterer) (*TokenPairFilterer, error) {
	contract, err := bindTokenPair(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &TokenPairFilterer{contract: contract}, nil
}

// bindTokenPair binds a generic wrapper to an already deployed contract.
func bindTokenPair(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(TokenPairABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenPair *TokenPairRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenPair.Contract.TokenPairCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenPair *TokenPairRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenPair.Contract.TokenPairTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenPair *TokenPairRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenPair.Contract.TokenPairTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_TokenPair *TokenPairCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _TokenPair.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_TokenPair *TokenPairTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _TokenPair.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_TokenPair *TokenPairTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _TokenPair.Contract.contract.Transact(opts, method, params...)
}

// LPPrice is a free data retrieval call binding the contract method 0x19110747.
//
// Solidity: function LPPrice(address lp, uint256 token0Price, uint256 token1Price) view returns(uint256)
func (_TokenPair *TokenPairCaller) LPPrice(opts *bind.CallOpts, lp common.Address, token0Price *big.Int, token1Price *big.Int) (*big.Int, error) {
	var out []interface{}
	err := _TokenPair.contract.Call(opts, &out, "LPPrice", lp, token0Price, token1Price)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// LPPrice is a free data retrieval call binding the contract method 0x19110747.
//
// Solidity: function LPPrice(address lp, uint256 token0Price, uint256 token1Price) view returns(uint256)
func (_TokenPair *TokenPairSession) LPPrice(lp common.Address, token0Price *big.Int, token1Price *big.Int) (*big.Int, error) {
	return _TokenPair.Contract.LPPrice(&_TokenPair.CallOpts, lp, token0Price, token1Price)
}

// LPPrice is a free data retrieval call binding the contract method 0x19110747.
//
// Solidity: function LPPrice(address lp, uint256 token0Price, uint256 token1Price) view returns(uint256)
func (_TokenPair *TokenPairCallerSession) LPPrice(lp common.Address, token0Price *big.Int, token1Price *big.Int) (*big.Int, error) {
	return _TokenPair.Contract.LPPrice(&_TokenPair.CallOpts, lp, token0Price, token1Price)
}

// GetLPReserves is a free data retrieval call binding the contract method 0xa1d28ac2.
//
// Solidity: function getLPReserves(address lp) view returns(uint256, uint256)
func (_TokenPair *TokenPairCaller) GetLPReserves(opts *bind.CallOpts, lp common.Address) (*big.Int, *big.Int, error) {
	var out []interface{}
	err := _TokenPair.contract.Call(opts, &out, "getLPReserves", lp)

	if err != nil {
		return *new(*big.Int), *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)
	out1 := *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)

	return out0, out1, err

}

// GetLPReserves is a free data retrieval call binding the contract method 0xa1d28ac2.
//
// Solidity: function getLPReserves(address lp) view returns(uint256, uint256)
func (_TokenPair *TokenPairSession) GetLPReserves(lp common.Address) (*big.Int, *big.Int, error) {
	return _TokenPair.Contract.GetLPReserves(&_TokenPair.CallOpts, lp)
}

// GetLPReserves is a free data retrieval call binding the contract method 0xa1d28ac2.
//
// Solidity: function getLPReserves(address lp) view returns(uint256, uint256)
func (_TokenPair *TokenPairCallerSession) GetLPReserves(lp common.Address) (*big.Int, *big.Int, error) {
	return _TokenPair.Contract.GetLPReserves(&_TokenPair.CallOpts, lp)
}

// GetLPTokens is a free data retrieval call binding the contract method 0x2e1ec945.
//
// Solidity: function getLPTokens(address lp) view returns(address, address)
func (_TokenPair *TokenPairCaller) GetLPTokens(opts *bind.CallOpts, lp common.Address) (common.Address, common.Address, error) {
	var out []interface{}
	err := _TokenPair.contract.Call(opts, &out, "getLPTokens", lp)

	if err != nil {
		return *new(common.Address), *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)
	out1 := *abi.ConvertType(out[1], new(common.Address)).(*common.Address)

	return out0, out1, err

}

// GetLPTokens is a free data retrieval call binding the contract method 0x2e1ec945.
//
// Solidity: function getLPTokens(address lp) view returns(address, address)
func (_TokenPair *TokenPairSession) GetLPTokens(lp common.Address) (common.Address, common.Address, error) {
	return _TokenPair.Contract.GetLPTokens(&_TokenPair.CallOpts, lp)
}

// GetLPTokens is a free data retrieval call binding the contract method 0x2e1ec945.
//
// Solidity: function getLPTokens(address lp) view returns(address, address)
func (_TokenPair *TokenPairCallerSession) GetLPTokens(lp common.Address) (common.Address, common.Address, error) {
	return _TokenPair.Contract.GetLPTokens(&_TokenPair.CallOpts, lp)
}

// GetTokenOut is a free data retrieval call binding the contract method 0x3a8befdc.
//
// Solidity: function getTokenOut(address router, address tokenIn, address tokenOut) view returns(uint256)
func (_TokenPair *TokenPairCaller) GetTokenOut(opts *bind.CallOpts, router common.Address, tokenIn common.Address, tokenOut common.Address) (*big.Int, error) {
	var out []interface{}
	err := _TokenPair.contract.Call(opts, &out, "getTokenOut", router, tokenIn, tokenOut)

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetTokenOut is a free data retrieval call binding the contract method 0x3a8befdc.
//
// Solidity: function getTokenOut(address router, address tokenIn, address tokenOut) view returns(uint256)
func (_TokenPair *TokenPairSession) GetTokenOut(router common.Address, tokenIn common.Address, tokenOut common.Address) (*big.Int, error) {
	return _TokenPair.Contract.GetTokenOut(&_TokenPair.CallOpts, router, tokenIn, tokenOut)
}

// GetTokenOut is a free data retrieval call binding the contract method 0x3a8befdc.
//
// Solidity: function getTokenOut(address router, address tokenIn, address tokenOut) view returns(uint256)
func (_TokenPair *TokenPairCallerSession) GetTokenOut(router common.Address, tokenIn common.Address, tokenOut common.Address) (*big.Int, error) {
	return _TokenPair.Contract.GetTokenOut(&_TokenPair.CallOpts, router, tokenIn, tokenOut)
}
