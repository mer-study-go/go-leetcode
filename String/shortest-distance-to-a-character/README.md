# [Shortest Distance to a Character](https://leetcode.com/problems/shortest-distance-to-a-character/)

## Description

Given a string S and a character C, return an array of integers representing the shortest distance from the character C in the string.

**Example 1:**

><strong>Input:</strong> S = "loveleetcode", C = 'e' <br />
><strong>Output:</strong> [3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0]

**Note**

1. `S` string length is in `[1, 10000].`
2. `C` is a single character, and guaranteed to be in string `S`.
3. All letters in `S` and `C` are lowercase.

## Hint
Min(number[c_from_left], number[c_from_right])

## Solution

### Intuition

For each index `S[i]`, let's try to find the distance to the next character `C` going left, and going right. The answer is the minimum of these two values.

### Algorithm

When going left to right, we'll remember the index `prev` of the last character `C` we've seen. Then the answer is `i-prev`.

When going right to left, we'll remeber the index `prev` of the last Character `C` we've senn. Then the answer is `prev-i`.

We take the minimum of these two answers to create our final answer.

### Solution in Go

```go
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
```
## Complexity Analysis 

* Time Complexity: O(N), where `N` is the length of `S`. We scan through the string twice. 
* Space Complexity: O(N), the size of `result`