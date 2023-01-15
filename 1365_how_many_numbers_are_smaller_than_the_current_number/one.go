package howManyNumbersAreSmallerThanTheCurrentNumber

func smallerNumbersThanCurrent(nums []int) []int {

	res := make(map[int]int)
	var solution []int
	ctr := 0
	for i := 0; i < len(nums); i++ {
		if _, ok := res[nums[i]]; ok {
			solution = append(solution, res[nums[i]])
		} else {
			for j := 0; j < len(nums); j++ {
				if nums[j] < nums[i] {
					ctr++
				}
			}
			res[nums[i]] = ctr
			solution = append(solution, ctr)
			ctr = 0
		}
	}
	return solution
}
