package main

import "fmt"

// Longest Substring Without Repeating Characters

func lenOfLongestSubstring(s string) int {

	if len(s) == 0 {
		return 0
	}

	seen := map[byte]int{}
	left, result := 0, 0

	for right := 0; right < len(s); right++ {
		if idx, exists := seen[s[right]]; exists && idx >= left {
			left = idx + 1
		} else {
			result = max(result, right-left+1)
		}

		seen[s[right]] = right
	}

	return result
}

func max(a, b int) int {
	if a > b {return a}
	return b
}

func main() {
	tests := []string{"abcabcbb", "bbbbb", "pwwkew", ""}

	for _, test := range tests {
		fmt.Printf("%q - %d\n", test, lenOfLongestSubstring(test))
	}
}