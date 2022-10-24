package sortingTheSentence

import "strings"

func sortSentence(s string) string {

	slice := strings.Split(s, " ")
	order := make([]string, len(slice))
	var solution string

	for i := 0; i < len(slice); i++ {
		order[int(slice[i][len(slice[i])-1]-49)] = slice[i][:len(slice[i])-1]
	}

	for j := 0; j < len(order); j++ {
		if j != len(order)-1 {
			solution += order[j] + " "
		} else {
			solution += order[j]
		}
	}
	return solution
}
