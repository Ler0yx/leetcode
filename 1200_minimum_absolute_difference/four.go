package minimumAbsoluteDifference

import (
	"math"
	"sort"
)

func abs(n int) int {
	if n < 0 {
		return -n
	}
	return n
}

func minimumAbsDifference(arr []int) [][]int {
	var min int = math.MaxInt
	var res [][]int
	for i := 0; i < len(arr); i++ {
		for j := 0; j < len(arr); j++ {
			if j == i {
				continue
			}
			if abs(arr[i]-arr[j]) < min {
				min = abs(arr[i] - arr[j])
			}
		}
	}
	for k := 0; k < len(arr); k++ {
		for l := k + 1; l < len(arr); l++ {
			if arr[k]-arr[l] == min {
				res = append(res, []int{arr[l], arr[k]})
				continue
			}
			if arr[l]-arr[k] == min {
				res = append(res, []int{arr[k], arr[l]})
			}
		}
	}
	sort.Slice(res, func(i, j int) bool {
		return res[i][0]+res[i][1] < res[j][0]+res[j][1]
	})
	return res
}
