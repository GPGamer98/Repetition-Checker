package main

import (
	"embed"

	"github.com/wailsapp/wails/v2"
	// "github.com/wailsapp/wails/v2/pkg/menu"
	// "github.com/wailsapp/wails/v2/pkg/menu/keys"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
)

//go:embed all:frontend/build
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	// AppMenu := menu.AppMenu()
	// AboutMenu := AppMenu.SubMenu.AddSubmenu("Informazioni")
	// AboutMenu.AddText("Informazioni", keys.CmdOrCtrl("i"), Info)

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "Contatore di Apparenze | Cogito Ergo Vet",
		Width:  1080,
		Height: 720,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 36, G: 36, B: 36, A: 1},
		OnStartup:        app.startup,
		OnDomReady: 	  app.domReady,
		Bind: []interface{}{
			app,
		},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
