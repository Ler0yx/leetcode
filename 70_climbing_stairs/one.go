package climbingStairs

import (
	"fmt"
)

// Calculates the amount of possibilities to climb a staircase with n amount of steps
// if you can either take 1 or 2 steps at a time.
func climbStairs(n int) int {
	x := 1
	y := 2

	if n == 1 {
		return x
	} else if n == 2 {
		return y
	} else {
		for i := 2; i < n; i++ {
			tmp := x + y
			x = y
			y = tmp
		}
	}

	return y
}

func main() {
	fmt.Println(climbStairs(3))
}
