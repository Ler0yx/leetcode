package uniqueMorseCodeWords

func uniqueMorseRepresentations(words []string) int {
	totalDifferences := make(map[string]int)
	transformations := []string{}
	alphabet := map[byte]int{
		97: 0, 98: 1, 99: 2, 100: 3, 101: 4, 102: 5, 103: 6, 104: 7, 105: 8, 106: 9, 107: 10, 108: 11,
		109: 12, 110: 13, 111: 14, 112: 15, 113: 16, 114: 17, 115: 18, 116: 19, 117: 20, 118: 21, 119: 22, 120: 23, 121: 24, 122: 25,
	}
	morse := []string{
		".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--",
		"-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--..",
	}

	for i := 0; i < len(words); i++ {
		tmp := ""
		for j := 0; j < len(words[i]); j++ {
			tmp += morse[alphabet[words[i][j]]]
		}
		transformations = append(transformations, tmp)
	}
	for i := 0; i < len(transformations); i++ {
		totalDifferences[transformations[i]] += 1
	}
	return len(totalDifferences)
}
