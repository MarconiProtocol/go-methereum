#!/bin/bash

mkdir -p /opt/marconi/etc/meth/datadir

./gmeth --networkid 161027 --port 13200 --rpc --rpcapi=db,eth,net,web3,personal --rpcport 10004 --datadir /opt/marconi/etc/meth/datadir/ --metrics --metrics.influxdb --metrics.influxdb.endpoint="http://35.227.152.19:8086"
