#!/bin/bash

source config.sh

warnPeer() {
  echo "Misconfigured config.sh, please check peer settings"
}

if [ -z $PEER_DATADIR ]; then
  warnPeer
fi

HOME=`eval echo "~$USER"`
PEER_DATADIR=$HOME/$PEER_DATADIR

rm -rf $PEER_DATADIR
./geth --datadir $PEER_DATADIR init genesis.json
./geth --datadir $PEER_DATADIR account new --password <(echo "neji")
