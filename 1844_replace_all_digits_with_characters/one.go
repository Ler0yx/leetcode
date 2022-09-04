package replaceAllDigitsWithCharacters

import "fmt"

func shift(c byte, x byte) byte {
	c += x - 48
	return c
}

func replaceDigits(s string) string {

	var solution string

	for i := 0; i < len(s); i++ {
		if s[i] < 95 {
			solution += fmt.Sprintf("%c", shift(s[i-1], s[i]))
		} else {
			solution += fmt.Sprintf("%c", s[i])
		}
	}
	return solution
}
