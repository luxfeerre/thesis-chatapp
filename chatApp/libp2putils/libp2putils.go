/*
* @Author: Joakim
* @Year: 2022
* @Description:
* File is used to create libp2p node and a DHT routing and discovery capabilities.
 */

package libp2putils

import (
	"context"
	"io"

	"fmt"
	"log"

	mrand "math/rand"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ipfs/go-datastore"
	leveldb "github.com/ipfs/go-ds-leveldb"
	libp2p "github.com/libp2p/go-libp2p"
	dht "github.com/libp2p/go-libp2p-kad-dht"
	"github.com/libp2p/go-libp2p/core/crypto"
	"github.com/libp2p/go-libp2p/core/host"
	"github.com/libp2p/go-libp2p/core/peer"
	discovery "github.com/libp2p/go-libp2p/p2p/discovery/routing"
	noise "github.com/libp2p/go-libp2p/p2p/security/noise"
	"github.com/luxfeerre/thesis-chatapp/chatapp/configuration"
	"github.com/luxfeerre/thesis-chatapp/chatapp/ethereumService"
	"github.com/luxfeerre/thesis-chatapp/chatapp/state"
	multiaddr "github.com/multiformats/go-multiaddr"
)

/*
* Description: Creates a host with a seed to generate a key
* that listens on the port specfified with the TCp protocol
* on a generated multi address using the noise portocol for security
* and default resrouce manager.
* @dev Return a Host struct and nill if success or an error message on failure.
* @return value of 'Host struct and error'
 */
func NewHost(seed int64, port int) (host.Host, error) {
	// Create variable to hold a random value to generate cryptographic keys
	var r io.Reader

	mrand.Seed(seed)
	r = mrand.New(mrand.NewSource(seed))

	// Generate private key to use an identity for the host
	private, _, err := crypto.GenerateEd25519Key(r)

	if err != nil {
		return nil, err
	}

	// Create new multiaddress which uses set port or random port depnding on value of port
	addrTcp, _ := multiaddr.NewMultiaddr(fmt.Sprintf("/ip4/0.0.0.0/tcp/%d", port))
	// Return a new host to the user
	return libp2p.New(
		libp2p.Identity(private),
		libp2p.Security(noise.ID, noise.New),
		libp2p.ListenAddrs(addrTcp),
		libp2p.DefaultTransports,
		libp2p.DefaultResourceManager,
	)
}

/*
* Description: Create a RoutingDiscovery struct populated with values needed to create a
* DHT routing table that can be used for discovery.
* Byt first creatin a new data store to hold values for the DHT router
* and then set the options to bootnode, the database, bootpeers, protocol prefix to make it uniqe
* and a function to handle if the DHT routing table it empty.
* Also DisableProviders it set since it wont need to use Internet connected bootstrap nodes.
* Then it creates the routing table and activates it.
* @dev Return a RoutingDiscovery struct which is a DHT bootnode routing table.
* @return value of 'RoutingDiscovery struct'
 */
func NewDht(ctxBackground context.Context, host host.Host, bootPeers []peer.AddrInfo,
	instance *state.State, transactionData *bind.TransactOpts,
	conf *configuration.Configuration) *discovery.RoutingDiscovery {
	// Create a datastore to hold value for the dht router
	var ds datastore.Batching
	ds, err := leveldb.NewDatastore("", nil)
	if err != nil {
		log.Printf("Database error\n")
	}

	// Create an array to hold dht options
	options := []dht.Option{
		// Set it to bootstrap node
		dht.Mode(dht.ModeServer),
		// Set a new database
		dht.Datastore(ds),
		// Set upp bootstrap peers
		dht.BootstrapPeers(bootPeers...),
		dht.DisableProviders(),
		dht.ProtocolPrefix("/test"),
		// Function used if DHT routing table is empty
		dht.BootstrapPeersFunc(func() []peer.AddrInfo {
			conf.DiscoveryPeers = make(configuration.AddressList, 0)

			_ = ethereumService.ModDhtList(instance, transactionData, fmt.Sprintf("%s/p2p/%s", host.Addrs()[0], host.ID().Pretty()))
			dhtList := ethereumService.GetDhtList(instance, transactionData)

			for _, peer := range dhtList {
				conf.DiscoveryPeers.Set(peer)
			}
			return conf.DiscoveryPeers
		}),
	}

	// Create dht node
	kaddht, err := dht.New(ctxBackground, host, options...)
	// Check if error accored and then return to erro to let calle handle it
	if err != nil {
		return nil
	}

	// Check if error accored and then return to erro to let calle handle it
	if err = kaddht.Bootstrap(ctxBackground); err != nil {
		return nil
	}

	// return routing table based on dht
	return discovery.NewRoutingDiscovery(kaddht)
}
