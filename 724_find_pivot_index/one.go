package findPivotIndex

func sum(s []int) int {

	var sum int

	for _, v := range s {
		sum += v
	}
	return sum
}

func pivotIndex(nums []int) int {

	for i, _ := range nums {
		if sum(nums[:i]) == sum(nums[i+1:]) {
			return i
		}
	}
	return -1
}
