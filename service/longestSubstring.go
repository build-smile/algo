package service

import (
	"fmt"
)

func LongestSubstringWithoutDuplicate(s string) int {
	seen := make(map[string]string)
	items := make([]string, 0)
	temp := ""
	for i := range s {
		for j := i; j < len(s); j++ {
			ch := fmt.Sprintf("%c", s[j])
			if _, ok := seen[ch]; ok {
				items = append(items, temp)
				temp = ""
				seen = make(map[string]string)
				break
			}
			temp += ch
			seen[ch] = ch
		}
		if temp != "" {
			items = append(items, temp)
		}

	}

	maxItem := 0
	for _, item := range items {
		if len(item) > maxItem {
			maxItem = len(item)
		}
	}

	return maxItem
}
