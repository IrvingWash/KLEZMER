package dirreader

import (
	"log"
	"os"
	"path/filepath"
)

type DirReader struct {
	path     string
	contents []string
}

func New(path string) *DirReader {
	return &DirReader{
		path:     path,
		contents: make([]string, 0),
	}
}

func (dirReader *DirReader) GetContents() []string {
	return dirReader.contents
}

func (dirReader *DirReader) GetPath() string {
	return dirReader.path
}

func (dirReader *DirReader) Read() {
	err := filepath.Walk(
		dirReader.path,
		func(path string, info os.FileInfo, err error) error {
			if err != nil {
				return err
			}

			dirReader.contents = append(dirReader.contents, path)

			return nil
		},
	)

	if err != nil {
		log.Fatal(err)
	}
}
