package main

import (
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
}
