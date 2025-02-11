#!/bin/bash

# To download
wget https://go.dev/dl/go1.23.6.linux-armv6l.tar.gz

sudo rm -rf /usr/local/go
sudo tar -C /usr/local -xzf go1.23.6.linux-armv6l.tar.gz

export PATH=$PATH:/usr/local/go/bin

go version

