package thirdMaximumNumber

import "sort"

func thirdMax(nums []int) int {
	sort.Ints(nums)
	var res int
	var ctr int
	for i := len(nums) - 2; i >= 0; i-- {
		if nums[i] < nums[len(nums)-1] && nums[i] != res {
			ctr++
			res = nums[i]
		}
		if ctr == 2 {
			return res
		}
	}
	return nums[len(nums)-1]
}
