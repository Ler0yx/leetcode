package main

import "fmt"

func main() {
	words := []string{"hello", "i", "am", "leetcode", "hello"}
	// target := "hello"
	// startIndex := 1
	lenW := len(words)

	for i := 0; i < len(words); i++ {
		fmt.Println((i + 1) % lenW)
	}
}
