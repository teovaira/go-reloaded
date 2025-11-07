package main

import (
	"strconv"
	"strings"
)

// applyCaseRules detects (up), (low) and (cap) markers
// and applies the appropriate transformation to the
// previous one or multiple words. Markers are removed from the final output
func ApplyCaseRules(words []string) []string {
	var result []string

	for i := 0; i < len(words); i++ {
		word := words[i]

		// Handle (up, n)
		if strings.HasPrefix(word, "(up,") && strings.HasSuffix(word, ")") {
			n := extractNumber(word)
			for j := 1; j <= n && len(result)-j >=0; j++ {
				result[len(result)-j] = strings.ToUpper(result[len(result)-j])
			}
			continue
		}

		// Handle (low, n)
		if strings.HasPrefix(word, "(low,") && strings.HasSuffix(word, ")") {
			n := extractNumber(word)
			for j := 1; j <= n && len(result)-j >= 0; j++ {
				result[len(result)-j] = strings.ToLower(result[len(result)-j])
			}
			continue
		}

		// Handle (cap, n)
		if strings.HasPrefix(word, "(cap,") && strings.HasSuffix(word, ")") {
			n := extractNumber(word)
			for j := 1; j <= n && len(result)-j >= 0; j++ {
				result[len(result)-j] = capitalize(result[len(result)-j])
			}
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
			result[len(result)-1] = capitalize(result[len(result)-1])
			continue
		}
		result = append(result, word)
	}
	return result
	
}

// extractNumber extracts the number from markers like (up, 3)
func extractNumber(token string) int {
	token = strings.TrimSuffix(strings.TrimPrefix(token, "("), ")")
	parts := strings.Split(token, ",")
	if len(parts) < 2 {
		return 1
	}
	numStr:= strings.TrimSpace(parts[1])
	n, err := strconv.Atoi(numStr)
	if err != nil {
		return 1
	}
	return n
}

// capitalize turns the first letter uppercase and the rest lowercase
func capitalize(word string) string {
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(string(word[0])) + strings.ToLower(string(word[1:]))
}