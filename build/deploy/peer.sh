#!/bin/bash

source ../etc/meth/config.sh

if [ -z $NETWORK_ID ] || [ -z $BOOTNODE_PORT ] || [ -z $PEER_PORT ] || [ -z $SYNC_MODE ] || [ -z $TESTNET_FLAG ]; then
  echo "***************************"
  echo "* Misconfigured config.sh *"
  echo "***************************"
  exit 1
fi

# TODO: accept this as param, passed by mcli
ABSOLUTE_PEERDIR="/opt/marconi/etc/meth/datadir/"

if [ "$TESTNET_FLAG" = true ]; then
./gmeth \
  --networkid $NETWORK_ID \
  --port $PEER_PORT \
  --rpc \
  --rpcapi="db,eth,net,web3,personal" \
  --rpcport $PEER_RPC_PORT \
  --datadir ${ABSOLUTE_PEERDIR} \
  --rpcaddr 0.0.0.0 \
  --syncmode $SYNC_MODE \
  --testnet
else
./gmeth \
  --networkid $NETWORK_ID \
  --port $PEER_PORT \
  --rpc \
  --rpcapi="db,eth,net,web3,personal" \
  --rpcport $PEER_RPC_PORT \
  --datadir ${ABSOLUTE_PEERDIR} \
  --rpcaddr 0.0.0.0 \
  --syncmode $SYNC_MODE
fi
