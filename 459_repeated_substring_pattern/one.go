package main

import (
	"strings"
)

func repeatedSubstringPattern(s string) bool {

	var subString string

	for i := 0; i < len(s)/2; i++ {
		subString += string(s[i])
		if len(s)%len(subString) == 0 {
			if strings.Repeat(subString, len(s)/len(subString)) == s {
				return true
			}
		}
	}
	return false
}
