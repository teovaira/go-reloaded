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
	"strconv"
	"strings"
)

func main() {
	// validation for correct number arguments
	if len(os.Args) != 3 {
		fmt.Println("Usage: go run . <input.txt> <output.txt>")
		os.Exit(1)
	}

	inputFile := os.Args[1]
	outputFile := os.Args[2]

	inputText, err := readInputFile(inputFile)
	if err != nil {
		fmt.Println("Error in reading the input file:", err)
		os.Exit(1)
	}

	processText(inputText)

	err = writeOutputFile(outputFile, inputText)
	if err != nil {
		fmt.Println("Error in writing the output file:", err)
		os.Exit(1)
	}
}

// readInputFile opens the given file and returns its content as a string
func readInputFile(path string) (string, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(data), nil // convert []byte to string and return it
}

// writeOutputFile writes the given content after all transformation rules applied into a file at the given path
func writeOutputFile(path, content string) error {
	err := os.WriteFile(path, []byte(content), 0o644)
	if err != nil {
		return err
	}
	return nil
}

//processText is the main text-processing pipeline
func processText(text string) string {
	 
	words := strings.Fields(text)

	fmt.Print("DEBUG tokens: ", words)

	return text
}

// convertHexAndBin looks for "(hex)" or "(bin)" 
// patterns and replaces the word before them with its
// decimal value
// func convertHexAndBin(words [string]) [string {

// }]
