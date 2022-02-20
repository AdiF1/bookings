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

// CampaignFactoryMetaData contains all meta data concerning the CampaignFactory contract.
var CampaignFactoryMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"minimum\",\"type\":\"uint256\"}],\"name\":\"createCampaign\",\"outputs\":[],\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"deployedCampaigns\",\"outputs\":[{\"internalType\":\"address\",\"name\":\"\",\"type\":\"address\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDeployedCampaigns\",\"outputs\":[{\"internalType\":\"address[]\",\"name\":\"\",\"type\":\"address[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x608060405234801561001057600080fd5b50610a72806100206000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c8063339d50a5146100465780634acb9d4f1461007f578063a3303a75146100d7575b600080fd5b6100636004803603602081101561005c57600080fd5b50356100f6565b604080516001600160a01b039092168252519081900360200190f35b61008761011d565b60408051602080825283518183015283519192839290830191858101910280838360005b838110156100c35781810151838201526020016100ab565b505050509050019250505060405180910390f35b6100f4600480360360208110156100ed57600080fd5b503561017f565b005b6000818154811061010357fe5b6000918252602090912001546001600160a01b0316905081565b6060600080548060200260200160405190810160405280929190818152602001828054801561017557602002820191906000526020600020905b81546001600160a01b03168152600190910190602001808311610157575b5050505050905090565b6000813360405161018f90610210565b9182526001600160a01b03166020820152604080519182900301906000f0801580156101bf573d6000803e3d6000fd5b50600080546001810182559080527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e5630180546001600160a01b0319166001600160a01b039092169190911790555050565b61081f8061021e8339019056fe608060405234801561001057600080fd5b5060405161081f38038061081f8339818101604052604081101561003357600080fd5b508051602090910151600180546001600160a01b0319166001600160a01b039092169190911790556002556107b28061006d6000396000f3fe6080604052600436106100865760003560e01c806382fde0931161005957806382fde093146101f85780638a9cfd551461021f578063937e09b1146102e0578063d7bb99ba146102f5578063d7d1bbdb146102fd57610086565b806302d05d3f1461008b57806303441006146100bc5780630a144391146100e857806381d12c581461012f575b600080fd5b34801561009757600080fd5b506100a0610327565b604080516001600160a01b039092168252519081900360200190f35b3480156100c857600080fd5b506100e6600480360360208110156100df57600080fd5b5035610336565b005b3480156100f457600080fd5b5061011b6004803603602081101561010b57600080fd5b50356001600160a01b03166103f8565b604080519115158252519081900360200190f35b34801561013b57600080fd5b506101596004803603602081101561015257600080fd5b503561040d565b6040518080602001868152602001856001600160a01b031681526020018415158152602001838152602001828103825287818151815260200191508051906020019080838360005b838110156101b95781810151838201526020016101a1565b50505050905090810190601f1680156101e65780820380516001836020036101000a031916815260200191505b50965050505050505060405180910390f35b34801561020457600080fd5b5061020d6104e6565b60408051918252519081900360200190f35b34801561022b57600080fd5b506100e66004803603606081101561024257600080fd5b81019060208101813564010000000081111561025d57600080fd5b82018360208201111561026f57600080fd5b8035906020019184600183028401116401000000008311171561029157600080fd5b91908080601f01602080910402602001604051908101604052809392919081815260200183838082843760009201919091525092955050823593505050602001356001600160a01b03166104ec565b3480156102ec57600080fd5b5061020d6105ea565b6100e66105f0565b34801561030957600080fd5b506100e66004803603602081101561032057600080fd5b5035610626565b6001546001600160a01b031681565b6001546001600160a01b0316331461034d57600080fd5b600080828154811061035b57fe5b9060005260206000209060050201905060026004548161037757fe5b0481600301541161038757600080fd5b6002810154600160a01b900460ff16156103a057600080fd5b600281015460018201546040516001600160a01b039092169181156108fc0291906000818181858888f193505050501580156103e0573d6000803e3d6000fd5b50600201805460ff60a01b1916600160a01b17905550565b60036020526000908152604090205460ff1681565b6000818154811061041a57fe5b60009182526020918290206005919091020180546040805160026001841615610100026000190190931692909204601f8101859004850283018501909152808252919350918391908301828280156104b35780601f10610488576101008083540402835291602001916104b3565b820191906000526020600020905b81548152906001019060200180831161049657829003601f168201915b5050505060018301546002840154600390940154929390926001600160a01b0382169250600160a01b90910460ff169085565b60045481565b6001546001600160a01b0316331461050357600080fd5b61050b6106af565b506040805160a08101825284815260208082018590526001600160a01b03841692820192909252600060608201819052608082018190528054600181018255908052815180519293849360059093027f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563019261058a92849201906106e9565b5060208201516001820155604082015160028201805460608501511515600160a01b0260ff60a01b196001600160a01b039094166001600160a01b0319909216919091179290921691909117905560809091015160039091015550505050565b60025481565b60025434116105fe57600080fd5b336000908152600360205260409020805460ff19166001908117909155600480549091019055565b600080828154811061063457fe5b6000918252602080832033845260039091526040909220546005909102909101915060ff1661066257600080fd5b33600090815260048201602052604090205460ff161561068157600080fd5b3360009081526004820160205260409020805460ff1916600190811790915560039091018054909101905550565b6040518060a00160405280606081526020016000815260200160006001600160a01b03168152602001600015158152602001600081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f1061072a57805160ff1916838001178555610757565b82800160010185558215610757579182015b8281111561075757825182559160200191906001019061073c565b50610763929150610767565b5090565b5b80821115610763576000815560010161076856fea2646970667358221220a0f818c99c8abf00022edbd0d78211606dcc860c82048c5ede33410e16a797f164736f6c634300060c0033a2646970667358221220c6b10c6e70f09074ed6cf71059479cb17a3ba57f990efb054037bc4d88e13ea264736f6c634300060c0033",
}

// CampaignFactoryABI is the input ABI used to generate the binding from.
// Deprecated: Use CampaignFactoryMetaData.ABI instead.
var CampaignFactoryABI = CampaignFactoryMetaData.ABI

// CampaignFactoryBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use CampaignFactoryMetaData.Bin instead.
var CampaignFactoryBin = CampaignFactoryMetaData.Bin

// DeployCampaignFactory deploys a new Ethereum contract, binding an instance of CampaignFactory to it.
func DeployCampaignFactory(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *CampaignFactory, error) {
	parsed, err := CampaignFactoryMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(CampaignFactoryBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CampaignFactory{CampaignFactoryCaller: CampaignFactoryCaller{contract: contract}, CampaignFactoryTransactor: CampaignFactoryTransactor{contract: contract}, CampaignFactoryFilterer: CampaignFactoryFilterer{contract: contract}}, nil
}

// CampaignFactory is an auto generated Go binding around an Ethereum contract.
type CampaignFactory struct {
	CampaignFactoryCaller     // Read-only binding to the contract
	CampaignFactoryTransactor // Write-only binding to the contract
	CampaignFactoryFilterer   // Log filterer for contract events
}

// CampaignFactoryCaller is an auto generated read-only Go binding around an Ethereum contract.
type CampaignFactoryCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CampaignFactoryTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CampaignFactoryTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CampaignFactoryFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CampaignFactoryFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CampaignFactorySession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CampaignFactorySession struct {
	Contract     *CampaignFactory  // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CampaignFactoryCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CampaignFactoryCallerSession struct {
	Contract *CampaignFactoryCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts          // Call options to use throughout this session
}

// CampaignFactoryTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CampaignFactoryTransactorSession struct {
	Contract     *CampaignFactoryTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts          // Transaction auth options to use throughout this session
}

// CampaignFactoryRaw is an auto generated low-level Go binding around an Ethereum contract.
type CampaignFactoryRaw struct {
	Contract *CampaignFactory // Generic contract binding to access the raw methods on
}

// CampaignFactoryCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CampaignFactoryCallerRaw struct {
	Contract *CampaignFactoryCaller // Generic read-only contract binding to access the raw methods on
}

// CampaignFactoryTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type CampaignFactoryTransactorRaw struct {
	Contract *CampaignFactoryTransactor // Generic write-only contract binding to access the raw methods on
}

// NewCampaignFactory creates a new instance of CampaignFactory, bound to a specific deployed contract.
func NewCampaignFactory(address common.Address, backend bind.ContractBackend) (*CampaignFactory, error) {
	contract, err := bindCampaignFactory(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &CampaignFactory{CampaignFactoryCaller: CampaignFactoryCaller{contract: contract}, CampaignFactoryTransactor: CampaignFactoryTransactor{contract: contract}, CampaignFactoryFilterer: CampaignFactoryFilterer{contract: contract}}, nil
}

// NewCampaignFactoryCaller creates a new read-only instance of CampaignFactory, bound to a specific deployed contract.
func NewCampaignFactoryCaller(address common.Address, caller bind.ContractCaller) (*CampaignFactoryCaller, error) {
	contract, err := bindCampaignFactory(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &CampaignFactoryCaller{contract: contract}, nil
}

// NewCampaignFactoryTransactor creates a new write-only instance of CampaignFactory, bound to a specific deployed contract.
func NewCampaignFactoryTransactor(address common.Address, transactor bind.ContractTransactor) (*CampaignFactoryTransactor, error) {
	contract, err := bindCampaignFactory(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &CampaignFactoryTransactor{contract: contract}, nil
}

// NewCampaignFactoryFilterer creates a new log filterer instance of CampaignFactory, bound to a specific deployed contract.
func NewCampaignFactoryFilterer(address common.Address, filterer bind.ContractFilterer) (*CampaignFactoryFilterer, error) {
	contract, err := bindCampaignFactory(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &CampaignFactoryFilterer{contract: contract}, nil
}

// bindCampaignFactory binds a generic wrapper to an already deployed contract.
func bindCampaignFactory(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(CampaignFactoryABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CampaignFactory *CampaignFactoryRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CampaignFactory.Contract.CampaignFactoryCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CampaignFactory *CampaignFactoryRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CampaignFactory.Contract.CampaignFactoryTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CampaignFactory *CampaignFactoryRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CampaignFactory.Contract.CampaignFactoryTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_CampaignFactory *CampaignFactoryCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _CampaignFactory.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_CampaignFactory *CampaignFactoryTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _CampaignFactory.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_CampaignFactory *CampaignFactoryTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _CampaignFactory.Contract.contract.Transact(opts, method, params...)
}

// DeployedCampaigns is a free data retrieval call binding the contract method 0x339d50a5.
//
// Solidity: function deployedCampaigns(uint256 ) view returns(address)
func (_CampaignFactory *CampaignFactoryCaller) DeployedCampaigns(opts *bind.CallOpts, arg0 *big.Int) (common.Address, error) {
	var out []interface{}
	err := _CampaignFactory.contract.Call(opts, &out, "deployedCampaigns", arg0)

	if err != nil {
		return *new(common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new(common.Address)).(*common.Address)

	return out0, err

}

// DeployedCampaigns is a free data retrieval call binding the contract method 0x339d50a5.
//
// Solidity: function deployedCampaigns(uint256 ) view returns(address)
func (_CampaignFactory *CampaignFactorySession) DeployedCampaigns(arg0 *big.Int) (common.Address, error) {
	return _CampaignFactory.Contract.DeployedCampaigns(&_CampaignFactory.CallOpts, arg0)
}

// DeployedCampaigns is a free data retrieval call binding the contract method 0x339d50a5.
//
// Solidity: function deployedCampaigns(uint256 ) view returns(address)
func (_CampaignFactory *CampaignFactoryCallerSession) DeployedCampaigns(arg0 *big.Int) (common.Address, error) {
	return _CampaignFactory.Contract.DeployedCampaigns(&_CampaignFactory.CallOpts, arg0)
}

// GetDeployedCampaigns is a free data retrieval call binding the contract method 0x4acb9d4f.
//
// Solidity: function getDeployedCampaigns() view returns(address[])
func (_CampaignFactory *CampaignFactoryCaller) GetDeployedCampaigns(opts *bind.CallOpts) ([]common.Address, error) {
	var out []interface{}
	err := _CampaignFactory.contract.Call(opts, &out, "getDeployedCampaigns")

	if err != nil {
		return *new([]common.Address), err
	}

	out0 := *abi.ConvertType(out[0], new([]common.Address)).(*[]common.Address)

	return out0, err

}

// GetDeployedCampaigns is a free data retrieval call binding the contract method 0x4acb9d4f.
//
// Solidity: function getDeployedCampaigns() view returns(address[])
func (_CampaignFactory *CampaignFactorySession) GetDeployedCampaigns() ([]common.Address, error) {
	return _CampaignFactory.Contract.GetDeployedCampaigns(&_CampaignFactory.CallOpts)
}

// GetDeployedCampaigns is a free data retrieval call binding the contract method 0x4acb9d4f.
//
// Solidity: function getDeployedCampaigns() view returns(address[])
func (_CampaignFactory *CampaignFactoryCallerSession) GetDeployedCampaigns() ([]common.Address, error) {
	return _CampaignFactory.Contract.GetDeployedCampaigns(&_CampaignFactory.CallOpts)
}

// CreateCampaign is a paid mutator transaction binding the contract method 0xa3303a75.
//
// Solidity: function createCampaign(uint256 minimum) returns()
func (_CampaignFactory *CampaignFactoryTransactor) CreateCampaign(opts *bind.TransactOpts, minimum *big.Int) (*types.Transaction, error) {
	return _CampaignFactory.contract.Transact(opts, "createCampaign", minimum)
}

// CreateCampaign is a paid mutator transaction binding the contract method 0xa3303a75.
//
// Solidity: function createCampaign(uint256 minimum) returns()
func (_CampaignFactory *CampaignFactorySession) CreateCampaign(minimum *big.Int) (*types.Transaction, error) {
	return _CampaignFactory.Contract.CreateCampaign(&_CampaignFactory.TransactOpts, minimum)
}

// CreateCampaign is a paid mutator transaction binding the contract method 0xa3303a75.
//
// Solidity: function createCampaign(uint256 minimum) returns()
func (_CampaignFactory *CampaignFactoryTransactorSession) CreateCampaign(minimum *big.Int) (*types.Transaction, error) {
	return _CampaignFactory.Contract.CreateCampaign(&_CampaignFactory.TransactOpts, minimum)
}
