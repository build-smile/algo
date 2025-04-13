package service

import (
	"testing"
)

func TestLongestSubstringWithoutDuplicate(t *testing.T) {
	tests := []struct {
		name     string
		input    string
		expected int
	}{
		{
			name:     "empty string",
			input:    "",
			expected: 0,
		},
		{
			name:     "single character",
			input:    "a",
			expected: 1,
		},
		{
			name:     "all unique characters",
			input:    "abcdef",
			expected: 6,
		},
		{
			name:     "all duplicate characters",
			input:    "aaaaaa",
			expected: 1,
		},
		{
			name:     "substring in middle",
			input:    "abcabcbb",
			expected: 3,
		},
		{
			name:     "substring at the end",
			input:    "pwwkew",
			expected: 3,
		},
		{
			name:     "multiple repeating patterns",
			input:    "aabcaabcda",
			expected: 4,
		},
		{
			name:     "special characters",
			input:    "!@#$%^&*()",
			expected: 10,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := LongestSubstringWithoutDuplicate(tt.input)
			if result != tt.expected {
				t.Errorf("LongestSubstringWithoutDuplicate(%q) = %d, want %d", tt.input, result, tt.expected)
			}
		})
	}
}
