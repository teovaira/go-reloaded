package transform

import "strings"

// FixArticles checks every occurrence of "a" in the text and changes it
// to "an" if the next word starts with a vowel (a, e, i, o, u) or 'h'.
// The comparison is case-insensitive, and the loop stops before the last
// word to prevent out-of-range errors.
func FixArticles(words []string) []string {
	for i := 0; i < len(words)-1; i++ { // stop before the last word
		// Strip leading quotes to check if this is an article
		currentLower := strings.ToLower(strings.Trim(words[i], "\"'"))
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

		// HasPrefix handles empty strings by returning false, no panic.
		if strings.HasPrefix(trimmed, "a") ||
			strings.HasPrefix(trimmed, "e") ||
			strings.HasPrefix(trimmed, "i") ||
			strings.HasPrefix(trimmed, "o") ||
			strings.HasPrefix(trimmed, "u") ||
			strings.HasPrefix(trimmed, "h") {
			// Preserve original case and any leading quotes
			// Replace "a" or "A" with "an" or "An" while keeping quotes
			if strings.Contains(words[i], "A") {
				words[i] = strings.Replace(words[i], "A", "An", 1)
			} else {
				words[i] = strings.Replace(words[i], "a", "an", 1)
			}
		}
	}
	return words
}
