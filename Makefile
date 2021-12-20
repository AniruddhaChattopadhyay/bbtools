.PHONY: all fetch extract checking_installation install_bbtools clean
all:fetch extract checking_installation install_bbtools clean

SHELL := /bin/bash

fetch:
		@echo "Fetching the go installer"
		curl -OL https://golang.org/dl/go1.17.5.linux-amd64.tar.gz

extract:
		@echo "installing go in /usr/local"
		sudo tar -C /usr/local -xvf go1.17.5.linux-amd64.tar.gz

checking_installation:
		@echo "checking go version"
		go version

install_bbtools:
		@echo "checking go version"
		/usr/local/go/bin/go install
clean:
		@echo "Cleaning up..."
		rm *.gz