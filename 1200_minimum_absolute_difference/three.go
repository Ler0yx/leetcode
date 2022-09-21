package minimumAbsoluteDifference

import (
	"math"
	"sort"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minimumAbsDifference(arr []int) [][]int {

	var minDiff = math.MaxInt
	solution := make([][]int, 0)

	sort.Ints(arr)

	for i := 0; i < len(arr)-1; i++ {
		if abs(arr[i]-arr[i+1]) < minDiff {
			minDiff = abs(arr[i] - arr[i+1])
		}
	}

	for i := 0; i < len(arr)-1; i++ {
		if abs(arr[i]-arr[i+1]) == minDiff {
			tmp := []int{arr[i], arr[i+1]}
			solution = append(solution, tmp)
		}
	}
	return solution
}
