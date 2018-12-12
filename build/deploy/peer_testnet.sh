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
  --networkid 179109 \
  --port $PEER_PORT \
  --rpc \
  --rpcapi="db,eth,net,web3,personal" \
  --rpcport $PEER_RPC_PORT \
  --datadir ${ABSOLUTE_PEERDIR} \
  --bootnodes enode://3e112c527fc6e672808de5129d6d493b77191706d59dda984496c34b49109ae09ff1f35398aa6f31e44cb3ab65ec91443e94ab7151d1bed618933acae34069a1@54.39.180.40:30301 \
  --testnet
