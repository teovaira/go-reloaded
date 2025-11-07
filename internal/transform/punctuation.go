package transform

import (
	"bytes"
	"strings"
	"unicode"
)

// ApplyPunctuationRules fixes spacing for . , ! ? : ; according to the spec.
//
// Rules implemented:
// 1) Remove spaces before punctuation (. , ! ? : ;)
// 2) Add one space after punctuation if next is a letter, digit, quote, or '('
// 3) Treat multi-punctuation groups (..., !!, !?, etc.) as one unit
// 4) Never add or remove spaces around parentheses except when '(' follows punctuation
//
// Example:
// Input:  "I was sitting over there ,and then BAMM !!"
// Output: "I was sitting over there, and then BAMM!!"
func ApplyPunctuationRules(text string) string {
	var b bytes.Buffer
	runes := []rune(text)
	length := len(runes)

	for i := 0; i < length; i++ {
		r := runes[i]

		// --- Rule 1: remove spaces before punctuation
		if unicode.IsSpace(r) && i+1 < length {
			next := runes[i+1]
			if strings.ContainsRune(".,!?;:", next) {
				continue // skip writing this space
			}
		}

		// --- Write current rune
		b.WriteString(string(r))

		// --- Rule 3: detect punctuation groups (like ... or !?)
		if strings.ContainsRune(".,!?;:", r) {
			// Peek ahead to see if next is punctuation too
			if i+1 < length && strings.ContainsRune(".,!?;:", runes[i+1]) {
				// Don't insert space yet — still inside group
				continue
			}

			// --- Rule 2: add a space after punctuation if next is word, quote, or '('
			if i+1 < length {
				next := runes[i+1]
				if unicode.IsLetter(next) || unicode.IsDigit(next) || next == '"' || next == '\'' || next == '(' {
					b.WriteByte(' ')
				}
			}
		}
	}

	// Preserve line breaks while normalizing spaces within each line.
	// We avoid a global strings.Fields to keep '\n' structure intact.
	built := b.String()
	lines := strings.Split(built, "\n")
	for i := range lines {
		// Collapse runs of whitespace on each line to single spaces
		// (leading/trailing spaces on the line are trimmed by Fields).
		lines[i] = strings.Join(strings.Fields(lines[i]), " ")
	}
	out := strings.Join(lines, "\n")
	return strings.TrimSpace(out)
}
