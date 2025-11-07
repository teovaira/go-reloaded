package main

import (
    "strconv"
    "strings"
    "unicode"
)

// extractNumber extracts the number from markers like (up, 3)
func ExtractNumber(token string) int {
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
