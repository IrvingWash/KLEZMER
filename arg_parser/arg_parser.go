package argparser

import (
	"os"
	"strings"

	"github.com/IrvingWash/klezmer/utils"
)

type ArgParser struct {
	MusicPath string
	Decide    bool
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
		MusicPath: musicPath,
		Decide:    decide,
	}
}
