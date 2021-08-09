package common_test

import (
	"context"
	"fmt"
	"testing"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	c "github.com/idirall22/matic.go/common"
	"github.com/idirall22/matic.go/network"
	"github.com/stretchr/testify/require"
)

func TestClientSDK(t *testing.T) {

	pk, err := crypto.HexToECDSA(privatekey)
	require.NoError(t, err)

	from := common.HexToAddress(address)
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
		PrivateKey:          pk,
	}

	SDKClient, err := c.NewSDKClient(context.Background(), opts)
	require.NoError(t, err)
	require.NotNil(t, SDKClient)

	b, err := SDKClient.BalanceOfERC20(
		from,
		common.HexToAddress("0x0000000000000000000000000000000000001010"),
		&c.CallOption{
			Parent: false,
		},
	)

	require.NoError(t, err)
	fmt.Println(b.String())
}
