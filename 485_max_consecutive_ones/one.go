package maxConsecutiveOnes

func findMaxConsecutiveOnes(nums []int) int {
	var max int
	var tmp int
	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			tmp++
		} else {
			if tmp > max {
				max = tmp
			}
			tmp = 0
		}
	}
	if tmp > max {
		max = tmp
	}
	return max
}
