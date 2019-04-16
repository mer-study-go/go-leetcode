package main

import "fmt"

func main() {
	triangle := [][]int{
		[]int{2},
		[]int{3, 4},
		[]int{6, 5, 7},
		[]int{4, 1, 8, 3},
	}
	fmt.Println(minimumTotal(triangle))
}

func minimumTotal(triangle [][]int) int {
	// state: dp[x] = minimum path value from xth level to bottom, dp[0] means the top level to bottom
	size := len(triangle)
	dp := make([]int, size)

	// initialize: from the bottom level, the minimum of the pathsum is the number itself on the bottom level
	for i := 0; i < size; i++ {
		dp[i] = triangle[size-1][i]
	}

	// bottom up, from the second row from the bottom
	for i := size - 2; i >= 0; i-- {
		for j := 0; j <= i; j++ {
			dp[j] = min(dp[j], dp[j+1]) + triangle[i][j]
		}
	}

	return dp[0]
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}
