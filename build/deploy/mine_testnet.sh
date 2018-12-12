#!/bin/bash

ABSOLUTE_MINERDIR=${1}

if [ -z ${ABSOLUTE_MINERDIR} ]; then
  echo "Missing path to mining data directory"
fi

source config.sh

if [ -z $BOOTNODE_ENODE_HASH ] || [ -z $BOOTNODE_IP ] || [ -z $BOOTNODE_PORT ]; then
  echo "Misconfigured config.sh, please check bootnode settings"
fi

if [ -z $MINER_PORT ] || [ -z $MINER_ETHERBASE ]; then
  echo "Misconfigured config.sh, please check miner settings"
fi

./geth --datadir ${ABSOLUTE_MINERDIR} \
  --port $MINER_PORT \
  --rpc \
  --rpcport $MINER_RPC_PORT \
  --rpcapi="db,eth,net,web3,personal" \
  --mine --minerthreads 1 \
  --bootnodes enode://3e112c527fc6e672808de5129d6d493b77191706d59dda984496c34b49109ae09ff1f35398aa6f31e44cb3ab65ec91443e94ab7151d1bed618933acae34069a1@54.39.180.40:30301 \
  --testnet \
  --networkid 179109 \
  --etherbase $MINER_ETHERBASE \
  --rpcaddr "0.0.0.0" 
