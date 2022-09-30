package containsDuplicate

func containsDuplicate(nums []int) bool {

	numsMap := make(map[int]int)

	for i := 0; i < len(nums); i++ {
		if _, ok := numsMap[nums[i]]; ok {
			return true
		}
		numsMap[nums[i]] = 1
	}
	return false
}
