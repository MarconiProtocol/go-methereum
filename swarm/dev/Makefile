.PHONY: build cluster test

default: build

build:
	go build -o bin/swarm    gitlab.neji.vm.tc/marconi/go-ethereum/cmd/swarm
	go build -o bin/geth     gitlab.neji.vm.tc/marconi/go-ethereum/cmd/geth
	go build -o bin/bootnode gitlab.neji.vm.tc/marconi/go-ethereum/cmd/bootnode

cluster: build
	scripts/boot-cluster.sh

test:
	go test -v gitlab.neji.vm.tc/marconi/go-ethereum/swarm/...
