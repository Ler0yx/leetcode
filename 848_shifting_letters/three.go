package shiftingLetters

func shiftAmount3(arr []int) []int {

	solution := make([]int, len(arr))

	for i := 0; i < len(arr); i++ {
		if arr[i] > 25 {
			arr[i] -= (arr[i] / 26) * 26
		}
		solution[i] = arr[i]
	}
	return solution
}

func shiftingLetters3(s string, shifts []int) string {

	byteSlice := []byte(s)
	var solution string
	var shiftsFinal = shiftAmount(shifts)
	var tmp int

	for i := 0; i < len(byteSlice); i++ {
		for _, v := range shiftsFinal[i:] {
			tmp += v
		}

		if tmp > 122 {
			tmp -= (tmp / 26) * 26
		}
		byteSlice[i] = byte(tmp)
	}
	solution = string(byteSlice)
	return solution
}
