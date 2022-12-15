package minimumSubsequenceInNonIncreasingOrder

import "sort"

func minSubsequence(nums []int) []int {
	var sum int
	var sum2 int
	var solution []int
	sort.Ints(nums)

	for i := 0; i < len(nums); i++ {
		sum += nums[i]
	}
	for j := len(nums) - 1; j >= 0; j-- {
		solution = append(solution, nums[j])
		sum -= nums[j]
		sum2 += nums[j]
		if sum2 > sum {
			return solution
		}
	}

	return []int{}
}
