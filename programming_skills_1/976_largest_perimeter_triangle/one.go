package main

import (
	"fmt"
	"sort"
)

func largestPerimeter(nums []int) int {

	sort.Ints(nums)

	for i := len(nums) - 1; i > 1; i-- {
		if nums[i] < nums[i-1]+nums[i-2] {
			return nums[i] + nums[i-1] + nums[i-2]
		}
	}
	return 0
}

func main() {

	nums := []int{1, 100, 99999, 99999, 99999, 99999, 99999, 3, 5, 4567, 345, 345, 1}

	fmt.Println(largestPerimeter(nums))
}
