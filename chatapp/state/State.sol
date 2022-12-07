/*
* @Author: Joakim
* @Year: 2022
* @Description:
* File is used to create the smart contract called State
* that is used to store values or retrive values from the blockchain
* smart contract.
 */

// SPDX-License-Identifier: MIT
pragma solidity ^0.8.16;

contract State {

    struct User {
        // Ethereum wallet msg.sender identity
        address owner;
        // A name, for example John Smith
        string identity;
        // peer ID derived from pubsub private/public key
        string peer_id;
        // IPFS address of peer
        string addr;
    }

    // Store dht boot nodes that join the network
    string[] public dhtList;

    // Initalize to the time to live for the dhtList to when the contract is first upploaded
    uint dhtListTtl = block.timestamp;

    // Storage for registered users
    User[] private user;

    // First mapps the peer id to a username
    mapping(string => string) private identityMapping;
    // Next maps the username to an address
    mapping(string => string) private addressMapping;
    // Last maps the senders identity to a struct of its registered values
    mapping(address => User) private ownerMapping;

     /*
     * Description: Used to register a user if they do not allready exist
     * @dev Return value of error and a transaction
     * @return value of 'Error or transaction'
     */
    function register(string calldata identity, string calldata peer_id, string calldata addr) public payable {
        // First it verifies that the user it not registered
        if (!verification()) {
            user.push(User(msg.sender, identity, peer_id, addr));
        }
        // Then it adds the user data to the hash maps
        identityMapping[peer_id] = identity;
        addressMapping[identity] = addr;
        ownerMapping[msg.sender] = User(msg.sender, identity, peer_id, addr);

        // Then last to the DHT bootstrap list
        modDhtList(addr);
    }

    /* 
     * Description: Function used to verify if a user exists and the view keyword ensures it is readonly.
     * @dev Return value of true if user exists and fale otherwsie
     * @return value of 'bool'
     */
    function verification() public view returns(bool) {
        if (keccak256(abi.encodePacked(ownerMapping[msg.sender].owner)) == keccak256(abi.encodePacked(msg.sender))) {
            return true;
        }
        return false;
    }

     /*
     * Description: Function used to retrive a address for the usersname and the view keyword ensures it is readonly.
     * @dev Return value string of users address
     * @return value of 'string'
     */
    function getUserAddress(string memory name) external view returns(string memory) {
        return addressMapping[name];
    }
    
     /*
     * Description: Function is used to retrive a username mapped to a users CID and the view keyword makes sure it is readonly.
     * @dev Return value of Username in the form a string
     * @return value of 'string'
     */
    function getUserName(string memory peer_id) external view returns(string memory) {
        return identityMapping[peer_id];
    }

     /*
     * Description: Function gets the dhtList which contains bootstrap nodes
     * used to join the DHT routing and discovery network.
     * @dev Return value of string list with Dht bootstrap nodes in it
     * @return value of 'list of strings'
     */
    function getDhtList() public view returns(string[] memory) {
        return dhtList;
    }

    /*
     * Description: Functions sets a time to live value by subtracting
     * it by the previouse value to gain a counter and if it lived longer then 60 second
     * then the boot list is reset and a new start time is set.
     * Addresses are pushed to the end of the list.
     * Else it checks if the list is longer then 3 nodes and if true it pops the oldest value.
     * Then it pushes the new bootnode to the front of the list.
     * This is to ensure that only active bootnodes are used and that the load is spread between peers. 
     * @dev Return transaction data and error string
     * @return value of 'transaction data and error'
     */
    function modDhtList(string calldata addr) public payable {
        uint checkTtl = block.timestamp - dhtListTtl;

        if (checkTtl > 60) {
            dhtList = new string[](0);
            dhtListTtl = block.timestamp;
            dhtList.push(addr);
        }
        else {
            if (dhtList.length > 2) {
                dhtList.pop();
            }
            dhtList.push(addr);
        }
    }
}