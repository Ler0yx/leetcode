package minimumOperationsToMakeTheArrayIncreasing

func minOperations(nums []int) int {
	var ops int

	for i := 0; i < len(nums)-1; i++ {
		if nums[i+1] <= nums[i] {
			ops += nums[i] - nums[i+1] + 1
			nums[i+1] = nums[i] + 1
		}
	}
	return ops
}
