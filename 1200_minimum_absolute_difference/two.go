// package minimumAbsoluteDifference

// import "sort"

// func abs(x int) int {
// 	if x < 0 {
// 		return -x
// 	}
// 	return x
// }

// //still too slow
// func minimumAbsDifference(arr []int) [][]int {

// 	var minDiffSlice []int
// 	solution := make([][]int, 0)

// 	sort.Ints(arr)

// 	minDiffSlice = append(minDiffSlice, abs(arr[0]-arr[1]))
// 	minDiffSlice = append(minDiffSlice, abs(arr[len(arr)-2]-arr[len(arr)-1]))
// 	minDiffSlice = append(minDiffSlice, abs(arr[0]-arr[len(arr)-1]))
// 	if len(arr)%2 == 0 {
// 		minDiffSlice = append(minDiffSlice, abs(arr[len(arr)/2-1]-arr[len(arr)/2]))
// 	} else {
// 		minDiffSlice = append(minDiffSlice, abs(arr[len(arr)/2-1]-arr[len(arr)/2]))
// 		minDiffSlice = append(minDiffSlice, abs(arr[len(arr)/2]-arr[len(arr)/2+1]))
// 	}

// 	var minDiff = minDiffSlice[0]

// 	for i := 1; i < len(minDiffSlice); i++ {
// 		if minDiffSlice[i] < minDiff {
// 			minDiff = minDiffSlice[i]
// 		}
// 	}

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
