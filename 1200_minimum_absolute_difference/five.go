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

	sort.Ints(arr)
	for i := 0; i < len(arr)-1; i++ {
		if abs(arr[i]-arr[i+1]) < min {
			min = abs(arr[i] - arr[i+1])
		}
	}

	for k := 0; k < len(arr); k++ {
		for l := k + 1; l < len(arr); l++ {
			if arr[l]-arr[k] == min {
				res = append(res, []int{arr[k], arr[l]})
			} else if arr[l]-arr[k] > min {
				break
			}
		}
	}

	sort.Slice(res, func(i, j int) bool {
		return res[i][0]+res[i][1] < res[j][0]+res[j][1]
	})
	return res
}
