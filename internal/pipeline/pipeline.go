package pipeline

import (
	"strings"

	"go-reloaded/internal/tokenizer"
	"go-reloaded/internal/transform"
)

// ProcessText is the main text-processing pipeline.
// It processes text line-by-line to preserve newlines while applying
// all transformation stages in the correct order.
func ProcessText(text string) string {
	// Preserve newlines by processing line-by-line
	lines := strings.Split(text, "\n")
	output := make([]string, 0, len(lines))

	for _, line := range lines {
		if line == "" {
			output = append(output, "")
			continue
		}
		// Tokenize and apply token-level transforms
		words := tokenizer.Tokenize(line)
		words = transform.ConvertHexAndBin(words)
		words = transform.FixArticles(words)
		words = transform.ApplyCaseRules(words)

		// Rebuild the line
		rebuiltLine := strings.Join(words, " ")
		// Final spacing and quotes per line
		rebuiltLine = transform.ApplyPunctuationRules(rebuiltLine)
		rebuiltLine = transform.FixQuotes(rebuiltLine)
		output = append(output, rebuiltLine)
	}
	return strings.Join(output, "\n")
}
