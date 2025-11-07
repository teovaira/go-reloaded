package transform

import (
	"fmt"
	"strconv"
	"strings"
)

const (
	// BaseHexadecimal is the numeric base for hexadecimal numbers
	BaseHexadecimal = 16
	// BaseBinary is the numeric base for binary numbers
	BaseBinary = 2
)

// ConvertHexAndBin scans the tokenized text for patterns like "<number> (hex)" or "<number> (bin)"
// and replaces the numeric word before them with its decimal equivalent.
//
// It supports punctuation directly following the markers, such as (hex). or (bin),
// since punctuation normalization guarantees they are separate tokens or adjacent punctuation.
//
// Example:
// Input tokens: ["1E", "(hex)", "files", "were", "added", "."]
// Output:       ["30", "files", "were", "added", "."]
//
// Input tokens: ["1111", "(bin),", "add", "1E", "(hex)."]
// Output:       ["15", "add", "30", "."]
func ConvertHexAndBin(words []string) []string {
	var result []string

	for i := 0; i < len(words); i++ {
		word := words[i]

		// Defensive check: look ahead to the next token if available
		if i+1 < len(words) {
			// Trim possible trailing punctuation from the next token
			next := strings.Trim(words[i+1], ".,!?;:")

			switch next {
			case "(hex)":
				// Attempt hexadecimal conversion
				value, err := strconv.ParseInt(word, BaseHexadecimal, 64)
				if err == nil {
					result = append(result, fmt.Sprintf("%d", value))
				} else {
					result = append(result, word) // keep original if not a valid hex number
				}
				i++ // skip the "(hex)" token
				continue

			case "(bin)":
				// Attempt binary conversion
				value, err := strconv.ParseInt(word, BaseBinary, 64)
				if err == nil {
					result = append(result, fmt.Sprintf("%d", value))
				} else {
					result = append(result, word)
				}
				i++ // skip the "(bin)" token
				continue
			}
		}

		// Normal word (not followed by a marker)
		result = append(result, word)
	}

	return result
}
