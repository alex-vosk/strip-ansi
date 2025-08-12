package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestMakeInputOutputFiles(t *testing.T) {
	t.Run("StdIn and StdOut", func(t *testing.T) {
		args := Args{InputName: stdInOut, OutputName: stdInOut}
		input, output, err := makeInputOutputFiles(args)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		if input != os.Stdin {
			t.Error("expected stdin for input")
		}
		if output != os.Stdout {
			t.Error("expected stdout for output")
		}
	})

	t.Run("File Input and StdOut", func(t *testing.T) {
		tempDir := t.TempDir()
		inputFile, err := os.Create(filepath.Join(tempDir, "input.txt"))
		if err != nil {
			t.Fatalf("failed to create temp input file: %v", err)
		}
		inputFile.Close()

		args := Args{InputName: inputFile.Name(), OutputName: stdInOut}
		input, output, err := makeInputOutputFiles(args)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		defer input.Close()

		if input.Name() != inputFile.Name() {
			t.Errorf("expected input file %s, got %s", inputFile.Name(), input.Name())
		}
		if output != os.Stdout {
			t.Error("expected stdout for output")
		}
	})

	t.Run("StdIn and File Output", func(t *testing.T) {
		tempDir := t.TempDir()
		outputFilePath := filepath.Join(tempDir, "output.txt")

		args := Args{InputName: stdInOut, OutputName: outputFilePath}
		input, output, err := makeInputOutputFiles(args)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		defer output.Close()

		if input != os.Stdin {
			t.Error("expected stdin for input")
		}
		if output.Name() != outputFilePath {
			t.Errorf("expected output file %s, got %s", outputFilePath, output.Name())
		}
	})

	t.Run("File Input and File Output", func(t *testing.T) {
		tempDir := t.TempDir()
		inputFile, err := os.Create(filepath.Join(tempDir, "input.txt"))
		if err != nil {
			t.Fatalf("failed to create temp input file: %v", err)
		}
		inputFile.Close()
		outputFilePath := filepath.Join(tempDir, "output.txt")

		args := Args{InputName: inputFile.Name(), OutputName: outputFilePath}
		input, output, err := makeInputOutputFiles(args)
		if err != nil {
			t.Fatalf("unexpected error: %v", err)
		}
		defer input.Close()
		defer output.Close()

		if input.Name() != inputFile.Name() {
			t.Errorf("expected input file %s, got %s", inputFile.Name(), input.Name())
		}
		if output.Name() != outputFilePath {
			t.Errorf("expected output file %s, got %s", outputFilePath, output.Name())
		}
	})

	t.Run("Non-existent Input File", func(t *testing.T) {
		args := Args{InputName: "non-existent-file.txt", OutputName: stdInOut}
		_, _, err := makeInputOutputFiles(args)
		if err == nil {
			t.Fatal("expected an error for non-existent input file, but got nil")
		}
	})
}
