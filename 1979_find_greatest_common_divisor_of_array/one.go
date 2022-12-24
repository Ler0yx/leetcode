package findGreatestCommonDivisorOfArray

import "math"

func findGCD(nums []int) int {

	var smallest int = math.MaxInt
	var biggest int = math.MinInt
	for i := 0; i < len(nums); i++ {
		if nums[i] < smallest {
			smallest = nums[i]
		}
		if nums[i] > biggest {
			biggest = nums[i]
		}
	}
	for j := smallest; j > 0; j-- {
		if smallest%j == 0 && biggest%j == 0 {
			return j
		}
	}
	return 0
}
