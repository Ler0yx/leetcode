package shiftingLetters

import "fmt"

func shiftingLetters2(s string, shifts []int) string {
	byteSlice := []byte(s)

	var solution string
	var shiftAmount int

	for i := 0; i < len(byteSlice); i++ {

		for _, v := range shifts[i:] {
			if v > 25 {
				v -= (v / 26) * 26
			}
			shiftAmount += v
		}
		fmt.Println(shiftAmount)
		if shiftAmount > 25 {
			shiftAmount -= (shiftAmount / 26) * 26
		}

		tmp := int(byteSlice[i]) + shiftAmount
		if tmp > 122 {
			tmp -= 26
		}

		byteSlice[i] = byte(tmp)
		shiftAmount = 0

	}
	solution = string(byteSlice)
	return solution
}
