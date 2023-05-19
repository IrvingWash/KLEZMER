package argparser

import (
	"os"
	"strings"

	"github.com/IrvingWash/klezmer/utils"
)

type ArgParser struct {
	musicPath string
	decide    bool
}

func New() *ArgParser {
	args := os.Args[1:]

	var musicPath string
	var decide bool

	for _, arg := range args {
		if strings.HasPrefix(arg, "music-path=") {
			musicPath = utils.ParseTilde(
				strings.Split(arg, "=")[1],
			)
		}

		if arg == "decide" {
			decide = true
		}
	}

	return &ArgParser{
		musicPath: musicPath,
		decide:    decide,
	}
}

func (argParser *ArgParser) GetMusicPath() string {
	return argParser.musicPath
}

func (argParser *ArgParser) GetDecide() bool {
	return argParser.decide
}
