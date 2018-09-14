#!/bin/bash

source config.sh

warnBootnode() {
  echo "Misconfigured config.sh, please check bootnode settings"
}

warnPeer() {
  echo "Misconfigured config.sh, please check peer settings"
}

if [ -z $BOOTNODE_ENODE_HASH ] || [ -z $BOOTNODE_IP ] || [ -z $BOOTNODE_PORT ]; then
  warnBootnode
fi

if [ -z $PEER_PORT ] || [ -z $PEER_DATADIR ]; then
  warnPeer
fi

HOME=`eval echo "~$USER"`
PEER_DATADIR=$HOME/$PEER_DATADIR

./geth --networkid 161027 \
  --port $PEER_PORT \
  --rpc \
  --rpcapi="db,eth,net,web3,personal" \
  --rpcport $PEER_RPC_PORT \
  --datadir $PEER_DATADIR \
  --bootnodes enode://$BOOTNODE_ENODE_HASH@$BOOTNODE_IP:$BOOTNODE_PORT

