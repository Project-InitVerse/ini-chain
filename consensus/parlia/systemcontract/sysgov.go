package systemcontract

import (
	"PureChain/common"
	"PureChain/consensus/parlia/vmcaller"
	"PureChain/core"
	"PureChain/core/state"
	"PureChain/core/types"
	"PureChain/log"
	"PureChain/params"
	"math"
	"math/big"
)

var (
	govAdmin        = common.HexToAddress("0xce930537a2148b8dc43899ff2e9bcbee0e801c54")
	govAdminTestnet = common.HexToAddress("0xce930537a2148b8dc43899ff2e9bcbee0e801c54")
)

const (
	govCode = "0x608060405234801561001057600080fd5b50600436106101425760003560e01c8063741579b1116100b8578063e3377eb91161007c578063e3377eb914610361578063ec0cb3361461024d578063f3b1cc67146103f6578063f851a440146103fe578063fb48270c14610406578063fbb847e11461040e57610142565b8063741579b1146102eb5780639001eed8146102f3578063c4d66de8146102fb578063c967f90f14610321578063e08b1d381461034057610142565b8063267822471161010a57806326782247146102745780632e4f67e41461024d5780633656de211461029857806344f99900146102b55780634fb9e9b7146102bd57806371a1bb75146102e357610142565b806303fab4f61461014757806305b8481014610161578063158ef93e1461023157806315de360e1461024d578063232e5ffc14610255575b600080fd5b61014f610416565b60408051918252519081900360200190f35b6101846004803603602081101561017757600080fd5b503563ffffffff16610423565b60405180878152602001868152602001856001600160a01b03168152602001846001600160a01b0316815260200183815260200180602001828103825283818151815260200191508051906020019080838360005b838110156101f15781810151838201526020016101d9565b50505050905090810190601f16801561021e5780820380516001836020036101000a031916815260200191505b5097505050505050505060405180910390f35b6102396105bc565b604080519115158252519081900360200190f35b61014f6105c5565b6102726004803603602081101561026b57600080fd5b50356105cc565b005b61027c6107a4565b604080516001600160a01b039092168252519081900360200190f35b610184600480360360208110156102ae57600080fd5b50356107b3565b61027c61081f565b610272600480360360208110156102d357600080fd5b50356001600160a01b0316610825565b61027c6108c0565b61014f6108c6565b61014f6108d2565b6102726004803603602081101561031157600080fd5b50356001600160a01b03166108e0565b61032961095d565b6040805161ffff9092168252519081900360200190f35b610348610962565b6040805163ffffffff9092168252519081900360200190f35b610272600480360360a081101561037757600080fd5b8135916001600160a01b03602082013581169260408301359091169160608101359181019060a0810160808201356401000000008111156103b757600080fd5b8201836020820111156103c957600080fd5b803590602001918460018302840111640100000000831117156103eb57600080fd5b509092509050610968565b61014f610cf8565b61027c610cff565b610272610d13565b61014f610dcd565b68056bc75e2d6310000081565b600080600080600060606003805490508763ffffffff1610610481576040805162461bcd60e51b8152602060048201526012602482015271496e646578206f7574206f662072616e676560701b604482015290519081900360640190fd5b610489610dd3565b60038863ffffffff168154811061049c57fe5b60009182526020918290206040805160c08101825260069390930290910180548352600180820154848601526002808301546001600160a01b039081168686015260038401541660608601526004830154608086015260058301805485516101009482161594909402600019011691909104601f81018790048702830187019094528382529394919360a086019391929091908301828280156105805780601f1061055557610100808354040283529160200191610580565b820191906000526020600020905b81548152906001019060200180831161056357829003601f168201915b5050509190925250508151602083015160408401516060850151608086015160a090960151939e929d50909b5099509297509550909350505050565b60005460ff1681565b6201518081565b33411461060d576040805162461bcd60e51b815260206004820152600a6024820152694d696e6572206f6e6c7960b01b604482015290519081900360640190fd5b60005b6003548110156107a057816003828154811061062857fe5b9060005260206000209060060201600001541415610798576003546000190181146107055760038054600019810190811061065f57fe5b90600052602060002090600602016003828154811061067a57fe5b6000918252602090912082546006909202019081556001808301548183015560028084015481840180546001600160a01b039283166001600160a01b03199182161790915560038087015490860180549190931691161790556004808501549084015560058085018054610701949286019391926101009082161502600019011604610e1b565b5050505b600380548061071057fe5b600082815260208120600660001990930192830201818155600181018290556002810180546001600160a01b0319908116909155600382018054909116905560048101829055906107646005830182610ea0565b5050905560405182907fc2946e69de813a7cede502a3b315aa221abf9fcca5c7134b0ae6b2c3857cf63d90600090a26107a0565b600101610610565b5050565b6001546001600160a01b031681565b60008060008060006060600280549050871061080a576040805162461bcd60e51b8152602060048201526011602482015270125908191bd95cc81b9bdd08195e1a5cdd607a1b604482015290519081900360640190fd5b610812610dd3565b6002888154811061049c57fe5b61c00681565b60005461010090046001600160a01b03163314610876576040805162461bcd60e51b815260206004820152600a60248201526941646d696e206f6e6c7960b01b604482015290519081900360640190fd5b600180546001600160a01b0319166001600160a01b0383169081179091556040517faefcaa6215f99fe8c2f605dd268ee4d23a5b596bbca026e25ce8446187f4f1ba90600090a250565b61c00581565b670de0b6b3a764000081565b69010f0cf064dd5920000081565b60005460ff161561092e576040805162461bcd60e51b8152602060048201526013602482015272105b1c9958591e481a5b9a5d1a585b1a5e9959606a1b604482015290519081900360640190fd5b6000805460ff196001600160a01b0390931661010002610100600160a81b031990911617919091166001179055565b601581565b60035490565b60005461010090046001600160a01b031633146109b9576040805162461bcd60e51b815260206004820152600a60248201526941646d696e206f6e6c7960b01b604482015290519081900360640190fd5b6002546109c4610dd3565b6040518060c00160405280838152602001898152602001886001600160a01b03168152602001876001600160a01b0316815260200186815260200185858080601f0160208091040260200160405190810160405280939291908181526020018383808284376000920182905250939094525050600280546001810182559152825160069091027f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ace81019182556020808501517f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5acf83015560408501517f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ad0830180546001600160a01b039283166001600160a01b03199182161790915560608701517f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ad18501805491909316911617905560808501517f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ad283015560a085015180519596508695939450610b7b937f405787fa12a823e0f2b7631cc41b3ba8828b3321ca811111fa75cd3aa3bb5ad390930192910190610ee7565b505060038054600181018255600091909152825160069091027fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85b81019182556020808501517fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85c83015560408501517fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85d830180546001600160a01b039283166001600160a01b03199182161790915560608701517fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85e8501805491909316911617905560808501517fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f85f83015560a08501518051869550610cc0937fc2575a0e9e593c00f959f8c92f12db2869c3395a3b0502d05e2516446f71f86001929190910190610ee7565b50506040518391507f2f28cf6eab3be78ec5322050b7c7ce47adc6f2cf957c0a7b7c6d893fcec891d990600090a25050505050505050565b6206270081565b60005461010090046001600160a01b031681565b6001546001600160a01b03163314610d63576040805162461bcd60e51b815260206004820152600e60248201526d4e65772061646d696e206f6e6c7960901b604482015290519081900360640190fd5b60018054600080546001600160a01b03808416610100908102610100600160a81b0319909316929092178084556001600160a01b03199094169094556040519204909216917f7ce7ec0b50378fb6c0186ffb5f48325f6593fcb4ca4386f21861af3129188f5c91a2565b60025490565b6040518060c00160405280600081526020016000815260200160006001600160a01b0316815260200160006001600160a01b0316815260200160008152602001606081525090565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610e545780548555610e90565b82800160010185558215610e9057600052602060002091601f016020900482015b82811115610e90578254825591600101919060010190610e75565b50610e9c929150610f55565b5090565b50805460018160011615610100020316600290046000825580601f10610ec65750610ee4565b601f016020900490600052602060002090810190610ee49190610f55565b50565b828054600181600116156101000203166002900490600052602060002090601f016020900481019282601f10610f2857805160ff1916838001178555610e90565b82800160010185558215610e90579182015b82811115610e90578251825591602001919060010190610f3a565b5b80821115610e9c5760008155600101610f5656fea26469706673582212204394637396bd82424a3d120c757c8ea7de997303eaca9a8e579475b91e8e0d1564736f6c634300060c0033"
)

type hardForkSysGov struct {
}

func (s *hardForkSysGov) GetName() string {
	return SysGovContractName
}

func (s *hardForkSysGov) Update(config *params.ChainConfig, height *big.Int, state *state.StateDB) (err error) {
	contractCode := common.FromHex(govCode)

	//write govCode to sys contract
	state.SetCode(SysGovContractAddr, contractCode)
	log.Debug("Write code to system contract account", "addr", SysGovContractAddr.String(), "code", govCode)

	return
}

func (s *hardForkSysGov) getAdminByChainId(chainId *big.Int) common.Address {
	if chainId.Cmp(params.MainnetChainConfig.ChainID) == 0 {
		return govAdmin
	}

	return govAdminTestnet
}

func (s *hardForkSysGov) Execute(state *state.StateDB, header *types.Header, chainContext core.ChainContext, config *params.ChainConfig) (err error) {

	method := "initialize"
	data, err := GetInteractiveABI()[SysGovContractName].Pack(method, s.getAdminByChainId(config.ChainID))
	if err != nil {
		log.Error("Can't pack data for initialize", "error", err)
		return err
	}

	msg := types.NewMessage(header.Coinbase, &SysGovContractAddr, 0, new(big.Int), math.MaxUint64, new(big.Int), data, nil, false)
	vmcaller.ExecuteMsg(msg, state, header, chainContext, config)
	//context := core.NewEVMContext(msg, header, chainContext, nil)
	//evm := vm.NewEVM(context, state, config, vm.Config{})
	//
	//_, _, err = evm.Call(vm.AccountRef(msg.From()), *msg.To(), msg.Data(), msg.Gas(), msg.Value())

	return
}
