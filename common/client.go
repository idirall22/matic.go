package common

import (
	"context"
	"crypto/ecdsa"
	"errors"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/ethclient"
)

const EXTRAGASFORPROXYCALL = 1000000

type Client struct {
	parent               *ethclient.Client
	child                *ethclient.Client
	parentDefaultOptions Options
	childDefaultOptions  Options
	pk                   *ecdsa.PrivateKey
}

func NewClient(
	ctx context.Context,
	opts InitializationOptions,
) (*Client, error) {

	parent, err := ethclient.DialContext(ctx, opts.ParentProviderURL)
	if err != nil {
		return nil, errors.New("Could not connect to parent chain: " + err.Error())
	}

	child, err := ethclient.DialContext(ctx, opts.MaticProviderURL)
	if err != nil {
		return nil, errors.New("Could not connect to child chain: " + err.Error())
	}
	return &Client{
		parent:               parent,
		child:                child,
		parentDefaultOptions: opts.ParentDefaultOptions,
		childDefaultOptions:  opts.MaticDefaultOptions,
		pk:                   opts.PrivateKey,
	}, nil
}

func (c *Client) SetWallet(pk *ecdsa.PrivateKey) {
	c.pk = pk
}

func (c *Client) GetBackend(parent bool) bind.ContractBackend {
	if parent {
		return c.parent
	}
	return c.child
}

func (c *Client) GetParentWeb3() *ethclient.Client {
	return c.parent
}

func (c *Client) GetMaticWeb3() *ethclient.Client {
	return c.child
}

func (c *Client) SetParentDefaultOptions(parentDefaultOptions Options) {
	c.parentDefaultOptions = parentDefaultOptions
}

func (c *Client) SetMaticDefaultOptions(childDefaultOptions Options) {
	c.childDefaultOptions = childDefaultOptions
}

// 85252280932
