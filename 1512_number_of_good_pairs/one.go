package numberOfGoodPairs

func numIdenticalPairs(nums []int) int {

    var counter int
    for i := 0; i < len(nums); i++ {
        for j := i+1; j < len(nums); j++ {
            if nums[i] == nums[j] {
                counter += 1
            }
        }
    }
    return counter
}