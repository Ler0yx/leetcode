package minimumAbsoluteDifference

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func minimumAbsDifference(arr []int) [][]int {

	var minDiff int
	solution := make([][]int, 0)

	for i := 0; i < len(arr)-1; i++ {
		if abs(arr[i]-arr[i+1]) < minDiff {
			minDiff = abs(arr[i] - arr[i+1])
		}
	}

	for i := 0; i < len(arr)-1; i++ {
		for j := i + 1; j < len(arr); j++ {
			if abs(arr[i]-arr[j]) == minDiff {
				tmp := []int{arr[i], arr[j]}
				solution = append(solution, tmp)
				break
			}
		}
	}
	return solution
}
