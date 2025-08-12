package main

import (
	"os"
)

func makeInputOutputFiles(args Args) (input *os.File, output *os.File, err error) {
	if args.InputName != stdInOut {
		input, err = os.Open(args.InputName)
		if err != nil {
			return
		}
	} else {
		input = os.Stdin
	}

	if args.OutputName != stdInOut {
		output, err = os.Create(args.OutputName)
		if err != nil {
			return
		}
	} else {
		output = os.Stdout
	}
	return
}
