package main

import (
	"testing"

	"go-reloaded/internal/pipeline"
)

func TestProcessTextIntegration(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "hex and bin conversion",
			input:    "Simply add 42 (hex) and 10 (bin) and you will see the result is 68.",
			expected: "Simply add 66 and 2 and you will see the result is 68.",
		},
		{
			name:     "article correction",
			input:    "There is no greater agony than bearing a untold story inside you.",
			expected: "There is no greater agony than bearing an untold story inside you.",
		},
		{
			name:     "punctuation spacing",
			input:    "Punctuation tests are ... kinda boring ,what do you think ?",
			expected: "Punctuation tests are... kinda boring, what do you think?",
		},
		{
			name:     "uppercase transformation",
			input:    "We are learning go (up) today!",
			expected: "We are learning GO today!",
		},
		{
			name:     "lowercase transformation",
			input:    "This should become lowercase (low).",
			expected: "This should become lowercase.",
		},
		{
			name:     "capitalize transformation",
			input:    "please capitalize this (cap).",
			expected: "please capitalize This.",
		},
		{
			name:     "capitalize multiple words",
			input:    "welcome to the brooklyn bridge (cap, 3).",
			expected: "welcome to The Brooklyn Bridge.",
		},
		{
			name:     "uppercase multiple words",
			input:    "this will go up two words (up, 2) in a row.",
			expected: "this will go up TWO WORDS in a row.",
		},
		{
			name:     "quote handling",
			input:    "He said: ' hello '",
			expected: "He said: 'hello'",
		},
		{
			name:     "complex mix",
			input:    "it (cap) was the best of times, it was the worst of times (up)",
			expected: "It was the best of times, it was the worst of TIMES",
		},
		{
			name:     "article with h",
			input:    "a honor to meet a hero",
			expected: "an honor to meet an hero",
		},
		{
			name:     "multiple conversions",
			input:    "Values: 1E (hex) and FF (hex) and A (hex)",
			expected: "Values: 30 and 255 and 10",
		},
		{
			name:     "empty string",
			input:    "",
			expected: "",
		},
		{
			name:     "only spaces",
			input:    "   ",
			expected: "",
		},
		{
			name:     "newline preservation",
			input:    "First line\nSecond line",
			expected: "First line\nSecond line",
		},
		{
			name:     "multiple newlines",
			input:    "Line 1\n\nLine 3",
			expected: "Line 1\n\nLine 3",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := pipeline.ProcessText(tt.input)
			if result != tt.expected {
				t.Errorf("ProcessText(%q)\n  got: %q\n want: %q", tt.input, result, tt.expected)
			}
		})
	}
}
