package common

import (
	"context"
	"errors"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type SDKClient struct {
	client       *Client
	baseContract *BaseContract
}

func NewSDKClient(ctx context.Context, opts InitializationOptions) (*SDKClient, error) {
	client, err := NewClient(ctx, opts)
	if err != nil {
		return &SDKClient{}, err
	}

	baseContract := NewBaseContract(opts.Network, client)
	return &SDKClient{
		client:       client,
		baseContract: baseContract,
	}, nil
}

func (s *SDKClient) BalanceOfERC20(
	userAddress,
	token common.Address,
	opts *CallOption,
) (*big.Int, error) {
	if (token == common.Address{}) || (userAddress == common.Address{}) {
		return new(big.Int), errors.New("token address or user address is missing")
	}
	contract, err := s.baseContract.GetPOSERC20TokenContract(token, opts.Parent)
	if err != nil {
		return new(big.Int), err
	}
	var out []interface{}

	if opts == nil {
		opts.CallOpts = s.baseContract.client.childDefaultOptions.Caller
	}

	err = contract.Call(opts.CallOpts, &out, "balanceOf", userAddress)
	if err != nil {
		return new(big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new(big.Int)).(*big.Int)
	return &out0, err
}

func (s *SDKClient) BalanceOfERC721(
	userAddress,
	token common.Address,
	opts *CallOption,
) (*big.Int, error) {
	if (token == common.Address{}) || (userAddress == common.Address{}) {
		return new(big.Int), errors.New("token address or user address is missing")
	}
	contract, err := s.baseContract.GetERC721TokenContract(token, opts.Parent)
	if err != nil {
		return new(big.Int), err
	}
	var out []interface{}

	if opts == nil {
		opts.CallOpts = s.baseContract.client.childDefaultOptions.Caller
	}

	err = contract.Call(opts.CallOpts, &out, "balanceOf", userAddress)
	if err != nil {
		return new(big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new(big.Int)).(*big.Int)

	return &out0, err
}

func (s *SDKClient) TokenOfOwnerByIndexERC721(
	userAddress, token common.Address,
	index *big.Int,
	opts *CallOption,
) (*big.Int, error) {

	if (token == common.Address{}) || (userAddress == common.Address{}) {
		return new(big.Int), errors.New("token address or user address is missing")
	}

	contract, err := s.baseContract.GetERC721TokenContract(token, opts.Parent)
	if err != nil {
		return new(big.Int), err
	}
	var out []interface{}

	if opts == nil {
		opts.CallOpts = s.baseContract.client.childDefaultOptions.Caller
	}

	err = contract.Call(opts.CallOpts, &out, "tokenOfOwnerByIndex", userAddress, index)
	if err != nil {
		return new(big.Int), err
	}
	out0 := *abi.ConvertType(out[0], new(big.Int)).(*big.Int)

	return &out0, err
}

func (s *SDKClient) TransferERC20Tokens(
	token common.Address,
	amount *big.Int,
	to common.Address,
	opts *TransactOption,
) (*types.Transaction, error) {

	if (token == common.Address{}) || (to == common.Address{}) {
		return new(types.Transaction), errors.New("token address or user address is missing")
	}

	contract, err := s.baseContract.GetERC20TokenContract(token, opts.Parent)
	if err != nil {
		return new(types.Transaction), err
	}

	if opts == nil {
		opts.TransactOpts = s.baseContract.client.childDefaultOptions.Transactor
	}

	tx, err := contract.Transact(opts.TransactOpts, "transfer", to, amount)
	if err != nil {
		return new(types.Transaction), err
	}
	return tx, err
}

func (s *SDKClient) TransferERC721Tokens(
	token common.Address,
	to common.Address,
	opts *TransactOption,
) (*types.Transaction, error) {

	if (token == common.Address{}) || (to == common.Address{}) {
		return new(types.Transaction), errors.New("token address or user address is missing")
	}

	contract, err := s.baseContract.GetERC721TokenContract(token, opts.Parent)
	if err != nil {
		return new(types.Transaction), err
	}

	if opts == nil {
		opts.TransactOpts = s.baseContract.client.childDefaultOptions.Transactor
	}

	tx, err := contract.Transact(opts.TransactOpts, "transfer", to)
	if err != nil {
		return new(types.Transaction), err
	}
	return tx, err
}

func (s *SDKClient) TransferMaticEth(
	to common.Address,
	amount *big.Int,
	opts *TransactOption,
) (*types.Transaction, error) {

	if (to == common.Address{}) || (amount.Cmp(big.NewInt(0)) == 0) {
		return new(types.Transaction), errors.New("token address or user address is missing")
	}

	contract, err := s.baseContract.GetChildMaticContract()
	if err != nil {
		return new(types.Transaction), err
	}

	if opts == nil {
		opts.TransactOpts = s.baseContract.client.childDefaultOptions.Transactor
	}

	tx, err := contract.Transact(opts.TransactOpts, "transfer", to, amount)
	if err != nil {
		return new(types.Transaction), err
	}
	return tx, err
}
