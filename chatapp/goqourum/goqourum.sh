#! /bin/bash

geth --verbosity 1 console --datadir node_poa --usb --http --http.addr xxx.xxx.xxx.xxx --http.port 8552 --http.api eth,web3,personal,net  --allow-insecure-unlock --unlock "0" --mine --miner.threads=8 --snapshot=false --syncmode full 
