package transform

import (
	"strconv"
	"strings"
	"unicode"
)

// ParseMarkerCount extracts the count from markers like (up, 3) or (low, 2).
// Returns (n, true) when a valid positive integer is present; otherwise (0, false).
// Invalid counts (negative, zero, or malformed) return (0, false).
func ParseMarkerCount(token string) (int, bool) {
	token = strings.TrimSuffix(strings.TrimPrefix(token, "("), ")")
	parts := strings.Split(token, ",")
	if len(parts) < 2 {
		return 0, false
	}
	numStr := strings.TrimSpace(parts[1])
	n, err := strconv.Atoi(numStr)
	if err != nil || n <= 0 {
		return 0, false
	}
	return n, true
}

// Capitalize turns the first letter uppercase and the rest lowercase.
func Capitalize(word string) string {
	if word == "" {
		return word
	}
	runes := []rune(word)
	runes[0] = unicode.ToUpper(runes[0])
	for i := 1; i < len(runes); i++ {
		runes[i] = unicode.ToLower(runes[i])
	}
	return string(runes)
}

// ApplyCaseRules detects (up), (low) and (cap) markers
// and applies the appropriate transformation to the
// previous one or multiple words. Markers are removed from the final output.
func ApplyCaseRules(words []string) []string {
	var result []string

	for i := 0; i < len(words); i++ {
		word := words[i]

		// Handle (up, n)
		if strings.HasPrefix(word, "(up,") && strings.HasSuffix(word, ")") {
			if n, ok := ParseMarkerCount(word); ok {
				for j := 1; j <= n && len(result)-j >= 0; j++ {
					result[len(result)-j] = strings.ToUpper(result[len(result)-j])
				}
				// consume marker
				continue
			}
			// invalid marker: keep as literal
			result = append(result, word)
			continue
		}

		// Handle (low, n)
		if strings.HasPrefix(word, "(low,") && strings.HasSuffix(word, ")") {
			if n, ok := ParseMarkerCount(word); ok {
				for j := 1; j <= n && len(result)-j >= 0; j++ {
					result[len(result)-j] = strings.ToLower(result[len(result)-j])
				}
				continue
			}
			result = append(result, word)
			continue
		}

		// Handle (cap, n)
		if strings.HasPrefix(word, "(cap,") && strings.HasSuffix(word, ")") {
			if n, ok := ParseMarkerCount(word); ok {
				for j := 1; j <= n && len(result)-j >= 0; j++ {
					result[len(result)-j] = Capitalize(result[len(result)-j])
				}
				continue
			}
			result = append(result, word)
			continue
		}

		// Handle (up)
		if word == "(up)" && len(result) > 0 {
			result[len(result)-1] = strings.ToUpper(result[len(result)-1])
			continue
		}

		// Handle (low)
		if word == "(low)" && len(result) > 0 {
			result[len(result)-1] = strings.ToLower(result[len(result)-1])
			continue
		}

		// Handle (cap)
		if word == "(cap)" && len(result) > 0 {
			result[len(result)-1] = Capitalize(result[len(result)-1])
			continue
		}
		result = append(result, word)
	}
	return result

}
