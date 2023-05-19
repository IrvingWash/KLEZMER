package musicreader

import (
	"strings"

	dirreader "github.com/IrvingWash/klezmer/utils/dir_reader"
)

type MusicReader struct {
	Music []string
}

func New(path string) *MusicReader {
	dirReader := dirreader.New(path)

	dirReader.Read()

	music := convertDirContents(*dirReader)

	return &MusicReader{
		Music: music,
	}
}

func convertDirContents(dirReader dirreader.DirReader) []string {
	music := make([]string, 0)

	for _, dir := range dirReader.Contents {
		if dir == dirReader.Path {
			continue
		}

		if strings.HasSuffix(dir, "DS_Store") {
			continue
		}

		music = append(music, dir)
	}

	return music
}
