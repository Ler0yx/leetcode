package main

import (
	"fmt"
	"strings"
)

func areAlmostEqual(s1 string, s2 string) bool {
	s1Slice := strings.Split(s1, "")
	s2Slice := strings.Split(s2, "")
	var diff []int

	s1Map := make(map[int]string)
	s2Map := make(map[int]string)

	for i := 0; i < len(s1Slice); i++ {
		if s1Slice[i] != s2Slice[i] {
			diff = append(diff, i)
			if len(diff) > 2 {
				return false
			}
			s1Map[i] = s1Slice[i]
			s2Map[i] = s2Slice[i]
		}
	}

	if len(diff) == 0 {
		return true
	} else if len(diff) == 2 {
		if s1Map[diff[0]] == s2Map[diff[1]] && s1Map[diff[1]] == s2Map[diff[0]] {
			return true
		}
	}
	return false
}

func main() {
	fmt.Println(areAlmostEqual("kelbbcd", "eklbdcb"))
}
