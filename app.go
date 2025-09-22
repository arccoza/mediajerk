package main

import (
	"context"
	"fmt"

	"github.com/wailsapp/wails/v2/pkg/runtime"
)

// App struct
type App struct {
	ctx context.Context
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{}
}

// startup is called when the app starts. The context is saved
// so we can call the runtime methods
func (a *App) startup(ctx context.Context) {
	a.ctx = ctx
}

// Greet returns a greeting for the given name
func (a *App) Greet(name string) string {
	return fmt.Sprintf("Hello %s, It's show time!", name)
}

func (a *App) SelectFiles() ([]string, error) {
	files, err := runtime.OpenMultipleFilesDialog(a.ctx,
		runtime.OpenDialogOptions{
			Title: "Select Media Files",
			Filters: []runtime.FileFilter{
				{
					DisplayName: "Video Files",
					Pattern:     "*.mp4;*.mkv;*.avi;*.mov;*.wmv;*.flv;*.webm",
				},
				{
					DisplayName: "All Files",
					Pattern:     "*.*",
				},
			},
		})
	return files, err
}
