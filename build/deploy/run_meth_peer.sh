#!/bin/bash

usage() {
  echo "usage: $0 : enode_url"
  echo "  enode_url | full enode url: enode://12345...@xxx.xxx.xxx.xxx"
}

if [ -z $1 ]; then 
  usage
fi

ENODE_URL=$1

./geth --networkid 161027 \
  --port 30304 \
  --rpc \
  --rpcapi="db,eth,net,web3,personal" \
  --rpcport 10004 \
  --datadir ~/.ethereum-30304 \
  --bootnodes $ENODE_URL:30301
