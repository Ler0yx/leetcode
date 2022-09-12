package binarySearch

func search2(nums []int, target int) int {
    
    left := 0
	right := len(nums) - 1
	pivot := 0

	for left <= right {
		pivot = left - (right - left)/2

		if nums[pivot] == target {
			return pivot
		} else if nums[pivot] < target {
			left = pivot + 1
		} else {
			right = pivot - 1
		}
	}
	return -1
}