package findpivotindex

func pivotIndex(nums []int) int {

	if len(nums) == 0 {
		return -1
	}

	sumTotal := 0
	for index := range nums {
		sumTotal += nums[index]
	}

	sumSoFar := 0
	for index := range nums {
		if sumSoFar == sumTotal-sumSoFar-nums[index] {
			return index
		}
		sumSoFar += nums[index]
	}

	return -1
}
