// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package state

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

// StateMetaData contains all meta data concerning the State contract.
var StateMetaData = &bind.MetaData{
	ABI: "[{\"inputs\":[{\"internalType\":\"uint256\",\"name\":\"\",\"type\":\"uint256\"}],\"name\":\"dhtList\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"getDhtList\",\"outputs\":[{\"internalType\":\"string[]\",\"name\":\"\",\"type\":\"string[]\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"name\",\"type\":\"string\"}],\"name\":\"getUserAddress\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"peer_id\",\"type\":\"string\"}],\"name\":\"getUserName\",\"outputs\":[{\"internalType\":\"string\",\"name\":\"\",\"type\":\"string\"}],\"stateMutability\":\"view\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"addr\",\"type\":\"string\"}],\"name\":\"modDhtList\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[{\"internalType\":\"string\",\"name\":\"identity\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"peer_id\",\"type\":\"string\"},{\"internalType\":\"string\",\"name\":\"addr\",\"type\":\"string\"}],\"name\":\"register\",\"outputs\":[],\"stateMutability\":\"payable\",\"type\":\"function\"},{\"inputs\":[],\"name\":\"verification\",\"outputs\":[{\"internalType\":\"bool\",\"name\":\"\",\"type\":\"bool\"}],\"stateMutability\":\"view\",\"type\":\"function\"}]",
	Bin: "0x60806040524260015534801561001457600080fd5b50610e75806100246000396000f3fe6080604052600436106100705760003560e01c80634cd08d031161004e5780634cd08d03146100ed5780634ffe2a8b146101025780637228c6e014610127578063cf1c23c11461014757600080fd5b806307f2cc01146100755780631a64b8e8146100a05780634985e85c146100cd575b600080fd5b34801561008157600080fd5b5061008a61015a565b6040516100979190610954565b60405180910390f35b3480156100ac57600080fd5b506100c06100bb3660046109b6565b610233565b60405161009791906109cf565b3480156100d957600080fd5b506100c06100e83660046109ff565b6102df565b6101006100fb366004610af9565b61038f565b005b34801561010e57600080fd5b5061011761068a565b6040519015158152602001610097565b34801561013357600080fd5b506100c06101423660046109ff565b610713565b610100610155366004610b93565b610725565b60606000805480602002602001604051908101604052809291908181526020016000905b8282101561022a57838290600052602060002001805461019d90610bd5565b80601f01602080910402602001604051908101604052809291908181526020018280546101c990610bd5565b80156102165780601f106101eb57610100808354040283529160200191610216565b820191906000526020600020905b8154815290600101906020018083116101f957829003601f168201915b50505050508152602001906001019061017e565b50505050905090565b6000818154811061024357600080fd5b90600052602060002001600091509050805461025e90610bd5565b80601f016020809104026020016040519081016040528092919081815260200182805461028a90610bd5565b80156102d75780601f106102ac576101008083540402835291602001916102d7565b820191906000526020600020905b8154815290600101906020018083116102ba57829003601f168201915b505050505081565b60606004826040516102f19190610c0f565b9081526020016040518091039020805461030a90610bd5565b80601f016020809104026020016040519081016040528092919081815260200182805461033690610bd5565b80156103835780601f1061035857610100808354040283529160200191610383565b820191906000526020600020905b81548152906001019060200180831161036657829003601f168201915b50505050509050919050565b61039761068a565b6104e15760026040518060800160405280336001600160a01b0316815260200188888080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505090825250604080516020601f890181900481028201810190925287815291810191908890889081908401838280828437600092019190915250505090825250604080516020601f8701819004810282018101909252858152918101919086908690819084018382808284376000920182905250939094525050835460018082018655948252602091829020845160049092020180546001600160a01b0319166001600160a01b03909216919091178155908301519293909290830191506104b39082610c71565b50604082015160028201906104c89082610c71565b50606082015160038201906104dd9082610c71565b5050505b8585600386866040516104f5929190610d31565b90815260200160405180910390209182610510929190610d41565b50818160048888604051610525929190610d31565b90815260200160405180910390209182610540929190610d41565b506040518060800160405280336001600160a01b0316815260200187878080601f016020809104026020016040519081016040528093929190818152602001838380828437600092019190915250505090825250604080516020601f880181900481028201810190925286815291810191908790879081908401838280828437600092019190915250505090825250604080516020601f8601819004810282018101909252848152918101919085908590819084018382808284376000920182905250939094525050338152600560209081526040909120835181546001600160a01b0319166001600160a01b0390911617815590830151909150600182019061064a9082610c71565b506040820151600282019061065f9082610c71565b50606082015160038201906106749082610c71565b509050506106828282610725565b505050505050565b6040516bffffffffffffffffffffffff193360601b16602082015260009060340160408051601f19818403018152828252805160209182012033600090815260058352929092205460601b6bffffffffffffffffffffffff19169083015290603401604051602081830303815290604052805190602001200361070d5750600190565b50600090565b60606003826040516102f19190610c0f565b6000600154426107359190610e02565b9050603c8111156107c557604080516000808252602082019092529061076b565b60608152602001906001900390816107565790505b5080516107809160009160209091019061083f565b5042600190815560008054918201815580527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563016107bf838583610d41565b50505050565b600054600210156108005760008054806107e1576107e1610e29565b6001900381819060005260206000200160006107fd9190610895565b90555b600080546001810182559080527f290decd9548b62a8d60345a988386fc84ba6bc95484008f6362f93160ef3e563016107bf838583610d41565b505050565b828054828255906000526020600020908101928215610885579160200282015b8281111561088557825182906108759082610c71565b509160200191906001019061085f565b506108919291506108d2565b5090565b5080546108a190610bd5565b6000825580601f106108b1575050565b601f0160209004906000526020600020908101906108cf91906108ef565b50565b808211156108915760006108e68282610895565b506001016108d2565b5b8082111561089157600081556001016108f0565b60005b8381101561091f578181015183820152602001610907565b50506000910152565b60008151808452610940816020860160208601610904565b601f01601f19169290920160200192915050565b6000602080830181845280855180835260408601915060408160051b870101925083870160005b828110156109a957603f19888603018452610997858351610928565b9450928501929085019060010161097b565b5092979650505050505050565b6000602082840312156109c857600080fd5b5035919050565b6020815260006109e26020830184610928565b9392505050565b634e487b7160e01b600052604160045260246000fd5b600060208284031215610a1157600080fd5b813567ffffffffffffffff80821115610a2957600080fd5b818401915084601f830112610a3d57600080fd5b813581811115610a4f57610a4f6109e9565b604051601f8201601f19908116603f01168101908382118183101715610a7757610a776109e9565b81604052828152876020848701011115610a9057600080fd5b826020860160208301376000928101602001929092525095945050505050565b60008083601f840112610ac257600080fd5b50813567ffffffffffffffff811115610ada57600080fd5b602083019150836020828501011115610af257600080fd5b9250929050565b60008060008060008060608789031215610b1257600080fd5b863567ffffffffffffffff80821115610b2a57600080fd5b610b368a838b01610ab0565b90985096506020890135915080821115610b4f57600080fd5b610b5b8a838b01610ab0565b90965094506040890135915080821115610b7457600080fd5b50610b8189828a01610ab0565b979a9699509497509295939492505050565b60008060208385031215610ba657600080fd5b823567ffffffffffffffff811115610bbd57600080fd5b610bc985828601610ab0565b90969095509350505050565b600181811c90821680610be957607f821691505b602082108103610c0957634e487b7160e01b600052602260045260246000fd5b50919050565b60008251610c21818460208701610904565b9190910192915050565b601f82111561083a57600081815260208120601f850160051c81016020861015610c525750805b601f850160051c820191505b8181101561068257828155600101610c5e565b815167ffffffffffffffff811115610c8b57610c8b6109e9565b610c9f81610c998454610bd5565b84610c2b565b602080601f831160018114610cd45760008415610cbc5750858301515b600019600386901b1c1916600185901b178555610682565b600085815260208120601f198616915b82811015610d0357888601518255948401946001909101908401610ce4565b5085821015610d215787850151600019600388901b60f8161c191681555b5050505050600190811b01905550565b8183823760009101908152919050565b67ffffffffffffffff831115610d5957610d596109e9565b610d6d83610d678354610bd5565b83610c2b565b6000601f841160018114610da15760008515610d895750838201355b600019600387901b1c1916600186901b178355610dfb565b600083815260209020601f19861690835b82811015610dd25786850135825560209485019460019092019101610db2565b5086821015610def5760001960f88860031b161c19848701351681555b505060018560011b0183555b5050505050565b81810381811115610e2357634e487b7160e01b600052601160045260246000fd5b92915050565b634e487b7160e01b600052603160045260246000fdfea264697066735822122040623768a54edd7f9218aa77447045e3a6088a6d49ae9e820dce2a312cb3383364736f6c63430008110033",
}

// StateABI is the input ABI used to generate the binding from.
// Deprecated: Use StateMetaData.ABI instead.
var StateABI = StateMetaData.ABI

// StateBin is the compiled bytecode used for deploying new contracts.
// Deprecated: Use StateMetaData.Bin instead.
var StateBin = StateMetaData.Bin

// DeployState deploys a new Ethereum contract, binding an instance of State to it.
func DeployState(auth *bind.TransactOpts, backend bind.ContractBackend) (common.Address, *types.Transaction, *State, error) {
	parsed, err := StateMetaData.GetAbi()
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	if parsed == nil {
		return common.Address{}, nil, nil, errors.New("GetABI returned nil")
	}

	address, tx, contract, err := bind.DeployContract(auth, *parsed, common.FromHex(StateBin), backend)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &State{StateCaller: StateCaller{contract: contract}, StateTransactor: StateTransactor{contract: contract}, StateFilterer: StateFilterer{contract: contract}}, nil
}

// State is an auto generated Go binding around an Ethereum contract.
type State struct {
	StateCaller     // Read-only binding to the contract
	StateTransactor // Write-only binding to the contract
	StateFilterer   // Log filterer for contract events
}

// StateCaller is an auto generated read-only Go binding around an Ethereum contract.
type StateCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateTransactor is an auto generated write-only Go binding around an Ethereum contract.
type StateTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type StateFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// StateSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type StateSession struct {
	Contract     *State            // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StateCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type StateCallerSession struct {
	Contract *StateCaller  // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts // Call options to use throughout this session
}

// StateTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type StateTransactorSession struct {
	Contract     *StateTransactor  // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// StateRaw is an auto generated low-level Go binding around an Ethereum contract.
type StateRaw struct {
	Contract *State // Generic contract binding to access the raw methods on
}

// StateCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type StateCallerRaw struct {
	Contract *StateCaller // Generic read-only contract binding to access the raw methods on
}

// StateTransactorRaw is an auto generated low-level write-only Go binding around an Ethereum contract.
type StateTransactorRaw struct {
	Contract *StateTransactor // Generic write-only contract binding to access the raw methods on
}

// NewState creates a new instance of State, bound to a specific deployed contract.
func NewState(address common.Address, backend bind.ContractBackend) (*State, error) {
	contract, err := bindState(address, backend, backend, backend)
	if err != nil {
		return nil, err
	}
	return &State{StateCaller: StateCaller{contract: contract}, StateTransactor: StateTransactor{contract: contract}, StateFilterer: StateFilterer{contract: contract}}, nil
}

// NewStateCaller creates a new read-only instance of State, bound to a specific deployed contract.
func NewStateCaller(address common.Address, caller bind.ContractCaller) (*StateCaller, error) {
	contract, err := bindState(address, caller, nil, nil)
	if err != nil {
		return nil, err
	}
	return &StateCaller{contract: contract}, nil
}

// NewStateTransactor creates a new write-only instance of State, bound to a specific deployed contract.
func NewStateTransactor(address common.Address, transactor bind.ContractTransactor) (*StateTransactor, error) {
	contract, err := bindState(address, nil, transactor, nil)
	if err != nil {
		return nil, err
	}
	return &StateTransactor{contract: contract}, nil
}

// NewStateFilterer creates a new log filterer instance of State, bound to a specific deployed contract.
func NewStateFilterer(address common.Address, filterer bind.ContractFilterer) (*StateFilterer, error) {
	contract, err := bindState(address, nil, nil, filterer)
	if err != nil {
		return nil, err
	}
	return &StateFilterer{contract: contract}, nil
}

// bindState binds a generic wrapper to an already deployed contract.
func bindState(address common.Address, caller bind.ContractCaller, transactor bind.ContractTransactor, filterer bind.ContractFilterer) (*bind.BoundContract, error) {
	parsed, err := abi.JSON(strings.NewReader(StateABI))
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(address, parsed, caller, transactor, filterer), nil
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_State *StateRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _State.Contract.StateCaller.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_State *StateRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _State.Contract.StateTransactor.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_State *StateRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _State.Contract.StateTransactor.contract.Transact(opts, method, params...)
}

// Call invokes the (constant) contract method with params as input values and
// sets the output to result. The result type might be a single field for simple
// returns, a slice of interfaces for anonymous returns and a struct for named
// returns.
func (_State *StateCallerRaw) Call(opts *bind.CallOpts, result *[]interface{}, method string, params ...interface{}) error {
	return _State.Contract.contract.Call(opts, result, method, params...)
}

// Transfer initiates a plain transaction to move funds to the contract, calling
// its default method if one is available.
func (_State *StateTransactorRaw) Transfer(opts *bind.TransactOpts) (*types.Transaction, error) {
	return _State.Contract.contract.Transfer(opts)
}

// Transact invokes the (paid) contract method with params as input values.
func (_State *StateTransactorRaw) Transact(opts *bind.TransactOpts, method string, params ...interface{}) (*types.Transaction, error) {
	return _State.Contract.contract.Transact(opts, method, params...)
}

// DhtList is a free data retrieval call binding the contract method 0x1a64b8e8.
//
// Solidity: function dhtList(uint256 ) view returns(string)
func (_State *StateCaller) DhtList(opts *bind.CallOpts, arg0 *big.Int) (string, error) {
	var out []interface{}
	err := _State.contract.Call(opts, &out, "dhtList", arg0)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// DhtList is a free data retrieval call binding the contract method 0x1a64b8e8.
//
// Solidity: function dhtList(uint256 ) view returns(string)
func (_State *StateSession) DhtList(arg0 *big.Int) (string, error) {
	return _State.Contract.DhtList(&_State.CallOpts, arg0)
}

// DhtList is a free data retrieval call binding the contract method 0x1a64b8e8.
//
// Solidity: function dhtList(uint256 ) view returns(string)
func (_State *StateCallerSession) DhtList(arg0 *big.Int) (string, error) {
	return _State.Contract.DhtList(&_State.CallOpts, arg0)
}

// GetDhtList is a free data retrieval call binding the contract method 0x07f2cc01.
//
// Solidity: function getDhtList() view returns(string[])
func (_State *StateCaller) GetDhtList(opts *bind.CallOpts) ([]string, error) {
	var out []interface{}
	err := _State.contract.Call(opts, &out, "getDhtList")

	if err != nil {
		return *new([]string), err
	}

	out0 := *abi.ConvertType(out[0], new([]string)).(*[]string)

	return out0, err

}

// GetDhtList is a free data retrieval call binding the contract method 0x07f2cc01.
//
// Solidity: function getDhtList() view returns(string[])
func (_State *StateSession) GetDhtList() ([]string, error) {
	return _State.Contract.GetDhtList(&_State.CallOpts)
}

// GetDhtList is a free data retrieval call binding the contract method 0x07f2cc01.
//
// Solidity: function getDhtList() view returns(string[])
func (_State *StateCallerSession) GetDhtList() ([]string, error) {
	return _State.Contract.GetDhtList(&_State.CallOpts)
}

// GetUserAddress is a free data retrieval call binding the contract method 0x4985e85c.
//
// Solidity: function getUserAddress(string name) view returns(string)
func (_State *StateCaller) GetUserAddress(opts *bind.CallOpts, name string) (string, error) {
	var out []interface{}
	err := _State.contract.Call(opts, &out, "getUserAddress", name)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetUserAddress is a free data retrieval call binding the contract method 0x4985e85c.
//
// Solidity: function getUserAddress(string name) view returns(string)
func (_State *StateSession) GetUserAddress(name string) (string, error) {
	return _State.Contract.GetUserAddress(&_State.CallOpts, name)
}

// GetUserAddress is a free data retrieval call binding the contract method 0x4985e85c.
//
// Solidity: function getUserAddress(string name) view returns(string)
func (_State *StateCallerSession) GetUserAddress(name string) (string, error) {
	return _State.Contract.GetUserAddress(&_State.CallOpts, name)
}

// GetUserName is a free data retrieval call binding the contract method 0x7228c6e0.
//
// Solidity: function getUserName(string peer_id) view returns(string)
func (_State *StateCaller) GetUserName(opts *bind.CallOpts, peer_id string) (string, error) {
	var out []interface{}
	err := _State.contract.Call(opts, &out, "getUserName", peer_id)

	if err != nil {
		return *new(string), err
	}

	out0 := *abi.ConvertType(out[0], new(string)).(*string)

	return out0, err

}

// GetUserName is a free data retrieval call binding the contract method 0x7228c6e0.
//
// Solidity: function getUserName(string peer_id) view returns(string)
func (_State *StateSession) GetUserName(peer_id string) (string, error) {
	return _State.Contract.GetUserName(&_State.CallOpts, peer_id)
}

// GetUserName is a free data retrieval call binding the contract method 0x7228c6e0.
//
// Solidity: function getUserName(string peer_id) view returns(string)
func (_State *StateCallerSession) GetUserName(peer_id string) (string, error) {
	return _State.Contract.GetUserName(&_State.CallOpts, peer_id)
}

// Verification is a free data retrieval call binding the contract method 0x4ffe2a8b.
//
// Solidity: function verification() view returns(bool)
func (_State *StateCaller) Verification(opts *bind.CallOpts) (bool, error) {
	var out []interface{}
	err := _State.contract.Call(opts, &out, "verification")

	if err != nil {
		return *new(bool), err
	}

	out0 := *abi.ConvertType(out[0], new(bool)).(*bool)

	return out0, err

}

// Verification is a free data retrieval call binding the contract method 0x4ffe2a8b.
//
// Solidity: function verification() view returns(bool)
func (_State *StateSession) Verification() (bool, error) {
	return _State.Contract.Verification(&_State.CallOpts)
}

// Verification is a free data retrieval call binding the contract method 0x4ffe2a8b.
//
// Solidity: function verification() view returns(bool)
func (_State *StateCallerSession) Verification() (bool, error) {
	return _State.Contract.Verification(&_State.CallOpts)
}

// ModDhtList is a paid mutator transaction binding the contract method 0xcf1c23c1.
//
// Solidity: function modDhtList(string addr) payable returns()
func (_State *StateTransactor) ModDhtList(opts *bind.TransactOpts, addr string) (*types.Transaction, error) {
	return _State.contract.Transact(opts, "modDhtList", addr)
}

// ModDhtList is a paid mutator transaction binding the contract method 0xcf1c23c1.
//
// Solidity: function modDhtList(string addr) payable returns()
func (_State *StateSession) ModDhtList(addr string) (*types.Transaction, error) {
	return _State.Contract.ModDhtList(&_State.TransactOpts, addr)
}

// ModDhtList is a paid mutator transaction binding the contract method 0xcf1c23c1.
//
// Solidity: function modDhtList(string addr) payable returns()
func (_State *StateTransactorSession) ModDhtList(addr string) (*types.Transaction, error) {
	return _State.Contract.ModDhtList(&_State.TransactOpts, addr)
}

// Register is a paid mutator transaction binding the contract method 0x4cd08d03.
//
// Solidity: function register(string identity, string peer_id, string addr) payable returns()
func (_State *StateTransactor) Register(opts *bind.TransactOpts, identity string, peer_id string, addr string) (*types.Transaction, error) {
	return _State.contract.Transact(opts, "register", identity, peer_id, addr)
}

// Register is a paid mutator transaction binding the contract method 0x4cd08d03.
//
// Solidity: function register(string identity, string peer_id, string addr) payable returns()
func (_State *StateSession) Register(identity string, peer_id string, addr string) (*types.Transaction, error) {
	return _State.Contract.Register(&_State.TransactOpts, identity, peer_id, addr)
}

// Register is a paid mutator transaction binding the contract method 0x4cd08d03.
//
// Solidity: function register(string identity, string peer_id, string addr) payable returns()
func (_State *StateTransactorSession) Register(identity string, peer_id string, addr string) (*types.Transaction, error) {
	return _State.Contract.Register(&_State.TransactOpts, identity, peer_id, addr)
}
