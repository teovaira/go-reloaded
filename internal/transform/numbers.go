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
			// Trim possible trailing punctuation and quotes from the next token
			next := strings.Trim(words[i+1], ".,!?;:\"'")

			switch next {
			case "(hex)":
				// Strip quotes for parsing, but preserve them in output
				cleanWord := strings.Trim(word, "\"'")
				prefix := strings.TrimSuffix(word, cleanWord)
				suffix := strings.TrimPrefix(word, prefix+cleanWord)

				// Attempt hexadecimal conversion
				value, err := strconv.ParseInt(cleanWord, BaseHexadecimal, 64)
				if err == nil {
					// Successfully converted - preserve any quotes around the converted number
					result = append(result, prefix+fmt.Sprintf("%d", value)+suffix)
					i++ // skip the "(hex)" token
					continue
				}
				// If conversion failed, keep both word and marker unchanged
				// Fall through to append word normally (marker will be added in next iteration)

			case "(bin)":
				// Strip quotes for parsing, but preserve them in output
				cleanWord := strings.Trim(word, "\"'")
				prefix := strings.TrimSuffix(word, cleanWord)
				suffix := strings.TrimPrefix(word, prefix+cleanWord)

				// Attempt binary conversion
				value, err := strconv.ParseInt(cleanWord, BaseBinary, 64)
				if err == nil {
					// Successfully converted - preserve any quotes around the converted number
					result = append(result, prefix+fmt.Sprintf("%d", value)+suffix)
					i++ // skip the "(bin)" token
					continue
				}
				// If conversion failed, keep both word and marker unchanged
				// Fall through to append word normally (marker will be added in next iteration)
			}
		}

		// Normal word (not followed by a marker)
		result = append(result, word)
	}

	return result
}
