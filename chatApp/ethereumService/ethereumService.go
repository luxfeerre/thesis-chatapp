/*
* @Author: Joakim
* @Year: 2022
* @Description:
* File is used to interact with the Ethereum blockchain smart contract
* service.
 */

package ethereumService

import (
	"context"
	"crypto/ecdsa"
	"io/ioutil"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/accounts/keystore"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"

	state "github.com/luxfeerre/thesis-chatapp/chatapp/state"
)

/*
* Description: Read a wallet file and store it in jsonBytes
* Then decrypt the key with the users password and return the key.
* @dev Return a dereferenced pointer to a Key struct from the Ethereum keystore libary.
* @return value of 'key struct'
 */
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

/*
* Description: Creates a new keystore under the wallet directory and add a new account to it.
* @dev Return nothing.
* @return value of 'nnothing'
 */
func CreateKeyStore(password string) {
	keyStore := keystore.NewKeyStore("./wallet", keystore.StandardScryptN, keystore.StandardScryptP)

	_, err := keyStore.NewAccount(password)
	if err != nil {
		log.Fatal(err)
	}
}

/*
* Description: Open or creates a keystore through walletExists function and return a path to it
* and then reads the key from it throguh ReadKeyStore.
* @dev Return a dereferenced pointer to a Key struct from the Ethereum keystore libary.
* @return value of 'key struct'
 */
func UserWallet(password string) *keystore.Key {
	fileName := walletExists(password)
	key := ReadKeyStore(fileName, password)
	return key
}

/*
* Description: Either creates or returns an active path to a user wallet.
* @dev Return a path to an activ wallet for a user.
* @return value of 'string'
 */
func walletExists(password string) string {
	dir, err := ioutil.ReadDir("./wallet")
	if err != nil {
		CreateKeyStore(password)
		dir, err = ioutil.ReadDir("./wallet")
		if err != nil {
			panic(err)
		}
	}

	return "./wallet/" + dir[0].Name()
}

/*
* Description: Connectes to an Ethereum full node through RPC and returns
* the connection to the user.
* @dev Return a struct to an RPC client connected to Ethereum full node.
* @return value of 'Pointer to struct of rpc client'
 */
func ConnectClient(address string) *ethclient.Client {
	client, err := ethclient.Dial(address)

	if err != nil {
		log.Fatal((err))
	}

	return client
}

/*
* Description: Sets up transaction data used to interact with a Payable function
* on the Ethereum blockchain smart contract.
* @dev Return a struct to Options used for a transaction with Ethereum blockchain.
* @return value of 'Pointer to a struct of transaction options'
 */
func Transactor(client *ethclient.Client, key *keystore.Key, ctxBackground context.Context) *bind.TransactOpts {
	// Extract priavte key
	privateKey := key.PrivateKey

	// Extract public key and store is as ecdsa format
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		log.Fatal("cannot assert type: publicKey is not of type *ecdsa.PublicKey")
	}

	// Generate a common address used to identify the user
	fromAddress := crypto.PubkeyToAddress(*publicKeyECDSA)
	// Generate a nonce from it with the current context
	nonce, err := client.PendingNonceAt(context.Background(), fromAddress)
	if err != nil {
		log.Fatal(err)
	}

	// Generate a chain id used in the transaction data
	chainID := big.NewInt(101)

	// Create the transaction data struct signed by the private key for the specific chain
	transactionData, err := bind.NewKeyedTransactorWithChainID(privateKey, chainID)
	if err != nil {
		log.Fatal(err)
	}

	// Populate the transaction data
	transactionData.Nonce = big.NewInt(int64(nonce))
	transactionData.Value = big.NewInt(0) // in wei
	transactionData.GasLimit = uint64(0)  // in units
	transactionData.GasPrice = big.NewInt(0)
	transactionData.From = fromAddress
	transactionData.Context = ctxBackground

	return transactionData
}

/*
* Description: Register a user by using transaction data and the given input that
* is needed to populate the smart contracts registration service.
* @dev Return a error which is nil if success and otherwise contins an error message.
* @return value of 'error'
 */
func Register(instance *state.State, transactionData *bind.TransactOpts, name string, peerId string, address string) error {
	_, err := instance.Register(transactionData, name, peerId, address)
	return err
}

/*
* Description: Creates a new state connection with an address to a contract to a specific ethereum full node.
* @dev Return a connection in the contract through a State struct and error information which is nil if succesfull or error message on fail.
* @return value of 'State struct and error'
 */
func StateInstance(client *ethclient.Client) (*state.State, error) {
	address := common.HexToAddress("0x99ddD1DF9C9719294e8cD34B1FFCC6B03CfFeBB0")
	return state.NewState(address, client)
}

/*
* Description: Get username information from the smart contract by seraching for a peer id and
* setting call options from the transaction data.
* @dev Return a username and nil if success or error message on failure.
* @return value of 'String and error'
 */
func GetUserName(instance *state.State, transactionData *bind.TransactOpts, peerId string) (string, error) {
	callOpts := bind.CallOpts{}
	callOpts.Pending = true
	callOpts.From = transactionData.From
	callOpts.Context = transactionData.Context

	return instance.GetUserName(&callOpts, peerId)
}

/*
* Description: Get address information from the smart contract by seraching for a username and
* setting call options from the transaction data.
* @dev Return a Address and nil if success or error message on failure.
* @return value of 'String and error'
 */
func GetUserAddress(instance *state.State, transactionData *bind.TransactOpts, userName string) (string, error) {
	callOpts := bind.CallOpts{}
	callOpts.Pending = true
	callOpts.From = transactionData.From
	callOpts.Context = transactionData.Context

	return instance.GetUserAddress(&callOpts, userName)
}

/*
* Description: Search all registered user accounts on ethereum address used for identity with
* call opts from transaction data and return a boolean value.
* @dev Return a bool if user exists and otherwise false.
* @return value of 'bool'
 */
func Verify(instance *state.State, transactionData *bind.TransactOpts) bool {
	callOpts := bind.CallOpts{}
	callOpts.Pending = true
	callOpts.From = transactionData.From
	callOpts.Context = transactionData.Context

	exists, err := instance.Verification(&callOpts)

	if err != nil {
		return false
	}

	return exists
}

/*
* Description: Takes an instance of the smart contract and transaction data
* and then query the conttract to get a list of active dht bootstrap nodes.
* @dev Return a string slice of activ DHT bootstrap nodes on the blockchain.
* @return value of 'String slice'
 */
func GetDhtList(instance *state.State, transactionData *bind.TransactOpts) []string {
	callOpts := bind.CallOpts{}
	callOpts.Pending = true
	callOpts.From = transactionData.From
	callOpts.Context = transactionData.Context

	dhtAddressList, _ := instance.GetDhtList(&callOpts)

	return dhtAddressList
}

/*
* Description: Takes an instance of the smart contract and transaction data
* and address string and adds it to a DHT bootstrap list on the ethereum smart contract.
* @dev Return a nil value on success and error message on fail.
* @return value of 'error'
 */
func ModDhtList(instance *state.State, transactionData *bind.TransactOpts, addr string) error {

	_, err := instance.ModDhtList(transactionData, addr)

	if err != nil {
		return err
	}

	return nil
}
