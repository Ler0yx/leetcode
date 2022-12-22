package longestContinuousIncreasingSubsequence

func findLengthOfLCIS(nums []int) int {
	maxLen, tmp := 1, 1
	for i, num := range nums {
		switch {
		case num < nums[i+1]:
			tmp++
		case tmp > maxLen:
			maxLen = tmp
			tmp = 1
		default:
			tmp = 1
		}
	}
	return maxLen
}
