package main

import (
	"fmt"
	"strconv"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func getDecimalValue(head *ListNode) int {

	var solutionString string

	for head != nil {
		solutionString += fmt.Sprintf("%d", head.Val)
		head = head.Next
	}

	solution, _ := strconv.ParseInt(solutionString, 2, 0)
	return int(solution)
}

func main() {

}
