package main

import (
	"flag"

	"github.com/asticode/go-astilectron"
	"github.com/asticode/go-astilectron-bootstrap"
	"github.com/asticode/go-astilog"
	"github.com/pkg/errors"
)

// Constants
const htmlAbout = `Consensus Seed Generator`

// Vars
var (
	AppName string
	BuiltAt string
	debug   = flag.Bool("d", false, "enables the debug mode")
	w       *astilectron.Window
)

func main() {

	// Init
	flag.Parse()
	astilog.FlagInit()

	// Run bootstrap
	astilog.Debugf("Running app built at %s", BuiltAt)
	if err := bootstrap.Run(bootstrap.Options{
		Asset:    Asset,
		AssetDir: AssetDir,
		AstilectronOptions: astilectron.Options{
			AppName:            AppName,
			AppIconDarwinPath:  "resources/icon.icns",
			AppIconDefaultPath: "resources/icon.png",
		},
		Debug:         *debug,
		RestoreAssets: RestoreAssets,
		MenuOptions: []*astilectron.MenuItemOptions{
			{
				Label: astilectron.PtrStr("File"),
				SubMenu: []*astilectron.MenuItemOptions{
					{Role: astilectron.MenuItemRoleQuit},
				},
			},
			{
				Label: astilectron.PtrStr("Edit"),
				SubMenu: []*astilectron.MenuItemOptions{
					{Role: astilectron.MenuItemRoleCopy},
				},
			},
		},
		Windows: []*bootstrap.Window{{
			Homepage:       "index.html",
			MessageHandler: handleMessages,
			Options: &astilectron.WindowOptions{
				BackgroundColor: astilectron.PtrStr("#333"),
				Center:          astilectron.PtrBool(true),
				Height:          astilectron.PtrInt(600),
				Width:           astilectron.PtrInt(550),
				Resizable:       astilectron.PtrBool(false),
				Fullscreenable:  astilectron.PtrBool(false),
				Maximizable:     astilectron.PtrBool(false),
				MinimizeOnClose: astilectron.PtrBool(false),
			},
		}},
	}); err != nil {
		astilog.Fatal(errors.Wrap(err, "running bootstrap failed"))
	}
}
