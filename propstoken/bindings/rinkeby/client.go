package rinkeby

import (
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/rpc"
)

// PropsClient ...
type Client struct {
	Token     *PropsTokenRinkeby
	RPC *ethclient.Client
}

// NewPropsTokenClient init global token client var for use
func NewPropsTokenClient(contractAddr, rpcAddr string) (*Client, error) {
	rpcClient, err := rpc.DialHTTP(rpcAddr)
	if err != nil {
		return nil, err
	}

	conn := ethclient.NewClient(rpcClient)
	propsTokenClient, err := NewPropsTokenRinkeby(common.HexToAddress(contractAddr), conn)
	if err != nil {
		return nil, err
	}

	return &Client{propsTokenClient, conn}, nil
}
