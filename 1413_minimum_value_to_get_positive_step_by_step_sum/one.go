package minimumValueToGetPositiveStepByStepSum

func abs (x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func minStartValue(nums []int) int {
    var val int = 1
    var sum int = 1
    for i := 0; i < len(nums); i++ {
        sum += nums[i]
        if sum < 1 {
            val += abs(sum)+1
            sum += abs(sum)+1
        }
    }
    return val
}