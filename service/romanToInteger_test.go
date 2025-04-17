package service

import (
	"testing"
)

func TestRomanToInt(t *testing.T) {
	tests := []struct {
		name   string
		input  string
		expect int
	}{
		{"Empty string", "", 0},
		{"I", "I", 1},
		{"III", "III", 3},
		{"IV", "IV", 4},
		{"V", "V", 5},
		{"IX", "IX", 9},
		{"LVIII", "LVIII", 58},
		{"MCMXCIV", "MCMXCIV", 1994},
		{"MMMMCMXCIX", "MMMMCMXCIX", 4999},
		{"Invalid IIIV", "IIIV", 0},
		{"Invalid MMMMC", "MMMMC", 0},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			actual := RomanToInt(tt.input)
			if actual != tt.expect {
				t.Errorf("RomanToInt(%q) = %v; want %v", tt.input, actual, tt.expect)
			}
		})
	}
}