package main

import "fmt"

func main() {

	nums := []int{2, 7, 11, 5}
	target := 9
	fmt.Println(twoSum(nums, target))

}

func twoSum(nums []int, target int) []int {
	dataset := make(map[int]int)

	for index, value := range nums {
		item, ok := dataset[target - value]
		if ok == true {
			return []int{item, index}
		}
		dataset[value] = index
	}

	return []int{}
}
