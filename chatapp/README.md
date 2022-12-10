# **Chatapp repository with set up instructions**

## **Depedencie list used for this project**
```
Operating System: Ubuntu 20.04
Golang: 18.x or 19.x
Ethereum Node: Goquorum 22.7.1
Puppeth
Docker
```

## **How to use it**

## **Environment set up**

<details><summary>Operating System set up</summary><p>
	
### **Operating System set up**
	
#### Install Golang:
	
Download the golang package from: https://go.dev/doc/install
	
Then navigate to the Downloads folder:
	
```
cd ~/Downloads/
```
	
Followed by executing the command with your golang package:
	
```
sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.19.1.linux-amd64.tar.gz
```
	
Then set the path for the golang package:
	
```
echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bash_aliases
```
	
Test the isntallation by using the command:
	
```
go version
```
	
#### Install Docker:
	
Follow the instructions on https://docs.docker.com/engine/install/ubuntu/.
	
Then optionaly add a user other then root to the Docker group so it can execute docker
by following the instructions here: https://docs.docker.com/engine/install/linux-postinstall/
	
Then test the docker application out by running:
	
```
docker run hello-world
```
</p>
</details>

<details><summary>Ethereum full node set up</summary>
<p>
	
### **Ethereum full node set up**

Be sure to install Ubuntu 20.04 or higher and the dependencies described above and then install the software from this part.

#### Install Goquorum by following these steps:

Download Goquorum from this site: https://github.com/ConsenSys/quorum/releases

Then pack up the package by using the tar command on 
your package as shown below by using the command:
```
tar xfv geth_v22.7.1_linux_amd64.tar.gz
```
Then move the geth package /usr/local/bin by using the command:
```
sudo mv geth /usr/local/bin
```
Then try out the geht package by writing:
```
geth version
```
### Setup a genesis file and start the full node:

Then setup the genesis block by using puppeth.
So first the repository for the main branch of Ethereum geth is added with the command below:
```
sudo add-apt-repository -y ppa:ethereum/ethereum && sudo apt update
```
Then install puppeth with the command:
```
sudo apt install puppeth -y
```
After this step setup an Ethereum POA chain by first:

First download the content from this repository: https://github.com/luxfeerre/thesis-chatapp/tree/main/chatapp/goqourum
Downloading the content with the directory from this repository which includes a support program main.go which can be used to get the private key for the account which we will create next. This key is needed when later uploading the smart contract to the chain.
Thier is also the program goqourum.sh which will be used to start up the Ethereum full node.

Then navigate to the new folder with the command:
```
cd goqourum
```
Executing the geth command to create a new account:
```
geth account new --datadir node_poa
```
This creates a new folder that contains the wallet for the new account.

Executing the puppeth command:
```
puppeth
```
* Set the Network Name to **node_poa** 
* Set consensus alogrithm to **Clique(PoA)**
* Set block creation time to **1 second** 
* Set the signer account to the account created in the previouse step 
* Set it to be a prefounded account and the network id to **101**(Otherwise the chain id will have to be changed in the program).
* Then choose to write out the **genesis file with the default name**.

Then modify the script goqourum.sh which can be found in the repository.
Change the --http.addr xxx.xxx.xxx.xxx to a valid IPv4 address of your Ethereum node.
Then execute the script goqourum.sh to start the node from the directory where the genesis file was created.

</p>
</details>

<details><summary>Upload smart contract</summary>
<p>
### **Upload smart contract**
Use the Ethereum full node or setup a new node and be sure to install Ubuntu 20.04 or higher and the dependencies described above and then install the software from this part.

Then download the entaier repository at: https://github.com/luxfeerre/thesis-chatapp/tree/main/chatapp

Then naviage to the chatapp folder through the command:
```
cd chatapp
```
Then use the script: 
```
bash recomplineSmartContract.sh
```
This will download two docker images and execute them to compile the smart contract state/State.sol.
It will also generate the Golang code used by the contract and the application.

Then on the Ethereum full node go to the directory goqourum.
Thier use the command:
```
go run main.go
```
This will print out the private key for the account which will be used to deploy the smart contract to the private Ethereum network.

Then on the node used to deploy the contract make sure you are in the chatapp directory.
Then modify the file utils/contract_deploy.go and change line with this content:

```
privateKey, err := crypto.HexToECDSA("811283a34a4429e520dc7f78bfc8be83fc756a6f79e823c91733b90210cd39f5")
```

Where "811283a34a4429e520dc7f78bfc8be83fc756a6f79e823c91733b90210cd39f5" is changed to the private key for your Ethereum network which was printed ou in the main program. 
Also chane the line with:
```
client, err := ethclient.Dial("http://192.168.2.12:8552")
```

To your Ethereum full node IPv4 address.

Make sure that the Ethereum full node is running and then use the shell script to deploy the contract with the command:
```
bash deployContract.sh
```
This writes out two lines and use the first one which is the address for the smart contract on your private blockchain.

Add this to the file /ethereumService/ethereumService.go and change the function:

```
func StateInstance(client *ethclient.Client) (*state.State, error) {
	address := common.HexToAddress("0x99ddD1DF9C9719294e8cD34B1FFCC6B03CfFeBB0")
	return state.NewState(address, client)
}
```

Where the 0x99ddD1DF9C9719294e8cD34B1FFCC6B03CfFeBB0should be changed to your smart contracts address.
</p>
</details>

<details><summary>Remote Tracer set up</summary>
<p>
### **Remote Tracer set up**

Set this up on the same node as the Ethereum full node or on another node.

Be sure to install Ubuntu 20.04 or higher and the dependencies described above and then install the software from this part.

</p>
</details>
