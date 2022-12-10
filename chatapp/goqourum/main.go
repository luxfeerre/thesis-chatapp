package main

import (
	"fmt"
	"io/ioutil"
	"log"

	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/ethereum/go-ethereum/crypto"
)

func keyStorePath() string {
	path := "./node_poa/keystore/"
	dir, err := ioutil.ReadDir(path)
	if err != nil {
		panic(err)
	}

	return "./wallet/" + dir[0].Name()
}

func ReadKeyStore(fileName string, password string) *keystore.Key {

	jsonBytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		log.Fatal(err)
	}

	key, err := keystore.DecryptKey(jsonBytes, password)

	if err != nil {
		log.Fatal(err)
	}

	return key
}

func main() {
	keyPath := keyStorePath()
	key := ReadKeyStore(keyPath, "")
	
	privateKeyBytes := crypto.FromECDSA(key.PrivateKey)

	fmt.Println(hexutil.Encode(privateKeyBytes)[2:])
}
