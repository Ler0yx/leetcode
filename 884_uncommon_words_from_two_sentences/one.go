package uncommonWordsFromTwoSentences

import "strings"

func uncommonFromSentences(s1 string, s2 string) []string {
	var res []string
	s1Map := make(map[string]int)
	s2Map := make(map[string]int)
	s1Slice := strings.Split(s1, " ")
	s2Slice := strings.Split(s2, " ")
	var maxLen int = len(s1Slice)
	if len(s2Slice) > maxLen {
		maxLen = len(s2Slice)
	}

	for i := 0; i < maxLen; i++ {
		if i < len(s1Slice) {
			s1Map[s1Slice[i]] += 1
		}
		if i < len(s2Slice) {
			s2Map[s2Slice[i]] += 1
		}
	}

	for i := 0; i < maxLen; i++ {
		if i < len(s2Slice) {
			if s2Map[s2Slice[i]] == 1 && s1Map[s2Slice[i]] == 0 {
				res = append(res, s2Slice[i])
			}
		}
		if i < len(s1Slice) {
			if s1Map[s1Slice[i]] == 1 && s2Map[s1Slice[i]] == 0 {
				res = append(res, s1Slice[i])
			}
		}
	}
	return res
}
