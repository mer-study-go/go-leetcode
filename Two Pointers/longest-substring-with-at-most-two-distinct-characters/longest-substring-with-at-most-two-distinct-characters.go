package main

import "fmt"

func main() {
	s := "eceba"
	fmt.Println(lengthOfLongestSubstringTwoDistinct(s))
}

func lengthOfLongestSubstringTwoDistinct(s string) int {
	i := 0
	j := -1
	maxLen := 0
	for k := 1; k < len(s); k++ {
		if s[k] == s[k-1] {
			continue
		}
		if j >= 0 && s[k] != s[j] {
			maxLen = max(maxLen, k - i)
			i = j+1
		}
		j = k - 1
	}
	return max(maxLen, len(s) - i)
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}
