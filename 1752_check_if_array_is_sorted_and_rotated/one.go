package checkIfArrayIsSortedAndRotated

import "math"

func check(nums []int) bool {
	var min int = math.MaxInt
	var minIdx int
	for i := len(nums) - 1; i >= 0; i-- {
		if nums[i] < min {
			min = nums[i]
			for k := i; k >= 0; k-- {
				if nums[k] == min {
					minIdx = k
				} else {
					break
				}
			}
		}
	}
	nums = append(nums[minIdx:], nums[:minIdx]...)
	for j := 0; j < len(nums)-1; j++ {
		if nums[j] > nums[j+1] {
			return false
		}
	}
	return true
}
