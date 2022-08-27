package main

import "fmt"

func signFunc(x int) int {
	if x > 0 {
		return 1
	} else if x < 0 {
		return -1
	} else {
		return 0
	}
}

func arraySign(nums []int) int {
	product := nums[0]
	for i := 1; i < len(nums); i++ {
		if nums[i] == 0 {
			return 0
		} else if nums[i] < 0 {
			product *= -1
		}
	}
	return signFunc(product)
}

func main() {

	nums := []int{1, 2, -3, 4, 5}

	fmt.Println(arraySign(nums))
}
