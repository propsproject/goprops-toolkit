package propstoken

import (
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/propsproject/go-utils/propstoken/tokengen"
)

//DelegatedTransfer ...
type DelegatedTransfer struct {
	Signature string `json:"signature" gencodec:"required"`
	To        string `json:"to" gencodec:"required"`
	Amount    int    `json:"amount" gencodec:"required"`
	Fee       int    `json:"fee" gencodec:"required"`
	Nonce     string `json:"nonce"    gencodec:"required"`
	From      string `json:"from" gencodec:"required"`
}

// NewPropsTokenClient init global token client var for use
func NewPropsTokenClient(contractAddr, ipcPath string) (*tokengen.PropsToken, error) {
	conn, err := ethclient.Dial(ipcPath)
	if err != nil {
		return nil, err
	}

	propsTokenClient, err := tokengen.NewPropsToken(common.HexToAddress(contractAddr), conn)
	if err != nil {
		return nil, err
	}

	_, err = propsTokenClient.Name(nil)
	if err != nil {
		return nil, err
	}

	return propsTokenClient, nil
}
