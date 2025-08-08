package main

import (
	"flag"
)

const stdInOut = "-"

type Args struct {
	InputName  string
	OutputName string
}

func parseCliArguments() Args {
	rv := Args{}

	flag.StringVar(&rv.InputName, "input", stdInOut, "input file name; use '-' for stdin")
	flag.StringVar(&rv.InputName, "i", stdInOut, "(short) "+"input file name; use '-' for stdin")
	flag.StringVar(&rv.OutputName, "output", stdInOut, "output filename; use '-' for stdout")
	flag.StringVar(&rv.OutputName, "o", stdInOut, "(short) "+"output filename; use '-' for stdout")
	flag.Parse()

	return rv
}
