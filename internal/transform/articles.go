package transform

import "strings"

// FixArticles checks every occurrence of "a" in the text and changes it
// to "an" if the next word starts with a vowel (a, e, i, o, u) or 'h'.
// The comparison is case-insensitive, and the loop stops before the last
// word to prevent out-of-range errors.
func FixArticles(words []string) []string {
	for i := 0; i < len(words)-1; i++ { // stop before the last word
		currentLower := strings.ToLower(words[i])
		if currentLower != "a" {
			continue
		}

		// Peek next token and ignore leading quotes when deciding
		nextLower := strings.ToLower(words[i+1])
		// Strip ASCII quotes at the start (common in our tokenization)
		trimmed := strings.TrimLeft(nextLower, "'\"")
		if trimmed == "" {
			continue
		}

		if strings.HasPrefix(trimmed, "a") ||
			strings.HasPrefix(trimmed, "e") ||
			strings.HasPrefix(trimmed, "i") ||
			strings.HasPrefix(trimmed, "o") ||
			strings.HasPrefix(trimmed, "u") ||
			strings.HasPrefix(trimmed, "h") {
			words[i] = "an"
		}
	}
	return words
}
