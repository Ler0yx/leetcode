package intersectionOfMultipleArrays

import "sort"

func intersection(nums [][]int) []int {
	integers := make(map[int]int)
	keys := nums[0]
	var res []int
	for i := 0; i < len(nums); i++ {
		for j := 0; j < len(nums[i]); j++ {
			integers[nums[i][j]] += 1
		}
	}
	for k := 0; k < len(keys); k++ {
		if integers[keys[k]] == len(nums) {
			res = append(res, keys[k])
		}
	}
	sort.Ints(res)
	return res
}
