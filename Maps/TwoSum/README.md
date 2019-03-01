# [1. Two Sum](https://leetcode.com/problems/two-sum/)

## Description

Given an array of integers, return indices of the two numbers such that they add up to a specific target.

You may assume that each input would have exactly one solution, and you may not use the same element twice.

Example:

```
Given nums = [2, 7, 11, 15], target = 9,

Because nums[0] + nums[1] = 2 + 7 = 9,
return [0, 1].
```

## Hint

```
a + b = target
```

```
a = target - b
```

## Solution

**O(n) runtime, O(n) space - Map**

We could reduce the runtime complexity of looking up a value to O(1) using a hash map that maps a value to its index. 

**Go Solution**

```go
func twoSum(nums []int, target int) []int {
	dataset := map[int]int{}
	for index, value := range nums {
		x := target - nums[index]
		item, ok := dataset[x]
		if ok == true {
			return []int{item, index}
		}
		dataset[value] = index
	}
	return []int{}
}
```

**Java Solution**

```java
class Solution {
    public int[] twoSum(int[] nums, int target) {
        Map<Integer, Integer> map = new HashMap<>();
        for (int i = 0; i < nums.length; i++) {
            int x = nums[i];
            if (map.containsKey(target - x)) {
                return new int[]{map.get(target - x), i};
            }
            map.put(x, i);
        }
        return new int[]{};
    }
}
```

## Follow up: 
What if the given input is already sorted in ascending order? See Question [Two Sum II - Input array is sorted](../../Two%20Pointers/README.md)