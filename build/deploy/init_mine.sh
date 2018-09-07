#!/bin/bash

source config.sh

warnMiner() {
  echo "Misconfigured config.sh, please check miner settings"
}

if [ -z $MINER_DATADIR ]; then
  warnMiner
fi


./geth --datadir $MINER_DATADIR init genesis.json
./geth --datadir $MINER_DATADIR account new --password <(echo "neji")
