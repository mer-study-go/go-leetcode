package main

import (
	"fmt"
	"strings"
)

func main() {
	target := "A man, a plan, a canal: Panama"
	fmt.Println(isPalindrome(target))
}

func isPalindrome(s string) bool {
	if len(s) == 0 {
		return true
	}
	s = strings.ToLower(s)
	head := 0
	tail := len(s) - 1
	for head < tail {
		for head < tail && !isChar(s[head]) {
			head++
		}
		for head < tail && !isChar(s[tail]) {
			tail--
		}
		if s[head] != s[tail] {
			return false
		}
		head++
		tail--
	}
	return true
}

func isChar(c byte) bool {
	if ('a' <= c && c <= 'z') || ('0' <= c && c <= '9') {
		return true
	}
	return false
}
