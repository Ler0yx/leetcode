package countingWordsWithAGivenPrefix

func prefixCount(words []string, pref string) int {

	var counter int

	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words[i]); j++ {
			if words[i][j] != pref[j] {
				break
			}
			if j == len(pref)-1 {
				counter += 1
				break
			}
		}
	}
	return counter
}
