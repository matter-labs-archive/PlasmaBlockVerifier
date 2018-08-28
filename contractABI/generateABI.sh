#!/bin/sh
cat ../contracts/build/abi | jq .abi > ./abi
go run $GOPATH/src/github.com/ethereum/go-ethereum/cmd/abigen/main.go --abi ./abi --pkg contractABI --type PlasmaParent --out plasmaParent.go
cat ../contracts/build/contracts/PlasmaBlockStorage.json | jq .abi > ./storageABI
go run $GOPATH/src/github.com/ethereum/go-ethereum/cmd/abigen/main.go --abi ./storageABI --pkg contractABI --type PlasmaBlockStorage --out plasmaBlockStorage.go