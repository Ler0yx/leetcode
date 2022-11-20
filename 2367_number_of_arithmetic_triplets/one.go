package numberOfArithmeticTriplets

func arithmeticTriplets(nums []int, diff int) int {
    
    var counter int

    for i := 0; i < len(nums)-2; i++ {
        for j := i+1; j < len(nums)-1; j++ {
            for k := i+2; k < len(nums); k++ {
                if nums[j] - nums[i] == diff && nums[k] - nums[j] == diff {
                    counter++
                }
            }
        }
    }
    return counter
}