package main

import (
	"fmt"
	b "leetcode/1109_corporate_flight_bookings"
)

func main() {

	//bookings = [[1,2,10],[2,3,20],[2,5,25]], n = 5
	// bookings := [][]int{{1, 2, 10}, {2, 3, 20}, {2, 5, 25}}
	bookings := []int{2, 2, 30, 2, 2, 45}
	bookings2 := []int{2, 2, 30, 2, 2, 45}

	fmt.Println(bookings == bookings2)
}
