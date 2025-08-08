package main

import (
	"fmt"
	"os"
)

func makeInputOutputFiles(args Args) (input *os.File, output *os.File) {
	// default return values
	input = os.Stdin
	output = os.Stdout

	if args.InputName != stdInOut {
		file, err := os.Open(args.InputName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}
		input = file
	}

	if args.OutputName != stdInOut {
		file, err := os.Create(args.OutputName)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Error: %s\n", err)
			os.Exit(1)
		}
		output = file
	}
	return
}
