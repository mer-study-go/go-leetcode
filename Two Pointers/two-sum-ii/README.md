# [Two Sum II](https://leetcode.com/problems/two-sum-ii-input-array-is-sorted/)

## Description

Given an array of integers that is already <b><i>sorted in ascending order</i></b>, find two numbers such that they add up to a specific target number.

The function twoSum should return indices of the two numbers such that they add up to the target, where index1 must be less than index2.

**Note:**

* Your returned answers (both index1 and index2) are not zero-based.
* You may assume that each input would have exactly one solution and you may not use the same element twice.

**Example:**

```
Input: numbers = [2,7,11,5], target = 9
Output: [1,2]
Explanation: The sum of 2 and 7 is 9. Therefore index1 = 1, index2 = 2.
```

## Hint
Two Pointers

## Solution 
Of course we could still apply the [Map](../../Maps/TwoSum/README.md) approach, but it costs us O(n) extra space, plus it does not make use of the fact that the input is already sorted. 

**O(n) runtime, O(1) space - Two pointers**

Let's assume we have two indices pointing to the i<sup>th</sup> and j<sup>th</sup> elements, A<sub>i</sub> and A<sub>j</sub> respectively. 

1. A<sub>i</sub> + A<sub>j</sub> > target. Increasing i isn't going to help us, as it makes the sum even bigger. Therefore we should decrement j. 
2. A<sub>i</sub> + A<sub>j</sub> < target. Decreasing j isn't going to help us, as it makes the sum even smaller. Therefore we should increment i. 
3. A<sub>i</sub> + A<sub>j</sub> == target. We have found the answer. 

**Go Solution**
```go
func twoSum(numbers []int, target int) []int {
	left := 0
	right := len(numbers) - 1

	for left < right  {
		if (numbers[left] + numbers[right] == target) {
			return []int {left+1, right+1}
		} else if (numbers[left] + numbers[right] < target) {
			left++;
		} else {
			right--;
		}
	}

	return []int{}
}
```

**Java Solution**

```java
class Solution {
    public int[] twoSum(int[] numbers, int target) {
        int left = 0;
        int right = numbers.length - 1;
        
        while (left < right) {
            if (numbers[left] + numbers[right] == target) {
                return new int[]{left+1, right+1};
            } else if (numbers[left] + numbers[right] < target) {
                left++;
            } else {
                right--;
            }
        }
        
        return new int[]{};
    }
}
```

## Related:

* [Two Sum](../../Maps/TwoSum/README.md)
