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
./geth --datadir $MINER_DATADIR account new --password <(echo "neji") > /tmp/account.js
ACCOUNT=`cat /tmp/account.js | cut -d "{" -f2 | cut -d "}" -f1`
cp genesis.json /tmp/newgenesis.json
sed -i "s/abc8a1591aa743ecfb4c0556aa91105e9edd1718/${ACCOUNT}/" /tmp/newgenesis.json
rm -rf $MINER_DATADIR/geth
./geth --datadir $MINER_DATADIR init /tmp/newgenesis.json
