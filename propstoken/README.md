# propstoken

Golang generated bindings for the PROPSToken contract. Packages are under the bindings 
directory. 

### Rinkeby

- package: [rinkeby](https://github.com/PROPSProject/goprops-toolkit/tree/feature/balances-api-parallel/propstoken/bindings/rinkeby)
- token address: [0xb6768cAFeE6C4CcE5A5643bC78DB09286A33a905](https://rinkeby.etherscan.io/address/0xb6768cafee6c4cce5a5643bc78db09286a33a905)
- compiler: 0.4.24+commit.e67f0147.Emscripten.clang
- generated from abi: 
  > [
      	{
      		"constant": false,
      		"inputs": [
      			{
      				"name": "_spender",
      				"type": "address"
      			},
      			{
      				"name": "_value",
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
      				"name": "_gasPrice",
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
      				"name": "_gasPrice",
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
      		"inputs": [
      			{
      				"name": "_who",
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
      				"name": "_to",
      				"type": "address"
      			},
      			{
      				"name": "_value",
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
      		"constant": true,
      		"inputs": [
      			{
      				"name": "_owner",
      				"type": "address"
      			},
      			{
      				"name": "_spender",
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
      	}
      ]