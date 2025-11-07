package main

import (
	"strconv"
	"strings"
	"unicode"
)

// ExtractNumber extracts the count from markers like (up, 3).
// Returns (n, true) when a valid positive integer is present; otherwise (0, false).
func ExtractNumber(token string) (int, bool) {
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
