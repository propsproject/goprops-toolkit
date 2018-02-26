package propstoken

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/propsproject/go-utils/propstoken/tokengen"
)

//DelegatedTransfer ...
type DelegatedTransfer struct {
	Signature string  `json:"signature" gencodec:"required"`
	From      string  `json:"from" gencodec:"required"`
	To        string  `json:"to" gencodec:"required"`
	Amount    float64 `json:"amount" gencodec:"required"`
	Fee       float64 `json:"fee" gencodec:"required"`
	Nonce     int64   `json:"nonce"    gencodec:"required"`
}

// PropsClient ...
type PropsClient struct {
	Token     *tokengen.PropsToken
	RPCClient *ethclient.Client
}

// NewPropsTokenClient init global token client var for use
func NewPropsTokenClient(contractAddr, rpcAddr string) (*PropsClient, error) {
	rpcClient, err := rpc.DialHTTP(rpcAddr)
	if err != nil {
		return nil, err
	}

	conn := ethclient.NewClient(rpcClient)

	propsTokenClient, err := tokengen.NewPropsToken(common.HexToAddress(contractAddr), conn)
	if err != nil {
		return nil, err
	}

	_, err = propsTokenClient.Name(nil)
	if err != nil {
		return nil, err
	}

	return &PropsClient{propsTokenClient, conn}, nil
}
