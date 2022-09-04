package shiftingLetters

import "fmt"

func shiftingLetters(s string, shifts []int) string {
	byteSlice := []byte(s)

	var solution string

	for i := 0; i < len(byteSlice); i++ {
		fmt.Println("1: ", byteSlice[i])
		for _, v := range shifts[i:] {
			if v > 25 {
				v -= (v / 26) * 26
			}

			fmt.Println("2: ", byteSlice[i])

			tmp := int(byteSlice[i]) + v
			if tmp > 122 {
				tmp -= 26
			}

			byteSlice[i] = byte(tmp)
		}
	}
	solution = string(byteSlice)
	return solution
}