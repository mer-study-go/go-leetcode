package main 

import "fmt"

func main() {
	numbers := []int{2, 7, 11, 15}
	target := 9
	fmt.Println(twoSum(numbers, target))
}

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