package longestContinuousIncreasingSubsequence

func findLengthOfLCIS(nums []int) int {
    var maxLen int = 1
    var tmp int = 1

    for i := 0; i < len(nums)-1; i++ {
        if nums[i] < nums[i+1] {
            tmp++
        } else if tmp > maxLen {
            maxLen = tmp
            tmp = 1
        } else {
            tmp = 1
        }
        if i == len(nums)-2 && tmp > maxLen {
            maxLen = tmp
        }
    }
    return maxLen
}