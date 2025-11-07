package main

import (
	"strconv"
	"strings"
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
	if len(word) == 0 {
		return word
	}
	return strings.ToUpper(string(word[0])) + strings.ToLower(string(word[1:]))
}