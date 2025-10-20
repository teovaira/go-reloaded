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
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input.txt> <output.txt>")
		return
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	fmt.Println("Input file:", inputFile)
	fmt.Println("Output file:", outputFile)
}
