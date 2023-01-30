package countSquareSumTriples

func cSquares(n int) (map[int]bool, int) {
	cSquares := make(map[int]bool)
	var i int
	for ; i <= n; i++ {
		cSquares[i*i] = true
	}
	return cSquares, i - 1
}

func countTriples(n int) int {
	var res int
	cSquares, max := cSquares(n)
	for i := 1; i*i+1 < max*max; i++ {
		for j := 1; ; j++ {
			if i*i+j*j > max*max {
				break
			} else if cSquares[i*i+j*j] {
				res++
			}
		}
	}
	return res
}
