package main

import (
	"errors"
	"os"
	"testing"
)

func TestMakeInputOutputFiles(t *testing.T) {
	t.Run("default file", func(t *testing.T) {
		f, err := makeFile(stdInOut, os.Stderr, nil)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if f != os.Stderr {
			t.Error("expected file for input")
		}
	})

	t.Run("make file - success", func(t *testing.T) {
		const fileName = "/some/file/name"
		actualFileName := ""
		f, err := makeFile(fileName, os.Stderr, func(fn string) (*os.File, error) {
			actualFileName = fn
			return os.Stdin, nil
		})

		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if f != os.Stdin {
			t.Error("expected file for input")
		}
		if actualFileName != fileName {
			t.Error("unexpected filename for input")
		}
	})

	t.Run("make file - error", func(t *testing.T) {
		const fileName = "/some/file/name"
		const errMsg = "Can't open file"
		actualFileName := ""
		f, err := makeFile(fileName, os.Stderr, func(fn string) (*os.File, error) {
			actualFileName = fn
			return nil, errors.New(errMsg)
		})

		if err == nil || err.Error() != errMsg {
			t.Fatalf("unexpected error: %v", err)
		}
		if f != nil {
			t.Error("expected file for input")
		}
		if actualFileName != fileName {
			t.Error("unexpected filename for input")
		}
	})
}
