package checkIfAll1sAreAtLeastLengthKPlacesAway

func kLengthApart(nums []int, k int) bool {
    
    for i := 0; i < len(nums); i++ {
        if nums[i] == 1 {
            for j := i; j < len(nums); j++ {
                if j < i+k && nums[j] == 1 {
                    return false
                }
            }
        i += k
        }   
    }
    return true
}