package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/pborman/ansi"
)

func main() {

	args := parseCliArguments()

	// create input file
	input, err := makeFile(args.InputName, os.Stdin, os.Open)
	exitOnError(err)
	defer input.Close()

	// created output file
	output, err := makeFile(args.OutputName, os.Stdout, os.Create)
	exitOnError(err)
	defer output.Close()

	// process the input and write it to the output
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		text := scanner.Text()
		clean, _ := ansi.Strip([]byte(text))
		fmt.Fprintf(output, "%s\n", clean)
	}
}

func exitOnError(err error) {
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %s", err)
		os.Exit(1)
	}
}
