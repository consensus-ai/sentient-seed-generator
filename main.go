package main

import (
	"github.com/zserge/webview"
)

type SeedGenerator struct {
	Seed    string `json:"seed"`
	Address string `json:"address"`
}

func (c *SeedGenerator) Generate() {
	wallet, err := GenerateWallet()
	if err != nil {
		return
	}

	c.Seed = wallet.SeedPhrase
	c.Address = wallet.FirstPubAddr
}

func main() {
	w := webview.New(webview.Settings{
		Title:     "Sentient Seed Generator",
		Width:     550, // px
		Height:    600, // px
		Resizable: false,
		Debug:     false,
	})
	defer w.Exit()

	w.Dispatch(func() {
		// pre-generate seed to be displayed on startup
		seedGenerator := SeedGenerator{}
		seedGenerator.Generate()

		// Inject controlle
		w.Bind("seedGenerator", &seedGenerator)

		// Inject CSS
		w.InjectCSS(string(MustAsset("assets/main.css")))

		// Inject react framework and app UI code
		loadUIFramework(w)
	})
	w.Run()
}
