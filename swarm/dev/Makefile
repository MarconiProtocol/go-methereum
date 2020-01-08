.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    github.com/MarconiProtocol/go-methereum/cmd/swarm
	go build -o bin/geth     github.com/MarconiProtocol/go-methereum/cmd/geth
	go build -o bin/bootnode github.com/MarconiProtocol/go-methereum/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v github.com/MarconiProtocol/go-methereum/swarm/...
