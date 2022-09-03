package convert1dArrayInto2dArray

func Construct2DArray(original []int, m int, n int) [][]int {

	if len(original) != m*n {
		return [][]int{}
	}

	solution := make([][]int, m)
	k := 0

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			solution[i] = append(solution[i], original[k])
			k++
		}
	}
	return solution
}
