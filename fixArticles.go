package main

import "strings"

// fixArticles checks every occurrence of "a" in the text and changes it
// to "an" if the next word starts with a vowel (a, e, i, o, u) or 'h'.
// The comparison is case-insensitive, and the loop stops before the last
// word to prevent out-of-range errors.
func FixArticles(words []string) []string {
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