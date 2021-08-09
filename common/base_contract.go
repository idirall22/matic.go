package common

import (
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/idirall22/matic.go/network"
)

const MaticChildTokenAddress string = "0x0000000000000000000000000000000000001010"

type BaseContract struct {
	network *network.Network
	client  *Client
}

func NewBaseContract(Network *network.Network, client *Client) *BaseContract {
	return &BaseContract{
		network: Network,
		client:  client,
	}
}

func (c *BaseContract) GetERC20TokenContract(token common.Address, parent bool) (*bind.BoundContract, error) {
	return GetContract(c.client.GetBackend(parent), c.network, token, "ChildERC20", "plasma")
}

func (c *BaseContract) GetERC721TokenContract(token common.Address, parent bool) (*bind.BoundContract, error) {
	return GetContract(c.client.GetBackend(parent), c.network, token, "ChildERC721", "plasma")
}

func (c *BaseContract) GetChildMaticContract() (*bind.BoundContract, error) {
	return GetContract(c.client.GetBackend(false), c.network, common.HexToAddress(MaticChildTokenAddress), "MRC20", "plasma")
}

func (c *BaseContract) GetPOSERC20TokenContract(token common.Address, parent bool) (*bind.BoundContract, error) {
	return GetContract(c.client.GetBackend(parent), c.network, token, "ChildERC20", "pos")
}

func (c *BaseContract) GetPOSERC721TokenContract(token common.Address, parent bool) (*bind.BoundContract, error) {
	return GetContract(c.client.GetBackend(parent), c.network, token, "ChildERC721", "pos")
}

func (c *BaseContract) GetPOSERC1155TokenContract(token common.Address, parent bool) (*bind.BoundContract, error) {
	return GetContract(c.client.GetBackend(parent), c.network, token, "ChildERC1155", "pos")
}
