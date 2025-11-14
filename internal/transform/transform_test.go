package transform

import (
	"reflect"
	"testing"
)

func TestConvertHexAndBin(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "hex conversion",
			input:    []string{"1E", "(hex)", "files"},
			expected: []string{"30", "files"},
		},
		{
			name:     "bin conversion",
			input:    []string{"10", "(bin)", "years"},
			expected: []string{"2", "years"},
		},
		{
			name:     "multiple hex conversions",
			input:    []string{"1A", "(hex)", "and", "FF", "(hex)"},
			expected: []string{"26", "and", "255"},
		},
		{
			name:     "invalid hex - keeps both word and marker",
			input:    []string{"XYZ", "(hex)", "test"},
			expected: []string{"XYZ", "(hex)", "test"},
		},
		{
			name:     "invalid bin - keeps both word and marker",
			input:    []string{"102", "(bin)", "test"},
			expected: []string{"102", "(bin)", "test"},
		},
		{
			name:     "no conversion markers",
			input:    []string{"hello", "world"},
			expected: []string{"hello", "world"},
		},
		{
			name:     "hex with punctuation",
			input:    []string{"1E", "(hex).", "test"},
			expected: []string{"30", "test"},
		},
		{
			name:     "bin with punctuation",
			input:    []string{"101", "(bin),", "test"},
			expected: []string{"5", "test"},
		},
		{
			name:     "marker at start with no previous word",
			input:    []string{"(hex)", "test"},
			expected: []string{"(hex)", "test"},
		},
		{
			name:     "zero hex",
			input:    []string{"0", "(hex)", "test"},
			expected: []string{"0", "test"},
		},
		{
			name:     "zero bin",
			input:    []string{"0", "(bin)", "test"},
			expected: []string{"0", "test"},
		},
		{
			name:     "large hex number",
			input:    []string{"FFFF", "(hex)", "large"},
			expected: []string{"65535", "large"},
		},
		{
			name:     "large bin number",
			input:    []string{"11111111", "(bin)", "large"},
			expected: []string{"255", "large"},
		},
		{
			name:     "mixed conversions",
			input:    []string{"1111", "(bin)", "add", "1E", "(hex)"},
			expected: []string{"15", "add", "30"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ConvertHexAndBin(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ConvertHexAndBin(%v)\n  got: %v\n want: %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestConvertHexAndBinEdgeCases(t *testing.T) {
	// Test empty slice
	result := ConvertHexAndBin([]string{})
	if len(result) != 0 {
		t.Errorf("Expected empty slice, got %v", result)
	}

	// Test single element
	result = ConvertHexAndBin([]string{"test"})
	if len(result) != 1 || result[0] != "test" {
		t.Errorf("Expected [test], got %v", result)
	}
}

func TestApplyCaseRules(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "uppercase single word",
			input:    []string{"hello", "(up)"},
			expected: []string{"HELLO"},
		},
		{
			name:     "lowercase single word",
			input:    []string{"WORLD", "(low)"},
			expected: []string{"world"},
		},
		{
			name:     "capitalize single word",
			input:    []string{"hello", "(cap)"},
			expected: []string{"Hello"},
		},
		{
			name:     "uppercase multiple words",
			input:    []string{"one", "two", "three", "(up, 2)"},
			expected: []string{"one", "TWO", "THREE"},
		},
		{
			name:     "lowercase multiple words",
			input:    []string{"ONE", "TWO", "THREE", "(low, 2)"},
			expected: []string{"ONE", "two", "three"},
		},
		{
			name:     "capitalize multiple words",
			input:    []string{"the", "new", "york", "times", "(cap, 4)"},
			expected: []string{"The", "New", "York", "Times"},
		},
		{
			name:     "marker at start - no previous word, keeps both",
			input:    []string{"(up)", "hello"},
			expected: []string{"(up)", "hello"},
		},
		{
			name:     "count exceeds available words",
			input:    []string{"only", "two", "(up, 10)"},
			expected: []string{"ONLY", "TWO"},
		},
		{
			name:     "invalid marker with negative count",
			input:    []string{"test", "(up, -1)"},
			expected: []string{"test", "(up, -1)"},
		},
		{
			name:     "invalid marker with zero count",
			input:    []string{"test", "(low, 0)"},
			expected: []string{"test", "(low, 0)"},
		},
		{
			name:     "invalid marker with empty count",
			input:    []string{"test", "(up, )"},
			expected: []string{"test", "(up, )"},
		},
		{
			name:     "no markers",
			input:    []string{"hello", "world"},
			expected: []string{"hello", "world"},
		},
		{
			name:     "capitalize already uppercase",
			input:    []string{"HELLO", "(cap)"},
			expected: []string{"Hello"},
		},
		{
			name:     "uppercase already uppercase",
			input:    []string{"HELLO", "(up)"},
			expected: []string{"HELLO"},
		},
		{
			name:     "multiple markers in sequence",
			input:    []string{"this", "is", "exciting", "(up, 2)"},
			expected: []string{"this", "IS", "EXCITING"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ApplyCaseRules(tt.input)
			if !reflect.DeepEqual(result, tt.expected) {
				t.Errorf("ApplyCaseRules(%v)\n  got: %v\n want: %v", tt.input, result, tt.expected)
			}
		})
	}
}

func TestCapitalize(t *testing.T) {
	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "Hello"},
		{"WORLD", "World"},
		{"tEsT", "Test"},
		{"a", "A"},
		{"", ""},
		{"SHOUTING", "Shouting"},
		{"αλφα", "Αλφα"}, // Greek Unicode
	}

	for _, tt := range tests {
		t.Run(tt.input, func(t *testing.T) {
			result := Capitalize(tt.input)
			if result != tt.expected {
				t.Errorf("Capitalize(%q) = %q, want %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestParseMarkerCount(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
		valid    bool
	}{
		{"valid count", "(up, 3)", 3, true},
		{"valid count with spaces", "(low, 5)", 5, true},
		{"negative count", "(up, -1)", 0, false},
		{"zero count", "(cap, 0)", 0, false},
		{"no count", "(up)", 0, false},
		{"empty count", "(up, )", 0, false},
		{"non-numeric", "(up, abc)", 0, false},
		{"large count", "(up, 100)", 100, true},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, valid := ParseMarkerCount(tt.input)
			if valid != tt.valid {
				t.Errorf("ParseMarkerCount(%q) valid = %v, want %v", tt.input, valid, tt.valid)
			}
			if valid && result != tt.expected {
				t.Errorf("ParseMarkerCount(%q) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}

func TestFixArticles(t *testing.T) {
	tests := []struct {
		name     string
		input    []string
		expected []string
	}{
		{
			name:     "a before vowel",
			input:    []string{"a", "apple"},
			expected: []string{"an", "apple"},
		},
		{
			name:     "a before h",
			input:    []string{"a", "house"},
			expected: []string{"an", "house"},
		},
		{
			name:     "a before consonant",
			input:    []string{"a", "book"},
			expected: []string{"a", "book"},
		},
		{
			name:     "A before vowel uppercase",
			input:    []string{"A", "apple"},
			expected: []string{"An", "apple"},
		},
		{
			name:     "multiple articles",
			input:    []string{"a", "apple", "and", "a", "orange"},
			expected: []string{"an", "apple", "and", "an", "orange"},
		},
		{
			name:     "article at end",
			input:    []string{"I", "saw", "a"},
			expected: []string{"I", "saw", "a"},
		},
		{
			name:     "all vowels",
			input:    []string{"a", "elephant", "a", "igloo", "a", "octopus", "a", "umbrella"},
			expected: []string{"an", "elephant", "an", "igloo", "an", "octopus", "an", "umbrella"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FixArticles(tt.input)
			for i := range result {
				if result[i] != tt.expected[i] {
					t.Errorf("FixArticles(%v)\n  got: %v\n want: %v", tt.input, result, tt.expected)
					break
				}
			}
		})
	}
}

func TestApplyPunctuationRules(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "space before comma",
			input:    "Hello ,world",
			expected: "Hello, world",
		},
		{
			name:     "space before period",
			input:    "End of sentence .",
			expected: "End of sentence.",
		},
		{
			name:     "ellipsis grouping",
			input:    "Wait ...",
			expected: "Wait...",
		},
		{
			name:     "multiple punctuation",
			input:    "What !? No way",
			expected: "What!? No way",
		},
		{
			name:     "no space after comma",
			input:    "Hello,world",
			expected: "Hello, world",
		},
		{
			name:     "complex punctuation",
			input:    "Hello , world ! This is amazing .",
			expected: "Hello, world! This is amazing.",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := ApplyPunctuationRules(tt.input)
			if result != tt.expected {
				t.Errorf("ApplyPunctuationRules(%q)\n  got: %q\n want: %q", tt.input, result, tt.expected)
			}
		})
	}
}

func TestFixQuotes(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "simple double quotes",
			input:    "He said \" hello \"",
			expected: "He said \"hello\"",
		},
		{
			name:     "simple single quotes",
			input:    "He said ' hello '",
			expected: "He said 'hello'",
		},
		{
			name:     "quotes with multiple words",
			input:    "He said ' hello world '",
			expected: "He said 'hello world'",
		},
		{
			name:     "no spaces around quotes",
			input:    "\"hello\"",
			expected: "\"hello\"",
		},
		{
			name:     "apostrophe in word",
			input:    "It's fine",
			expected: "It's fine",
		},
		{
			name:     "mixed quotes and apostrophes",
			input:    "It's John's \"book\"",
			expected: "It's John's \"book\"",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := FixQuotes(tt.input)
			if result != tt.expected {
				t.Errorf("FixQuotes(%q)\n  got: %q\n want: %q", tt.input, result, tt.expected)
			}
		})
	}
}
