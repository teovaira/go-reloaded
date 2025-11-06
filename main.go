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

	formattedText := applyPunctuationRules(words)

	finalText := fixQuotes(formattedText)

	fmt.Println("DEBUG tokens: ", finalText)

	return finalText

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

// applyPunctuationRules builds the final formatted text from tokens.
// It guarantees: no space before punctuation, exactly one between words
// and it skips visual test separators like '---'
func applyPunctuationRules(words []string) string {
// 	var result []string

// 	for _, word := range words {

// 		if strings.ContainsRune(".!?;,:", rune(word[0])) && len(word) == 1 {
// 			if len(result) > 0 {
// 				result[len(result)-1] += word
// 			} else {
// 				result = append(result, word)
// 			}
// 		} else {
// 			result = append(result, word)
// 		}
// 	}
// 	return strings.Join(result, " ")

	// Helper closure: small function defined inside another function.
	// It tells us if a token is a single-char punctuation mark we must
	// attach to the previous word without any space
	isPunct := func(w string) bool {
		if len(w) != 1 {
			return false
		}

		switch w[0] {
		case '.', ',', '!', '?', ';', ':':
			return true
		}
		return false
	}

	var b strings.Builder
	wroteAny := false // tracks if something has been written

	for _, w := range words {
		if w == "---" {
			continue
		}

		if isPunct(w) {
			b.WriteString(w) // attach to previous word punctuation without space
			wroteAny = true
			continue
		}

		// Normal word: if its not the first token, insert exactly one space first.
		if wroteAny {      
			b.WriteByte(' ')
			
		}
		b.WriteString(w)   // write the word itself
		wroteAny = true
	}
	return strings.TrimSpace(b.String())
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
	var b strings.Builder
	inQuotes := false
	var quoteType rune

	runes := []rune(text)

	for i := 0; i < len(runes); i++ {
		ch := runes[i]

		if ch == '"' || ch == '\'' {
			if !inQuotes {
				if b.Len() > 0 {
					trailingSpaces := 0
					for {
						r, size := utf8.DecodeLastRuneInString(b.String())
						if size == 0 || r != ' ' {
							break
						}
						trailingSpaces++
						b.Truncate(b.Len()-size)
					}
					if trailingSpaces > 0 {
						b.WriteByte(' ')
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
					r, size := utf8.DecodeLastRuneInString(b.String())
					if size == 0 ||  r != ' ' {
						break
					}
					b.Truncate(b.Len() -size)
				}

				b.WriteRune(ch)
				inQuotes = false

				if i+1 < len(runes) {
					next := runes[i+1]
					if unicode.IsLetter(next) || unicode.IsDigit(next) ||  next == '(' {
						b.WriteByte(' ')
					}
				}
				continue
			}
		}
		b.WriteRune(ch)
	}
	return strings.TrimSpace(b.String())
}