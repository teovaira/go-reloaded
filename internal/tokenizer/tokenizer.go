package tokenizer

import "strings"

// Tokenize splits the text into words while keeping punctuation
// as separate tokens and preserving markers like (up, 2).
//
// Example:
//   Input:  "hello world (up, 2) !"
//   Output: ["hello", "world", "(up, 2)", "!"]
func Tokenize(text string) []string {
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
