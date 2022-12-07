#! /bin/bash
# Used to check that the testground is helthy and use correct environment parameters between restarts

# Set up environment stats that needs to be higher for the test
sudo sysctl -w vm.max_map_count=120000;

# Control health of testground
testground healthcheck --runner local:docker --fix;
