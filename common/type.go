package common

import (
	"crypto/ecdsa"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/idirall22/matic.go/network"
)

type CallOption struct {
	Parent bool
	*bind.CallOpts
}

type TransactOption struct {
	Parent bool
	*bind.TransactOpts
}

type InitializationOptions struct {
	Network                     *network.Network
	ParentProviderURL           string
	MaticProviderURL            string
	ParentDefaultOptions        Options
	MaticDefaultOptions         Options
	Registry                    common.Address
	RootChain                   common.Address
	DepositManager              common.Address
	WithdrawManager             common.Address
	POSRootChainManager         common.Address
	POSERC20Predicate           common.Address
	POSERC721Predicate          common.Address
	POSERC1155Predicate         common.Address
	POSMintableERC1155Predicate common.Address
	RequestConcurrency          int
	PrivateKey                  *ecdsa.PrivateKey
}

type Options struct {
	Caller     *bind.CallOpts
	Transactor *bind.TransactOpts
	Filterer   *bind.FilterOpts
}
