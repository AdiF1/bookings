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

// BaseCampaignMetaData contains all meta data concerning the BaseCampaign contract.
var BaseCampaignMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minimum\",\"type\":\"uint256\"}],\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"approveRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"name\":\"approvers\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"approversCount\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contractBalance\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"contribute\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"}],\"name\":\"createRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"creator\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"index\",\"type\":\"uint256\"}],\"name\":\"finalizeRequest\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"minimumContribution\",\"outputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"requests\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"description\",\"type\":\"string\"},{\"internalType\":\"uint256\",\"name\":\"value\",\"type\":\"uint256\"},{\"internalType\":\"address\",\"name\":\"recipient\",\"type\":\"address\"},{\"internalType\":\"bool\",\"name\":\"complete\",\"type\":\"bool\"},{\"internalType\":\"uint256\",\"name\":\"approvalCount\",\"type\":\"uint256\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040524760015534801561001457600080fd5b506040516108353803806108358339818101604052602081101561003757600080fd5b5051600280546001600160a01b031916331790556003556107d88061005d6000396000f3fe6080604052600436106100915760003560e01c80638a9cfd55116100595780638a9cfd551461022a5780638b7afe2e146102eb578063937e09b114610300578063d7bb99ba14610315578063d7d1bbdb1461031d57610091565b806302d05d3f1461009657806303441006146100c75780630a144391146100f357806381d12c581461013a57806382fde09314610203575b600080fd5b3480156100a257600080fd5b506100ab610347565b604080516001600160a01b039092168252519081900360200190f35b3480156100d357600080fd5b506100f1600480360360208110156100ea57600080fd5b5035610356565b005b3480156100ff57600080fd5b506101266004803603602081101561011657600080fd5b50356001600160a01b0316610418565b604080519115158252519081900360200190f35b34801561014657600080fd5b506101646004803603602081101561015d57600080fd5b503561042d565b6040518080602001868152602001856001600160a01b031681526020018415158152602001838152602001828103825287818151815260200191508051906020019080838360005b838110156101c45781810151838201526020016101ac565b50505050905090810190601f1680156101f15780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b34801561020f57600080fd5b50610218610506565b60408051918252519081900360200190f35b34801561023657600080fd5b506100f16004803603606081101561024d57600080fd5b81019060208101813564010000000081111561026857600080fd5b82018360208201111561027a57600080fd5b8035906020019184600183028401116401000000008311171561029c57600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050823593505050602001356001600160a01b031661050c565b3480156102f757600080fd5b5061021861060a565b34801561030c57600080fd5b50610218610610565b6100f1610616565b34801561032957600080fd5b506100f16004803603602081101561034057600080fd5b503561064c565b6002546001600160a01b031681565b6002546001600160a01b0316331461036d57600080fd5b600080828154811061037b57fe5b9060005260206000209060050201905060026005548161039757fe5b048160030154116103a757600080fd5b6002810154600160a01b900460ff16156103c057600080fd5b600281015460018201546040516001600160a01b039092169181156108fc0291906000818181858888f19350505050158015610400573d6000803e3d6000fd5b50600201805460ff60a01b1916600160a01b17905550565b60046020526000908152604090205460ff1681565b6000818154811061043a57fe5b60009182526020918290206005919091020180546040805160026001841615610100026000190190931692909204601f8101859004850283018501909152808252919350918391908301828280156104d35780601f106104a8576101008083540402835291602001916104d3565b820191906000526020600020905b8154815290600101906020018083116104b657829003601f168201915b5050505060018301546002840154600390940154929390926001600160a01b0382169250600160a01b90910460ff169085565b60055481565b6002546001600160a01b0316331461052357600080fd5b61052b6106d5565b506040805160a08101825284815260208082018590526001600160a01b03841692820192909252600060608201819052608082018190528054600181018255908052815180519293849360059093027f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e56301926105aa928492019061070f565b5060208201516001820155604082015160028201805460608501511515600160a01b0260ff60a01b196001600160a01b039094166001600160a01b0319909216919091179290921691909117905560809091015160039091015550505050565b60015481565b60035481565b600354341161062457600080fd5b336000908152600460205260409020805460ff19166001908117909155600580549091019055565b600080828154811061065a57fe5b6000918252602080832033845260049091526040909220546005909102909101915060ff1661068857600080fd5b33600090815260048201602052604090205460ff16156106a757600080fd5b3360009081526004820160205260409020805460ff1916600190811790915560039091018054909101905550565b6040518060a00160405280606081526020016000815260200160006001600160a01b03168152602001600015158152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061075057805160ff191683800117855561077d565b8280016001018555821561077d579182015b8281111561077d578251825591602001919060010190610762565b5061078992915061078d565b5090565b5b80821115610789576000815560010161078e56fea264697066735822122088f94bb62f02d0fb38612f17c66321f7219c7fd3528324545c9b7a03faf6eca764736f6c634300060c0033",
}

// BaseCampaignABI is the input ABI used to generate the binding from.
// Deprecated: Use BaseCampaignMetaData.ABI instead.
var BaseCampaignABI = BaseCampaignMetaData.ABI

// BaseCampaignBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use BaseCampaignMetaData.Bin instead.
var BaseCampaignBin = BaseCampaignMetaData.Bin

// DeployBaseCampaign deploys a new Ethereum contract, binding an instance of BaseCampaign to it.
func DeployBaseCampaign(auth *bind.TransactOpts, backend bind.ContractBackend, minimum *big.Int) (common.Address, *types.Transaction, *BaseCampaign, error) {
	parsed, err := BaseCampaignMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(BaseCampaignBin), backend, minimum)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &BaseCampaign{BaseCampaignCaller: BaseCampaignCaller{contract: contract}, BaseCampaignTransactor: BaseCampaignTransactor{contract: contract}, BaseCampaignFilterer: BaseCampaignFilterer{contract: contract}}, nil
}

// BaseCampaign is an auto generated Go binding around an Ethereum contract.
type BaseCampaign struct {
	BaseCampaignCaller     // Read-only binding to the contract
	BaseCampaignTransactor // Write-only binding to the contract
	BaseCampaignFilterer   // Log filterer for contract events
}

// BaseCampaignCaller is an auto generated read-only Go binding around an Ethereum contract.
type BaseCampaignCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseCampaignTransactor is an auto generated write-only Go binding around an Ethereum contract.
type BaseCampaignTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseCampaignFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type BaseCampaignFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// BaseCampaignSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type BaseCampaignSession struct {
	Contract     *BaseCampaign     // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// BaseCampaignCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type BaseCampaignCallerSession struct {
	Contract *BaseCampaignCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts       // Call options to use throughout this session
}

// BaseCampaignTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type BaseCampaignTransactorSession struct {
	Contract     *BaseCampaignTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts       // Transaction auth options to use throughout this session
}

// BaseCampaignRaw is an auto generated low-level Go binding around an Ethereum contract.
type BaseCampaignRaw struct {
	Contract *BaseCampaign // Generic contract binding to access the raw methods on
}

// BaseCampaignCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type BaseCampaignCallerRaw struct {
	Contract *BaseCampaignCaller // Generic read-only contract binding to access the raw methods on
}

// BaseCampaignTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type BaseCampaignTransactorRaw struct {
	Contract *BaseCampaignTransactor // Generic write-only contract binding to access the raw methods on
}

// NewBaseCampaign creates a new instance of BaseCampaign, bound to a specific deployed contract.
func NewBaseCampaign(address common.Address, backend bind.ContractBackend) (*BaseCampaign, error) {
	contract, err := bindBaseCampaign(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &BaseCampaign{BaseCampaignCaller: BaseCampaignCaller{contract: contract}, BaseCampaignTransactor: BaseCampaignTransactor{contract: contract}, BaseCampaignFilterer: BaseCampaignFilterer{contract: contract}}, nil
}

// NewBaseCampaignCaller creates a new read-only instance of BaseCampaign, bound to a specific deployed contract.
func NewBaseCampaignCaller(address common.Address, caller bind.ContractCaller) (*BaseCampaignCaller, error) {
	contract, err := bindBaseCampaign(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &BaseCampaignCaller{contract: contract}, nil
}

// NewBaseCampaignTransactor creates a new write-only instance of BaseCampaign, bound to a specific deployed contract.
func NewBaseCampaignTransactor(address common.Address, transactor bind.ContractTransactor) (*BaseCampaignTransactor, error) {
	contract, err := bindBaseCampaign(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &BaseCampaignTransactor{contract: contract}, nil
}

// NewBaseCampaignFilterer creates a new log filterer instance of BaseCampaign, bound to a specific deployed contract.
func NewBaseCampaignFilterer(address common.Address, filterer bind.ContractFilterer) (*BaseCampaignFilterer, error) {
	contract, err := bindBaseCampaign(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &BaseCampaignFilterer{contract: contract}, nil
}

// bindBaseCampaign binds a generic wrapper to an already deployed contract.
func bindBaseCampaign(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(BaseCampaignABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseCampaign *BaseCampaignRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseCampaign.Contract.BaseCampaignCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseCampaign *BaseCampaignRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseCampaign.Contract.BaseCampaignTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseCampaign *BaseCampaignRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseCampaign.Contract.BaseCampaignTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_BaseCampaign *BaseCampaignCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _BaseCampaign.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_BaseCampaign *BaseCampaignTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseCampaign.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_BaseCampaign *BaseCampaignTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _BaseCampaign.Contract.contract.Transact(opts, method, params...)
}

// Approvers is a free data retrieval call binding the contract method 0x0a144391.
//
// Solidity: function approvers(address ) view returns(bool)
func (_BaseCampaign *BaseCampaignCaller) Approvers(opts *bind.CallOpts, arg0 common.Address) (bool, error) {
	var out []interface{}
	err := _BaseCampaign.contract.Call(opts, &out, "approvers", arg0)

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Approvers is a free data retrieval call binding the contract method 0x0a144391.
//
// Solidity: function approvers(address ) view returns(bool)
func (_BaseCampaign *BaseCampaignSession) Approvers(arg0 common.Address) (bool, error) {
	return _BaseCampaign.Contract.Approvers(&_BaseCampaign.CallOpts, arg0)
}

// Approvers is a free data retrieval call binding the contract method 0x0a144391.
//
// Solidity: function approvers(address ) view returns(bool)
func (_BaseCampaign *BaseCampaignCallerSession) Approvers(arg0 common.Address) (bool, error) {
	return _BaseCampaign.Contract.Approvers(&_BaseCampaign.CallOpts, arg0)
}

// ApproversCount is a free data retrieval call binding the contract method 0x82fde093.
//
// Solidity: function approversCount() view returns(uint256)
func (_BaseCampaign *BaseCampaignCaller) ApproversCount(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BaseCampaign.contract.Call(opts, &out, "approversCount")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ApproversCount is a free data retrieval call binding the contract method 0x82fde093.
//
// Solidity: function approversCount() view returns(uint256)
func (_BaseCampaign *BaseCampaignSession) ApproversCount() (*big.Int, error) {
	return _BaseCampaign.Contract.ApproversCount(&_BaseCampaign.CallOpts)
}

// ApproversCount is a free data retrieval call binding the contract method 0x82fde093.
//
// Solidity: function approversCount() view returns(uint256)
func (_BaseCampaign *BaseCampaignCallerSession) ApproversCount() (*big.Int, error) {
	return _BaseCampaign.Contract.ApproversCount(&_BaseCampaign.CallOpts)
}

// ContractBalance is a free data retrieval call binding the contract method 0x8b7afe2e.
//
// Solidity: function contractBalance() view returns(uint256)
func (_BaseCampaign *BaseCampaignCaller) ContractBalance(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BaseCampaign.contract.Call(opts, &out, "contractBalance")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// ContractBalance is a free data retrieval call binding the contract method 0x8b7afe2e.
//
// Solidity: function contractBalance() view returns(uint256)
func (_BaseCampaign *BaseCampaignSession) ContractBalance() (*big.Int, error) {
	return _BaseCampaign.Contract.ContractBalance(&_BaseCampaign.CallOpts)
}

// ContractBalance is a free data retrieval call binding the contract method 0x8b7afe2e.
//
// Solidity: function contractBalance() view returns(uint256)
func (_BaseCampaign *BaseCampaignCallerSession) ContractBalance() (*big.Int, error) {
	return _BaseCampaign.Contract.ContractBalance(&_BaseCampaign.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_BaseCampaign *BaseCampaignCaller) Creator(opts *bind.CallOpts) (common.Address, error) {
	var out []interface{}
	err := _BaseCampaign.contract.Call(opts, &out, "creator")

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_BaseCampaign *BaseCampaignSession) Creator() (common.Address, error) {
	return _BaseCampaign.Contract.Creator(&_BaseCampaign.CallOpts)
}

// Creator is a free data retrieval call binding the contract method 0x02d05d3f.
//
// Solidity: function creator() view returns(address)
func (_BaseCampaign *BaseCampaignCallerSession) Creator() (common.Address, error) {
	return _BaseCampaign.Contract.Creator(&_BaseCampaign.CallOpts)
}

// MinimumContribution is a free data retrieval call binding the contract method 0x937e09b1.
//
// Solidity: function minimumContribution() view returns(uint256)
func (_BaseCampaign *BaseCampaignCaller) MinimumContribution(opts *bind.CallOpts) (*big.Int, error) {
	var out []interface{}
	err := _BaseCampaign.contract.Call(opts, &out, "minimumContribution")

	if err != nil {
		return *new(*big.Int), err
	}

	out0 := *abi.ConvertType(out[0], new(*big.Int)).(**big.Int)

	return out0, err

}

// MinimumContribution is a free data retrieval call binding the contract method 0x937e09b1.
//
// Solidity: function minimumContribution() view returns(uint256)
func (_BaseCampaign *BaseCampaignSession) MinimumContribution() (*big.Int, error) {
	return _BaseCampaign.Contract.MinimumContribution(&_BaseCampaign.CallOpts)
}

// MinimumContribution is a free data retrieval call binding the contract method 0x937e09b1.
//
// Solidity: function minimumContribution() view returns(uint256)
func (_BaseCampaign *BaseCampaignCallerSession) MinimumContribution() (*big.Int, error) {
	return _BaseCampaign.Contract.MinimumContribution(&_BaseCampaign.CallOpts)
}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) view returns(string description, uint256 value, address recipient, bool complete, uint256 approvalCount)
func (_BaseCampaign *BaseCampaignCaller) Requests(opts *bind.CallOpts, arg0 *big.Int) (struct {
	Description   string
	Value         *big.Int
	Recipient     common.Address
	Complete      bool
	ApprovalCount *big.Int
}, error) {
	var out []interface{}
	err := _BaseCampaign.contract.Call(opts, &out, "requests", arg0)

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
func (_BaseCampaign *BaseCampaignSession) Requests(arg0 *big.Int) (struct {
	Description   string
	Value         *big.Int
	Recipient     common.Address
	Complete      bool
	ApprovalCount *big.Int
}, error) {
	return _BaseCampaign.Contract.Requests(&_BaseCampaign.CallOpts, arg0)
}

// Requests is a free data retrieval call binding the contract method 0x81d12c58.
//
// Solidity: function requests(uint256 ) view returns(string description, uint256 value, address recipient, bool complete, uint256 approvalCount)
func (_BaseCampaign *BaseCampaignCallerSession) Requests(arg0 *big.Int) (struct {
	Description   string
	Value         *big.Int
	Recipient     common.Address
	Complete      bool
	ApprovalCount *big.Int
}, error) {
	return _BaseCampaign.Contract.Requests(&_BaseCampaign.CallOpts, arg0)
}

// ApproveRequest is a paid mutator transaction binding the contract method 0xd7d1bbdb.
//
// Solidity: function approveRequest(uint256 index) returns()
func (_BaseCampaign *BaseCampaignTransactor) ApproveRequest(opts *bind.TransactOpts, index *big.Int) (*types.Transaction, error) {
	return _BaseCampaign.contract.Transact(opts, "approveRequest", index)
}

// ApproveRequest is a paid mutator transaction binding the contract method 0xd7d1bbdb.
//
// Solidity: function approveRequest(uint256 index) returns()
func (_BaseCampaign *BaseCampaignSession) ApproveRequest(index *big.Int) (*types.Transaction, error) {
	return _BaseCampaign.Contract.ApproveRequest(&_BaseCampaign.TransactOpts, index)
}

// ApproveRequest is a paid mutator transaction binding the contract method 0xd7d1bbdb.
//
// Solidity: function approveRequest(uint256 index) returns()
func (_BaseCampaign *BaseCampaignTransactorSession) ApproveRequest(index *big.Int) (*types.Transaction, error) {
	return _BaseCampaign.Contract.ApproveRequest(&_BaseCampaign.TransactOpts, index)
}

// Contribute is a paid mutator transaction binding the contract method 0xd7bb99ba.
//
// Solidity: function contribute() payable returns()
func (_BaseCampaign *BaseCampaignTransactor) Contribute(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _BaseCampaign.contract.Transact(opts, "contribute")
}

// Contribute is a paid mutator transaction binding the contract method 0xd7bb99ba.
//
// Solidity: function contribute() payable returns()
func (_BaseCampaign *BaseCampaignSession) Contribute() (*types.Transaction, error) {
	return _BaseCampaign.Contract.Contribute(&_BaseCampaign.TransactOpts)
}

// Contribute is a paid mutator transaction binding the contract method 0xd7bb99ba.
//
// Solidity: function contribute() payable returns()
func (_BaseCampaign *BaseCampaignTransactorSession) Contribute() (*types.Transaction, error) {
	return _BaseCampaign.Contract.Contribute(&_BaseCampaign.TransactOpts)
}

// CreateRequest is a paid mutator transaction binding the contract method 0x8a9cfd55.
//
// Solidity: function createRequest(string description, uint256 value, address recipient) returns()
func (_BaseCampaign *BaseCampaignTransactor) CreateRequest(opts *bind.TransactOpts, description string, value *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _BaseCampaign.contract.Transact(opts, "createRequest", description, value, recipient)
}

// CreateRequest is a paid mutator transaction binding the contract method 0x8a9cfd55.
//
// Solidity: function createRequest(string description, uint256 value, address recipient) returns()
func (_BaseCampaign *BaseCampaignSession) CreateRequest(description string, value *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _BaseCampaign.Contract.CreateRequest(&_BaseCampaign.TransactOpts, description, value, recipient)
}

// CreateRequest is a paid mutator transaction binding the contract method 0x8a9cfd55.
//
// Solidity: function createRequest(string description, uint256 value, address recipient) returns()
func (_BaseCampaign *BaseCampaignTransactorSession) CreateRequest(description string, value *big.Int, recipient common.Address) (*types.Transaction, error) {
	return _BaseCampaign.Contract.CreateRequest(&_BaseCampaign.TransactOpts, description, value, recipient)
}

// FinalizeRequest is a paid mutator transaction binding the contract method 0x03441006.
//
// Solidity: function finalizeRequest(uint256 index) returns()
func (_BaseCampaign *BaseCampaignTransactor) FinalizeRequest(opts *bind.TransactOpts, index *big.Int) (*types.Transaction, error) {
	return _BaseCampaign.contract.Transact(opts, "finalizeRequest", index)
}

// FinalizeRequest is a paid mutator transaction binding the contract method 0x03441006.
//
// Solidity: function finalizeRequest(uint256 index) returns()
func (_BaseCampaign *BaseCampaignSession) FinalizeRequest(index *big.Int) (*types.Transaction, error) {
	return _BaseCampaign.Contract.FinalizeRequest(&_BaseCampaign.TransactOpts, index)
}

// FinalizeRequest is a paid mutator transaction binding the contract method 0x03441006.
//
// Solidity: function finalizeRequest(uint256 index) returns()
func (_BaseCampaign *BaseCampaignTransactorSession) FinalizeRequest(index *big.Int) (*types.Transaction, error) {
	return _BaseCampaign.Contract.FinalizeRequest(&_BaseCampaign.TransactOpts, index)
}
