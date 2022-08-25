package main

import "fmt"

func hammingWeight(num uint32) int {
	weight := 0
	a := fmt.Sprintf("%b", num)
	for _, v := range a {
		if v == '1' {
			weight += 1
		}
	}
	return weight
}

func main() {
	fmt.Println(hammingWeight(00000000000000000000000000001011))
}
