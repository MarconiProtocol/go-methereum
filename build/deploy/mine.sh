#!/bin/bash

usage() {
  echo "usage: $0 enode_url"
  echo " enode_url | full enode url: enode://12345...@xxx.xxx.xxx.xxx"
}

if [ -z $1 ]; then
  usage
fi

ENODE_URL=$1

./geth --datadir ~/.ethereum-30100 \
  --port 30100 \
  --rpc \
  --rpcport 10100 \
  --networkid 161027 \
  --mine --minerthreads 1 \
  --bootnodes $ENODE_URL:30301 \
  --etherbase 0x21928b6a4a3b768202e3704c957fa1ba8a0c6331 \
