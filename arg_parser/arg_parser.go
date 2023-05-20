package argparser

import (
	"os"
	"strings"

	"github.com/IrvingWash/klezmer/utils"
)

type ArgParser struct {
	musicPath string
	decide    bool
	backup    bool
}

func New() *ArgParser {
	args := os.Args[1:]

	var musicPath string
	var decide bool
	var backup bool

	for _, arg := range args {
		if strings.HasPrefix(arg, "music-path=") {
			musicPath = utils.ParseTilde(
				strings.Split(arg, "=")[1],
			)
		}

		if arg == "decide" {
			decide = true
		}

		if arg == "backup" {
			backup = true
		}
	}

	return &ArgParser{
		musicPath: musicPath,
		decide:    decide,
		backup:    backup,
	}
}

func (argParser *ArgParser) GetMusicPath() string {
	return argParser.musicPath
}

func (argParser *ArgParser) GetDecide() bool {
	return argParser.decide
}

func (argParser *ArgParser) GetBackup() bool {
	return argParser.backup
}
