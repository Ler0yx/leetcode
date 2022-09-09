package moveZeroes

func moveZeroes(nums []int) {

	for i := 0; i < len(nums)-1; i++ {
		if nums[i] == 0 {
			if nums[i+1] != 0 {
				nums[i] = nums[i+1]
				nums[i+1] = 0
			} else {
				for d := i; d < len(nums); d++ {
					if nums[d] != 0 {
						nums[i] = nums[d]
						nums[d] = 0
						break
					}
				}
			}
		}
	}
}
