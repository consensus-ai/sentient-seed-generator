package main

import (
	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
)

// handleMessages handles messages
func handleMessages(_ *astilectron.Window, m bootstrap.MessageIn) (interface{}, error) {
	var payload WalletSeed
	var err error
	switch m.Name {
	case "generate":
		wallet, err := GenerateWallet()
		if err != nil {
			return nil, err
		}

		payload = WalletSeed{
			Seed:      wallet.SeedPhrase,
			Addresses: []string{wallet.FirstPubAddr},
		}
	}

	return &payload, err
}

type WalletSeed struct {
	Seed      string   `json:"seed"`
	Addresses []string `json:"addresses"`
}
