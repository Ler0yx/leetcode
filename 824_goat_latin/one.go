package goatLatin

import (
	"regexp"
	"strings"
)

func toGoatLatin(sentence string) string {
	words := strings.Split(sentence, " ")
	res := []string{}
	for i, word := range words {
		match, _ := regexp.MatchString("[AEIOUaeiou]", string(word[0]))
		switch {
		case match:
			word += "ma"

		default:
			tmp := word
			word = string(tmp[1:]) + string(tmp[0]) + "ma"
		}
		word += strings.Repeat("a", i+1)
		res = append(res, word)
	}
	return strings.Join(res, " ")
}
