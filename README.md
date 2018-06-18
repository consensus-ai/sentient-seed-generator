# User Guide

This is a fully offline wallet generator for Sentient Network. It will generate the following two values:

* Backup Seed
  * *ATTENTION!* Never give this seed to anyone or type it on the web
  * Make sure to store it in a safe place offline
  * This seed can be used to restore your wallet using the official wallet software
  * Make sure to follow instructions for verifying the official wallet software signature before you type this seed into it
* Public Address
  * This is the first public address that is generated using the seed above
  * You can receive funds by sharing this address with the payer
  * If you're one of the participants who will be included in the genesis block, you will need to copy this address into your Consensus account

# Development Setup

## 1. Install the repo and dependencies
```bash
go get -u github.com/consensus-ai/sentient-seed-generator/...
go get -u github.com/asticode/go-astilectron/...
go get -u github.com/asticode/go-astilectron-bootstrap/...
go get -u github.com/asticode/go-astilectron-bundler/...
```

## 2. Bundle the app for your environment
```bash
cd $GOPATH/src/github.com/consensus-ai/sentient-seed-generator
astilectron-bundler -v
```

## 3. Open the app
Open the bundle for your OS located in the `output` directory.
