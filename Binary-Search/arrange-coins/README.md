# [Arrange Coins](https://leetcode.com/problems/arranging-coins/)

## Description

You have a total of n coins that you want to form in a staircase shape, where every k-th row must have exactly k coins. 

Given n, find the total number of full staircase rows that can be formed. 

n is a non-negative integer and fits within the range of a 32-bit signed integer. 

**Example 1:**

```
n = 5
The coins can form the following rows:
¤
¤ ¤
¤ ¤
Because the 3rd row is incomplete, we return 2.
```

**Example 2:**

```
n = 8
The coins can form the following rows:
¤
¤ ¤
¤ ¤ ¤
¤ ¤
Because the 4th row is incomplete, we return 3.
```

## Hint

Binary Search

## Solution 

Note: The sum of coins in a full staircase is `n*(n+1)/2`.

### Algorithm

1. Take middle number
2. Return `middle` if `sum(middle)` is equal to the number. 

`sum(middle)` < `number` or `sum(middle)` > `number` :
* Make rightmost index as `middle-1` if `sum(middle)` > `number`
* Make leftmost index as `middle+1` if `sum(middle)` < `number`

### Solution In Java

```java
public int arrangeCoins(int n) {
    return binarySearch((long)n, 1, (long)n);
}

public int binarySearch(long n, long start, long end) {
    while (start + 1 < end) {
        long mid = start + (end - start) / 2;
        if (sum(mid) <= n) {
            start = mid;
        } else {
            end = mid;
        }
    }

    if (sum(end) == n) {
        return (int)end;
    }

    return (int)start;
}

public long sum(long n) {
    return n * (n + 1) / 2;
}
```

```go
func arrangeCoins(n int) int {
    return helper(n, 1, n)
}

func helper(n, start, end int) int {
    for start + 1 < end {
        mid := start + (end - start) / 2
        if sum(mid) <= n {
            start = mid
        } else {
            end = mid
        }
    }
    
    if sum(end) == n {
         return end
    }
    
    return start
}

func sum(n int) int {
    return n * (n + 1) / 2
}
```

## Complexity Analysis

As it's a binary search way of solving the problem.
* Time Complexity: O(log(n)). 
* Space Complexity: O(1). We only need two variables (`start`, `end`) for the current solutions. 
