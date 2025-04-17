package service

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

func isHasNextVal(s string, i int) bool {
	return i+1 < len(s)
}
