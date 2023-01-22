package minimumOperationsToMakeTheArrayIncreasing

func minOperations(nums []int) int {
	var res int
	for i := 1; i < len(nums); i++ {
		if nums[i-1] >= nums[i] {
			res += nums[i-1] - nums[i] + 1
			nums[i] = nums[i-1] + 1
		}
	}
	return res
}
