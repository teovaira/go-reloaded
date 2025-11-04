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

	outputText := processText(inputText)

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

// processText is the main text-processing pipeline
// it sequentially applies all transformation stages
func processText(text string) string {
	
	words := tokenize(text)

	// fmt.Printf("debug raw tokens: %#v", words)

	words = convertHexAndBin(words)

	words = fixArticles(words)

	words = applyCaseRules(words)

	fmt.Println("DEBUG tokens: ", words)

	return strings.Join(words, " ")

}

// convertHexAndBin looks for "(hex)" or "(bin)" 
// patterns and replaces the word before them with its
// decimal value
func convertHexAndBin(words []string) []string {
	var result []string
	
	for i:= 0; i < len(words); i++ {
		word:= words[i]

		if word == "(hex)" && i > 0 {
			prev:= words[i-1]

			value, err := strconv.ParseInt(prev, 16, 64)
			if err == nil {
				result[len(result)-1] = fmt.Sprint(value)
			}
			continue
		}

		if word == "(bin)" && i> 0 {
			prev:= words[i-1]

			value, err := strconv.ParseInt(prev, 2, 64)
			if err == nil {
				result[len(result)-1] = fmt.Sprint(value)
			}
			continue
		}
		result = append(result, word)
	}
	return result

}

// fixArticles checks every occurrence of "a" in the text and changes it
// to "an" if the next word starts with a vowel (a, e, i, o, u) or 'h'.
// The comparison is case-insensitive, and the loop stops before the last
// word to prevent out-of-range errors.
func fixArticles(words []string) []string {
	for i := 0; i < len(words)-1; i++ { //stop before the last word
		current := strings.ToLower(words[i])
		next := strings.ToLower(words[i+1])

		if current == "a" {
			
			if strings.HasPrefix(next, "a") ||
			 strings.HasPrefix(next, "e") ||
			 strings.HasPrefix(next, "i") ||
			 strings.HasPrefix(next, "o") ||
			 strings.HasPrefix(next, "u") ||
			 strings.HasPrefix(next, "h") {
			 words[i] = "an"
			}
		}
	}
	return words
}

// applyCaseRules detects (up), (low) and (cap) markers 
// and applies the appropriate transformation to the 
// previous one or multiple words. Markers are removed from the final output
func applyCaseRules(words []string) []string {
	var result []string

	for i := 0; i < len(words); i++ {
		word := words[i]

		// Handle (up, n)
		if strings.HasPrefix(word, "(up,") && strings.HasSuffix(word, ")") {
			n := extractNumber(word)
			for j := 1; j <= n && len(result)-j >=0; j++ {
				result[len(result)-j] = strings.ToUpper(result[len(result)-j])
			}
			continue
		}

		// Handle (low, n)
		if strings.HasPrefix(word, "(low,") && strings.HasSuffix(word, ")") {
			n := extractNumber(word)
			for j := 1; j <= n && len(result)-j >= 0; j++ {
				result[len(result)-j] = strings.ToLower(result[len(result)-j])
			}
			continue
		}

		// Handle (cap, n)
		if strings.HasPrefix(word, "(cap,") && strings.HasSuffix(word, ")") {
			n := extractNumber(word)
			for j := 1; j <= n && len(result)-j >= 0; j++ {
				result[len(result)-j] = capitalize(result[len(result)-j])
			}
			continue
		}

		// Handle (up)
		if word == "(up)" && len(result) > 0 {
			result[len(result)-1] = strings.ToUpper(result[len(result)-1])
			continue
		}

		// Handle (low)
		if word == "(low)" && len(result) > 0 {
			result[len(result)-1] = strings.ToLower(result[len(result)-1])
			continue
		}

		// Handle (cap)
		if word == "(cap)" && len(result) > 0 {
			result[len(result)-1] = capitalize(result[len(result)-1])
			continue
		}
		result = append(result, word)
	}
	return result
	
}

// extractNumber extracts the number from markers like (up, 3)
func extractNumber(token string) int {
	token = strings.TrimSuffix(strings.TrimPrefix(token, "("), ")")
	parts := strings.Split(token, ",")
	if len(parts) < 2 {
		return 1
	}
	numStr:= strings.TrimSpace(parts[1])
	n, err := strconv.Atoi(numStr)
	if err != nil {
		return 1
	}
	return n
}

// capitalize turns the first letter uppercase and the rest lowercase
func capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(string(word[0])) + strings.ToLower(string(word[1:]))
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