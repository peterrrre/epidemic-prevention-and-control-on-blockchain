// Code generated - DO NOT EDIT.
// This file is a generated binding and any manual changes will be lost.

package contract

import (
	"math/big"
	"strings"

	
	"github.com/accounts/abi"
	"github.com/accounts/abi/bind"
	
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

// CheckpointOracleABI is the input ABI used to generate the binding from.
const CheckpointOracleABI = "[{\"constant\":true,\"inputs\":[],\"name\":\"GetAllAdmin\",\"outputs\":[{\"name\":\"\",\"type\":\"address[]\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":true,\"inputs\":[],\"name\":\"GetLatestCheckpoint\",\"outputs\":[{\"name\":\"\",\"type\":\"uint64\"},{\"name\":\"\",\"type\":\"bytes32\"},{\"name\":\"\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"view\",\"type\":\"function\"},{\"constant\":false,\"inputs\":[{\"name\":\"_recentNumber\",\"type\":\"uint256\"},{\"name\":\"_recentHash\",\"type\":\"bytes32\"},{\"name\":\"_hash\",\"type\":\"bytes32\"},{\"name\":\"_sectionIndex\",\"type\":\"uint64\"},{\"name\":\"v\",\"type\":\"uint8[]\"},{\"name\":\"r\",\"type\":\"bytes32[]\"},{\"name\":\"s\",\"type\":\"bytes32[]\"}],\"name\":\"SetCheckpoint\",\"outputs\":[{\"name\":\"\",\"type\":\"bool\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"function\"},{\"inputs\":[{\"name\":\"_adminlist\",\"type\":\"address[]\"},{\"name\":\"_sectionSize\",\"type\":\"uint256\"},{\"name\":\"_processConfirms\",\"type\":\"uint256\"},{\"name\":\"_threshold\",\"type\":\"uint256\"}],\"payable\":false,\"stateMutability\":\"nonpayable\",\"type\":\"constructor\"},{\"anonymous\":false,\"inputs\":[{\"indexed\":true,\"name\":\"index\",\"type\":\"uint64\"},{\"indexed\":false,\"name\":\"checkpointHash\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"v\",\"type\":\"uint8\"},{\"indexed\":false,\"name\":\"r\",\"type\":\"bytes32\"},{\"indexed\":false,\"name\":\"s\",\"type\":\"bytes32\"}],\"name\":\"NewCheckpointVote\",\"type\":\"event\"}]"

// CheckpointOracleBin is the compiled bytecode used for deploying new contracts.
const CheckpointOracleBin = `0x608060405234801561001057600080fd5b506040516108153803806108158339818101604052608081101561003357600080fd5b81019080805164010000000081111561004b57600080fd5b8201602081018481111561005e57600080fd5b815185602082028301116401000000008211171561007b57600080fd5b505060208201516040830151606090930151919450925060005b84518110156101415760016000808784815181106100af57fe5b60200260200101516001600160a01b03166001600160a01b0316815260200190815260200160002060006101000a81548160ff02191690831515021790555060018582815181106100fc57fe5b60209081029190910181015182546001808201855560009485529290932090920180546001600160a01b0319166001600160a01b039093169290921790915501610095565b50600592909255600655600755506106b78061015e6000396000f3fe608060405234801561001057600080fd5b50600436106100415760003560e01c806345848dfc146100465780634d6a304c1461009e578063d459fc46146100cf575b600080fd5b61004e6102b0565b60408051602080825283518183015283519192839290830191858101910280838360005b8381101561008a578181015183820152602001610072565b505050509050019250505060405180910390f35b6100a661034f565b6040805167ffffffffffffffff9094168452602084019290925282820152519081900360600190f35b61029c600480360360e08110156100e557600080fd5b81359160208101359160408201359167ffffffffffffffff6060820135169181019060a08101608082013564010000000081111561012257600080fd5b82018360208201111561013457600080fd5b8035906020019184602083028401116401000000008311171561015657600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092959493602081019350359150506401000000008111156101a657600080fd5b8201836020820111156101b857600080fd5b803590602001918460208302840111640100000000831117156101da57600080fd5b919080806020026020016040519081016040528093929190818152602001838360200280828437600092019190915250929594936020810193503591505064010000000081111561022a57600080fd5b82018360208201111561023c57600080fd5b8035906020019184602083028401116401000000008311171561025e57600080fd5b91908080602002602001604051908101604052809392919081815260200183836020028082843760009201919091525092955061036a945050505050565b604080519115158252519081900360200190f35b6060806001805490506040519080825280602002602001820160405280156102e2578160200160208202803883390190505b50905060005b60015481101561034957600181815481106102ff57fe5b9060005260206000200160009054906101000a90046001600160a01b031682828151811061032957fe5b6001600160a01b03909216602092830291909101909101526001016102e8565b50905090565b60025460045460035467ffffffffffffffff90921691909192565b3360009081526020819052604081205460ff1661038657600080fd5b8688401461039357600080fd5b82518451146103a157600080fd5b81518451146103af57600080fd5b6006546005548660010167ffffffffffffffff1602014310156103d457506000610677565b60025467ffffffffffffffff90811690861610156103f457506000610677565b60025467ffffffffffffffff8681169116148015610426575067ffffffffffffffff8516151580610426575060035415155b1561043357506000610677565b8561044057506000610677565b60408051601960f81b6020808301919091526000602183018190523060601b60228401526001600160c01b031960c08a901b166036840152603e8084018b905284518085039091018152605e909301909352815191012090805b86518110156106715760006001848984815181106104b457fe5b60200260200101518985815181106104c857fe5b60200260200101518986815181106104dc57fe5b602002602001015160405160008152602001604052604051808581526020018460ff1660ff1681526020018381526020018281526020019450505050506020604051602081039080840390855afa15801561053b573d6000803e3d6000fd5b505060408051601f1901516001600160a01b03811660009081526020819052919091205490925060ff16905061057057600080fd5b826001600160a01b0316816001600160a01b03161161058e57600080fd5b8092508867ffffffffffffffff167fce51ffa16246bcaf0899f6504f473cd0114f430f566cef71ab7e03d3dde42a418b8a85815181106105ca57fe5b60200260200101518a86815181106105de57fe5b60200260200101518a87815181106105f257fe5b6020026020010151604051808581526020018460ff1660ff16815260200183815260200182815260200194505050505060405180910390a260075482600101106106685750505060048790555050436003556002805467ffffffffffffffff191667ffffffffffffffff86161790556001610677565b5060010161049a565b50600080fd5b97965050505050505056fea265627a7a723058207f6a191ce575596a2f1e907c8c0a01003d16b69fb2c4f432d10878e8c0a99a0264736f6c634300050a0032`

// DeployCheckpointOracle deploys a new Ethereum contract, binding an instance of CheckpointOracle to it.
func DeployCheckpointOracle(auth *bind.TransactOpts, backend bind.ContractBackend, _adminlist []common.Address, _sectionSize *big.Int, _processConfirms *big.Int, _threshold *big.Int) (common.Address, *types.Transaction, *CheckpointOracle, error) {
	parsed, err := abi.JSON(strings.NewReader(CheckpointOracleABI))
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	address, tx, contract, err := bind.DeployContract(auth, parsed, common.FromHex(CheckpointOracleBin), backend, _adminlist, _sectionSize, _processConfirms, _threshold)
	if err != nil {
		return common.Address{}, nil, nil, err
	}
	return address, tx, &CheckpointOracle{CheckpointOracleCaller: CheckpointOracleCaller{contract: contract}, CheckpointOracleTransactor: CheckpointOracleTransactor{contract: contract}, CheckpointOracleFilterer: CheckpointOracleFilterer{contract: contract}}, nil
}

// CheckpointOracle is an auto generated Go binding around an Ethereum contract.
type CheckpointOracle struct {
	CheckpointOracleCaller     // Read-only binding to the contract
	CheckpointOracleTransactor // Write-only binding to the contract
	CheckpointOracleFilterer   // Log filterer for contract events
}

// CheckpointOracleCaller is an auto generated read-only Go binding around an Ethereum contract.
type CheckpointOracleCaller struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CheckpointOracleTransactor is an auto generated write-only Go binding around an Ethereum contract.
type CheckpointOracleTransactor struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CheckpointOracleFilterer is an auto generated log filtering Go binding around an Ethereum contract events.
type CheckpointOracleFilterer struct {
	contract *bind.BoundContract // Generic contract wrapper for the low level calls
}

// CheckpointOracleSession is an auto generated Go binding around an Ethereum contract,
// with pre-set call and transact options.
type CheckpointOracleSession struct {
	Contract     *CheckpointOracle // Generic contract binding to set the session for
	CallOpts     bind.CallOpts     // Call options to use throughout this session
	TransactOpts bind.TransactOpts // Transaction auth options to use throughout this session
}

// CheckpointOracleCallerSession is an auto generated read-only Go binding around an Ethereum contract,
// with pre-set call options.
type CheckpointOracleCallerSession struct {
	Contract *CheckpointOracleCaller // Generic contract caller binding to set the session for
	CallOpts bind.CallOpts           // Call options to use throughout this session
}

// CheckpointOracleTransactorSession is an auto generated write-only Go binding around an Ethereum contract,
// with pre-set transact options.
type CheckpointOracleTransactorSession struct {
	Contract     *CheckpointOracleTransactor // Generic contract transactor binding to set the session for
	TransactOpts bind.TransactOpts           // Transaction auth options to use throughout this session
}

// CheckpointOracleRaw is an auto generated low-level Go binding around an Ethereum contract.
type CheckpointOracleRaw struct {
	Contract *CheckpointOracle // Generic contract binding to access the raw methods on
}

// CheckpointOracleCallerRaw is an auto generated low-level read-only Go binding around an Ethereum contract.
type CheckpointOracleCallerRaw struct {
	Contract *CheckpointOracleCaller // Generic read-only contract binding to access the raw methods on
}
