package rules

import "testing"

func TestIsWhitespace(t *testing.T) {
	tests := []struct {
		input    byte
		expected bool
	}{
		{' ', true},
		{'\t', true},
		{'\n', true},
		{'\r', true},
		{'a', false},
		{'1', false},
		{'-', false},
		{0, false},
		{',', false},
	}

	for _, test := range tests {
		result := isWhitespace(test.input)
		if result != test.expected {
			t.Errorf("isWhitespace(%q) = %v; want %v", test.input, result, test.expected)
		}
	}
}
