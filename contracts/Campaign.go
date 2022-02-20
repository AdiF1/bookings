// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package adf2

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
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minimum\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"_creator\",\"type\":\"address\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"approveRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"approvers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"approversCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contribute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"createRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"finalizeRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumContribution\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"requests\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"complete\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"approvalCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b5060405161081f38038061081f8339818101604052604081101561003357600080fd5b508051602090910151600180546001600160a01b0319166001600160a01b039092169190911790556002556107b28061006d6000396000f3fe6080604052600436106100865760003560e01c806382fde0931161005957806382fde093146101f85780638a9cfd551461021f578063937e09b1146102e0578063d7bb99ba146102f5578063d7d1bbdb146102fd57610086565b806302d05d3f1461008b57806303441006146100bc5780630a144391146100e857806381d12c581461012f575b600080fd5b34801561009757600080fd5b506100a0610327565b604080516001600160a01b039092168252519081900360200190f35b3480156100c857600080fd5b506100e6600480360360208110156100df57600080fd5b5035610336565b005b3480156100f457600080fd5b5061011b6004803603602081101561010b57600080fd5b50356001600160a01b03166103f8565b604080519115158252519081900360200190f35b34801561013b57600080fd5b506101596004803603602081101561015257600080fd5b503561040d565b6040518080602001868152602001856001600160a01b031681526020018415158152602001838152602001828103825287818151815260200191508051906020019080838360005b838110156101b95781810151838201526020016101a1565b50505050905090810190601f1680156101e65780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b34801561020457600080fd5b5061020d6104e6565b60408051918252519081900360200190f35b34801561022b57600080fd5b506100e66004803603606081101561024257600080fd5b81019060208101813564010000000081111561025d57600080fd5b82018360208201111561026f57600080fd5b8035906020019184600183028401116401000000008311171561029157600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050823593505050602001356001600160a01b03166104ec565b3480156102ec57600080fd5b5061020d6105ea565b6100e66105f0565b34801561030957600080fd5b506100e66004803603602081101561032057600080fd5b5035610626565b6001546001600160a01b031681565b6001546001600160a01b0316331461034d57600080fd5b600080828154811061035b57fe5b9060005260206000209060050201905060026004548161037757fe5b0481600301541161038757600080fd5b6002810154600160a01b900460ff16156103a057600080fd5b600281015460018201546040516001600160a01b039092169181156108fc0291906000818181858888f193505050501580156103e0573d6000803e3d6000fd5b50600201805460ff60a01b1916600160a01b17905550565b60036020526000908152604090205460ff1681565b6000818154811061041a57fe5b60009182526020918290206005919091020180546040805160026001841615610100026000190190931692909204601f8101859004850283018501909152808252919350918391908301828280156104b35780601f10610488576101008083540402835291602001916104b3565b820191906000526020600020905b81548152906001019060200180831161049657829003601f168201915b5050505060018301546002840154600390940154929390926001600160a01b0382169250600160a01b90910460ff169085565b60045481565b6001546001600160a01b0316331461050357600080fd5b61050b6106af565b506040805160a08101825284815260208082018590526001600160a01b03841692820192909252600060608201819052608082018190528054600181018255908052815180519293849360059093027f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563019261058a92849201906106e9565b5060208201516001820155604082015160028201805460608501511515600160a01b0260ff60a01b196001600160a01b039094166001600160a01b0319909216919091179290921691909117905560809091015160039091015550505050565b60025481565b60025434116105fe57600080fd5b336000908152600360205260409020805460ff19166001908117909155600480549091019055565b600080828154811061063457fe5b6000918252602080832033845260039091526040909220546005909102909101915060ff1661066257600080fd5b33600090815260048201602052604090205460ff161561068157600080fd5b3360009081526004820160205260409020805460ff1916600190811790915560039091018054909101905550565b6040518060a00160405280606081526020016000815260200160006001600160a01b03168152602001600015158152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061072a57805160ff1916838001178555610757565b82800160010185558215610757579182015b8281111561075757825182559160200191906001019061073c565b50610763929150610767565b5090565b5b80821115610763576000815560010161076856fea2646970667358221220a0f818c99c8abf00022edbd0d78211606dcc860c82048c5ede33410e16a797f164736f6c634300060c0033",
}

// CampaignABI is the input ABI used to generate the binding from.
// Deprecated: Use CampaignMetaData.ABI instead.
var CampaignABI = CampaignMetaData.ABI

// CampaignBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CampaignMetaData.Bin instead.
var CampaignBin = CampaignMetaData.Bin

// DeployCampaign deploys a new Ethereum contract, binding an instance of Campaign to it.
func DeployCampaign(auth *bind.TransactOpts, backend bind.ContractBackend, minimum *big.Int, _creator common.Address) (common.Address, *types.Transaction, *Campaign, error) {
	parsed, err := CampaignMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CampaignBin), backend, minimum, _creator)
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

// Approvers is a free data retrieval call binding the contract method 0x0a144391.
//
// Solidity: function approvers(address ) view returns(bool)
func (_Campaign *CampaignCaller) Approvers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _Campaign.contract.Call(opts, &out, "approvers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Approvers is a free data retrieval call binding the contract method 0x0a144391.
//
// Solidity: function approvers(address ) view returns(bool)
func (_Campaign *CampaignSession) Approvers(arg0 common.Address) (bool, error) {
	return _Campaign.Contract.Approvers(&_Campaign.CallOpts, arg0)
}

// Approvers is a free data retrieval call binding the contract method 0x0a144391.
//
// Solidity: function approvers(address ) view returns(bool)
func (_Campaign *CampaignCallerSession) Approvers(arg0 common.Address) (bool, error) {
	return _Campaign.Contract.Approvers(&_Campaign.CallOpts, arg0)
}

// ApproversCount is a free data retrieval call binding the contract method 0x82fde093.
//
// Solidity: function approversCount() view returns(uint256)
func (_Campaign *CampaignCaller) ApproversCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _Campaign.contract.Call(opts, &out, "approversCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ApproversCount is a free data retrieval call binding the contract method 0x82fde093.
//
// Solidity: function approversCount() view returns(uint256)
func (_Campaign *CampaignSession) ApproversCount() (*big.Int, error) {
	return _Campaign.Contract.ApproversCount(&_Campaign.CallOpts)
}

// ApproversCount is a free data retrieval call binding the contract method 0x82fde093.
//
// Solidity: function approversCount() view returns(uint256)
func (_Campaign *CampaignCallerSession) ApproversCount() (*big.Int, error) {
	return _Campaign.Contract.ApproversCount(&_Campaign.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_Campaign *CampaignCaller) Creator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _Campaign.contract.Call(opts, &out, "creator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_Campaign *CampaignSession) Creator() (common.Address, error) {
	return _Campaign.Contract.Creator(&_Campaign.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_Campaign *CampaignCallerSession) Creator() (common.Address, error) {
	return _Campaign.Contract.Creator(&_Campaign.CallOpts)
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

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) view returns(string description, uint256 value, address recipient, bool complete, uint256 approvalCount)
func (_Campaign *CampaignCaller) Requests(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Description   string
	Value         *big.Int
	Recipient     common.Address
	Complete      bool
	ApprovalCount *big.Int
}, error) {
	var out []interface{}
	err := _Campaign.contract.Call(opts, &out, "requests", arg0)

	outstruct := new(struct {
		Description   string
		Value         *big.Int
		Recipient     common.Address
		Complete      bool
		ApprovalCount *big.Int
	})
	if err != nil {
		return *outstruct, err
	}

	outstruct.Description = *abi.ConvertType(out[0], new(string)).(*string)
	outstruct.Value = *abi.ConvertType(out[1], new(*big.Int)).(**big.Int)
	outstruct.Recipient = *abi.ConvertType(out[2], new(common.Address)).(*common.Address)
	outstruct.Complete = *abi.ConvertType(out[3], new(bool)).(*bool)
	outstruct.ApprovalCount = *abi.ConvertType(out[4], new(*big.Int)).(**big.Int)

	return *outstruct, err

}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) view returns(string description, uint256 value, address recipient, bool complete, uint256 approvalCount)
func (_Campaign *CampaignSession) Requests(arg0 *big.Int) (struct {
	Description   string
	Value         *big.Int
	Recipient     common.Address
	Complete      bool
	ApprovalCount *big.Int
}, error) {
	return _Campaign.Contract.Requests(&_Campaign.CallOpts, arg0)
}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) view returns(string description, uint256 value, address recipient, bool complete, uint256 approvalCount)
func (_Campaign *CampaignCallerSession) Requests(arg0 *big.Int) (struct {
	Description   string
	Value         *big.Int
	Recipient     common.Address
	Complete      bool
	ApprovalCount *big.Int
}, error) {
	return _Campaign.Contract.Requests(&_Campaign.CallOpts, arg0)
}

// ApproveRequest is a paid mutator transaction binding the contract method 0xd7d1bbdb.
//
// Solidity: function approveRequest(uint256 index) returns()
func (_Campaign *CampaignTransactor) ApproveRequest(opts *bind.TransactOpts, index *big.Int) (*types.Transaction, error) {
	return _Campaign.contract.Transact(opts, "approveRequest", index)
}

// ApproveRequest is a paid mutator transaction binding the contract method 0xd7d1bbdb.
//
// Solidity: function approveRequest(uint256 index) returns()
func (_Campaign *CampaignSession) ApproveRequest(index *big.Int) (*types.Transaction, error) {
	return _Campaign.Contract.ApproveRequest(&_Campaign.TransactOpts, index)
}

// ApproveRequest is a paid mutator transaction binding the contract method 0xd7d1bbdb.
//
// Solidity: function approveRequest(uint256 index) returns()
func (_Campaign *CampaignTransactorSession) ApproveRequest(index *big.Int) (*types.Transaction, error) {
	return _Campaign.Contract.ApproveRequest(&_Campaign.TransactOpts, index)
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

// CreateRequest is a paid mutator transaction binding the contract method 0x8a9cfd55.
//
// Solidity: function createRequest(string description, uint256 value, address recipient) returns()
func (_Campaign *CampaignTransactor) CreateRequest(opts *bind.TransactOpts, description string, value *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Campaign.contract.Transact(opts, "createRequest", description, value, recipient)
}

// CreateRequest is a paid mutator transaction binding the contract method 0x8a9cfd55.
//
// Solidity: function createRequest(string description, uint256 value, address recipient) returns()
func (_Campaign *CampaignSession) CreateRequest(description string, value *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Campaign.Contract.CreateRequest(&_Campaign.TransactOpts, description, value, recipient)
}

// CreateRequest is a paid mutator transaction binding the contract method 0x8a9cfd55.
//
// Solidity: function createRequest(string description, uint256 value, address recipient) returns()
func (_Campaign *CampaignTransactorSession) CreateRequest(description string, value *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _Campaign.Contract.CreateRequest(&_Campaign.TransactOpts, description, value, recipient)
}

// FinalizeRequest is a paid mutator transaction binding the contract method 0x03441006.
//
// Solidity: function finalizeRequest(uint256 index) returns()
func (_Campaign *CampaignTransactor) FinalizeRequest(opts *bind.TransactOpts, index *big.Int) (*types.Transaction, error) {
	return _Campaign.contract.Transact(opts, "finalizeRequest", index)
}

// FinalizeRequest is a paid mutator transaction binding the contract method 0x03441006.
//
// Solidity: function finalizeRequest(uint256 index) returns()
func (_Campaign *CampaignSession) FinalizeRequest(index *big.Int) (*types.Transaction, error) {
	return _Campaign.Contract.FinalizeRequest(&_Campaign.TransactOpts, index)
}

// FinalizeRequest is a paid mutator transaction binding the contract method 0x03441006.
//
// Solidity: function finalizeRequest(uint256 index) returns()
func (_Campaign *CampaignTransactorSession) FinalizeRequest(index *big.Int) (*types.Transaction, error) {
	return _Campaign.Contract.FinalizeRequest(&_Campaign.TransactOpts, index)
}
