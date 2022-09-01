package main

import "fmt"

func freqAlphabets2(s string) string {

	var solutionSlice []byte
	var solution string

	for i := 0; i < len(s); i++ {
		if i < len(s)-2 && s[i+2] == '#' {
			if s[i] == '1' {
				solutionSlice = append(solutionSlice, s[i+1]+58)
				i += 2
			} else {
				solutionSlice = append(solutionSlice, s[i+1]+68)
				i += 2
			}
		} else {
			solutionSlice = append(solutionSlice, s[i]+48)
		}
	}
	solution = fmt.Sprintf("%s", solutionSlice)
	return solution
}

func main() {
	fmt.Println(freqAlphabets2("1326#"))
}
