package dirreader

import (
	"log"
	"os"
	"path/filepath"
)

type DirReader struct {
	Path     string
	Contents []string
}

func New(path string) *DirReader {
	return &DirReader{
		Path:     path,
		Contents: make([]string, 0),
	}
}

func (dirReader *DirReader) Read() {
	err := filepath.Walk(
		dirReader.Path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			dirReader.Contents = append(dirReader.Contents, path)

			return nil
		},
	)

	if err != nil {
		log.Fatal(err)
	}
}
