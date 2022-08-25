package main

import (
	"fmt"
)

func subtractProductAndSum(n int) int {
	a := fmt.Sprintf("%d", n)
	product := 1
	sum := 0
	for _, v := range a {
		product *= int(v - '0')
		sum += int(v - '0')
	}
	return product - sum
}

func main() {
	fmt.Println(subtractProductAndSum(123))
}
