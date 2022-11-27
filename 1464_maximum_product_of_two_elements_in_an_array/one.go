package maximumProductOfTwoElementsInAnArray

func maxProduct(nums []int) int {
	var max1 int
	var max2 int

	if nums[0] > nums[1] {
		max1 = nums[0]
		max2 = nums[1]
	} else {
		max1 = nums[1]
		max2 = nums[0]
	}

	for i := 2; i < len(nums); i++ {
		if nums[i] > max1 {
			if max1 > max2 {
				max2 = max1
			}
			max1 = nums[i]
			continue
		}
		if nums[i] > max2 {
			max2 = nums[i]
		}
	}
	return (max1 - 1) * (max2 - 1)
}
