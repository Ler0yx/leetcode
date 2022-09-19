package runningSumOf1DArray

func runningSum(nums []int) []int {

	var solution []int
	var sum int

	for _, v := range nums {
		sum += v
		solution = append(solution, sum)
	}
	return solution
}
