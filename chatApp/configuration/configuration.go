/*
* @Author: Joakim
* @Year: 2022
* @Description:
* File used to setup the configuration for the program by reading in data
* to variables trhough the flag strcture.
 */

package configuration

import (
	"flag"
	"strings"

	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/libp2p/go-libp2p-core/peer"
	"github.com/luxfeerre/thesis-chatapp/chatapp/ethereumService"
	"github.com/luxfeerre/thesis-chatapp/chatapp/state"
)

// Used to create a struct which contains values used by the program
type Configuration struct {
	Port           int
	Seed           int64
	DiscoveryPeers AddressList
	UserName       string
	Password       string
	Node           string
	Room           string
}

// Used to create a new type of slice
type AddressList []peer.AddrInfo

/*
* Description: This iterates over an address list of peers
* and concatantes it into a string.
* @dev Return a concatnated string
* @return value of 'string'
 */
func (al *AddressList) String() string {
	strs := make([]string, len(*al))
	for i, address := range *al {
		strs[i] = address.String()
	}
	return strings.Join(strs, ",")
}

/*
* Description: Sets a new multi-address to the addresslist of peers.
* @dev return error on failure nil on success.
* @return value of 'error'
 */
func (al *AddressList) Set(value string) error {
	address, err := peer.AddrInfoFromString(value)
	if err != nil {
		return err
	}
	*al = append(*al, *address)
	return nil
}

/*
* Description: Sets up a strcut with either pre configured values or read in values from the user.
* @dev return a pointer to a Configuration struct.
* @return value of 'Pointer to configuration struct'
 */
func Config(seed int64, name string) *Configuration {
	configuration := Configuration{}
	flag.Int64Var(&configuration.Seed, "seed", seed, "Seed value for generating a determenistik privatekey")
	flag.Var(&configuration.DiscoveryPeers, "peer", "Peer multiaddress for peer discovery")
	flag.IntVar(&configuration.Port, "port", 0, "")
	flag.StringVar(&configuration.UserName, "username", name, "username used for registration to use in chat. will try and use wallet if empty")
	flag.StringVar(&configuration.Password, "password", "secret", "password used to either create or open a wallet")
	flag.StringVar(&configuration.Node, "node", "http://192.168.2.12:8552", "Node use to connect to ethereum network")
	flag.StringVar(&configuration.Room, "room", "test", "Room name for pubsub")
	flag.Parse()

	return &configuration
}

/*
* Description: Sets up a boot node list by using the Ethereum blockchain smart contract and
* adds them to the address list of peers in the Configuration struct.
* @dev return nothing.
* @return value of 'nothing'
 */
func ConfBootNodeList(instance *state.State, transactionData *bind.TransactOpts, conf *Configuration) {
	dhtList := ethereumService.GetDhtList(instance, transactionData)
	for _, peer := range dhtList {
		conf.DiscoveryPeers.Set(peer)
	}
}
