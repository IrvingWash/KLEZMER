package main

import (
	argparser "github.com/IrvingWash/klezmer/arg_parser"
	dirreader "github.com/IrvingWash/klezmer/utils/dir_reader"
)

func main() {
	args := argparser.New()

	dirReader := dirreader.New(args.MusicPath)

	dirReader.Read()
}
