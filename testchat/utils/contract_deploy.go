/*
* @Author: Joakim
* @Year: 2022
* @Description:
* File is used to create a deployer of the
* smart contract aginst a specific ethereum node.
* Used the instructions from the Go developer book at
* https://goethereumbook.org/en/.
 */

package main

import (
	"context"
	"crypto/ecdsa"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	state "github.com/luxfeerre/thesis-chatapp/chatapp/state" // for demo
)

func contract_deploy() {
	// Change this to your specific ethereum full node
	client, err := ethclient.Dial("http://192.168.2.12:8552")
	if err != nil {
		log.Fatal(err)
	}

	// Change to the specific Ethereum account used to deploy the smart contract
	privateKey, err := crypto.HexToECDSA("811283a34a4429e520dc7f78bfc8be83fc756a6f79e823c91733b90210cd39f5")
	if err != nil {
		log.Fatal(err)
	}

	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	chainID := big.NewInt(101)

	auth, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)

	if err != nil {
		log.Fatal(err)
	}

	auth.Nonce = big.NewInt(int64(nonce))
	auth.Value = big.NewInt(0) // in wei
	auth.GasLimit = uint64(0)  // in units
	auth.GasPrice = big.NewInt(0)

	// Change this to the Deploy function which was genereated from the solidity code
	address, tx, instance, err := state.DeployState(auth, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println(address.Hex())   // Print smart contract address
	fmt.Println(tx.Hash().Hex()) // Print the hash of it

	_ = instance
}

// Used to run the program through go run contract_deploy.go
func main() {
	contract_deploy()
}
