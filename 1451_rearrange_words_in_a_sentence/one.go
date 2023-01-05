package rearrangeWordsInASentence

import (
	"fmt"
	"math"
	"strings"
	"unicode"
	"unicode/utf8"
)

func arrangeWords(text string) string {
	var max int = math.MinInt
	var res string
	lengths := make(map[int][]string)
	r, n := utf8.DecodeRuneInString(string(text[0]))
	text = string(unicode.ToLower(r)) + text[n:]
	slice := strings.Split(text, " ")

	for i := 0; i < len(slice); i++ {
		if len(slice[i]) > max {
			max = len(slice[i])
		}
		lengths[len(slice[i])] = append(lengths[len(slice[i])], slice[i])
	}
	fmt.Println(lengths)

	for j := 1; j <= max; j++ {
		if _, ok := lengths[j]; ok {
			tmp := strings.Join(lengths[j], " ")
			res += tmp + " "
		}
	}
	r, n = utf8.DecodeRuneInString(string(res))
	res = string(unicode.ToUpper(r)) + res[n:]
	return strings.TrimRight(res, " ")
}
