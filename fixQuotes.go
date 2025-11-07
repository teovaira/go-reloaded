package main

import (
	"bytes"
	"strings"
	"unicode"
	"unicode/utf8"
)

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
//   - If multiple spaces exist before the opening quote, collapse them to exactly one.
//   - We do not delete the single space before the opening quote.
//
// 2) Remove spaces immediately inside quotes:
//   - After opening quote.
//   - Before closing quote.
//
// 3) After a closing quote, insert one space if the next rune is a letter, a digit or '('.
//   - Do not insert a space if the next rune is punctuation (.,!?:;) or whitespace
//
// 4) Handles both " and ' quotes, and only closes with the same quote type opened
// 5) Unicode safe: uses utf8.DecodeLastRuneInString + Builder. for in-place trimming.
func FixQuotes(text string) string {
	var b bytes.Buffer // Efficient builder for UTF-8 bytes
	inQuotes := false
	var quoteType rune // Tracks opening quote type

	runes := []rune(text)

	for i := 0; i < len(runes); i++ {
		ch := runes[i]

		// Treat apostrophes inside words as literal, not as quotes
		if (ch == '\'' || ch == '’') && i-1 >= 0 && i+1 < len(runes) &&
			(unicode.IsLetter(runes[i-1]) || unicode.IsDigit(runes[i-1])) &&
			(unicode.IsLetter(runes[i+1]) || unicode.IsDigit(runes[i+1])) {
			b.WriteRune(ch)
			continue
		}

		// Handle quote characters (straight and curly)
		if ch == '"' || ch == '\'' || ch == '“' || ch == '”' || ch == '‘' || ch == '’' {
			// Normalize curly quotes to straight for consistency in spacing rules
			isSingle := (ch == '\'' || ch == '‘' || ch == '’')
			normalized := ch
			if ch == '“' || ch == '”' {
				normalized = '"'
			}
			if ch == '‘' || ch == '’' {
				normalized = '\''
			}

			if !inQuotes {
				// Collapse multiple spaces before opening quote to exactly one
				if b.Len() > 0 {
					trailingSpaces := 0
					for b.Len() > 0 {
						r, size := utf8.DecodeLastRune(b.Bytes())
						if size == 0 || !unicode.IsSpace(r) {
							break
						}
						trailingSpaces++
						b.Truncate(b.Len() - size)
					}
					if trailingSpaces > 0 {
						b.WriteRune(' ')
					}
				}

				b.WriteRune(normalized)
				inQuotes = true
				quoteType = normalized

				// Remove spaces immediately after opening quote
				for i+1 < len(runes) && unicode.IsSpace(runes[i+1]) {
					i++
				}
				continue
			}

			// Closing quote (match type)
			if inQuotes && ((isSingle && quoteType == '\'') || (!isSingle && quoteType == '"')) {
				// Remove spaces before closing quote
				for b.Len() > 0 {
					r, size := utf8.DecodeLastRune(b.Bytes())
					if size == 0 || !unicode.IsSpace(r) {
						break
					}
					b.Truncate(b.Len() - size)
				}

				b.WriteRune(quoteType)
				inQuotes = false

				// Add a space after closing quote if followed by a letter/digit or '('
				if i+1 < len(runes) {
					next := runes[i+1]
					if unicode.IsLetter(next) || unicode.IsDigit(next) || next == '(' {
						b.WriteRune(' ')
					}
				}
				continue
			}
		}

		// Default: copy character as-is
		b.WriteRune(ch)
	}
	return strings.TrimSpace(b.String())
}
