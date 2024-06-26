package backend

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3" // Add this line
	"log"
	"os"
	"path/filepath"
)

type FileIndexer struct {
	DB *sql.DB
}

func NewFileIndexer() *FileIndexer {
	database, err := sql.Open("sqlite3", "./fileindexer.db")
	if err != nil {
		log.Fatal(err)
	}

	statement, err := database.Prepare("CREATE TABLE IF NOT EXISTS files (id INTEGER PRIMARY KEY, path TEXT)")
	if err != nil {
		log.Fatal(err)
	}
	statement.Exec()

	return &FileIndexer{DB: database}
}

func (fi *FileIndexer) IndexFiles() error {
	err := filepath.Walk(".", func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() {
			statement, err := fi.DB.Prepare("INSERT INTO files (path) VALUES (?)")
			if err != nil {
				return err
			}
			statement.Exec(path)
		}
		return nil
	})

	if err != nil {
		return err
	}

	return nil
}

func (fi *FileIndexer) SearchFiles(query string) ([]string, error) {
	rows, err := fi.DB.Query("SELECT path FROM files WHERE path LIKE ?", "%"+query+"%")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var results []string
	for rows.Next() {
		var path string
		rows.Scan(&path)
		results = append(results, path)
	}

	return results, nil
}
