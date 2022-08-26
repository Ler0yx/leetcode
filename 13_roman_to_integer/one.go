package main

import "fmt"

func romanToInt(s string) int {

	sum := 0

	for i := 0; i < len(s); i++ {
		switch {
		case s[i] == 'I':
			if i < len(s)-1 && s[i+1] == 'V' {
				sum += 4
				i += 1
				continue
			}
			if i < len(s)-1 && s[i+1] == 'X' {
				sum += 9
				i += 1
				continue
			}
			sum += 1
		case s[i] == 'V':
			sum += 5
		case s[i] == 'X':
			if i < len(s)-1 && s[i+1] == 'L' {
				sum += 40
				i += 1
				continue
			}
			if i < len(s)-1 && s[i+1] == 'C' {
				sum += 90
				i += 1
				continue
			}
			sum += 10
		case s[i] == 'L':
			sum += 50
		case s[i] == 'C':
			if i < len(s)-1 && s[i+1] == 'D' {
				sum += 400
				i += 1
				continue
			}
			if i < len(s)-1 && s[i+1] == 'M' {
				sum += 900
				i += 1
				continue
			}
			sum += 100
		case s[i] == 'D':
			sum += 500
		case s[i] == 'M':
			sum += 1000
		}
	}
	return sum
}

func main() {
	fmt.Println(romanToInt("LVIII"))
}
