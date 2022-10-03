package keyboardRow

func findWords2(words []string) []string {

	firstRow := map[byte]bool{113: true, 119: true, 101: true, 114: true, 116: true, 121: true, 117: true, 105: true, 111: true, 112: true, 81: true, 87: true, 69: true, 82: true, 84: true, 89: true, 85: true, 73: true, 79: true, 80: true}
	secondRow := map[byte]bool{97: true, 115: true, 100: true, 102: true, 103: true, 104: true, 106: true, 107: true, 108: true, 65: true, 83: true, 68: true, 70: true, 71: true, 72: true, 74: true, 75: true, 76: true}
	thirdRow := map[byte]bool{122: true, 120: true, 99: true, 118: true, 98: true, 110: true, 109: true, 90: true, 88: true, 67: true, 86: true, 66: true, 78: true, 77: true}
	solution := []string{}

	for i := range words {
		switch {
		case firstRow[words[i][0]]:
			for j := 0; j < len(words[i]); j++ {
				if !firstRow[words[i][j]] {
					break
				}
				if j == len(words[i])-1 {
					solution = append(solution, words[i])
				}
			}

		case secondRow[words[i][0]]:
			for j := 0; j < len(words[i]); j++ {
				if !secondRow[words[i][j]] {
					break
				}
				if j == len(words[i])-1 {
					solution = append(solution, words[i])
				}
			}

		case thirdRow[words[i][0]]:
			for j := 0; j < len(words[i]); j++ {
				if !thirdRow[words[i][j]] {
					break
				}
				if j == len(words[i])-1 {
					solution = append(solution, words[i])
				}
			}
		}
	}
	return solution
}
