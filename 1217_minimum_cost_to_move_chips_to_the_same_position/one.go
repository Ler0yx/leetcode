package minimumCostToMoveChipsToTheSamePosition

func minCostToMoveChips(position []int) int {

	var (
		even int
		odd  int
	)

	for _, v := range position {
		if v%2 == 0 {
			even += 1
		} else {
			odd += 1
		}
	}

	if even > odd {
		return odd
	}

	return even
}
