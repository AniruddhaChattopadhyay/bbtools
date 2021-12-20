.PHONY: all fetch extract set_path source_profile checking_installation install_bbtools clean
all:fetch extract set_path source_profile checking_installation install_bbtools clean

SHELL := /bin/bash

fetch:
		@echo "Fetching the go installer"
		curl -OL https://golang.org/dl/go1.17.5.linux-amd64.tar.gz

extract:
		@echo "installing go in /usr/local"
		sudo tar -C /usr/local -xvf go1.17.5.linux-amd64.tar.gz

set_path:
		@echo "setting path to go binary in ~/.profile"
		echo "export PATH=$$PATH:/usr/local/go/bin" >> ~/.profile
		echo "export PATH=$$PATH:$$HOME/go/bin/bbtools" >> ~/.profile
		
source_profile:
		@echo "Using source to reload ~/.profile"
		source ~/.profile

checking_installation:
		@echo "checking go version"
		go version

install_bbtools:
		@echo "checking go version"
		go install
clean:
		@echo "Cleaning up..."
		rm *.gz