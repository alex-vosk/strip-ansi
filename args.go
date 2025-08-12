package main

import (
	"flag"
)

const stdInOut = "-"
const inputDesc = "input file name; use '-' for stdin"
const outputDesc = "output filename; use '-' for stdout"

type Args struct {
	InputName  string
	OutputName string
}

func parseCliArguments() Args {
	rv := Args{}

	flag.StringVar(&rv.InputName, "input", stdInOut, inputDesc)
	flag.StringVar(&rv.InputName, "i", stdInOut, inputDesc)
	flag.StringVar(&rv.OutputName, "output", stdInOut, outputDesc)
	flag.StringVar(&rv.OutputName, "o", stdInOut, outputDesc)
	flag.Parse()

	return rv
}
