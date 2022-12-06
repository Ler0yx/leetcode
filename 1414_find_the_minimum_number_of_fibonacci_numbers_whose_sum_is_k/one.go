package findTheMinimumNumberOfFibonaciNumbersWhoseSumIsK

import "fmt"

func calcFibonacciNumbers(k int) []int {
	fNums := make([]int, 2)
	fNums[0] = 1
	fNums[1] = 1

	for i := 2; ; i++ {
		tmp := fNums[i-1] + fNums[i-2]
		if tmp <= k {
			fNums = append(fNums, tmp)
		} else {
			break
		}
	}
	return fNums
}

func findMinFibonacciNumbers(k int) int {

	var counter int
	fNums := calcFibonacciNumbers(k)

	fmt.Println(fNums)

	for i := len(fNums) - 1; k > 0; i-- {
		if i > 1 {
			if fNums[i] > k {
				continue
			}
			k -= fNums[i]
			counter++
		} else {
			return counter + k
		}
	}
	return counter
}
