package common

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/idirall22/matic.go/network"
)

func GetContract(
	backend bind.ContractBackend,
	network *network.Network,
	token common.Address,
	name, typ string,
) (*bind.BoundContract, error) {
	content, err := network.GetAbi(name, typ)
	if err != nil {
		return nil, err
	}
	return bind.NewBoundContract(token, content, backend, backend, backend), nil
}
