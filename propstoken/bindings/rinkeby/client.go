package rinkeby

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

// PropsClient ...
type Client struct {
	Token   *PropsTokenRinkeby
	RPC     *ethclient.Client
	Address string
}

// NewPropsTokenClient init global token client var for use
func NewPropsTokenClient(contractAddr, rpcAddr string) (*Client, error) {
	rpcClient, err := ethclient.Dial(rpcAddr)
	if err != nil {
		return nil, err
	}
	address := common.HexToAddress(contractAddr)
	propsTokenClient, err := NewPropsTokenRinkeby(address, rpcClient)
	if err != nil {
		return nil, err
	}

	return &Client{Token: propsTokenClient, RPC: rpcClient, Address: address.String()}, nil
}
