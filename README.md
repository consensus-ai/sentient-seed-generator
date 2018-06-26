# User Guide

This is a wallet generator for the Sentient Network that can be run offline. It will generate the following two values:

If you're using one of the official releases, make sure you verify the signature before running it locally.

* Backup Seed
  * This seed can be used to restore your wallet using the official wallet software;
  * *ATTENTION!* Never give this seed to anyone, or type it on the website;
  * Make sure to store it in a safe place offline.
* Public Address
  * This is the first public address that is generated using the seed above;
  * You can receive funds by sharing this address with a payer.

When generating your wallet, you can go fully offline (disable your Wi-Fi, unplug your Ethernet cord, etc).


# Development Setup

This project uses https://github.com/zserge/webview

## 1. Install the repo and dependencies
```bash
go get -u github.com/consensus-ai/sentient-seed-generator/...
go get -u github.com/zserge/webview/...
go get -u github.com/NebulousLabs/entropy-mnemonics/...
go get -u github.com/NebulousLabs/fastrand/...
go get -u github.com/NebulousLabs/merkletree/...
go get -u github.com/gopherjs/gopherjs/...
```

## 2. Build the native apps
If you're on Linux, you will have to install `gtk+-3.0` and `webkit2gtk-4.0` like so:
```bash
sudo apt-get install build-essential libgtk-3-dev
sudo apt-get install webkit2gtk-4.0
```

To build the app, run the following
```bash
cd $GOPATH/src/github.com/consensus-ai/sentient-seed-generator
go build
```

## 2.1 Build the web app
```bash
cd $GOPATH/src/github.com/consensus-ai/sentient-seed-generator/gopherjs
gopherjs build main.go
```

## 3. Open the app
Just execute the binary that was built in the previous step, or the sentient-seed-generator.html file

## 4. Making changes

All the static assets such as JS/CSS files are embedded into the binary in the form of bindata compiled using:
https://github.com/jteeuwen/go-bindata

If you make changes to any of the JS/CSS files, you will have to re-create the bindata like so:
```bash
go-bindata -o assets.go assets/react/... assets/main.css
```
