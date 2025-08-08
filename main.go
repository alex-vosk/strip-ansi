package main

import (
	"bufio"
	"fmt"

	"github.com/pborman/ansi"
)

func main() {

	args := parseCliArguments()

	input, output := makeInputOutputFiles(args)
	defer input.Close()
	defer output.Close()

	// process the input and write it to the output
	scanner := bufio.NewScanner(input)
	for scanner.Scan() {
		text := scanner.Text()
		clean, _ := ansi.Strip([]byte(text))
		fmt.Fprintf(output, "%s\n", clean)
	}
}
