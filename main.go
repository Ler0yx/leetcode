package main

import "fmt"

func main() {

	nums := []int{123, 234}
	bnums := []byte(nums[0], nums[1])
	fmt.Println(byte(nums))
}
