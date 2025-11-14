package tokenizer

import (
	"reflect"
	"testing"
)

func TestTokenize(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected []string
	}{
		{
			name:     "simple words",
			input:    "hello world",
			expected: []string{"hello", "world"},
		},
		{
			name:     "words with comma",
			input:    "hello, world",
			expected: []string{"hello", ",", "world"},
		},
		{
			name:     "words with period",
			input:    "hello world.",
			expected: []string{"hello", "world", "."},
		},
		{
			name:     "marker with parentheses",
			input:    "test (up)",
			expected: []string{"test", "(up)"},
		},
		{
			name:     "marker with count",
			input:    "test (up, 3)",
			expected: []string{"test", "(up, 3)"},
		},
		{
			name:     "multiple punctuation",
			input:    "Wait... really?",
			expected: []string{"Wait", ".", ".", ".", "really", "?"},
		},
		{
			name:     "quotes",
			input:    "He said \"hello\"",
			expected: []string{"He", "said", "\"hello\""},
		},
		{
			name:     "single quotes",
			input:    "It's fine",
			expected: []string{"It's", "fine"},
		},
		{
			name:     "hex marker",
			input:    "1E (hex)",
			expected: []string{"1E", "(hex)"},
		},
		{
			name:     "bin marker",
			input:    "101 (bin)",
			expected: []string{"101", "(bin)"},
		},
		{
			name:     "empty string",
			input:    "",
			expected: []string{},
		},
		{
			name:     "only spaces",
			input:    "   ",
			expected: []string{},
		},
		{
			name:     "multiple spaces between words",
			input:    "hello    world",
			expected: []string{"hello", "world"},
		},
		{
			name:     "marker at start",
			input:    "(up) hello",
			expected: []string{"(up)", "hello"},
		},
		{
			name:     "complex sentence",
			input:    "it (cap) was the best of times, it was",
			expected: []string{"it", "(cap)", "was", "the", "best", "of", "times", ",", "it", "was"},
		},
		// Edge cases: Punctuation as word boundaries
		{
			name:     "punctuation in middle of word",
			input:    "hel!lo",
			expected: []string{"hel", "!", "lo"},
		},
		{
			name:     "word without spaces between punctuation",
			input:    "hello!world",
			expected: []string{"hello", "!", "world"},
		},
		{
			name:     "URL-like pattern with dots",
			input:    "example.com",
			expected: []string{"example", ".", "com"},
		},
		{
			name:     "multiple consecutive punctuation",
			input:    "hello!!!",
			expected: []string{"hello", "!", "!", "!"},
		},
		{
			name:     "punctuation between numbers",
			input:    "3.14",
			expected: []string{"3", ".", "14"},
		},
		{
			name:     "mixed punctuation without spaces",
			input:    "hello,world!test",
			expected: []string{"hello", ",", "world", "!", "test"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Tokenize(tt.input)
			// Handle nil vs empty slice comparison
			if len(result) == 0 && len(tt.expected) == 0 {
				return
			}
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("Tokenize(%q)\n  got: %v\n want: %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestTokenizePreservesMarkers(t *testing.T) {
	input := "hello (up, 2) world (low) test"
	result := Tokenize(input)

	// Check that markers are kept intact
	hasUpMarker := false
	hasLowMarker := false

	for _, token := range result {
		if token == "(up, 2)" {
			hasUpMarker = true
		}
		if token == "(low)" {
			hasLowMarker = true
		}
	}

	if !hasUpMarker {
		t.Error("Expected to find (up, 2) marker in tokens")
	}
	if !hasLowMarker {
		t.Error("Expected to find (low) marker in tokens")
	}
}
