# [Triangle](https://leetcode.com/problems/triangle/)

## Description

Given a triangle, find the minimum path sum from top to bottom. Each step you may move to adjacent numbers on the row below.

For example, given the following triange

```
[
     [2],
    [3,4],
   [6,5,7],
  [4,1,8,3]
]
```

The minimum path sum from top to bottom is `11` (i.e., **2 + 3 + 5 + 1** = 11.)

**Note:**

Bonus point if you are able to do this using only O(n) extra space, where n is the total number of rows in the triangle.

## Hint

Dynamic Programming

## Solution

`Bottom-up` DP: We start from the nodes on the bottom row; the min pathsums for these nodes are the values of the nodes themsleves.
From there, the min pathsum at the `ith` node on the `kth` row would be smaller one of the pathsum of its two children plus the value of itself.

```
minpath[k][i] = min(minpath[k+1][i], minpath[k+1][i+1]) + triangle[k][i]
```

Or even better, since the row minpath[k+1] would be useless after minpath[k] is computed, we can simply set minpath as a 1D array, and iteratively update itself: 

For the kth level:

```
minpath[i] = min(minpath[i], minpath[i+1]) + triangle[k][i];
```

### Solution in Java

```java
public int minimumTotal(List<List<Integer>> triangle) {
  // state: dp[x] = minimum path value from xth level to bottom, dp[0] means the top level to bottom
  int n = triangle.size();
  int[] dp = new int[n];

  // initialize: from the bottom level, the minimum of the pathsum is the number itself on the bottom level
  for (int i = 0; i < n; i++) {
    dp[i] = triangle.get(n-1).get(i);
  }

  // bottom up
  for (int i = n - 2; i >= 0; i--) {
    for (int j = 0; j <= i; j++) {
      dp[j] = Math.min(dp[j], dp[j+1]) + triangle.get(i).get(j);
    }
  }
  return dp[0];
}
```

### Solution in Go

```go
func minimumTotal(triangle [][]int) int {

    size := len(triangle)
    dp := make([]int, size)
    
    for i := 0; i < size; i++ {
        dp[i] = triangle[size-1][i]
    }
    
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
```