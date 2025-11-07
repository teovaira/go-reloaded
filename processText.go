package main

import (
	"strings"
)

// ProcessText is the main text-processing pipeline.
// It processes text line-by-line to preserve newlines while applying
// all transformation stages in the correct order.
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
