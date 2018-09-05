package rinkeby

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
)

// PropsClient ...
type Client struct {
	Token   *PropsTokenRinkeby
	RPC     *ethclient.Client
	Address string
}

// NewPropsTokenClient init global token client var for use
func NewPropsTokenClient(contractAddr, rpcAddr, origin string) (*Client, error) {
	rawRPC, err := rpc.DialWebsocket(context.Background(), rpcAddr, origin)
	if err != nil {
		return nil, err
	}

	rpcClient := ethclient.NewClient(rawRPC)

	address := common.HexToAddress(contractAddr)
	propsTokenClient, err := NewPropsTokenRinkeby(address, rpcClient)
	if err != nil {
		return nil, err
	}

	return &Client{Token: propsTokenClient, RPC: rpcClient, Address: address.String()}, nil
}
