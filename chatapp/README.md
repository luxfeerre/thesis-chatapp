# **chatapp code repository**

## **How to use it**

## **Depedencie list used for this project**
Operating System: Ubuntu 20.04
Golang: 18.x or 19.x
Ethereum Node: Goquorum 22.7.1
Puppeth

## **Environment set up**

### **Operating System**
Install Golang:

Download the golang package from: https://go.dev/doc/install

Then navigate to the Downloads folder:

**cd ~/Downloads/**

Followed by executing the command with your golang package:

**sudo rm -rf /usr/local/go && sudo tar -C /usr/local -xzf go1.19.1.linux-amd64.tar.gz**

Then set the path for the golang package:

**echo "export PATH=$PATH:/usr/local/go/bin" >> ~/.bash_aliases**

Test the isntallation by using the command:

**go version**

### **Ethereum Node**
Install Goquorum by following these steps:

Download Goquorum from this site: https://github.com/ConsenSys/quorum/releases

Then untar the package by applying the tar command on 
your package as shown below by using the command:

**tar xfv geth_v22.7.1_linux_amd64.tar.gz**

Then move the geth package /usr/local/bin by using the command:

**sudo mv geth /usr/local/bin**

Then try out the geht package by writing:

**geth version**

Then setup the genesis block by using puppeth.
So first the repository for the main branch of Ethereum geth is added with the command below:

**sudo add-apt-repository -y ppa:ethereum/ethereum && sudo apt update**

Then install puppeth with the command:

**sudo apt install puppeth -y**

After this step setup an Ethereum POA chain by first:

Executing the geth command to create a new account:

**geth account new --datadir node_poa**

Executing the puppeth command:

**puppeth**
Then follow the instructions and choose the Clique consensus alogrith with 
the account create in the previouse step and the time it should take to generate a block
was 1 second for this application and write out the genesis block to a file.

Then modify the script goqourum.sh which can be found in the repository.
Change the --http.addr xxx.xxx.xxx.xxx to a valid IPv4 address of your Ethereum node.
Then execute the script goqourum.sh to start the node:

**goqourum.sh**

### **Upload smart contract**


### **Testground setup**

### **Remote Tracer set up**

## **Program configuration**
