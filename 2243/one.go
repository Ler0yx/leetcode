package main

import (
	"fmt"
	"strconv"
	"strings"
)

func digitSum(s string, k int) string {

	if len(s) <= k {
		return s
	}

	var stringSlice []string
	var intSlice []int
	var i int

	stringSlice = strings.Split(s, "")

	//Converts the string slice to an int slice
	for i := 0; i < len(stringSlice); i++ {
		tmp, _ := strconv.Atoi(stringSlice[i])
		intSlice = append(intSlice, tmp)
	}

loop:
	solutionSlice := []int{}
	for i < len(intSlice) {
		tmpSum := 0
		for j := 0; j < k; j++ {
			if i+j < len(intSlice) {
				tmpSum += intSlice[i+j]
			} else {
				break
			}
		}
		solutionSlice = append(solutionSlice, tmpSum)
		i += k
	}

	if len(solutionSlice) > k {
		intSlice = solutionSlice
		i = 0
		goto loop
	}

	stringSlice = []string{}
	fmt.Println(solutionSlice)
	for i := 0; i < len(solutionSlice); i++ {
		stringSlice = append(stringSlice, strconv.Itoa(solutionSlice[i]))
	}

	fmt.Println(strings.Join(stringSlice, ""))
	return strings.Join(stringSlice, "")
}

func main() {
	digitSum("01234567890", 2)
}
