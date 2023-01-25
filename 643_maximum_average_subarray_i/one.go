package maximumAverageSubarrayI

import "math"

func findMaxAverage(nums []int, k int) float64 {
	var sum int
	var maxAvg float64 = math.MinInt
	for i := 0; i < len(nums)-k+1; i++ {
		for j := i; j < i+k; j++ {
			sum += nums[j]
		}
		if float64(sum)/float64(k) > maxAvg {
			maxAvg = float64(sum) / float64(k)
		}
		sum = 0
	}
	return maxAvg
}
