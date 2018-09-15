#!/bin/bash

source config.sh

warnBootnode() {
  echo "Misconfigured config.sh, please check bootnode settings"
}

warnMiner() {
  echo "Misconfigured config.sh, please check miner settings"
}

if [ -z $BOOTNODE_ENODE_HASH ] || [ -z $BOOTNODE_IP ] || [ -z $BOOTNODE_PORT ]; then
  warnBootnode
fi

if [ -z $MINER_PORT ] || [ -z $MINER_DATADIR ] || [ -z $MINER_ETHERBASE ]; then
  warnMiner
fi

HOME=`eval echo "~$USER"`
MINER_DATADIR=$HOME/$MINER_DATADIR

./geth --datadir $MINER_DATADIR \
  --port $MINER_PORT \
  --rpc \
  --rpcport $MINER_RPC_PORT \
  --rpcapi="db,eth,net,web3,personal" \
  --networkid 161027 \
  --mine --minerthreads 1 \
  --bootnodes enode://$BOOTNODE_ENODE_HASH@$BOOTNODE_IP:$BOOTNODE_PORT \
  --etherbase $MINER_ETHERBASE 

