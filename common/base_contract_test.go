package common_test

import (
	"context"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	c "github.com/idirall22/matic.go/common"
	"github.com/idirall22/matic.go/network"
	"github.com/stretchr/testify/require"
)

func TestBaseContract(t *testing.T) {
	from := common.HexToAddress("0x5a6eade27773eab2e4038ec7cf388781eca02f63")
	tokenAddressTest := common.HexToAddress("0x7d1afa7b718fb893db30a3abc0cfc608aacfebb0")
	network := network.NewNetwork("testnet", "mumbai")

	opts := c.InitializationOptions{
		Network:           network,
		ParentProviderURL: parentURL,
		MaticProviderURL:  childURL,
		ParentDefaultOptions: c.Options{
			Caller: &bind.CallOpts{
				From: from,
			},
			Transactor: &bind.TransactOpts{
				From: from,
			},
			Filterer: &bind.FilterOpts{},
		},
		MaticDefaultOptions: c.Options{},
		PrivateKey:          nil,
	}
	client, err := c.NewClient(context.Background(), opts)
	require.NoError(t, err)
	require.NotNil(t, client)

	baseContract := c.NewBaseContract(network, client)

	contract, err := baseContract.GetERC20TokenContract(tokenAddressTest, true)
	require.NoError(t, err)
	require.NotNil(t, contract)

	contract, err = baseContract.GetERC721TokenContract(tokenAddressTest, true)
	require.NoError(t, err)
	require.NotNil(t, contract)

	contract, err = baseContract.GetChildMaticContract()
	require.NoError(t, err)
	require.NotNil(t, contract)

	contract, err = baseContract.GetPOSERC20TokenContract(tokenAddressTest, true)
	require.NoError(t, err)
	require.NotNil(t, contract)

	contract, err = baseContract.GetPOSERC721TokenContract(tokenAddressTest, true)
	require.NoError(t, err)
	require.NotNil(t, contract)

	contract, err = baseContract.GetPOSERC1155TokenContract(tokenAddressTest, true)
	require.NoError(t, err)
	require.NotNil(t, contract)
}
