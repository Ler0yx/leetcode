package circularSentence

import "strings"

func isCircularSentence(sentence string) bool {

	var words []string
	words = strings.Split(sentence, " ")
	if words[len(words)-1][len(words[len(words)-1])-1] != words[0][0] {
		return false
	}
	for i := len(words) - 1; i > 0; i-- {
		if words[i][0] != words[i-1][len(words[i-1])-1] {
			return false
		}
	}
	return true
}
