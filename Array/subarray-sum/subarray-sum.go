package main

import "fmt"

func subarraySum(nums []int) []int {

	var keyValue = make(map[int]int)
	sum := 0
	keyValue[sum] = -1
	for i := 0; i < len(nums); i++ {
		sum += nums[i]
		if item, ok := keyValue[sum]; ok {
			return []int{item+1, i}
		}
		keyValue[sum] = i
	}
	return []int{}
}

func main() {
	nums := []int{-3, 1, 2, -3, 4}
	output := subarraySum(nums)

	fmt.Println(output)
}