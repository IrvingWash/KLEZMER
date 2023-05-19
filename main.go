package main

import (
	"fmt"

	argparser "github.com/IrvingWash/klezmer/arg_parser"
	musicreader "github.com/IrvingWash/klezmer/music_reader"
)

func main() {
	args := argparser.New()

	musicReader := musicreader.New(args.MusicPath)

	fmt.Println(musicReader.Music)
}
