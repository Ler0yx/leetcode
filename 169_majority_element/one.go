package majorityElement

func majorityElement(nums []int) int {

	counts := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		counts[nums[i]] += 1
		if counts[nums[i]] > len(nums)/2 {
			return nums[i]
		}
	}
	return 0
}
