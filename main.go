package main

import (
	"fmt"

	argparser "github.com/IrvingWash/klezmer/arg_parser"
	musicreader "github.com/IrvingWash/klezmer/music_reader"
	userconfig "github.com/IrvingWash/klezmer/user_config"
)

func main() {
	args := argparser.New()

	userConfig := userconfig.New()

	if args.GetMusicPath() != "" {
		userConfig.SetMusicPath(args.GetMusicPath())
	}

	musicReader := musicreader.New(userConfig.GetMusicPath())

	if args.GetDecide() {
		musicReader.GetRandom()

		return
	}

	if args.GetBackup() {
		for _, mus := range musicReader.GetMusic() {
			fmt.Println(mus)
		}
	}
}
