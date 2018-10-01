#!/bin/bash

source config.sh

warnMiner() {
  echo "Misconfigured config.sh, please check miner settings"
}

if [ -z $MINER_DATADIR ]; then
  warnMiner
fi

HOME=`eval echo "~$USER"`
MINER_DATADIR=$HOME/$MINER_DATADIR

rm -rf $MINER_DATADIR
./geth --datadir $MINER_DATADIR init genesis.json
