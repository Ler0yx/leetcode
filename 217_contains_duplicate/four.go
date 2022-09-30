package containsDuplicate

func containsDuplicate4(nums []int) bool {

	numsMap := make(map[int]bool)

	for i := 0; i < len(nums); i++ {
		if numsMap[nums[i]] {
			return true
		}
		numsMap[nums[i]] = true
	}
	return false
}
