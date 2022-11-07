package missingNumber

import "sort"

func missingNumber(nums []int) int {
	sort.Ints(nums)
	var i int

	for ; i < len(nums); i++ {
		if i != nums[i] {
			return i
		}
	}
	return i
}
