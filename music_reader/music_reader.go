package musicreader

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	dirreader "github.com/IrvingWash/klezmer/utils/dir_reader"
)

type MusicReader struct {
	music []string
}

func New(path string) *MusicReader {
	dirReader := dirreader.New(path)

	dirReader.Read()

	music := convertDirContents(*dirReader)

	return &MusicReader{
		music: music,
	}
}

func (mr *MusicReader) GetRandom() {
	rand.Seed(time.Now().Unix())

	fmt.Println(mr.music[rand.Intn(len(mr.music))])
}

func convertDirContents(dirReader dirreader.DirReader) []string {
	music := make([]string, 0)

	for _, dir := range dirReader.GetContents() {
		if dir == dirReader.GetPath() {
			continue
		}

		if strings.HasSuffix(dir, "DS_Store") {
			continue
		}

		music = append(music, dir)
	}

	return music
}
