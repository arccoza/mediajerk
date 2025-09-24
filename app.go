package main

import (
	"context"
	"fmt"
	"mediajerk/backend/non"
	"os"
	"path/filepath"

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

func (a *App) SelectFiles(options FileDialogOptions) ([]FileInfo, error) {
	homeDir, _ := os.UserHomeDir()

	// Convert our FileFilter to runtime.FileFilter
	runtimeFilters := make([]runtime.FileFilter, len(options.Filters))
	for i, filter := range options.Filters {
		runtimeFilters[i] = runtime.FileFilter{
			DisplayName: filter.DisplayName,
			Pattern:     filter.Pattern,
		}
	}

	// Default filters if none provided
	if len(runtimeFilters) == 0 {
		runtimeFilters = []runtime.FileFilter{
			{
				DisplayName: "All Files",
				Pattern:     "*.*",
			},
		}
	}

	files, err := runtime.OpenMultipleFilesDialog(a.ctx,
		runtime.OpenDialogOptions{
			Title:            non.Zero(options.Title, "Select Files"),
			Filters:          runtimeFilters,
			DefaultDirectory: homeDir,
		})

	fileList := make([]FileInfo, 0, len(files))
	for _, path := range files {
		info, err := os.Stat(path)
		if err != nil {
			return nil, err
		}

		fileList = append(fileList, FileInfo{info.Name(), filepath.Ext(path), filepath.Dir(path), path, string(filepath.Separator), int(info.ModTime().UnixMilli())})
	}

	// fmt.Println(fileList)
	return fileList, err
}

func (a *App) FilepathJoin(elem ...string) string {
	return filepath.Join(elem...)
}

type FileFilter struct {
	DisplayName string `json:"displayName"`
	Pattern     string `json:"pattern"`
}

type FileDialogOptions struct {
	Title   string       `json:"title"`
	Filters []FileFilter `json:"filters"`
}

type FileInfo struct {
	Name         string `json:"name"`
	Ext          string `json:"ext"`
	Dir          string `json:"dir"`
	Path         string `json:"path"`
	Seperator    string `json:"seperator"`
	LastModified int    `json:"lastModified"`
}
