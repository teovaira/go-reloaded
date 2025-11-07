package main

import (
	"strings"
)

// processText is the main text-processing pipeline
// it sequentially applies all transformation stages
// func processText(text string) string {

// 	words := tokenize(text)

// 	fmt.Printf("debug raw tokens: %#v", words)

// 	words = convertHexAndBin(words)

// 	words = fixArticles(words)

// 	words = applyCaseRules(words)

// 	formattedText := applyPunctuationRules(words)

// 	finalText := fixQuotes(formattedText)

// 	fmt.Println("DEBUG tokens: ", finalText)

// 	return finalText

// }

func ProcessText(text string) string {
	// Preserve newlines by processing line-by-line
	lines := strings.Split(text, "\n")
	out := make([]string, 0, len(lines))

	for _, line := range lines {
		if line == "" {
			out = append(out, "")
			continue
		}
		// Tokenize and apply token-level transforms
		words := Tokenize(line)
		words = ConvertHexAndBin(words)
		words = FixArticles(words)
		words = ApplyCaseRules(words)

		// Rebuild the line
		rebuilt := strings.Join(words, " ")
		// Final spacing and quotes per line
		rebuilt = ApplyPunctuationRules(rebuilt)
		rebuilt = FixQuotes(rebuilt)
		out = append(out, rebuilt)
	}
	return strings.Join(out, "\n")
}
