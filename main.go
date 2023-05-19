package main

import (
	argparser "github.com/IrvingWash/klezmer/arg_parser"
	musicreader "github.com/IrvingWash/klezmer/music_reader"
)

func main() {
	args := argparser.New()

	musicReader := musicreader.New(args.MusicPath)

	if args.Decide {
		musicReader.GetRandom()

		return
	}
}
