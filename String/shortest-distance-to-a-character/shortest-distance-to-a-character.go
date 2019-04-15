package main

import "fmt"

func main() {
	S := "loveleetcode"
	C := byte('e')
	fmt.Println(shortestToChar(S, C))
}

func shortestToChar(S string, C byte) []int {

	n := len(S)
	result := make([]int, n)

	// how to decide the inital value for previousC?
	// Think about the case if no C in S exists, what should the distance be?
	previousC := -1 * n
	for i := 0; i < n; i++ {
		if S[i] == C {
			previousC = i
		}
		result[i] = i - previousC
	}

	// how to decide the inital value for previousC?
	// Think about the case if no C in S exists, what should the distance be?
	previousC = 2 * n
	for i := n - 1; i >= 0; i-- {
		if S[i] == C {
			previousC = i
		}
		result[i] = min(result[i], previousC-i)
	}

	return result
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
