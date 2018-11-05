package rinkeby

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"path/filepath"
	"io/ioutil"
	"fmt"
	"strings"
)

const AbiFilePath = "./resources/props-token-rinkeby.abi"

// PropsClient ...
type Client struct {
	Token   *PropsTokenRinkeby
	RPC     *ethclient.Client
	ABI     abi.ABI
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

func NewPropsTokenHTTPClient(contractAddr, rpcAddr string) (*Client, error) {
	rawRPC, err := rpc.DialHTTP(rpcAddr)
	if err != nil {
		return nil, err
	}

	rpcClient := ethclient.NewClient(rawRPC)

	address := common.HexToAddress(contractAddr)
	propsTokenClient, err := NewPropsTokenRinkeby(address, rpcClient)
	if err != nil {
		return nil, err
	}

	propsABI, err := readABI()
	if err != nil {
		return nil, err
	}

	return &Client{Token: propsTokenClient, RPC: rpcClient, ABI: propsABI, Address: address.String()}, nil
}

func readABI() (abi.ABI, error) {
	path, _ := filepath.Abs(AbiFilePath)
	file, err := ioutil.ReadFile(path)
	if err != nil {
		return abi.ABI{}, fmt.Errorf("unable to parse PROPS abi (%s)", err)
	}

	propsAbi, err := abi.JSON(strings.NewReader(string(file)))
	if err != nil {
		return abi.ABI{}, fmt.Errorf("unable to parse PROPS abi (%s)", err)
	}

	return propsAbi, nil
}