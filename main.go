package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/pborman/ansi"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		text := scanner.Text()
		clean, _ := ansi.Strip([]byte(text))
		fmt.Fprintf(os.Stdout, "%s\n", clean)
	}
}
