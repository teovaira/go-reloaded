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

	"go-reloaded/internal/fileio"
	"go-reloaded/internal/pipeline"
)

func main() {
	// validation for correct number arguments
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input.txt> <output.txt>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	inputText, err := fileio.ReadInputFile(inputFile)
	if err != nil {
		fmt.Println("Error in reading the input file:", err)
		os.Exit(1)
	}

	outputText := pipeline.ProcessText(inputText)

	err = fileio.WriteOutputFile(outputFile, outputText)
	if err != nil {
		fmt.Println("Error in writing the output file:", err)
		os.Exit(1)
	}
}
