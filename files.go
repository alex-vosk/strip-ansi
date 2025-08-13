package main

import (
	"os"
)

type fileMaker func(string) (*os.File, error)

func makeFile(filename string, defaultFile *os.File, fm fileMaker) (rv *os.File, err error) {
	if filename == stdInOut {
		return defaultFile, nil
	}
	rv, err = fm(filename)
	return
}
