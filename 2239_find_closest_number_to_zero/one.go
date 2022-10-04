package findClosestNumberToZero

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func findClosestNumber(nums []int) int {

	var minGreaterZero int = 1000000
	var minSmallerZero int = -1000000

	for i := 0; i < len(nums); i++ {
		if nums[i] <= 0 && nums[i] > minSmallerZero {
			minSmallerZero = nums[i]
		} else if nums[i] >= 0 && nums[i] < minGreaterZero {
			minGreaterZero = nums[i]
		}
	}

	if minGreaterZero == minSmallerZero {
		return minGreaterZero
	}

	var min int

	if abs(minGreaterZero) <= abs(minSmallerZero) && minGreaterZero < 1000000 {
		min = minGreaterZero
	} else {
		min = minSmallerZero
	}

	return min
}
