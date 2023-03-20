// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package abigen

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

// ChainConfigMetaData contains all meta data concerning the ChainConfig contract.
var ChainConfigMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"contractIStaking\",\"name\":\"stakingContract\",\"type\":\"address\"},{\"internalType\":\"contractISlashingIndicator\",\"name\":\"slashingIndicatorContract\",\"type\":\"address\"},{\"internalType\":\"contractISystemReward\",\"name\":\"systemRewardContract\",\"type\":\"address\"},{\"internalType\":\"contractIStakingPool\",\"name\":\"stakingPoolContract\",\"type\":\"address\"},{\"internalType\":\"contractIGovernance\",\"name\":\"governanceContract\",\"type\":\"address\"},{\"internalType\":\"contractIChainConfig\",\"name\":\"chainConfigContract\",\"type\":\"address\"},{\"internalType\":\"contractIRuntimeUpgrade\",\"name\":\"runtimeUpgradeContract\",\"type\":\"address\"},{\"internalType\":\"contractIDeployerProxy\",\"name\":\"deployerProxyContract\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint64\",\"name\":\"blockNumber\",\"type\":\"uint64\"}],\"name\":\"OnlyBlock\",\"type\":\"error\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"coinbase\",\"type\":\"address\"}],\"name\":\"OnlyCoinbase\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlyGovernance\",\"type\":\"error\"},{\"inputs\":[],\"name\":\"OnlySlashingIndicator\",\"type\":\"error\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"prevValue\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newValue\",\"type\":\"uint32\"}],\"name\":\"ActiveValidatorsLengthChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"prevValue\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newValue\",\"type\":\"uint32\"}],\"name\":\"EpochBlockIntervalChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"prevValue\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newValue\",\"type\":\"uint32\"}],\"name\":\"FelonyThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint8\",\"name\":\"version\",\"type\":\"uint8\"}],\"name\":\"Initialized\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prevValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"MinStakingAmountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"prevValue\",\"type\":\"uint256\"},{\"indexed\":false,\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"MinValidatorStakeAmountChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"prevValue\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newValue\",\"type\":\"uint32\"}],\"name\":\"MisdemeanorThresholdChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"prevValue\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newValue\",\"type\":\"uint32\"}],\"name\":\"UndelegatePeriodChanged\",\"type\":\"event\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"prevValue\",\"type\":\"uint32\"},{\"indexed\":false,\"internalType\":\"uint32\",\"name\":\"newValue\",\"type\":\"uint32\"}],\"name\":\"ValidatorJailEpochLengthChanged\",\"type\":\"event\"},{\"inputs\":[],\"name\":\"init\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"isInitialized\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes[]\",\"name\":\"data\",\"type\":\"bytes[]\"}],\"name\":\"multicall\",\"outputs\":[{\"internalType\":\"bytes[]\",\"name\":\"results\",\"type\":\"bytes[]\"}],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"bytes\",\"name\":\"delayedInitializer\",\"type\":\"bytes\"}],\"name\":\"useDelayedInitializer\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"activeValidatorsLength\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"epochBlockInterval\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"misdemeanorThreshold\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"felonyThreshold\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"validatorJailEpochLength\",\"type\":\"uint32\"},{\"internalType\":\"uint32\",\"name\":\"undelegatePeriod\",\"type\":\"uint32\"},{\"internalType\":\"uint256\",\"name\":\"minValidatorStakeAmount\",\"type\":\"uint256\"},{\"internalType\":\"uint256\",\"name\":\"minStakingAmount\",\"type\":\"uint256\"}],\"name\":\"initialize\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getActiveValidatorsLength\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"newValue\",\"type\":\"uint32\"}],\"name\":\"setActiveValidatorsLength\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getEpochBlockInterval\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"newValue\",\"type\":\"uint32\"}],\"name\":\"setEpochBlockInterval\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMisdemeanorThreshold\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"newValue\",\"type\":\"uint32\"}],\"name\":\"setMisdemeanorThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getFelonyThreshold\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"newValue\",\"type\":\"uint32\"}],\"name\":\"setFelonyThreshold\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getValidatorJailEpochLength\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"newValue\",\"type\":\"uint32\"}],\"name\":\"setValidatorJailEpochLength\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getUndelegatePeriod\",\"outputs\":[{\"internalType\":\"uint32\",\"name\":\"\",\"type\":\"uint32\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint32\",\"name\":\"newValue\",\"type\":\"uint32\"}],\"name\":\"setUndelegatePeriod\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinValidatorStakeAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"setMinValidatorStakeAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getMinStakingAmount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"newValue\",\"type\":\"uint256\"}],\"name\":\"setMinStakingAmount\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"}]",
}

// ChainConfigABI is the input ABI used to generate the binding from.
// Deprecated: Use ChainConfigMetaData.ABI instead.
var ChainConfigABI = ChainConfigMetaData.ABI

// ChainConfig is an auto generated Go binding around an Ethereum contract.
type ChainConfig struct {
	ChainConfigCaller     // Read-only binding to the contract
	ChainConfigTransactor // Write-only binding to the contract
	ChainConfigFilterer   // Log filterer for contract events
}

// ChainConfigCaller is an auto generated read-only Go binding around an Ethereum contract.
type ChainConfigCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainConfigTransactor is an auto generated write-only Go binding around an Ethereum contract.
type ChainConfigTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainConfigFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type ChainConfigFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// ChainConfigSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type ChainConfigSession struct {
	Contract     *ChainConfig      // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// ChainConfigCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type ChainConfigCallerSession struct {
	Contract *ChainConfigCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts      // Call options to use throughout this session
}

// ChainConfigTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type ChainConfigTransactorSession struct {
	Contract     *ChainConfigTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts      // Transaction auth options to use throughout this session
}

// ChainConfigRaw is an auto generated low-level Go binding around an Ethereum contract.
type ChainConfigRaw struct {
	Contract *ChainConfig // Generic contract binding to access the raw methods on
}

// ChainConfigCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type ChainConfigCallerRaw struct {
	Contract *ChainConfigCaller // Generic read-only contract binding to access the raw methods on
}

// ChainConfigTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type ChainConfigTransactorRaw struct {
	Contract *ChainConfigTransactor // Generic write-only contract binding to access the raw methods on
}

// NewChainConfig creates a new instance of ChainConfig, bound to a specific deployed contract.
func NewChainConfig(address common.Address, backend bind.ContractBackend) (*ChainConfig, error) {
	contract, err := bindChainConfig(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &ChainConfig{ChainConfigCaller: ChainConfigCaller{contract: contract}, ChainConfigTransactor: ChainConfigTransactor{contract: contract}, ChainConfigFilterer: ChainConfigFilterer{contract: contract}}, nil
}

// NewChainConfigCaller creates a new read-only instance of ChainConfig, bound to a specific deployed contract.
func NewChainConfigCaller(address common.Address, caller bind.ContractCaller) (*ChainConfigCaller, error) {
	contract, err := bindChainConfig(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &ChainConfigCaller{contract: contract}, nil
}

// NewChainConfigTransactor creates a new write-only instance of ChainConfig, bound to a specific deployed contract.
func NewChainConfigTransactor(address common.Address, transactor bind.ContractTransactor) (*ChainConfigTransactor, error) {
	contract, err := bindChainConfig(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &ChainConfigTransactor{contract: contract}, nil
}

// NewChainConfigFilterer creates a new log filterer instance of ChainConfig, bound to a specific deployed contract.
func NewChainConfigFilterer(address common.Address, filterer bind.ContractFilterer) (*ChainConfigFilterer, error) {
	contract, err := bindChainConfig(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &ChainConfigFilterer{contract: contract}, nil
}

// bindChainConfig binds a generic wrapper to an already deployed contract.
func bindChainConfig(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := ChainConfigMetaData.GetAbi()
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, *parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChainConfig *ChainConfigRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChainConfig.Contract.ChainConfigCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChainConfig *ChainConfigRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainConfig.Contract.ChainConfigTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChainConfig *ChainConfigRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChainConfig.Contract.ChainConfigTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_ChainConfig *ChainConfigCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _ChainConfig.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_ChainConfig *ChainConfigTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainConfig.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_ChainConfig *ChainConfigTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _ChainConfig.Contract.contract.Transact(opts, method, params...)
}

// GetActiveValidatorsLength is a free data retrieval call binding the contract method 0x32cc6f08.
//
// Solidity: function getActiveValidatorsLength() view returns(uint32)
func (_ChainConfig *ChainConfigCaller) GetActiveValidatorsLength(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ChainConfig.contract.Call(opts, &out, "getActiveValidatorsLength")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetActiveValidatorsLength is a free data retrieval call binding the contract method 0x32cc6f08.
//
// Solidity: function getActiveValidatorsLength() view returns(uint32)
func (_ChainConfig *ChainConfigSession) GetActiveValidatorsLength() (uint32, error) {
	return _ChainConfig.Contract.GetActiveValidatorsLength(&_ChainConfig.CallOpts)
}

// GetActiveValidatorsLength is a free data retrieval call binding the contract method 0x32cc6f08.
//
// Solidity: function getActiveValidatorsLength() view returns(uint32)
func (_ChainConfig *ChainConfigCallerSession) GetActiveValidatorsLength() (uint32, error) {
	return _ChainConfig.Contract.GetActiveValidatorsLength(&_ChainConfig.CallOpts)
}

// GetEpochBlockInterval is a free data retrieval call binding the contract method 0x346c90a8.
//
// Solidity: function getEpochBlockInterval() view returns(uint32)
func (_ChainConfig *ChainConfigCaller) GetEpochBlockInterval(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ChainConfig.contract.Call(opts, &out, "getEpochBlockInterval")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetEpochBlockInterval is a free data retrieval call binding the contract method 0x346c90a8.
//
// Solidity: function getEpochBlockInterval() view returns(uint32)
func (_ChainConfig *ChainConfigSession) GetEpochBlockInterval() (uint32, error) {
	return _ChainConfig.Contract.GetEpochBlockInterval(&_ChainConfig.CallOpts)
}

// GetEpochBlockInterval is a free data retrieval call binding the contract method 0x346c90a8.
//
// Solidity: function getEpochBlockInterval() view returns(uint32)
func (_ChainConfig *ChainConfigCallerSession) GetEpochBlockInterval() (uint32, error) {
	return _ChainConfig.Contract.GetEpochBlockInterval(&_ChainConfig.CallOpts)
}

// GetFelonyThreshold is a free data retrieval call binding the contract method 0xbe199738.
//
// Solidity: function getFelonyThreshold() view returns(uint32)
func (_ChainConfig *ChainConfigCaller) GetFelonyThreshold(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ChainConfig.contract.Call(opts, &out, "getFelonyThreshold")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetFelonyThreshold is a free data retrieval call binding the contract method 0xbe199738.
//
// Solidity: function getFelonyThreshold() view returns(uint32)
func (_ChainConfig *ChainConfigSession) GetFelonyThreshold() (uint32, error) {
	return _ChainConfig.Contract.GetFelonyThreshold(&_ChainConfig.CallOpts)
}

// GetFelonyThreshold is a free data retrieval call binding the contract method 0xbe199738.
//
// Solidity: function getFelonyThreshold() view returns(uint32)
func (_ChainConfig *ChainConfigCallerSession) GetFelonyThreshold() (uint32, error) {
	return _ChainConfig.Contract.GetFelonyThreshold(&_ChainConfig.CallOpts)
}

// GetMinStakingAmount is a free data retrieval call binding the contract method 0xeea9a01b.
//
// Solidity: function getMinStakingAmount() view returns(uint256)
func (_ChainConfig *ChainConfigCaller) GetMinStakingAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ChainConfig.contract.Call(opts, &out, "getMinStakingAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinStakingAmount is a free data retrieval call binding the contract method 0xeea9a01b.
//
// Solidity: function getMinStakingAmount() view returns(uint256)
func (_ChainConfig *ChainConfigSession) GetMinStakingAmount() (*big.Int, error) {
	return _ChainConfig.Contract.GetMinStakingAmount(&_ChainConfig.CallOpts)
}

// GetMinStakingAmount is a free data retrieval call binding the contract method 0xeea9a01b.
//
// Solidity: function getMinStakingAmount() view returns(uint256)
func (_ChainConfig *ChainConfigCallerSession) GetMinStakingAmount() (*big.Int, error) {
	return _ChainConfig.Contract.GetMinStakingAmount(&_ChainConfig.CallOpts)
}

// GetMinValidatorStakeAmount is a free data retrieval call binding the contract method 0x6f856847.
//
// Solidity: function getMinValidatorStakeAmount() view returns(uint256)
func (_ChainConfig *ChainConfigCaller) GetMinValidatorStakeAmount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _ChainConfig.contract.Call(opts, &out, "getMinValidatorStakeAmount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// GetMinValidatorStakeAmount is a free data retrieval call binding the contract method 0x6f856847.
//
// Solidity: function getMinValidatorStakeAmount() view returns(uint256)
func (_ChainConfig *ChainConfigSession) GetMinValidatorStakeAmount() (*big.Int, error) {
	return _ChainConfig.Contract.GetMinValidatorStakeAmount(&_ChainConfig.CallOpts)
}

// GetMinValidatorStakeAmount is a free data retrieval call binding the contract method 0x6f856847.
//
// Solidity: function getMinValidatorStakeAmount() view returns(uint256)
func (_ChainConfig *ChainConfigCallerSession) GetMinValidatorStakeAmount() (*big.Int, error) {
	return _ChainConfig.Contract.GetMinValidatorStakeAmount(&_ChainConfig.CallOpts)
}

// GetMisdemeanorThreshold is a free data retrieval call binding the contract method 0x9dbf97db.
//
// Solidity: function getMisdemeanorThreshold() view returns(uint32)
func (_ChainConfig *ChainConfigCaller) GetMisdemeanorThreshold(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ChainConfig.contract.Call(opts, &out, "getMisdemeanorThreshold")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetMisdemeanorThreshold is a free data retrieval call binding the contract method 0x9dbf97db.
//
// Solidity: function getMisdemeanorThreshold() view returns(uint32)
func (_ChainConfig *ChainConfigSession) GetMisdemeanorThreshold() (uint32, error) {
	return _ChainConfig.Contract.GetMisdemeanorThreshold(&_ChainConfig.CallOpts)
}

// GetMisdemeanorThreshold is a free data retrieval call binding the contract method 0x9dbf97db.
//
// Solidity: function getMisdemeanorThreshold() view returns(uint32)
func (_ChainConfig *ChainConfigCallerSession) GetMisdemeanorThreshold() (uint32, error) {
	return _ChainConfig.Contract.GetMisdemeanorThreshold(&_ChainConfig.CallOpts)
}

// GetUndelegatePeriod is a free data retrieval call binding the contract method 0x5e7b72ad.
//
// Solidity: function getUndelegatePeriod() view returns(uint32)
func (_ChainConfig *ChainConfigCaller) GetUndelegatePeriod(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ChainConfig.contract.Call(opts, &out, "getUndelegatePeriod")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetUndelegatePeriod is a free data retrieval call binding the contract method 0x5e7b72ad.
//
// Solidity: function getUndelegatePeriod() view returns(uint32)
func (_ChainConfig *ChainConfigSession) GetUndelegatePeriod() (uint32, error) {
	return _ChainConfig.Contract.GetUndelegatePeriod(&_ChainConfig.CallOpts)
}

// GetUndelegatePeriod is a free data retrieval call binding the contract method 0x5e7b72ad.
//
// Solidity: function getUndelegatePeriod() view returns(uint32)
func (_ChainConfig *ChainConfigCallerSession) GetUndelegatePeriod() (uint32, error) {
	return _ChainConfig.Contract.GetUndelegatePeriod(&_ChainConfig.CallOpts)
}

// GetValidatorJailEpochLength is a free data retrieval call binding the contract method 0x6cbe6cd8.
//
// Solidity: function getValidatorJailEpochLength() view returns(uint32)
func (_ChainConfig *ChainConfigCaller) GetValidatorJailEpochLength(opts *bind.CallOpts) (uint32, error) {
	var out []interface{}
	err := _ChainConfig.contract.Call(opts, &out, "getValidatorJailEpochLength")

	if err != nil {
		return *new(uint32), err
	}

	out0 := *abi.ConvertType(out[0], new(uint32)).(*uint32)

	return out0, err

}

// GetValidatorJailEpochLength is a free data retrieval call binding the contract method 0x6cbe6cd8.
//
// Solidity: function getValidatorJailEpochLength() view returns(uint32)
func (_ChainConfig *ChainConfigSession) GetValidatorJailEpochLength() (uint32, error) {
	return _ChainConfig.Contract.GetValidatorJailEpochLength(&_ChainConfig.CallOpts)
}

// GetValidatorJailEpochLength is a free data retrieval call binding the contract method 0x6cbe6cd8.
//
// Solidity: function getValidatorJailEpochLength() view returns(uint32)
func (_ChainConfig *ChainConfigCallerSession) GetValidatorJailEpochLength() (uint32, error) {
	return _ChainConfig.Contract.GetValidatorJailEpochLength(&_ChainConfig.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_ChainConfig *ChainConfigCaller) IsInitialized(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _ChainConfig.contract.Call(opts, &out, "isInitialized")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_ChainConfig *ChainConfigSession) IsInitialized() (bool, error) {
	return _ChainConfig.Contract.IsInitialized(&_ChainConfig.CallOpts)
}

// IsInitialized is a free data retrieval call binding the contract method 0x392e53cd.
//
// Solidity: function isInitialized() view returns(bool)
func (_ChainConfig *ChainConfigCallerSession) IsInitialized() (bool, error) {
	return _ChainConfig.Contract.IsInitialized(&_ChainConfig.CallOpts)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_ChainConfig *ChainConfigTransactor) Init(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _ChainConfig.contract.Transact(opts, "init")
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_ChainConfig *ChainConfigSession) Init() (*types.Transaction, error) {
	return _ChainConfig.Contract.Init(&_ChainConfig.TransactOpts)
}

// Init is a paid mutator transaction binding the contract method 0xe1c7392a.
//
// Solidity: function init() returns()
func (_ChainConfig *ChainConfigTransactorSession) Init() (*types.Transaction, error) {
	return _ChainConfig.Contract.Init(&_ChainConfig.TransactOpts)
}

// Initialize is a paid mutator transaction binding the contract method 0x8fe91cc6.
//
// Solidity: function initialize(uint32 activeValidatorsLength, uint32 epochBlockInterval, uint32 misdemeanorThreshold, uint32 felonyThreshold, uint32 validatorJailEpochLength, uint32 undelegatePeriod, uint256 minValidatorStakeAmount, uint256 minStakingAmount) returns()
func (_ChainConfig *ChainConfigTransactor) Initialize(opts *bind.TransactOpts, activeValidatorsLength uint32, epochBlockInterval uint32, misdemeanorThreshold uint32, felonyThreshold uint32, validatorJailEpochLength uint32, undelegatePeriod uint32, minValidatorStakeAmount *big.Int, minStakingAmount *big.Int) (*types.Transaction, error) {
	return _ChainConfig.contract.Transact(opts, "initialize", activeValidatorsLength, epochBlockInterval, misdemeanorThreshold, felonyThreshold, validatorJailEpochLength, undelegatePeriod, minValidatorStakeAmount, minStakingAmount)
}

// Initialize is a paid mutator transaction binding the contract method 0x8fe91cc6.
//
// Solidity: function initialize(uint32 activeValidatorsLength, uint32 epochBlockInterval, uint32 misdemeanorThreshold, uint32 felonyThreshold, uint32 validatorJailEpochLength, uint32 undelegatePeriod, uint256 minValidatorStakeAmount, uint256 minStakingAmount) returns()
func (_ChainConfig *ChainConfigSession) Initialize(activeValidatorsLength uint32, epochBlockInterval uint32, misdemeanorThreshold uint32, felonyThreshold uint32, validatorJailEpochLength uint32, undelegatePeriod uint32, minValidatorStakeAmount *big.Int, minStakingAmount *big.Int) (*types.Transaction, error) {
	return _ChainConfig.Contract.Initialize(&_ChainConfig.TransactOpts, activeValidatorsLength, epochBlockInterval, misdemeanorThreshold, felonyThreshold, validatorJailEpochLength, undelegatePeriod, minValidatorStakeAmount, minStakingAmount)
}

// Initialize is a paid mutator transaction binding the contract method 0x8fe91cc6.
//
// Solidity: function initialize(uint32 activeValidatorsLength, uint32 epochBlockInterval, uint32 misdemeanorThreshold, uint32 felonyThreshold, uint32 validatorJailEpochLength, uint32 undelegatePeriod, uint256 minValidatorStakeAmount, uint256 minStakingAmount) returns()
func (_ChainConfig *ChainConfigTransactorSession) Initialize(activeValidatorsLength uint32, epochBlockInterval uint32, misdemeanorThreshold uint32, felonyThreshold uint32, validatorJailEpochLength uint32, undelegatePeriod uint32, minValidatorStakeAmount *big.Int, minStakingAmount *big.Int) (*types.Transaction, error) {
	return _ChainConfig.Contract.Initialize(&_ChainConfig.TransactOpts, activeValidatorsLength, epochBlockInterval, misdemeanorThreshold, felonyThreshold, validatorJailEpochLength, undelegatePeriod, minValidatorStakeAmount, minStakingAmount)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_ChainConfig *ChainConfigTransactor) Multicall(opts *bind.TransactOpts, data [][]byte) (*types.Transaction, error) {
	return _ChainConfig.contract.Transact(opts, "multicall", data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_ChainConfig *ChainConfigSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _ChainConfig.Contract.Multicall(&_ChainConfig.TransactOpts, data)
}

// Multicall is a paid mutator transaction binding the contract method 0xac9650d8.
//
// Solidity: function multicall(bytes[] data) returns(bytes[] results)
func (_ChainConfig *ChainConfigTransactorSession) Multicall(data [][]byte) (*types.Transaction, error) {
	return _ChainConfig.Contract.Multicall(&_ChainConfig.TransactOpts, data)
}

// SetActiveValidatorsLength is a paid mutator transaction binding the contract method 0xc227a412.
//
// Solidity: function setActiveValidatorsLength(uint32 newValue) returns()
func (_ChainConfig *ChainConfigTransactor) SetActiveValidatorsLength(opts *bind.TransactOpts, newValue uint32) (*types.Transaction, error) {
	return _ChainConfig.contract.Transact(opts, "setActiveValidatorsLength", newValue)
}

// SetActiveValidatorsLength is a paid mutator transaction binding the contract method 0xc227a412.
//
// Solidity: function setActiveValidatorsLength(uint32 newValue) returns()
func (_ChainConfig *ChainConfigSession) SetActiveValidatorsLength(newValue uint32) (*types.Transaction, error) {
	return _ChainConfig.Contract.SetActiveValidatorsLength(&_ChainConfig.TransactOpts, newValue)
}

// SetActiveValidatorsLength is a paid mutator transaction binding the contract method 0xc227a412.
//
// Solidity: function setActiveValidatorsLength(uint32 newValue) returns()
func (_ChainConfig *ChainConfigTransactorSession) SetActiveValidatorsLength(newValue uint32) (*types.Transaction, error) {
	return _ChainConfig.Contract.SetActiveValidatorsLength(&_ChainConfig.TransactOpts, newValue)
}

// SetEpochBlockInterval is a paid mutator transaction binding the contract method 0xaf70fa2c.
//
// Solidity: function setEpochBlockInterval(uint32 newValue) returns()
func (_ChainConfig *ChainConfigTransactor) SetEpochBlockInterval(opts *bind.TransactOpts, newValue uint32) (*types.Transaction, error) {
	return _ChainConfig.contract.Transact(opts, "setEpochBlockInterval", newValue)
}

// SetEpochBlockInterval is a paid mutator transaction binding the contract method 0xaf70fa2c.
//
// Solidity: function setEpochBlockInterval(uint32 newValue) returns()
func (_ChainConfig *ChainConfigSession) SetEpochBlockInterval(newValue uint32) (*types.Transaction, error) {
	return _ChainConfig.Contract.SetEpochBlockInterval(&_ChainConfig.TransactOpts, newValue)
}

// SetEpochBlockInterval is a paid mutator transaction binding the contract method 0xaf70fa2c.
//
// Solidity: function setEpochBlockInterval(uint32 newValue) returns()
func (_ChainConfig *ChainConfigTransactorSession) SetEpochBlockInterval(newValue uint32) (*types.Transaction, error) {
	return _ChainConfig.Contract.SetEpochBlockInterval(&_ChainConfig.TransactOpts, newValue)
}

// SetFelonyThreshold is a paid mutator transaction binding the contract method 0xfcd6cb3e.
//
// Solidity: function setFelonyThreshold(uint32 newValue) returns()
func (_ChainConfig *ChainConfigTransactor) SetFelonyThreshold(opts *bind.TransactOpts, newValue uint32) (*types.Transaction, error) {
	return _ChainConfig.contract.Transact(opts, "setFelonyThreshold", newValue)
}

// SetFelonyThreshold is a paid mutator transaction binding the contract method 0xfcd6cb3e.
//
// Solidity: function setFelonyThreshold(uint32 newValue) returns()
func (_ChainConfig *ChainConfigSession) SetFelonyThreshold(newValue uint32) (*types.Transaction, error) {
	return _ChainConfig.Contract.SetFelonyThreshold(&_ChainConfig.TransactOpts, newValue)
}

// SetFelonyThreshold is a paid mutator transaction binding the contract method 0xfcd6cb3e.
//
// Solidity: function setFelonyThreshold(uint32 newValue) returns()
func (_ChainConfig *ChainConfigTransactorSession) SetFelonyThreshold(newValue uint32) (*types.Transaction, error) {
	return _ChainConfig.Contract.SetFelonyThreshold(&_ChainConfig.TransactOpts, newValue)
}

// SetMinStakingAmount is a paid mutator transaction binding the contract method 0x612d669e.
//
// Solidity: function setMinStakingAmount(uint256 newValue) returns()
func (_ChainConfig *ChainConfigTransactor) SetMinStakingAmount(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error) {
	return _ChainConfig.contract.Transact(opts, "setMinStakingAmount", newValue)
}

// SetMinStakingAmount is a paid mutator transaction binding the contract method 0x612d669e.
//
// Solidity: function setMinStakingAmount(uint256 newValue) returns()
func (_ChainConfig *ChainConfigSession) SetMinStakingAmount(newValue *big.Int) (*types.Transaction, error) {
	return _ChainConfig.Contract.SetMinStakingAmount(&_ChainConfig.TransactOpts, newValue)
}

// SetMinStakingAmount is a paid mutator transaction binding the contract method 0x612d669e.
//
// Solidity: function setMinStakingAmount(uint256 newValue) returns()
func (_ChainConfig *ChainConfigTransactorSession) SetMinStakingAmount(newValue *big.Int) (*types.Transaction, error) {
	return _ChainConfig.Contract.SetMinStakingAmount(&_ChainConfig.TransactOpts, newValue)
}

// SetMinValidatorStakeAmount is a paid mutator transaction binding the contract method 0xe1a2e863.
//
// Solidity: function setMinValidatorStakeAmount(uint256 newValue) returns()
func (_ChainConfig *ChainConfigTransactor) SetMinValidatorStakeAmount(opts *bind.TransactOpts, newValue *big.Int) (*types.Transaction, error) {
	return _ChainConfig.contract.Transact(opts, "setMinValidatorStakeAmount", newValue)
}

// SetMinValidatorStakeAmount is a paid mutator transaction binding the contract method 0xe1a2e863.
//
// Solidity: function setMinValidatorStakeAmount(uint256 newValue) returns()
func (_ChainConfig *ChainConfigSession) SetMinValidatorStakeAmount(newValue *big.Int) (*types.Transaction, error) {
	return _ChainConfig.Contract.SetMinValidatorStakeAmount(&_ChainConfig.TransactOpts, newValue)
}

// SetMinValidatorStakeAmount is a paid mutator transaction binding the contract method 0xe1a2e863.
//
// Solidity: function setMinValidatorStakeAmount(uint256 newValue) returns()
func (_ChainConfig *ChainConfigTransactorSession) SetMinValidatorStakeAmount(newValue *big.Int) (*types.Transaction, error) {
	return _ChainConfig.Contract.SetMinValidatorStakeAmount(&_ChainConfig.TransactOpts, newValue)
}

// SetMisdemeanorThreshold is a paid mutator transaction binding the contract method 0xd98e3ebf.
//
// Solidity: function setMisdemeanorThreshold(uint32 newValue) returns()
func (_ChainConfig *ChainConfigTransactor) SetMisdemeanorThreshold(opts *bind.TransactOpts, newValue uint32) (*types.Transaction, error) {
	return _ChainConfig.contract.Transact(opts, "setMisdemeanorThreshold", newValue)
}

// SetMisdemeanorThreshold is a paid mutator transaction binding the contract method 0xd98e3ebf.
//
// Solidity: function setMisdemeanorThreshold(uint32 newValue) returns()
func (_ChainConfig *ChainConfigSession) SetMisdemeanorThreshold(newValue uint32) (*types.Transaction, error) {
	return _ChainConfig.Contract.SetMisdemeanorThreshold(&_ChainConfig.TransactOpts, newValue)
}

// SetMisdemeanorThreshold is a paid mutator transaction binding the contract method 0xd98e3ebf.
//
// Solidity: function setMisdemeanorThreshold(uint32 newValue) returns()
func (_ChainConfig *ChainConfigTransactorSession) SetMisdemeanorThreshold(newValue uint32) (*types.Transaction, error) {
	return _ChainConfig.Contract.SetMisdemeanorThreshold(&_ChainConfig.TransactOpts, newValue)
}

// SetUndelegatePeriod is a paid mutator transaction binding the contract method 0x41d8a080.
//
// Solidity: function setUndelegatePeriod(uint32 newValue) returns()
func (_ChainConfig *ChainConfigTransactor) SetUndelegatePeriod(opts *bind.TransactOpts, newValue uint32) (*types.Transaction, error) {
	return _ChainConfig.contract.Transact(opts, "setUndelegatePeriod", newValue)
}

// SetUndelegatePeriod is a paid mutator transaction binding the contract method 0x41d8a080.
//
// Solidity: function setUndelegatePeriod(uint32 newValue) returns()
func (_ChainConfig *ChainConfigSession) SetUndelegatePeriod(newValue uint32) (*types.Transaction, error) {
	return _ChainConfig.Contract.SetUndelegatePeriod(&_ChainConfig.TransactOpts, newValue)
}

// SetUndelegatePeriod is a paid mutator transaction binding the contract method 0x41d8a080.
//
// Solidity: function setUndelegatePeriod(uint32 newValue) returns()
func (_ChainConfig *ChainConfigTransactorSession) SetUndelegatePeriod(newValue uint32) (*types.Transaction, error) {
	return _ChainConfig.Contract.SetUndelegatePeriod(&_ChainConfig.TransactOpts, newValue)
}

// SetValidatorJailEpochLength is a paid mutator transaction binding the contract method 0xc8652bd5.
//
// Solidity: function setValidatorJailEpochLength(uint32 newValue) returns()
func (_ChainConfig *ChainConfigTransactor) SetValidatorJailEpochLength(opts *bind.TransactOpts, newValue uint32) (*types.Transaction, error) {
	return _ChainConfig.contract.Transact(opts, "setValidatorJailEpochLength", newValue)
}

// SetValidatorJailEpochLength is a paid mutator transaction binding the contract method 0xc8652bd5.
//
// Solidity: function setValidatorJailEpochLength(uint32 newValue) returns()
func (_ChainConfig *ChainConfigSession) SetValidatorJailEpochLength(newValue uint32) (*types.Transaction, error) {
	return _ChainConfig.Contract.SetValidatorJailEpochLength(&_ChainConfig.TransactOpts, newValue)
}

// SetValidatorJailEpochLength is a paid mutator transaction binding the contract method 0xc8652bd5.
//
// Solidity: function setValidatorJailEpochLength(uint32 newValue) returns()
func (_ChainConfig *ChainConfigTransactorSession) SetValidatorJailEpochLength(newValue uint32) (*types.Transaction, error) {
	return _ChainConfig.Contract.SetValidatorJailEpochLength(&_ChainConfig.TransactOpts, newValue)
}

// UseDelayedInitializer is a paid mutator transaction binding the contract method 0xc529d2f3.
//
// Solidity: function useDelayedInitializer(bytes delayedInitializer) returns()
func (_ChainConfig *ChainConfigTransactor) UseDelayedInitializer(opts *bind.TransactOpts, delayedInitializer []byte) (*types.Transaction, error) {
	return _ChainConfig.contract.Transact(opts, "useDelayedInitializer", delayedInitializer)
}

// UseDelayedInitializer is a paid mutator transaction binding the contract method 0xc529d2f3.
//
// Solidity: function useDelayedInitializer(bytes delayedInitializer) returns()
func (_ChainConfig *ChainConfigSession) UseDelayedInitializer(delayedInitializer []byte) (*types.Transaction, error) {
	return _ChainConfig.Contract.UseDelayedInitializer(&_ChainConfig.TransactOpts, delayedInitializer)
}

// UseDelayedInitializer is a paid mutator transaction binding the contract method 0xc529d2f3.
//
// Solidity: function useDelayedInitializer(bytes delayedInitializer) returns()
func (_ChainConfig *ChainConfigTransactorSession) UseDelayedInitializer(delayedInitializer []byte) (*types.Transaction, error) {
	return _ChainConfig.Contract.UseDelayedInitializer(&_ChainConfig.TransactOpts, delayedInitializer)
}

// ChainConfigActiveValidatorsLengthChangedIterator is returned from FilterActiveValidatorsLengthChanged and is used to iterate over the raw logs and unpacked data for ActiveValidatorsLengthChanged events raised by the ChainConfig contract.
type ChainConfigActiveValidatorsLengthChangedIterator struct {
	Event *ChainConfigActiveValidatorsLengthChanged // Event containing the contract specifics and raw log

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
func (it *ChainConfigActiveValidatorsLengthChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainConfigActiveValidatorsLengthChanged)
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
		it.Event = new(ChainConfigActiveValidatorsLengthChanged)
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
func (it *ChainConfigActiveValidatorsLengthChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainConfigActiveValidatorsLengthChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainConfigActiveValidatorsLengthChanged represents a ActiveValidatorsLengthChanged event raised by the ChainConfig contract.
type ChainConfigActiveValidatorsLengthChanged struct {
	PrevValue uint32
	NewValue  uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterActiveValidatorsLengthChanged is a free log retrieval operation binding the contract event 0x1c4cfc6dcf4219ed649285020aedf5d064480d1acdf4b8c75b397abd5910f40c.
//
// Solidity: event ActiveValidatorsLengthChanged(uint32 prevValue, uint32 newValue)
func (_ChainConfig *ChainConfigFilterer) FilterActiveValidatorsLengthChanged(opts *bind.FilterOpts) (*ChainConfigActiveValidatorsLengthChangedIterator, error) {

	logs, sub, err := _ChainConfig.contract.FilterLogs(opts, "ActiveValidatorsLengthChanged")
	if err != nil {
		return nil, err
	}
	return &ChainConfigActiveValidatorsLengthChangedIterator{contract: _ChainConfig.contract, event: "ActiveValidatorsLengthChanged", logs: logs, sub: sub}, nil
}

// WatchActiveValidatorsLengthChanged is a free log subscription operation binding the contract event 0x1c4cfc6dcf4219ed649285020aedf5d064480d1acdf4b8c75b397abd5910f40c.
//
// Solidity: event ActiveValidatorsLengthChanged(uint32 prevValue, uint32 newValue)
func (_ChainConfig *ChainConfigFilterer) WatchActiveValidatorsLengthChanged(opts *bind.WatchOpts, sink chan<- *ChainConfigActiveValidatorsLengthChanged) (event.Subscription, error) {

	logs, sub, err := _ChainConfig.contract.WatchLogs(opts, "ActiveValidatorsLengthChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainConfigActiveValidatorsLengthChanged)
				if err := _ChainConfig.contract.UnpackLog(event, "ActiveValidatorsLengthChanged", log); err != nil {
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

// ParseActiveValidatorsLengthChanged is a log parse operation binding the contract event 0x1c4cfc6dcf4219ed649285020aedf5d064480d1acdf4b8c75b397abd5910f40c.
//
// Solidity: event ActiveValidatorsLengthChanged(uint32 prevValue, uint32 newValue)
func (_ChainConfig *ChainConfigFilterer) ParseActiveValidatorsLengthChanged(log types.Log) (*ChainConfigActiveValidatorsLengthChanged, error) {
	event := new(ChainConfigActiveValidatorsLengthChanged)
	if err := _ChainConfig.contract.UnpackLog(event, "ActiveValidatorsLengthChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChainConfigEpochBlockIntervalChangedIterator is returned from FilterEpochBlockIntervalChanged and is used to iterate over the raw logs and unpacked data for EpochBlockIntervalChanged events raised by the ChainConfig contract.
type ChainConfigEpochBlockIntervalChangedIterator struct {
	Event *ChainConfigEpochBlockIntervalChanged // Event containing the contract specifics and raw log

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
func (it *ChainConfigEpochBlockIntervalChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainConfigEpochBlockIntervalChanged)
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
		it.Event = new(ChainConfigEpochBlockIntervalChanged)
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
func (it *ChainConfigEpochBlockIntervalChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainConfigEpochBlockIntervalChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainConfigEpochBlockIntervalChanged represents a EpochBlockIntervalChanged event raised by the ChainConfig contract.
type ChainConfigEpochBlockIntervalChanged struct {
	PrevValue uint32
	NewValue  uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterEpochBlockIntervalChanged is a free log retrieval operation binding the contract event 0x33c8012b0f51f8c1a1e525ea046da837d0eb4fa7473cd863e0bfb73a4f475a5a.
//
// Solidity: event EpochBlockIntervalChanged(uint32 prevValue, uint32 newValue)
func (_ChainConfig *ChainConfigFilterer) FilterEpochBlockIntervalChanged(opts *bind.FilterOpts) (*ChainConfigEpochBlockIntervalChangedIterator, error) {

	logs, sub, err := _ChainConfig.contract.FilterLogs(opts, "EpochBlockIntervalChanged")
	if err != nil {
		return nil, err
	}
	return &ChainConfigEpochBlockIntervalChangedIterator{contract: _ChainConfig.contract, event: "EpochBlockIntervalChanged", logs: logs, sub: sub}, nil
}

// WatchEpochBlockIntervalChanged is a free log subscription operation binding the contract event 0x33c8012b0f51f8c1a1e525ea046da837d0eb4fa7473cd863e0bfb73a4f475a5a.
//
// Solidity: event EpochBlockIntervalChanged(uint32 prevValue, uint32 newValue)
func (_ChainConfig *ChainConfigFilterer) WatchEpochBlockIntervalChanged(opts *bind.WatchOpts, sink chan<- *ChainConfigEpochBlockIntervalChanged) (event.Subscription, error) {

	logs, sub, err := _ChainConfig.contract.WatchLogs(opts, "EpochBlockIntervalChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainConfigEpochBlockIntervalChanged)
				if err := _ChainConfig.contract.UnpackLog(event, "EpochBlockIntervalChanged", log); err != nil {
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

// ParseEpochBlockIntervalChanged is a log parse operation binding the contract event 0x33c8012b0f51f8c1a1e525ea046da837d0eb4fa7473cd863e0bfb73a4f475a5a.
//
// Solidity: event EpochBlockIntervalChanged(uint32 prevValue, uint32 newValue)
func (_ChainConfig *ChainConfigFilterer) ParseEpochBlockIntervalChanged(log types.Log) (*ChainConfigEpochBlockIntervalChanged, error) {
	event := new(ChainConfigEpochBlockIntervalChanged)
	if err := _ChainConfig.contract.UnpackLog(event, "EpochBlockIntervalChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChainConfigFelonyThresholdChangedIterator is returned from FilterFelonyThresholdChanged and is used to iterate over the raw logs and unpacked data for FelonyThresholdChanged events raised by the ChainConfig contract.
type ChainConfigFelonyThresholdChangedIterator struct {
	Event *ChainConfigFelonyThresholdChanged // Event containing the contract specifics and raw log

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
func (it *ChainConfigFelonyThresholdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainConfigFelonyThresholdChanged)
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
		it.Event = new(ChainConfigFelonyThresholdChanged)
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
func (it *ChainConfigFelonyThresholdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainConfigFelonyThresholdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainConfigFelonyThresholdChanged represents a FelonyThresholdChanged event raised by the ChainConfig contract.
type ChainConfigFelonyThresholdChanged struct {
	PrevValue uint32
	NewValue  uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterFelonyThresholdChanged is a free log retrieval operation binding the contract event 0x67da1e9c07e7b373ed5e18cc8355caf6dfe18ab4472ec575600a2172772c6204.
//
// Solidity: event FelonyThresholdChanged(uint32 prevValue, uint32 newValue)
func (_ChainConfig *ChainConfigFilterer) FilterFelonyThresholdChanged(opts *bind.FilterOpts) (*ChainConfigFelonyThresholdChangedIterator, error) {

	logs, sub, err := _ChainConfig.contract.FilterLogs(opts, "FelonyThresholdChanged")
	if err != nil {
		return nil, err
	}
	return &ChainConfigFelonyThresholdChangedIterator{contract: _ChainConfig.contract, event: "FelonyThresholdChanged", logs: logs, sub: sub}, nil
}

// WatchFelonyThresholdChanged is a free log subscription operation binding the contract event 0x67da1e9c07e7b373ed5e18cc8355caf6dfe18ab4472ec575600a2172772c6204.
//
// Solidity: event FelonyThresholdChanged(uint32 prevValue, uint32 newValue)
func (_ChainConfig *ChainConfigFilterer) WatchFelonyThresholdChanged(opts *bind.WatchOpts, sink chan<- *ChainConfigFelonyThresholdChanged) (event.Subscription, error) {

	logs, sub, err := _ChainConfig.contract.WatchLogs(opts, "FelonyThresholdChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainConfigFelonyThresholdChanged)
				if err := _ChainConfig.contract.UnpackLog(event, "FelonyThresholdChanged", log); err != nil {
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

// ParseFelonyThresholdChanged is a log parse operation binding the contract event 0x67da1e9c07e7b373ed5e18cc8355caf6dfe18ab4472ec575600a2172772c6204.
//
// Solidity: event FelonyThresholdChanged(uint32 prevValue, uint32 newValue)
func (_ChainConfig *ChainConfigFilterer) ParseFelonyThresholdChanged(log types.Log) (*ChainConfigFelonyThresholdChanged, error) {
	event := new(ChainConfigFelonyThresholdChanged)
	if err := _ChainConfig.contract.UnpackLog(event, "FelonyThresholdChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChainConfigInitializedIterator is returned from FilterInitialized and is used to iterate over the raw logs and unpacked data for Initialized events raised by the ChainConfig contract.
type ChainConfigInitializedIterator struct {
	Event *ChainConfigInitialized // Event containing the contract specifics and raw log

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
func (it *ChainConfigInitializedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainConfigInitialized)
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
		it.Event = new(ChainConfigInitialized)
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
func (it *ChainConfigInitializedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainConfigInitializedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainConfigInitialized represents a Initialized event raised by the ChainConfig contract.
type ChainConfigInitialized struct {
	Version uint8
	Raw     types.Log // Blockchain specific contextual infos
}

// FilterInitialized is a free log retrieval operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ChainConfig *ChainConfigFilterer) FilterInitialized(opts *bind.FilterOpts) (*ChainConfigInitializedIterator, error) {

	logs, sub, err := _ChainConfig.contract.FilterLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return &ChainConfigInitializedIterator{contract: _ChainConfig.contract, event: "Initialized", logs: logs, sub: sub}, nil
}

// WatchInitialized is a free log subscription operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ChainConfig *ChainConfigFilterer) WatchInitialized(opts *bind.WatchOpts, sink chan<- *ChainConfigInitialized) (event.Subscription, error) {

	logs, sub, err := _ChainConfig.contract.WatchLogs(opts, "Initialized")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainConfigInitialized)
				if err := _ChainConfig.contract.UnpackLog(event, "Initialized", log); err != nil {
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

// ParseInitialized is a log parse operation binding the contract event 0x7f26b83ff96e1f2b6a682f133852f6798a09c465da95921460cefb3847402498.
//
// Solidity: event Initialized(uint8 version)
func (_ChainConfig *ChainConfigFilterer) ParseInitialized(log types.Log) (*ChainConfigInitialized, error) {
	event := new(ChainConfigInitialized)
	if err := _ChainConfig.contract.UnpackLog(event, "Initialized", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChainConfigMinStakingAmountChangedIterator is returned from FilterMinStakingAmountChanged and is used to iterate over the raw logs and unpacked data for MinStakingAmountChanged events raised by the ChainConfig contract.
type ChainConfigMinStakingAmountChangedIterator struct {
	Event *ChainConfigMinStakingAmountChanged // Event containing the contract specifics and raw log

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
func (it *ChainConfigMinStakingAmountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainConfigMinStakingAmountChanged)
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
		it.Event = new(ChainConfigMinStakingAmountChanged)
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
func (it *ChainConfigMinStakingAmountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainConfigMinStakingAmountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainConfigMinStakingAmountChanged represents a MinStakingAmountChanged event raised by the ChainConfig contract.
type ChainConfigMinStakingAmountChanged struct {
	PrevValue *big.Int
	NewValue  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMinStakingAmountChanged is a free log retrieval operation binding the contract event 0x973f438cb6bc47d284033b6113687c6087f4fb7a3395b03597578ae1259bf23c.
//
// Solidity: event MinStakingAmountChanged(uint256 prevValue, uint256 newValue)
func (_ChainConfig *ChainConfigFilterer) FilterMinStakingAmountChanged(opts *bind.FilterOpts) (*ChainConfigMinStakingAmountChangedIterator, error) {

	logs, sub, err := _ChainConfig.contract.FilterLogs(opts, "MinStakingAmountChanged")
	if err != nil {
		return nil, err
	}
	return &ChainConfigMinStakingAmountChangedIterator{contract: _ChainConfig.contract, event: "MinStakingAmountChanged", logs: logs, sub: sub}, nil
}

// WatchMinStakingAmountChanged is a free log subscription operation binding the contract event 0x973f438cb6bc47d284033b6113687c6087f4fb7a3395b03597578ae1259bf23c.
//
// Solidity: event MinStakingAmountChanged(uint256 prevValue, uint256 newValue)
func (_ChainConfig *ChainConfigFilterer) WatchMinStakingAmountChanged(opts *bind.WatchOpts, sink chan<- *ChainConfigMinStakingAmountChanged) (event.Subscription, error) {

	logs, sub, err := _ChainConfig.contract.WatchLogs(opts, "MinStakingAmountChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainConfigMinStakingAmountChanged)
				if err := _ChainConfig.contract.UnpackLog(event, "MinStakingAmountChanged", log); err != nil {
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

// ParseMinStakingAmountChanged is a log parse operation binding the contract event 0x973f438cb6bc47d284033b6113687c6087f4fb7a3395b03597578ae1259bf23c.
//
// Solidity: event MinStakingAmountChanged(uint256 prevValue, uint256 newValue)
func (_ChainConfig *ChainConfigFilterer) ParseMinStakingAmountChanged(log types.Log) (*ChainConfigMinStakingAmountChanged, error) {
	event := new(ChainConfigMinStakingAmountChanged)
	if err := _ChainConfig.contract.UnpackLog(event, "MinStakingAmountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChainConfigMinValidatorStakeAmountChangedIterator is returned from FilterMinValidatorStakeAmountChanged and is used to iterate over the raw logs and unpacked data for MinValidatorStakeAmountChanged events raised by the ChainConfig contract.
type ChainConfigMinValidatorStakeAmountChangedIterator struct {
	Event *ChainConfigMinValidatorStakeAmountChanged // Event containing the contract specifics and raw log

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
func (it *ChainConfigMinValidatorStakeAmountChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainConfigMinValidatorStakeAmountChanged)
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
		it.Event = new(ChainConfigMinValidatorStakeAmountChanged)
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
func (it *ChainConfigMinValidatorStakeAmountChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainConfigMinValidatorStakeAmountChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainConfigMinValidatorStakeAmountChanged represents a MinValidatorStakeAmountChanged event raised by the ChainConfig contract.
type ChainConfigMinValidatorStakeAmountChanged struct {
	PrevValue *big.Int
	NewValue  *big.Int
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMinValidatorStakeAmountChanged is a free log retrieval operation binding the contract event 0x207082661d623a88e041ad2d52c2d4ddc719880c70c3ab44aa81accff9bd86ed.
//
// Solidity: event MinValidatorStakeAmountChanged(uint256 prevValue, uint256 newValue)
func (_ChainConfig *ChainConfigFilterer) FilterMinValidatorStakeAmountChanged(opts *bind.FilterOpts) (*ChainConfigMinValidatorStakeAmountChangedIterator, error) {

	logs, sub, err := _ChainConfig.contract.FilterLogs(opts, "MinValidatorStakeAmountChanged")
	if err != nil {
		return nil, err
	}
	return &ChainConfigMinValidatorStakeAmountChangedIterator{contract: _ChainConfig.contract, event: "MinValidatorStakeAmountChanged", logs: logs, sub: sub}, nil
}

// WatchMinValidatorStakeAmountChanged is a free log subscription operation binding the contract event 0x207082661d623a88e041ad2d52c2d4ddc719880c70c3ab44aa81accff9bd86ed.
//
// Solidity: event MinValidatorStakeAmountChanged(uint256 prevValue, uint256 newValue)
func (_ChainConfig *ChainConfigFilterer) WatchMinValidatorStakeAmountChanged(opts *bind.WatchOpts, sink chan<- *ChainConfigMinValidatorStakeAmountChanged) (event.Subscription, error) {

	logs, sub, err := _ChainConfig.contract.WatchLogs(opts, "MinValidatorStakeAmountChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainConfigMinValidatorStakeAmountChanged)
				if err := _ChainConfig.contract.UnpackLog(event, "MinValidatorStakeAmountChanged", log); err != nil {
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

// ParseMinValidatorStakeAmountChanged is a log parse operation binding the contract event 0x207082661d623a88e041ad2d52c2d4ddc719880c70c3ab44aa81accff9bd86ed.
//
// Solidity: event MinValidatorStakeAmountChanged(uint256 prevValue, uint256 newValue)
func (_ChainConfig *ChainConfigFilterer) ParseMinValidatorStakeAmountChanged(log types.Log) (*ChainConfigMinValidatorStakeAmountChanged, error) {
	event := new(ChainConfigMinValidatorStakeAmountChanged)
	if err := _ChainConfig.contract.UnpackLog(event, "MinValidatorStakeAmountChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChainConfigMisdemeanorThresholdChangedIterator is returned from FilterMisdemeanorThresholdChanged and is used to iterate over the raw logs and unpacked data for MisdemeanorThresholdChanged events raised by the ChainConfig contract.
type ChainConfigMisdemeanorThresholdChangedIterator struct {
	Event *ChainConfigMisdemeanorThresholdChanged // Event containing the contract specifics and raw log

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
func (it *ChainConfigMisdemeanorThresholdChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainConfigMisdemeanorThresholdChanged)
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
		it.Event = new(ChainConfigMisdemeanorThresholdChanged)
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
func (it *ChainConfigMisdemeanorThresholdChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainConfigMisdemeanorThresholdChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainConfigMisdemeanorThresholdChanged represents a MisdemeanorThresholdChanged event raised by the ChainConfig contract.
type ChainConfigMisdemeanorThresholdChanged struct {
	PrevValue uint32
	NewValue  uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterMisdemeanorThresholdChanged is a free log retrieval operation binding the contract event 0x5aa72ebd12c45515403eef36583106e321b8707946a6ae621f5f393ee0c9677b.
//
// Solidity: event MisdemeanorThresholdChanged(uint32 prevValue, uint32 newValue)
func (_ChainConfig *ChainConfigFilterer) FilterMisdemeanorThresholdChanged(opts *bind.FilterOpts) (*ChainConfigMisdemeanorThresholdChangedIterator, error) {

	logs, sub, err := _ChainConfig.contract.FilterLogs(opts, "MisdemeanorThresholdChanged")
	if err != nil {
		return nil, err
	}
	return &ChainConfigMisdemeanorThresholdChangedIterator{contract: _ChainConfig.contract, event: "MisdemeanorThresholdChanged", logs: logs, sub: sub}, nil
}

// WatchMisdemeanorThresholdChanged is a free log subscription operation binding the contract event 0x5aa72ebd12c45515403eef36583106e321b8707946a6ae621f5f393ee0c9677b.
//
// Solidity: event MisdemeanorThresholdChanged(uint32 prevValue, uint32 newValue)
func (_ChainConfig *ChainConfigFilterer) WatchMisdemeanorThresholdChanged(opts *bind.WatchOpts, sink chan<- *ChainConfigMisdemeanorThresholdChanged) (event.Subscription, error) {

	logs, sub, err := _ChainConfig.contract.WatchLogs(opts, "MisdemeanorThresholdChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainConfigMisdemeanorThresholdChanged)
				if err := _ChainConfig.contract.UnpackLog(event, "MisdemeanorThresholdChanged", log); err != nil {
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

// ParseMisdemeanorThresholdChanged is a log parse operation binding the contract event 0x5aa72ebd12c45515403eef36583106e321b8707946a6ae621f5f393ee0c9677b.
//
// Solidity: event MisdemeanorThresholdChanged(uint32 prevValue, uint32 newValue)
func (_ChainConfig *ChainConfigFilterer) ParseMisdemeanorThresholdChanged(log types.Log) (*ChainConfigMisdemeanorThresholdChanged, error) {
	event := new(ChainConfigMisdemeanorThresholdChanged)
	if err := _ChainConfig.contract.UnpackLog(event, "MisdemeanorThresholdChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChainConfigUndelegatePeriodChangedIterator is returned from FilterUndelegatePeriodChanged and is used to iterate over the raw logs and unpacked data for UndelegatePeriodChanged events raised by the ChainConfig contract.
type ChainConfigUndelegatePeriodChangedIterator struct {
	Event *ChainConfigUndelegatePeriodChanged // Event containing the contract specifics and raw log

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
func (it *ChainConfigUndelegatePeriodChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainConfigUndelegatePeriodChanged)
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
		it.Event = new(ChainConfigUndelegatePeriodChanged)
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
func (it *ChainConfigUndelegatePeriodChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainConfigUndelegatePeriodChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainConfigUndelegatePeriodChanged represents a UndelegatePeriodChanged event raised by the ChainConfig contract.
type ChainConfigUndelegatePeriodChanged struct {
	PrevValue uint32
	NewValue  uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterUndelegatePeriodChanged is a free log retrieval operation binding the contract event 0xb191e5acbef9e4b8ce0f17af112f8984f92833324657b89fe39768885f81b6ce.
//
// Solidity: event UndelegatePeriodChanged(uint32 prevValue, uint32 newValue)
func (_ChainConfig *ChainConfigFilterer) FilterUndelegatePeriodChanged(opts *bind.FilterOpts) (*ChainConfigUndelegatePeriodChangedIterator, error) {

	logs, sub, err := _ChainConfig.contract.FilterLogs(opts, "UndelegatePeriodChanged")
	if err != nil {
		return nil, err
	}
	return &ChainConfigUndelegatePeriodChangedIterator{contract: _ChainConfig.contract, event: "UndelegatePeriodChanged", logs: logs, sub: sub}, nil
}

// WatchUndelegatePeriodChanged is a free log subscription operation binding the contract event 0xb191e5acbef9e4b8ce0f17af112f8984f92833324657b89fe39768885f81b6ce.
//
// Solidity: event UndelegatePeriodChanged(uint32 prevValue, uint32 newValue)
func (_ChainConfig *ChainConfigFilterer) WatchUndelegatePeriodChanged(opts *bind.WatchOpts, sink chan<- *ChainConfigUndelegatePeriodChanged) (event.Subscription, error) {

	logs, sub, err := _ChainConfig.contract.WatchLogs(opts, "UndelegatePeriodChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainConfigUndelegatePeriodChanged)
				if err := _ChainConfig.contract.UnpackLog(event, "UndelegatePeriodChanged", log); err != nil {
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

// ParseUndelegatePeriodChanged is a log parse operation binding the contract event 0xb191e5acbef9e4b8ce0f17af112f8984f92833324657b89fe39768885f81b6ce.
//
// Solidity: event UndelegatePeriodChanged(uint32 prevValue, uint32 newValue)
func (_ChainConfig *ChainConfigFilterer) ParseUndelegatePeriodChanged(log types.Log) (*ChainConfigUndelegatePeriodChanged, error) {
	event := new(ChainConfigUndelegatePeriodChanged)
	if err := _ChainConfig.contract.UnpackLog(event, "UndelegatePeriodChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}

// ChainConfigValidatorJailEpochLengthChangedIterator is returned from FilterValidatorJailEpochLengthChanged and is used to iterate over the raw logs and unpacked data for ValidatorJailEpochLengthChanged events raised by the ChainConfig contract.
type ChainConfigValidatorJailEpochLengthChangedIterator struct {
	Event *ChainConfigValidatorJailEpochLengthChanged // Event containing the contract specifics and raw log

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
func (it *ChainConfigValidatorJailEpochLengthChangedIterator) Next() bool {
	// If the iterator failed, stop iterating
	if it.fail != nil {
		return false
	}
	// If the iterator completed, deliver directly whatever's available
	if it.done {
		select {
		case log := <-it.logs:
			it.Event = new(ChainConfigValidatorJailEpochLengthChanged)
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
		it.Event = new(ChainConfigValidatorJailEpochLengthChanged)
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
func (it *ChainConfigValidatorJailEpochLengthChangedIterator) Error() error {
	return it.fail
}

// Close terminates the iteration process, releasing any pending underlying
// resources.
func (it *ChainConfigValidatorJailEpochLengthChangedIterator) Close() error {
	it.sub.Unsubscribe()
	return nil
}

// ChainConfigValidatorJailEpochLengthChanged represents a ValidatorJailEpochLengthChanged event raised by the ChainConfig contract.
type ChainConfigValidatorJailEpochLengthChanged struct {
	PrevValue uint32
	NewValue  uint32
	Raw       types.Log // Blockchain specific contextual infos
}

// FilterValidatorJailEpochLengthChanged is a free log retrieval operation binding the contract event 0x0a677ce4509bf46fe9bdf65c86abe71921755a78494111b1caa25df328ffcd1c.
//
// Solidity: event ValidatorJailEpochLengthChanged(uint32 prevValue, uint32 newValue)
func (_ChainConfig *ChainConfigFilterer) FilterValidatorJailEpochLengthChanged(opts *bind.FilterOpts) (*ChainConfigValidatorJailEpochLengthChangedIterator, error) {

	logs, sub, err := _ChainConfig.contract.FilterLogs(opts, "ValidatorJailEpochLengthChanged")
	if err != nil {
		return nil, err
	}
	return &ChainConfigValidatorJailEpochLengthChangedIterator{contract: _ChainConfig.contract, event: "ValidatorJailEpochLengthChanged", logs: logs, sub: sub}, nil
}

// WatchValidatorJailEpochLengthChanged is a free log subscription operation binding the contract event 0x0a677ce4509bf46fe9bdf65c86abe71921755a78494111b1caa25df328ffcd1c.
//
// Solidity: event ValidatorJailEpochLengthChanged(uint32 prevValue, uint32 newValue)
func (_ChainConfig *ChainConfigFilterer) WatchValidatorJailEpochLengthChanged(opts *bind.WatchOpts, sink chan<- *ChainConfigValidatorJailEpochLengthChanged) (event.Subscription, error) {

	logs, sub, err := _ChainConfig.contract.WatchLogs(opts, "ValidatorJailEpochLengthChanged")
	if err != nil {
		return nil, err
	}
	return event.NewSubscription(func(quit <-chan struct{}) error {
		defer sub.Unsubscribe()
		for {
			select {
			case log := <-logs:
				// New log arrived, parse the event and forward to the user
				event := new(ChainConfigValidatorJailEpochLengthChanged)
				if err := _ChainConfig.contract.UnpackLog(event, "ValidatorJailEpochLengthChanged", log); err != nil {
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

// ParseValidatorJailEpochLengthChanged is a log parse operation binding the contract event 0x0a677ce4509bf46fe9bdf65c86abe71921755a78494111b1caa25df328ffcd1c.
//
// Solidity: event ValidatorJailEpochLengthChanged(uint32 prevValue, uint32 newValue)
func (_ChainConfig *ChainConfigFilterer) ParseValidatorJailEpochLengthChanged(log types.Log) (*ChainConfigValidatorJailEpochLengthChanged, error) {
	event := new(ChainConfigValidatorJailEpochLengthChanged)
	if err := _ChainConfig.contract.UnpackLog(event, "ValidatorJailEpochLengthChanged", log); err != nil {
		return nil, err
	}
	event.Raw = log
	return event, nil
}
