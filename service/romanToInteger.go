package service

import "fmt"

func RomanToInt(s string) int {
	romanMap := map[byte]int{
		'I': 1,
		'V': 5,
		'X': 10,
		'L': 50,
		'C': 100,
		'D': 500,
		'M': 1000,
	}

	result := 0

	for i := 0; i < len(s); i++ {
		if invalidateRoman(i, s) {
			return 0
		}
		val := romanMap[s[i]]
		if isHasNextVal(s, i) {
			nextVal := romanMap[s[i+1]]
			if nextVal > val {
				result -= val
				continue
			}
		}
		result += val
	}
	return result
}

func invalidateRoman(index int, s string) bool {
	if index < 3 {
		return false
	}
	if s[index-1] == s[index] && s[index-2] == s[index] && s[index-3] == s[index] {
		fmt.Printf("invalid format roman %c,%c,%c,%c \n", s[index], s[index-1], s[index-2], s[index-3])
		return true
	}
	return false
}

func isHasNextVal(s string, i int) bool {
	return i+1 < len(s)
}
