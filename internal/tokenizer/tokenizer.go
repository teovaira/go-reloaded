package tokenizer

import "strings"

// Tokenize splits the text into words while keeping punctuation
// as separate tokens and preserving markers like (up, 2).
//
// Example:
//
//	Input:  "hello world (up, 2) !"
//	Output: ["hello", "world", "(up, 2)", "!"]
//
// Design Decision: Punctuation as Word Boundaries
//
// This tokenizer treats all punctuation marks (.,!?;:) as word boundaries,
// meaning they always separate tokens. This design choice ensures:
//
//  1. Consistency: Punctuation always splits tokens, regardless of position
//  2. Predictability: Simple rule with no special cases
//  3. Compatibility: Works with ApplyPunctuationRules which attaches
//     punctuation to previous words
//
// Edge Cases:
//
//   - "hello!world" → ["hello", "!", "world"] (splits even without spaces)
//   - "hel!lo" → ["hel", "!", "lo"] (punctuation in middle splits word)
//   - "example.com" → ["example", ".", "com"] (URLs get split)
//   - "3.14" → ["3", ".", "14"] (decimals get split)
//
// This behavior is intentional. Natural language text rarely has punctuation
// in the middle of words, and this approach keeps the tokenizer simple while
// handling the specification's requirements for normal cases.
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
