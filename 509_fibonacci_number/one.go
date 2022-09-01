package fibonacciNumber

import "fmt"

//Calculates the Fibonacci number (F(n) = F(n - 1) + F(n - 2), for n > 1.) for the given Input.
func fib(n int) int {
	x := 0
	y := 1
	sum := 0

	if n == 0 {
		return x
	} else if n == 1 {
		return y
	} else {
		for i := 1; i < n; i++ {
			sum = x + y
			x = y
			y = sum
		}
	}
	return sum
}

func main() {
	fmt.Println(fib(2))
}
