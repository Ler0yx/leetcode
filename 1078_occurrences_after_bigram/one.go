package occurrencesAfterBigram

import "strings"

func findOcurrences(text string, first string, second string) []string {
	textSlice := strings.Split(text, " ")
	res := make([]string, 0)
	for i := 1; i < len(textSlice)-1; i++ {
		if textSlice[i] == second && textSlice[i-1] == first {
			res = append(res, textSlice[i+1])
		}
	}
	return res
}
