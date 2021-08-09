package common

import (
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type RootChain struct {
	client               *Client
	bigOne               *big.Int
	bigTwo               *big.Int
	checkPointIdInterval *big.Int
	baseContract         *BaseContract
	rootChain            *bind.BoundContract
	Options              InitializationOptions
}

func NewRootChain(client *Client, opts InitializationOptions) (*RootChain, error) {
	rootChain, err := GetContract(client.parent, opts.Network, opts.RootChain, "RootChain", "plasma")
	if err != nil {
		return nil, err
	}
	baseContract := NewBaseContract(opts.Network, client)
	return &RootChain{
		client:               client,
		baseContract:         baseContract,
		bigOne:               big.NewInt(1),
		bigTwo:               big.NewInt(2),
		checkPointIdInterval: big.NewInt(10000),
		rootChain:            rootChain,
	}, nil
}

func (r *RootChain) GetLastChildBlock() (*big.Int, error) {
	var out []interface{}
	err := r.rootChain.Call(r.Options.ParentDefaultOptions.Caller, &out, "getLastChildBlock")
	if err != nil {
		return new(big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new(big.Int)).(*big.Int)

	return &out0, err
}
