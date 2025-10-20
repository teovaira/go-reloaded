// -----------------------------------------------
// Project: go-reloaded
// File: main.go
// Author: Theodore Vairaktaris
// Description: Entry point for the go-reloaded project.
// This program reads a text file, applies transformations,
// and writes the modified content into another file.
// -----------------------------------------------

package main

import (
	"fmt"
	"os"
)

func main() {
	// validation for correct number of passwords
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input.txt> <output.txt>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	inputText, err := readInputFile(inputFile)
	if err != nil {
		fmt.Println("Error in reading input file:", err)
		return
	}

	fmt.Println("File content:\n", inputText)
	_ = outputFile
}

// readInputFile opens the given file and returns its content as a string
func readInputFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil // convert []byte to string and return it
}
