package main

import (
	"fmt"
	"strings"
)

func interpret(command string) string {

	var sS []string
	var result string

	sS = strings.Split(command, "")

	for i := 0; i < len(sS); i++ {

		if sS[i] == "G" {
			result += "G"
		} else if i+1 < len(sS) && sS[i] == "(" && sS[i+1] == ")" {
			result += "o"
		} else if i+3 < len(sS) && sS[i] == "(" && sS[i+1] == "a" && sS[i+2] == "l" && sS[i+3] == ")" {
			result += "al"
		} else {
			continue
		}
	}
	return result
}

func main() {
	fmt.Println(interpret("(al)G(al)()()G"))
}
