/*
* @Author: Joakim
* @Year: 2022
* @Description:
* Main program used to execute the proof of concept code togther with an external tracer,
* Ethereum full node with an activ smart contract with a live user.
 */

package main

import (
	"context"
	"time"

	"fmt"
	"log"

	configuration "github.com/blockchain/chatApp/configuration"
	ethereumService "github.com/blockchain/chatApp/ethereumService"
	libp2putils "github.com/blockchain/chatApp/libp2putils"
	"github.com/blockchain/chatApp/pubSubService"
	test "github.com/blockchain/chatApp/test"
	userInterface "github.com/blockchain/chatApp/userInterface"

	pubsub "github.com/libp2p/go-libp2p-pubsub"
)

/*
* Description: It Starts by creating a cancel context and then two variables to store a seed and a username.
* After it setts the seed and username trhough the config package,
* connexts to the Ethereum full nodes smart contract and creates a wallet with a key
* used to iteract with the blockchain.
* Then it creates a peer-to-peer host through libp2p and then registers at the Ehtereum blockchain.
* After it configures a bootstrap node list from the Ethereum blockchain.
* After that it sets up its own bootndoe capabilities through DHT.
* Then it sets up a gossip router used for the publish and subscribe message application.
* Followed by a tracer used for logging and keeping track of network propogation deley.
* Then it setts up a chatroom and a UI for the program.
* Then it will exit gracefully.
* @dev Return value nil if successfull
* @return value of 'nil'
 */

func runner() error {
	// Initalize a cancalble context so it can stop if certaint channels get closed
	ctxBackground, cancel := context.WithCancel(context.Background())
	defer cancel()

	// initalize variables to the default used to store seed and username
	var seed int64 = 0
	var name string = ""

	/*
	 * Initalize Confiugration struct and then set the values in the struct
	 * by either reading from the terminal or setting a predefined value
	 */
	conf := configuration.Config(seed, name)

	// Create an ethereum client which is used to connect to a full node
	client := ethereumService.ConnectClient(conf.Node)

	// Create instance of contract which is used to interact with it.
	instance, err := ethereumService.StateInstance(client)
	if err != nil {
		log.Fatal(err)
	}

	// Either create a new wallet or reads from existing to extract private key from it.
	key := ethereumService.UserWallet(conf.Password)

	// Set up transaction data to interact with the Ethereum blockchain
	transactionData := ethereumService.Transactor(client, key, ctxBackground)

	// Create a new libp2p host which will listen on a random tcp port when using 0
	// and gets a determenistik key depending on the seed value
	host, err := libp2putils.NewHost(conf.Seed, conf.Port)
	if err != nil {
		panic(err)
	}

	// Register as a new user in the Ethereum blockchain or join it by registering as a bootstrap node
	if len(conf.UserName) != 0 {
		if exists := ethereumService.Verify(instance, transactionData); exists {
			log.Fatal("Account allready exists for this wallet.")
		} else if err := ethereumService.Register(instance, transactionData, conf.UserName,
			string(host.ID().Pretty()), fmt.Sprintf("%s/p2p/%s", host.Addrs()[0], host.ID().Pretty())); err != nil {
			log.Println(err)
		}
		// Register as a bootstrap node if allready registered and valid and else print out an error message
	} else if exists := ethereumService.Verify(instance, transactionData); exists {
		ethereumService.ModDhtList(instance, transactionData, fmt.Sprintf("%s/p2p/%s", host.Addrs()[0], host.ID().Pretty()))
	} else {
		log.Fatal("You have not registered for the ethereum service yet with this account")
	}

	// Sleep for 1 second so the registered information can be read into the blockchain
	time.Sleep(1 * time.Second)

	// Read username directly from the blockchain registered for the account
	if conf.UserName, err = ethereumService.GetUserName(instance, transactionData, host.ID().Pretty()); err != nil {
		log.Println(err)
	}

	// Set boot node list after user has registered at the smart contract
	configuration.ConfBootNodeList(instance, transactionData, conf)

	// Setup a new Distributed Hash table which is used to discover peers through a boot node
	// Or if its set to be a boot node act as a connector for others nodes which just joibs the network
	kaddht := libp2putils.NewDht(ctxBackground, host, conf.DiscoveryPeers, instance, transactionData, conf)
	if err != nil {
		log.Println(err)
	}

	// Create a remote trace for pubsub
	tracer, err := test.PubsubTrace(ctxBackground, host)
	if err != nil {
		log.Println(err)
	}
	defer tracer.Close()

	/* Set a context, the host using the pubsub channel, a tracer for logging and params for the test
	   Set up discovery through kademila dht and set a max message size of 4913
	*/
	pubSubChannel, err := pubsub.NewGossipSub(
		ctxBackground,
		host,
		pubsub.WithEventTracer(tracer),
		pubsub.WithDiscovery(kaddht),
		pubsub.WithMaxMessageSize(4913),
	)
	if err != nil {
		panic(err)
	}

	// Set up a new chatroom for the user with the specific host ID and username and room
	chatR, err := pubSubService.JoinChatRoom(ctxBackground, pubSubChannel, host.ID(), conf.UserName, conf.Room)
	if err != nil {
		log.Println(err)
	}

	// Create a UI for the user so it can interact and send and rescive messages
	uInterface := userInterface.NewChatUI(chatR)
	if err = uInterface.Run(instance, transactionData); err != nil {
		return err
	}

	// When finshed return nil to exist gracefully
	return nil
}

func main() {
	runner()
}
