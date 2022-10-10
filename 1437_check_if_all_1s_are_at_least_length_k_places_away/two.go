package checkIfAll1sAreAtLeastLengthKPlacesAway

func kLengthApart2(nums []int, k int) bool {

	for i := 0; i < len(nums); i++ {
		if nums[i] == 1 {
			for j := i + 1; j <= i+k && j < len(nums); j++ {
				if nums[j] == 1 {
					return false
				}
			}
		}
	}
	return true
}
