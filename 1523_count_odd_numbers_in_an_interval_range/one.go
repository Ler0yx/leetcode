package countOddNumbersInAnIntervalRange

import "fmt"

func countOdds(low int, high int) int {

	var counter int

	for i := low; i <= high; i++ {
		if i%2 != 0 {
			counter += 1
		}
	}
	return counter
}

func main() {
	fmt.Println(countOdds(8, 10))
}
