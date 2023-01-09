# Testchat repository

It will use Testground to execute the main.go program found in this repository.


## **Depedencie list used for this project**

```
Operating System: Ubuntu 20.04
Golang: 18.x or 19.x
Docker
Testground 0.5.3
```

## **How to use it**

## **Environment set up**
First follow the instructions here to set up the environment: 

* https://github.com/luxfeerre/thesis-chatapp/blob/main/chatapp/README.md

Also, use the instructions in **Operating System** from the link to setup the ground dependencies for the Testground node.

Then do the same for the Metrics collection node, but on a seperate node.

Then continue to the subsection Testground or Metrics collection as shown below.

<details><summary>Testground</summary>
<p>

### Testground installation and set up instructions
Follow the instructions here to set up testground:

* https://github.com/testground/testground

Then copy this directory: testchat to the testground/plans/chatapp directory

Next step test the plat by executing ./test.sh 1

This will execute one node aginst a running testground instance.

</p>
</details>

<details><summary>Metrics collection</summary>
<p>

### Metrics collection installation and set up instructions

Follow the instructions here to set up the metrics collection node:

* https://github.com/libp2p/go-libp2p-pubsub-tracer


</p>
</details>
