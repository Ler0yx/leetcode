package maximumDifferenceBetweenIncreasingElements

func maximumDifference(nums []int) int {
    var diff int = math.MaxInt

    for i := 0; i < len(nums); i++ {
        for j := i+1; j < len(nums); j++ {
            if nums[i] - nums[j] < diff {
                diff = nums[i]-nums[j]
            }
        }
    }
    if diff < 0 {
        return -diff
    }
    return -1
}