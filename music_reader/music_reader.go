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

func (mr *MusicReader) GetMusic() []string {
	return mr.music
}

func convertDirContents(dirReader dirreader.DirReader) []string {
	music := make([]string, 0)

	rootPath := dirReader.GetPath()

	for _, dir := range dirReader.GetContents() {
		if dir == rootPath {
			continue
		}

		if strings.HasSuffix(dir, "DS_Store") {
			continue
		}

		music = append(music, strings.Split(dir, rootPath)[1])
	}

	return music
}
