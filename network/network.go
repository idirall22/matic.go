package network

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi"
)

type Network struct {
	name    string
	version string
}

func NewNetwork(name, version string) *Network {
	return &Network{
		name:    name,
		version: version,
	}
}

func (n *Network) GetAbi(name, typ string) (abi.ABI, error) {
	path := fmt.Sprintf("../networks/%s/%s/artifacts/%s/%s.json", n.name, n.version, typ, name)
	if _, err := os.Stat(path); os.IsNotExist(err) {
		return abi.ABI{}, err
	}
	data, err := ioutil.ReadFile(path)
	if err != nil {
		return abi.ABI{}, err
	}

	m := make(map[string]interface{})
	err = json.Unmarshal(data, &m)
	if err != nil {
		return abi.ABI{}, err
	}

	data, err = json.Marshal(m["abi"])
	if err != nil {
		return abi.ABI{}, err
	}

	a, err := abi.JSON(strings.NewReader(string(data)))
	return a, err
}
