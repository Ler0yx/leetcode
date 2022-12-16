package sumOfUniqueElements

func sumOfUnique(nums []int) int {
	vals := make(map[int]int)
	var solution []int
	var sum int
	for i := 0; i < len(nums); i++ {
		vals[nums[i]] += 1
	}
	for i := 0; i <= 100; i++ {
		if vals[i] == 1 {
			solution = append(solution, i)
		}
	}
	for i := 0; i < len(solution); i++ {
		sum += solution[i]
	}

	return sum
}
