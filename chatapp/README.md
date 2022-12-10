# **chatapp code repository**

## **How to use it**

## **Depedencie list used for this project**
Operating System: Ubuntu 20.04

Golang: 18.x or 19.x

Ethereum Node: Goquorum 22.7.1

Puppeth

Docker

## **Environment set up**

### **Operating System**
#### Install Golang:

Download the golang package from: https://go.dev/doc/install

Then navigate to the Downloads folder:

**cd ~/Downloads/**

Followed by executing the command with your golang package:

**sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.19.1.linux-amd64.tar.gz**

Then set the path for the golang package:

**echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bash_aliases**

Test the isntallation by using the command:

**go version**

#### Install Docker:

Follow the instructions on https://docs.docker.com/engine/install/ubuntu/.

Then optionaly add a user other then root to the Docker group so it can execute docker
by following the instructions here: https://docs.docker.com/engine/install/linux-postinstall/

Then test the docker application out by running:

**docker run hello-world**

### **Ethereum Node**

Be sure to install Ubuntu 20.04 or higher and the dependencies described above and then install the software from this part.

#### Install Goquorum by following these steps:

Download Goquorum from this site: https://github.com/ConsenSys/quorum/releases

Then untar the package by applying the tar command on 
your package as shown below by using the command:

**tar xfv geth_v22.7.1_linux_amd64.tar.gz**

Then move the geth package /usr/local/bin by using the command:

**sudo mv geth /usr/local/bin**

Then try out the geht package by writing:

**geth version**

### Setup a genesis file and start the full node:

Then setup the genesis block by using puppeth.
So first the repository for the main branch of Ethereum geth is added with the command below:

**sudo add-apt-repository -y ppa:ethereum/ethereum && sudo apt update**

Then install puppeth with the command:

**sudo apt install puppeth -y**

After this step setup an Ethereum POA chain by first:

First download the content from this repository: https://github.com/luxfeerre/thesis-chatapp/tree/main/chatapp/goqourum
Downloading the content with the directory from this repository which includes a support program main.go which can be used to get the private key for the account which we will create next. This key is needed when later uploading the smart contract to the chain.
Thier is also the program goqourum.sh which will be used to start up the Ethereum full node.

Then navigate to the new folder with the command:

**cd goqourum**

Executing the geth command to create a new account:

**geth account new --datadir node_poa**

This creates a new folder that contains the wallet for the new account.

Executing the puppeth command:

**puppeth**

First set the Network Name to **node_poa**, set consensus alogrithm to Clique(PoA),
set the signer account to the account created in the previouse step, set it to be a prefounded account,
block creation time to 1 second and the network id to 101.
Then choose to write out the genesis file.

Then modify the script goqourum.sh which can be found in the repository.
Change the --http.addr xxx.xxx.xxx.xxx to a valid IPv4 address of your Ethereum node.
Then execute the script goqourum.sh to start the node from the directory where the genesis file was created.

### **Upload smart contract**
Setup a new node and be sure to install Ubuntu 20.04 or higher and the dependencies described above and then install the software from this part.



### **Remote Tracer set up**

Set this up on the same node as the Ethereum full node or on another node.

Be sure to install Ubuntu 20.04 or higher and the dependencies described above and then install the software from this part.


