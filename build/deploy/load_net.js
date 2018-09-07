personal.unlockAccount(eth.accounts[0], "neji")
var network = eth.contract([
    {
      "constant": true,
      "inputs": [],
      "name": "active",
      "outputs": [
        {
          "name": "",
          "type": "bool"
        }
      ],
      "payable": false,
      "stateMutability": "view",
      "type": "function"
    },
    {
      "constant": false,
      "inputs": [],
      "name": "renounceOwnership",
      "outputs": [],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [],
      "name": "owner",
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
      "inputs": [],
      "name": "id",
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
          "name": "",
          "type": "uint256"
        }
      ],
      "name": "peers",
      "outputs": [
        {
          "name": "networkId",
          "type": "uint256"
        },
        {
          "name": "pubKeyHash",
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
          "name": "_newOwner",
          "type": "address"
        }
      ],
      "name": "transferOwnership",
      "outputs": [],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [],
      "name": "admin",
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
      "inputs": [
        {
          "name": "_id",
          "type": "uint256"
        },
        {
          "name": "_admin",
          "type": "address"
        }
      ],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "constructor"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "name": "networkId",
          "type": "uint256"
        },
        {
          "indexed": false,
          "name": "pubKeyHash",
          "type": "string"
        }
      ],
      "name": "PeerAdded",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "name": "networkId",
          "type": "uint256"
        },
        {
          "indexed": false,
          "name": "pubKeyHash",
          "type": "string"
        }
      ],
      "name": "PeerRemoved",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": false,
          "name": "networkId",
          "type": "uint256"
        },
        {
          "indexed": false,
          "name": "pubKeyHashMine",
          "type": "string"
        },
        {
          "indexed": false,
          "name": "pubKeyHashOther",
          "type": "string"
        }
      ],
      "name": "PeerRelationAdded",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "name": "previousOwner",
          "type": "address"
        }
      ],
      "name": "OwnershipRenounced",
      "type": "event"
    },
    {
      "anonymous": false,
      "inputs": [
        {
          "indexed": true,
          "name": "previousOwner",
          "type": "address"
        },
        {
          "indexed": true,
          "name": "newOwner",
          "type": "address"
        }
      ],
      "name": "OwnershipTransferred",
      "type": "event"
    },
    {
      "constant": false,
      "inputs": [
        {
          "name": "pubKeyHash",
          "type": "string"
        }
      ],
      "name": "addPeer",
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
          "name": "pubKeyHashMine",
          "type": "string"
        },
        {
          "name": "pubKeyHashOther",
          "type": "string"
        }
      ],
      "name": "addPeerRelation",
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
          "name": "pubKeyHash",
          "type": "string"
        }
      ],
      "name": "removePeer",
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
          "name": "pubKeyHash",
          "type": "string"
        }
      ],
      "name": "getPeerRelations",
      "outputs": [
        {
          "name": "ret",
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
          "name": "enable",
          "type": "bool"
        }
      ],
      "name": "updateNetworkState",
      "outputs": [],
      "payable": false,
      "stateMutability": "nonpayable",
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [
        {
          "name": "index",
          "type": "uint256"
        }
      ],
      "name": "getPeer",
      "outputs": [
        {
          "name": "networkId",
          "type": "uint256"
        },
        {
          "name": "pubKeyHash",
          "type": "string"
        },
        {
          "name": "neighbors",
          "type": "string"
        }
      ],
      "payable": false,
      "stateMutability": "view",
      "type": "function"
    },
    {
      "constant": true,
      "inputs": [],
      "name": "getPeerCount",
      "outputs": [
        {
          "name": "",
          "type": "uint256"
        }
      ],
      "payable": false,
      "stateMutability": "view",
      "type": "function"
    }
  ]).at("0xb2FC30C3391A0a5c77ec81a08600aB842F651925")
