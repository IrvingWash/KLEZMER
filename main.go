package main

import argparser "github.com/IrvingWash/klezmer/arg_parser"

func main() {
	args := argparser.New()

	println(args.MusicPath)
}
