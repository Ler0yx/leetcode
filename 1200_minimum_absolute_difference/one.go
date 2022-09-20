// package minimumAbsoluteDifference

// import (
// 	"math"
// 	"sort"
// )

// // Too slow
// func abs(x int) int {
// 	if x < 0 {
// 		return -x
// 	}
// 	return x
// }

// func minimumAbsDifference(arr []int) [][]int {

// 	var minDiff = math.MaxInt
// 	solution := make([][]int, 0)

// findminDiff:
// 	for i := 0; i < len(arr); i++ {
// 		for j := 0; j < len(arr); j++ {
// 			if i == j {
// 				continue
// 			} else if abs(arr[i]-arr[j]) < minDiff {
// 				minDiff = abs(arr[i] - arr[j])
// 			}
// 			if minDiff == 1 {
// 				break findminDiff
// 			}
// 		}
// 	}

// 	sort.Ints(arr)

// 	for i := 0; i < len(arr)-1; i++ {
// 		for j := i + 1; j < len(arr); j++ {
// 			if abs(arr[i]-arr[j]) == minDiff {
// 				tmp := []int{arr[i], arr[j]}
// 				solution = append(solution, tmp)
// 				break
// 			}
// 		}
// 	}
// 	return solution
// }
