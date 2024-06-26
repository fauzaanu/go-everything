package backend

import (
	"context"
)

// App struct
type App struct {
	ctx     context.Context
	indexer *FileIndexer
}

// NewApp creates a new App application struct
func NewApp() *App {
	return &App{
		indexer: NewFileIndexer(),
	}
}

// Startup is called at application Startup
func (a *App) Startup(ctx context.Context) {
	a.ctx = ctx
}

// IndexFiles indexes all files on the PC
func (a *App) IndexFiles() string {
	err := a.indexer.IndexFiles()
	if err != nil {
		return "Error indexing files: " + err.Error()
	}
	return "Indexing complete"
}

// SearchFiles searches for files in the database
func (a *App) SearchFiles(query string) []string {
	results, err := a.indexer.SearchFiles(query)
	if err != nil {
		return []string{"Error searching files: " + err.Error()}
	}
	return results
}
