package main

import "fmt"

func removeDuplicates(nums []int) int {
	slow := 0
	for _, value := range nums {
		if nums[slow] != value {
			slow++
			nums[slow] = value
		}
	}

	return slow + 1
}

func main() {
	nums := []int{0, 0, 1, 1, 1, 2, 2, 3, 3, 4}
	fmt.Println(removeDuplicates(nums))
}
