#!/bin/bash

echo "Downloading Go 1.9"

# Go to the $(whoami)'s home dir
cd ~/ && pwd

#et the latest version
wget https://dl.google.com/go/go1.9.2.linux-amd64.tar.gz

# Untar to the local directory and remove the tarball
tar -xzvf go1.9.2.linux-amd64.tar.gz && rm -f go1.9.2.linux-amd64.tar.gz

# Export these variables for after this script is complete
export PATH=$PATH:/home/$(whoami)/go/bin
export GOPATH=/home/$(whoami)/go/third-party
export GOBIN=/home/$(whoami)/go/bin

# Have the $(whoami) manually enter the variables
echo "Place the following in your .bashrc:
export PATH=$PATH:/home/$(whoami)/go/bin
export GOPATH=/home/$(whoami)/go/third-party
export GOBIN=/home/$(whoami)/go/bin
"


