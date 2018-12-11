#!/bin/bash

source config.sh

if [ -z $BOOTNODE_ENODE_HASH ] || [ -z $BOOTNODE_IP ] || [ -z $BOOTNODE_PORT ]; then
  echo "Misconfigured config.sh, please check bootnode settings"
fi

if [ -z $PEER_PORT ]; then
  echo "Misconfigured config.sh, please check peer settings"
fi

# TODO: accept this as param, passed by mcli
ABSOLUTE_PEERDIR="/opt/marconid/etc/meth/datadir/"

./geth \
  --port $PEER_PORT \
  --rpc \
  --rpcapi="db,eth,net,web3,personal" \
  --rpcport $PEER_RPC_PORT \
  --datadir ${ABSOLUTE_PEERDIR} \
  --testnet