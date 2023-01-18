package findAllNumbersDisappearedInAnArray

func findDisappearedNumbers(nums []int) []int {
	ints := make(map[int]bool)
	var res []int
	for i := 0; i < len(nums); i++ {
		ints[nums[i]] = true
	}
	for j := 1; j <= len(nums); j++ {
		if !ints[j] {
			res = append(res, j)
		}
	}
	return res
}
