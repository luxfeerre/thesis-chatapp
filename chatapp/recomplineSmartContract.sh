#!/bin/sh

# Use this when changes to the smart contract has been made to recompile it and generate new golang code.

# Set the current working diretory
currentDirectory=$(pwd)

docker run --rm -v $currentDirectory/state:/state ethereum/solc:stable /state/State.sol --bin --abi --optimize --ir-optimized --overwrite -o /state

sleep 1

docker run --rm -v $currentDirectory/state:/state ethereum/client-go:alltools-latest abigen --bin /state/State.bin --abi /state/State.abi --pkg state --type state --out /state/State.go
