package mergeStringsAlternately

import "strings"

func mergeAlternately(word1 string, word2 string) string {
	sS1 := strings.Split(word1, "")
	sS2 := strings.Split(word2, "")
	var mergedString string

	lenMax := len(sS1)
	if len(sS2) > lenMax {
		lenMax = len(sS2)
	}

	for i := 0; i < lenMax; i++ {

		if len(sS1) > i && len(sS2) > i {
			mergedString += sS1[i] + sS2[i]
		} else if i >= len(sS1) {
			mergedString += sS2[i]
		} else if i >= len(sS2) {
			mergedString += sS1[i]
		}

	}
	return mergedString
}
