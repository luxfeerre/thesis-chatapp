#! /bin/bash

# Used to deploy a test run with argument of how many nodes should be used
testground run single --plan=chatapp/testchat --testcase=chatapp/testchat --builder=docker:go --runner=local:docker --instances=$1 --disable-metrics true
