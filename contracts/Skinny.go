// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package Campaign

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
)

// CampaignMetaData contains all meta data concerning the Campaign contract.
var CampaignMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"creator\",\"type\":\"address\"}],\"name\":\"Campaign\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contribute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"manager\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumContribution\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040524760025534801561001457600080fd5b5061018d806100246000396000f3fe60806040526004361061004a5760003560e01c8063481c6a751461004f5780637d6c6ff71461008c5780638b7afe2e146100ce578063937e09b1146100f2578063d7bb99ba14610108575b600080fd5b34801561005b57600080fd5b5060005461006f906001600160a01b031681565b6040516001600160a01b0390911681526020015b60405180910390f35b34801561009857600080fd5b506100cc6100a736600461011b565b600080546001600160a01b0319166001600160a01b0392909216919091179055600155565b005b3480156100da57600080fd5b506100e460025481565b604051908152602001610083565b3480156100fe57600080fd5b506100e460015481565b6100cc600154341161011957600080fd5b565b6000806040838503121561012e57600080fd5b8235915060208301356001600160a01b038116811461014c57600080fd5b80915050925092905056fea2646970667358221220e89b805883eae3bd17dd0f3290719674788b860249cdcdc97613c9aaa996aa3d64736f6c634300080b0033",
}

// CampaignABI is the input ABI used to generate the binding from.
// Deprecated: Use CampaignMetaData.ABI instead.
var CampaignABI = CampaignMetaData.ABI

// CampaignBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CampaignMetaData.Bin instead.
var CampaignBin = CampaignMetaData.Bin

// DeployCampaign deploys a new Ethereum contract, binding an instance of Campaign to it.
func DeployCampaign(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *Campaign, error) {
	parsed, err := CampaignMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CampaignBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &Campaign{CampaignCaller: CampaignCaller{contract: contract}, CampaignTransactor: CampaignTransactor{contract: contract}, CampaignFilterer: CampaignFilterer{contract: contract}}, nil
}

// Campaign is an auto generated Go binding around an Ethereum contract.
type Campaign struct {
	CampaignCaller     // Read-only binding to the contract
	CampaignTransactor // Write-only binding to the contract
	CampaignFilterer   // Log filterer for contract events
}

// CampaignCaller is an auto generated read-only Go binding around an Ethereum contract.
type CampaignCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CampaignTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CampaignTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CampaignFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CampaignFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CampaignSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CampaignSession struct {
	Contract     *Campaign         // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CampaignCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CampaignCallerSession struct {
	Contract *CampaignCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts   // Call options to use throughout this session
}

// CampaignTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CampaignTransactorSession struct {
	Contract     *CampaignTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts   // Transaction auth options to use throughout this session
}

// CampaignRaw is an auto generated low-level Go binding around an Ethereum contract.
type CampaignRaw struct {
	Contract *Campaign // Generic contract binding to access the raw methods on
}

// CampaignCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CampaignCallerRaw struct {
	Contract *CampaignCaller // Generic read-only contract binding to access the raw methods on
}

// CampaignTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CampaignTransactorRaw struct {
	Contract *CampaignTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCampaign creates a new instance of Campaign, bound to a specific deployed contract.
func NewCampaign(address common.Address, backend bind.ContractBackend) (*Campaign, error) {
	contract, err := bindCampaign(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &Campaign{CampaignCaller: CampaignCaller{contract: contract}, CampaignTransactor: CampaignTransactor{contract: contract}, CampaignFilterer: CampaignFilterer{contract: contract}}, nil
}

// NewCampaignCaller creates a new read-only instance of Campaign, bound to a specific deployed contract.
func NewCampaignCaller(address common.Address, caller bind.ContractCaller) (*CampaignCaller, error) {
	contract, err := bindCampaign(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CampaignCaller{contract: contract}, nil
}

// NewCampaignTransactor creates a new write-only instance of Campaign, bound to a specific deployed contract.
func NewCampaignTransactor(address common.Address, transactor bind.ContractTransactor) (*CampaignTransactor, error) {
	contract, err := bindCampaign(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CampaignTransactor{contract: contract}, nil
}

// NewCampaignFilterer creates a new log filterer instance of Campaign, bound to a specific deployed contract.
func NewCampaignFilterer(address common.Address, filterer bind.ContractFilterer) (*CampaignFilterer, error) {
	contract, err := bindCampaign(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CampaignFilterer{contract: contract}, nil
}

// bindCampaign binds a generic wrapper to an already deployed contract.
func bindCampaign(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CampaignABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Campaign *CampaignRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Campaign.Contract.CampaignCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Campaign *CampaignRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Campaign.Contract.CampaignTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Campaign *CampaignRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Campaign.Contract.CampaignTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_Campaign *CampaignCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _Campaign.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_Campaign *CampaignTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Campaign.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_Campaign *CampaignTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _Campaign.Contract.contract.Transact(opts, method, params...)
}

// ContractBalance is a free data retrieval call binding the contract method 0x8b7afe2e.
//
// Solidity: function contractBalance() view returns(uint256)
func (_Campaign *CampaignCaller) ContractBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Campaign.contract.Call(opts, &out, "contractBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ContractBalance is a free data retrieval call binding the contract method 0x8b7afe2e.
//
// Solidity: function contractBalance() view returns(uint256)
func (_Campaign *CampaignSession) ContractBalance() (*big.Int, error) {
	return _Campaign.Contract.ContractBalance(&_Campaign.CallOpts)
}

// ContractBalance is a free data retrieval call binding the contract method 0x8b7afe2e.
//
// Solidity: function contractBalance() view returns(uint256)
func (_Campaign *CampaignCallerSession) ContractBalance() (*big.Int, error) {
	return _Campaign.Contract.ContractBalance(&_Campaign.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_Campaign *CampaignCaller) Manager(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Campaign.contract.Call(opts, &out, "manager")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_Campaign *CampaignSession) Manager() (common.Address, error) {
	return _Campaign.Contract.Manager(&_Campaign.CallOpts)
}

// Manager is a free data retrieval call binding the contract method 0x481c6a75.
//
// Solidity: function manager() view returns(address)
func (_Campaign *CampaignCallerSession) Manager() (common.Address, error) {
	return _Campaign.Contract.Manager(&_Campaign.CallOpts)
}

// MinimumContribution is a free data retrieval call binding the contract method 0x937e09b1.
//
// Solidity: function minimumContribution() view returns(uint256)
func (_Campaign *CampaignCaller) MinimumContribution(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Campaign.contract.Call(opts, &out, "minimumContribution")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumContribution is a free data retrieval call binding the contract method 0x937e09b1.
//
// Solidity: function minimumContribution() view returns(uint256)
func (_Campaign *CampaignSession) MinimumContribution() (*big.Int, error) {
	return _Campaign.Contract.MinimumContribution(&_Campaign.CallOpts)
}

// MinimumContribution is a free data retrieval call binding the contract method 0x937e09b1.
//
// Solidity: function minimumContribution() view returns(uint256)
func (_Campaign *CampaignCallerSession) MinimumContribution() (*big.Int, error) {
	return _Campaign.Contract.MinimumContribution(&_Campaign.CallOpts)
}

// Campaign is a paid mutator transaction binding the contract method 0x7d6c6ff7.
//
// Solidity: function Campaign(uint256 minimum, address creator) returns()
func (_Campaign *CampaignTransactor) Campaign(opts *bind.TransactOpts, minimum *big.Int, creator common.Address) (*types.Transaction, error) {
	return _Campaign.contract.Transact(opts, "Campaign", minimum, creator)
}

// Campaign is a paid mutator transaction binding the contract method 0x7d6c6ff7.
//
// Solidity: function Campaign(uint256 minimum, address creator) returns()
func (_Campaign *CampaignSession) Campaign(minimum *big.Int, creator common.Address) (*types.Transaction, error) {
	return _Campaign.Contract.Campaign(&_Campaign.TransactOpts, minimum, creator)
}

// Campaign is a paid mutator transaction binding the contract method 0x7d6c6ff7.
//
// Solidity: function Campaign(uint256 minimum, address creator) returns()
func (_Campaign *CampaignTransactorSession) Campaign(minimum *big.Int, creator common.Address) (*types.Transaction, error) {
	return _Campaign.Contract.Campaign(&_Campaign.TransactOpts, minimum, creator)
}

// Contribute is a paid mutator transaction binding the contract method 0xd7bb99ba.
//
// Solidity: function contribute() payable returns()
func (_Campaign *CampaignTransactor) Contribute(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _Campaign.contract.Transact(opts, "contribute")
}

// Contribute is a paid mutator transaction binding the contract method 0xd7bb99ba.
//
// Solidity: function contribute() payable returns()
func (_Campaign *CampaignSession) Contribute() (*types.Transaction, error) {
	return _Campaign.Contract.Contribute(&_Campaign.TransactOpts)
}

// Contribute is a paid mutator transaction binding the contract method 0xd7bb99ba.
//
// Solidity: function contribute() payable returns()
func (_Campaign *CampaignTransactorSession) Contribute() (*types.Transaction, error) {
	return _Campaign.Contract.Contribute(&_Campaign.TransactOpts)
}
