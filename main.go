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
	"bytes"
	"fmt"
	"os"
	"strconv"
	"strings"
	"unicode"
	"unicode/utf8"
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

	outputText := ProcessText(inputText)

	err = writeOutputFile(outputFile, outputText)
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

// writeOutputFile writes the transformed content after all transformation rules applied into a file at the given path
func writeOutputFile(path, content string) error {
	err := os.WriteFile(path, []byte(content), 0o644)
	if err != nil {
		return err
	}
	return nil
}









// tokenize splits the text into words while keeping punctuation
// as separate tokens and preserving markers like (up, 2).
func tokenize(text string) []string {
	var tokens []string
	current := ""
	inParentheses := false // new flag to check if we are inside commands

	for _, r := range text {
		ch := string(r)

		switch {
		case r == '(':
			if current != "" {
				tokens = append(tokens, current)
				current = ""
			}
			inParentheses = true
			current += ch

		case r == ')':
			current += ch
			inParentheses = false
			tokens = append(tokens, current)
			current = ""

		case strings.ContainsRune(" \n\t", r) && !inParentheses:
			if current != "" {
				tokens = append(tokens, current)
				current = ""
			}

		case strings.ContainsRune(".,!?;:", r) && !inParentheses:
			if current != "" {
				tokens = append(tokens, current)
				current = ""
			}
			tokens = append(tokens, ch)

		default:
			current += ch
		}
	}
	if current != "" {
			tokens = append(tokens, current)
	}
	return tokens
}







//fixQuotes cleans spacing around quotation marks
// It ensures proper placement and spacing for opening and closing quotes
// func fixQuotes(text string) string {
// 	var b strings.Builder
// 	inQuotes := false
// 	runes := []rune(text)

// 	for i := 0; i < len(runes); i++ {
// 		ch := runes[i]

		// Case 1: it's a quotation mark
		// if ch == '"' {
			// Trim any space before opening quote
			// if !inQuotes {
				// If last written char was a space, remove it
				// if b.Len() > 0 {
				// 	last := b.String()[b.Len()-1]
				// 	if last == ' ' {
						// Overwrite last space (we 'll rebuild the string)
				// 		temp := b.String()[:b.Len()-1]
				// 		b.Reset()
				// 		b.WriteString(temp)
				// 	}
				// }
				// Add the opening quote
				// b.WriteRune(ch)
				// inQuotes = true
				// Skip any spaces immediately after the quote
			// 	for i+1 < len(runes) && runes[i+1] == ' ' {
			// 		i++
			// 	}
			// 	continue
			// }

			// Case 2: closing quote
			// if inQuotes {
				// Remove spaces before the closing quote
				// for b.Len() > 0 && b.String()[b.Len()-1] == ' ' {
				// 	temp := b.String()[:b.Len()-1]
				// 	b.Reset()
				// 	b.WriteString(temp)
				// }
				// b.WriteRune(ch)
				// inQuotes = false
				// Add one space after closing quote if next is a letter
		// 		if i+1 < len(runes) && runes[i+1] != '.' && runes[i+1] != ',' && runes[i+1] != ' ' {
		// 			b.WriteRune(' ')
		// 		}
		// 		continue
		// 	}
		// }
		// Default: copy character, copy as is
	// 	b.WriteRune(ch)
	// }

	// Final cleanup: remove extra spaces at ends
// 	return strings.TrimSpace(b.String())
// }




// fixQuotes cleans spacing around both "double" and "single" quotation marks
//
// Rules implemented:
// 1) Keep one space before an opening quote if there was any space before it.
//    - If multiple spaces exist before the opening quote, collapse them to exactly one.
//    - We do not delete the single space before the opening quote.
// 2) Remove spaces immediately inside quotes:
//    - After opening quote.
//    - Before closing quote.
// 3) After a closing quote, insert one space if the next rune is a letter, a digit or '('.
//    - Do not insert a space if the next rune is punctuation (.,!?:;) or whitespace
// 4) Handles both " and ' quotes, and only closes with the same quote type opened
// 5) Unicode safe: uses utf8.DecodeLastRuneInString + Builder. for in-place trimming.
func fixQuotes(text string) string {
	var b bytes.Buffer // Efficient builder for UTF-8 bytes
	inQuotes := false
	var quoteType rune // Stores the opening quote (" or ')

	runes := []rune(text)

	for i := 0; i < len(runes); i++ {
		ch := runes[i]

		if ch == '"' || ch == '\'' {
			if !inQuotes {
				if b.Len() > 0 {
					trailingSpaces := 0
					for b.Len() > 0{
						r, size := utf8.DecodeLastRune(b.Bytes())
						if size == 0 || !unicode.IsSpace(r) {
							break
						}
						trailingSpaces++
						b.Truncate(b.Len()-size)
					}
					if trailingSpaces > 0 {
						b.WriteRune(' ')
					}
				}

				b.WriteRune(ch)
				inQuotes = true
				quoteType = ch

				for i+1 < len(runes) && runes[i+1] == ' ' {
					i++
				}
				continue
			}

			if inQuotes && ch == quoteType {

				for b.Len() > 0 {
					r, size := utf8.DecodeLastRune(b.Bytes())
					if size == 0 ||  !unicode.IsSpace(r) {
						break
					}
					b.Truncate(b.Len() -size)
				}

				b.WriteRune(ch)
				inQuotes = false

				if i+1 < len(runes) {
					next := runes[i+1]
					if unicode.IsLetter(next) || unicode.IsDigit(next) ||  next == '(' {
						b.WriteRune(' ')
					}
				}
				continue
			}
		}
		b.WriteRune(ch)
	}
	return strings.TrimSpace(b.String())
}