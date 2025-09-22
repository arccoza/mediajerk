package main

import (
	"embed"
	"fmt"
	"mediajerk/backend/tmdb"

	"github.com/wailsapp/wails/v2"
	"github.com/wailsapp/wails/v2/pkg/options"
	"github.com/wailsapp/wails/v2/pkg/options/assetserver"
	"github.com/wailsapp/wails/v2/pkg/options/linux"
	"github.com/wailsapp/wails/v2/pkg/options/windows"
)

//go:embed all:frontend/dist
var assets embed.FS

func main() {
	// Create an instance of the app structure
	app := NewApp()

	cl := tmdb.NewClient("eyJhbGciOiJIUzI1NiJ9.eyJhdWQiOiIxMTliM2M1M2JjNzY5N2M0NWQzZGQ1ZmYzYzE3ZDFjMCIsIm5iZiI6MTQ3OTM1NDkyOS43MzEsInN1YiI6IjU4MmQyYTMxOTI1MTQxMDk1ZDAwMjY2OSIsInNjb3BlcyI6WyJhcGlfcmVhZCJdLCJ2ZXJzaW9uIjoxfQ.r-UXqOQP4LUSX2At2uV-RKMFT9AltvyudBmLmN9mpDw")
	if result, err := cl.SearchTVByQuery("Scooby-Doo where are you"); err == nil {
		fmt.Println("result:", result)
	} else {
		fmt.Println("error:", err)
	}

	// Create application with options
	err := wails.Run(&options.App{
		Title:  "mediajerk",
		Width:  1280,
		Height: 768,
		// Frameless: true,
		AssetServer: &assetserver.Options{
			Assets: assets,
		},
		BackgroundColour: &options.RGBA{R: 27, G: 38, B: 54, A: 1},
		OnStartup:        app.startup,
		Bind: []interface{}{
			app,
		},
		Windows: &windows.Options{
			DisableFramelessWindowDecorations: false,
		},
		Linux: &linux.Options{},
	})

	if err != nil {
		println("Error:", err.Error())
	}
}
