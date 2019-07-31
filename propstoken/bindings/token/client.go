package propstoken

import (
	"context"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/ethereum/go-ethereum/rpc"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"fmt"
	"strings"
)

var AbiString string

// PropsClient ...
type Client struct {
	Token   *PropsToken
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
	propsTokenClient, err := NewPropsToken(address, rpcClient)
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
	propsTokenClient, err := NewPropsToken(address, rpcClient)
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
	propsAbi, err := abi.JSON(strings.NewReader(AbiString))
	if err != nil {
		return abi.ABI{}, fmt.Errorf("unable to parse PROPS abi (%s)", err)
	}

	return propsAbi, nil
}

func init()  {
	AbiString = `[
    {
      "constant": false,
      "inputs": [
        {
          "name": "_controller",
          "type": "address"
        }
      ],
      "name": "updateController",
      "outputs": [],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [],
      "name": "name",
      "outputs": [
        {
          "name": "",
          "type": "string"
        }
      ],
      "payable": false,
      "stateMutability": "view",
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "spender",
          "type": "address"
        },
        {
          "name": "value",
          "type": "uint256"
        }
      ],
      "name": "approve",
      "outputs": [
        {
          "name": "",
          "type": "bool"
        }
      ],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_entityType",
          "type": "uint8"
        },
        {
          "name": "_rewardsDay",
          "type": "uint256"
        }
      ],
      "name": "getEntities",
      "outputs": [
        {
          "name": "",
          "type": "address[]"
        }
      ],
      "payable": false,
      "stateMutability": "view",
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_token",
          "type": "address"
        },
        {
          "name": "_to",
          "type": "address"
        },
        {
          "name": "_value",
          "type": "uint256"
        },
        {
          "name": "_fee",
          "type": "uint256"
        },
        {
          "name": "_nonce",
          "type": "uint256"
        }
      ],
      "name": "getTransferPreSignedHash",
      "outputs": [
        {
          "name": "",
          "type": "bytes32"
        }
      ],
      "payable": false,
      "stateMutability": "pure",
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "_signature",
          "type": "bytes"
        },
        {
          "name": "_to",
          "type": "address"
        },
        {
          "name": "_value",
          "type": "uint256"
        },
        {
          "name": "_fee",
          "type": "uint256"
        },
        {
          "name": "_nonce",
          "type": "uint256"
        }
      ],
      "name": "transferPreSigned",
      "outputs": [
        {
          "name": "",
          "type": "bool"
        }
      ],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_token",
          "type": "address"
        },
        {
          "name": "_spender",
          "type": "address"
        },
        {
          "name": "_addedValue",
          "type": "uint256"
        },
        {
          "name": "_fee",
          "type": "uint256"
        },
        {
          "name": "_nonce",
          "type": "uint256"
        }
      ],
      "name": "getIncreaseAllowancePreSignedHash",
      "outputs": [
        {
          "name": "",
          "type": "bytes32"
        }
      ],
      "payable": false,
      "stateMutability": "pure",
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "_name",
          "type": "uint8"
        },
        {
          "name": "_value",
          "type": "uint256"
        },
        {
          "name": "_rewardsDay",
          "type": "uint256"
        }
      ],
      "name": "updateParameter",
      "outputs": [],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [],
      "name": "totalSupply",
      "outputs": [
        {
          "name": "",
          "type": "uint256"
        }
      ],
      "payable": false,
      "stateMutability": "view",
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "from",
          "type": "address"
        },
        {
          "name": "to",
          "type": "address"
        },
        {
          "name": "value",
          "type": "uint256"
        }
      ],
      "name": "transferFrom",
      "outputs": [
        {
          "name": "",
          "type": "bool"
        }
      ],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [],
      "name": "maxTotalSupply",
      "outputs": [
        {
          "name": "",
          "type": "uint256"
        }
      ],
      "payable": false,
      "stateMutability": "view",
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [],
      "name": "decimals",
      "outputs": [
        {
          "name": "",
          "type": "uint8"
        }
      ],
      "payable": false,
      "stateMutability": "view",
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "spender",
          "type": "address"
        },
        {
          "name": "addedValue",
          "type": "uint256"
        }
      ],
      "name": "increaseAllowance",
      "outputs": [
        {
          "name": "",
          "type": "bool"
        }
      ],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [],
      "name": "canTransferBeforeStartTime",
      "outputs": [
        {
          "name": "",
          "type": "address"
        }
      ],
      "payable": false,
      "stateMutability": "view",
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_name",
          "type": "uint8"
        },
        {
          "name": "_rewardsDay",
          "type": "uint256"
        }
      ],
      "name": "getParameter",
      "outputs": [
        {
          "name": "",
          "type": "uint256"
        }
      ],
      "payable": false,
      "stateMutability": "view",
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_token",
          "type": "address"
        },
        {
          "name": "_spender",
          "type": "address"
        },
        {
          "name": "_subtractedValue",
          "type": "uint256"
        },
        {
          "name": "_fee",
          "type": "uint256"
        },
        {
          "name": "_nonce",
          "type": "uint256"
        }
      ],
      "name": "getDecreaseAllowancePreSignedHash",
      "outputs": [
        {
          "name": "",
          "type": "bytes32"
        }
      ],
      "payable": false,
      "stateMutability": "pure",
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "_rewardsDay",
          "type": "uint256"
        },
        {
          "name": "_rewardsHash",
          "type": "bytes32"
        },
        {
          "name": "_applications",
          "type": "address[]"
        },
        {
          "name": "_amounts",
          "type": "uint256[]"
        }
      ],
      "name": "submitDailyRewards",
      "outputs": [],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "_signature",
          "type": "bytes"
        },
        {
          "name": "_spender",
          "type": "address"
        },
        {
          "name": "_value",
          "type": "uint256"
        },
        {
          "name": "_fee",
          "type": "uint256"
        },
        {
          "name": "_nonce",
          "type": "uint256"
        }
      ],
      "name": "approvePreSigned",
      "outputs": [
        {
          "name": "",
          "type": "bool"
        }
      ],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [],
      "name": "rewardsStartTimestamp",
      "outputs": [
        {
          "name": "",
          "type": "uint256"
        }
      ],
      "payable": false,
      "stateMutability": "view",
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "owner",
          "type": "address"
        }
      ],
      "name": "balanceOf",
      "outputs": [
        {
          "name": "",
          "type": "uint256"
        }
      ],
      "payable": false,
      "stateMutability": "view",
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "_rewardsDay",
          "type": "uint256"
        },
        {
          "name": "_applications",
          "type": "address[]"
        }
      ],
      "name": "setApplications",
      "outputs": [],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_token",
          "type": "address"
        },
        {
          "name": "_spender",
          "type": "address"
        },
        {
          "name": "_value",
          "type": "uint256"
        },
        {
          "name": "_fee",
          "type": "uint256"
        },
        {
          "name": "_nonce",
          "type": "uint256"
        }
      ],
      "name": "getApprovePreSignedHash",
      "outputs": [
        {
          "name": "",
          "type": "bytes32"
        }
      ],
      "payable": false,
      "stateMutability": "pure",
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [],
      "name": "symbol",
      "outputs": [
        {
          "name": "",
          "type": "string"
        }
      ],
      "payable": false,
      "stateMutability": "view",
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "_entityType",
          "type": "uint8"
        },
        {
          "name": "_name",
          "type": "bytes32"
        },
        {
          "name": "_rewardsAddress",
          "type": "address"
        },
        {
          "name": "_sidechainAddress",
          "type": "address"
        }
      ],
      "name": "updateEntity",
      "outputs": [],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "spender",
          "type": "address"
        },
        {
          "name": "subtractedValue",
          "type": "uint256"
        }
      ],
      "name": "decreaseAllowance",
      "outputs": [
        {
          "name": "",
          "type": "bool"
        }
      ],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "_token",
          "type": "address"
        },
        {
          "name": "_from",
          "type": "address"
        },
        {
          "name": "_to",
          "type": "address"
        },
        {
          "name": "_value",
          "type": "uint256"
        },
        {
          "name": "_fee",
          "type": "uint256"
        },
        {
          "name": "_nonce",
          "type": "uint256"
        }
      ],
      "name": "getTransferFromPreSignedHash",
      "outputs": [
        {
          "name": "",
          "type": "bytes32"
        }
      ],
      "payable": false,
      "stateMutability": "pure",
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "to",
          "type": "address"
        },
        {
          "name": "value",
          "type": "uint256"
        }
      ],
      "name": "transfer",
      "outputs": [
        {
          "name": "",
          "type": "bool"
        }
      ],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "_signature",
          "type": "bytes"
        },
        {
          "name": "_from",
          "type": "address"
        },
        {
          "name": "_to",
          "type": "address"
        },
        {
          "name": "_value",
          "type": "uint256"
        },
        {
          "name": "_fee",
          "type": "uint256"
        },
        {
          "name": "_nonce",
          "type": "uint256"
        }
      ],
      "name": "transferFromPreSigned",
      "outputs": [
        {
          "name": "",
          "type": "bool"
        }
      ],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "_applicationAddress",
          "type": "address"
        },
        {
          "name": "_userId",
          "type": "bytes32"
        },
        {
          "name": "_to",
          "type": "address"
        },
        {
          "name": "_amount",
          "type": "uint256"
        }
      ],
      "name": "settle",
      "outputs": [],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [],
      "name": "transfersStartTime",
      "outputs": [
        {
          "name": "",
          "type": "uint256"
        }
      ],
      "payable": false,
      "stateMutability": "view",
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "owner",
          "type": "address"
        },
        {
          "name": "spender",
          "type": "address"
        }
      ],
      "name": "allowance",
      "outputs": [
        {
          "name": "",
          "type": "uint256"
        }
      ],
      "payable": false,
      "stateMutability": "view",
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "_rewardsDay",
          "type": "uint256"
        },
        {
          "name": "_validators",
          "type": "address[]"
        }
      ],
      "name": "setValidators",
      "outputs": [],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "_controller",
          "type": "address"
        },
        {
          "name": "_minSecondsBetweenDays",
          "type": "uint256"
        },
        {
          "name": "_rewardsStartTimestamp",
          "type": "uint256"
        }
      ],
      "name": "initializePostRewardsUpgrade1",
      "outputs": [],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "_signature",
          "type": "bytes"
        },
        {
          "name": "_spender",
          "type": "address"
        },
        {
          "name": "_subtractedValue",
          "type": "uint256"
        },
        {
          "name": "_fee",
          "type": "uint256"
        },
        {
          "name": "_nonce",
          "type": "uint256"
        }
      ],
      "name": "decreaseAllowancePreSigned",
      "outputs": [
        {
          "name": "",
          "type": "bool"
        }
      ],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [],
      "name": "controller",
      "outputs": [
        {
          "name": "",
          "type": "address"
        }
      ],
      "payable": false,
      "stateMutability": "view",
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "_signature",
          "type": "bytes"
        },
        {
          "name": "_spender",
          "type": "address"
        },
        {
          "name": "_addedValue",
          "type": "uint256"
        },
        {
          "name": "_fee",
          "type": "uint256"
        },
        {
          "name": "_nonce",
          "type": "uint256"
        }
      ],
      "name": "increaseAllowancePreSigned",
      "outputs": [
        {
          "name": "",
          "type": "bool"
        }
      ],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "name": "rewardsDay",
          "type": "uint256"
        },
        {
          "indexed": true,
          "name": "rewardsHash",
          "type": "bytes32"
        },
        {
          "indexed": true,
          "name": "validator",
          "type": "address"
        }
      ],
      "name": "DailyRewardsSubmitted",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "name": "rewardsDay",
          "type": "uint256"
        },
        {
          "indexed": true,
          "name": "rewardsHash",
          "type": "bytes32"
        },
        {
          "indexed": false,
          "name": "numOfApplications",
          "type": "uint256"
        },
        {
          "indexed": false,
          "name": "amount",
          "type": "uint256"
        }
      ],
      "name": "DailyRewardsApplicationsMinted",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "name": "rewardsDay",
          "type": "uint256"
        },
        {
          "indexed": true,
          "name": "rewardsHash",
          "type": "bytes32"
        },
        {
          "indexed": false,
          "name": "numOfValidators",
          "type": "uint256"
        },
        {
          "indexed": false,
          "name": "amount",
          "type": "uint256"
        }
      ],
      "name": "DailyRewardsValidatorsMinted",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "name": "id",
          "type": "address"
        },
        {
          "indexed": true,
          "name": "entityType",
          "type": "uint8"
        },
        {
          "indexed": false,
          "name": "name",
          "type": "bytes32"
        },
        {
          "indexed": false,
          "name": "rewardsAddress",
          "type": "address"
        },
        {
          "indexed": true,
          "name": "sidechainAddress",
          "type": "address"
        }
      ],
      "name": "EntityUpdated",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "name": "param",
          "type": "uint8"
        },
        {
          "indexed": false,
          "name": "newValue",
          "type": "uint256"
        },
        {
          "indexed": false,
          "name": "oldValue",
          "type": "uint256"
        },
        {
          "indexed": false,
          "name": "rewardsDay",
          "type": "uint256"
        }
      ],
      "name": "ParameterUpdated",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "name": "validatorsList",
          "type": "address[]"
        },
        {
          "indexed": true,
          "name": "rewardsDay",
          "type": "uint256"
        }
      ],
      "name": "ValidatorsListUpdated",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "name": "applicationsList",
          "type": "address[]"
        },
        {
          "indexed": true,
          "name": "rewardsDay",
          "type": "uint256"
        }
      ],
      "name": "ApplicationsListUpdated",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "name": "newController",
          "type": "address"
        }
      ],
      "name": "ControllerUpdated",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "name": "applicationId",
          "type": "address"
        },
        {
          "indexed": true,
          "name": "userId",
          "type": "bytes32"
        },
        {
          "indexed": true,
          "name": "to",
          "type": "address"
        },
        {
          "indexed": false,
          "name": "amount",
          "type": "uint256"
        },
        {
          "indexed": false,
          "name": "rewardsAddress",
          "type": "address"
        }
      ],
      "name": "Settlement",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "name": "from",
          "type": "address"
        },
        {
          "indexed": true,
          "name": "to",
          "type": "address"
        },
        {
          "indexed": true,
          "name": "delegate",
          "type": "address"
        },
        {
          "indexed": false,
          "name": "amount",
          "type": "uint256"
        },
        {
          "indexed": false,
          "name": "fee",
          "type": "uint256"
        }
      ],
      "name": "TransferPreSigned",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "name": "from",
          "type": "address"
        },
        {
          "indexed": true,
          "name": "to",
          "type": "address"
        },
        {
          "indexed": true,
          "name": "delegate",
          "type": "address"
        },
        {
          "indexed": false,
          "name": "amount",
          "type": "uint256"
        },
        {
          "indexed": false,
          "name": "fee",
          "type": "uint256"
        }
      ],
      "name": "ApprovalPreSigned",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "name": "from",
          "type": "address"
        },
        {
          "indexed": true,
          "name": "to",
          "type": "address"
        },
        {
          "indexed": false,
          "name": "value",
          "type": "uint256"
        }
      ],
      "name": "Transfer",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "name": "owner",
          "type": "address"
        },
        {
          "indexed": true,
          "name": "spender",
          "type": "address"
        },
        {
          "indexed": false,
          "name": "value",
          "type": "uint256"
        }
      ],
      "name": "Approval",
      "type": "event"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "name",
          "type": "string"
        },
        {
          "name": "symbol",
          "type": "string"
        },
        {
          "name": "decimals",
          "type": "uint8"
        }
      ],
      "name": "initialize",
      "outputs": [],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "_holder",
          "type": "address"
        },
        {
          "name": "_controller",
          "type": "address"
        },
        {
          "name": "_minSecondsBetweenDays",
          "type": "uint256"
        },
        {
          "name": "_rewardsStartTimestamp",
          "type": "uint256"
        }
      ],
      "name": "initialize",
      "outputs": [],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    }
  ]`
}