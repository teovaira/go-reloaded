package main

import (
	
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
			n := ExtractNumber(word)
			for j := 1; j <= n && len(result)-j >=0; j++ {
				result[len(result)-j] = strings.ToUpper(result[len(result)-j])
			}
			continue
		}

		// Handle (low, n)
		if strings.HasPrefix(word, "(low,") && strings.HasSuffix(word, ")") {
			n := ExtractNumber(word)
			for j := 1; j <= n && len(result)-j >= 0; j++ {
				result[len(result)-j] = strings.ToLower(result[len(result)-j])
			}
			continue
		}

		// Handle (cap, n)
		if strings.HasPrefix(word, "(cap,") && strings.HasSuffix(word, ")") {
			n := ExtractNumber(word)
			for j := 1; j <= n && len(result)-j >= 0; j++ {
				result[len(result)-j] = Capitalize(result[len(result)-j])
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
			result[len(result)-1] = Capitalize(result[len(result)-1])
			continue
		}
		result = append(result, word)
	}
	return result
	
}



