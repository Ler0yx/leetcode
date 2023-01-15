package determineIfStringHalvesAreAlike

import "regexp"

func halvesAreAlike(s string) bool {
	a := s[:len(s)/2]
	b := s[len(s)/2:]
	reg := regexp.MustCompile(`[aeiouAEIOU]`)
	vowelsA := reg.FindAllString(a, -1)
	vowelsB := reg.FindAllString(b, -1)
	if len(vowelsA) == len(vowelsB) {
		return true
	}
	return false
}
