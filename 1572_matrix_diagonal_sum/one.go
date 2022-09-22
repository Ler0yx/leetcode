package matrixDiagonalSum

func diagonalSum(mat [][]int) int {

	var sum int
	e := len(mat) - 1

	for i, row := range mat {
		for d := i; d < len(row); {
			sum += row[d]
			if i != e {
				sum += row[e]
			}
			break

		}
		e--
	}
	return sum
}
