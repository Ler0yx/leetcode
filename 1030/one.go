func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func allCellsDistOrder(rows int, cols int, rCenter int, cCenter int) [][]int {

	matrix := make(map[int][]int)
	points := []int{}
	var i int
	results := [][]int{}

	for j := 0; j < rows; j++ {
		for k := 0; k < cols; k++ {
			l := abs(rCenter-j) + abs(cCenter-k)
			points = append(points, j, k, l)
			matrix[i] = points
			i++
			points = []int{}
		}
	}
	fmt.Println(matrix)
	for i := 0; i < len(matrix); i++ {
		if matrix[i][2] == i {
			tmp := []int{matrix[i][0], matrix[i][1]}
			results = append(results, tmp)
		}
	}
	return results
}