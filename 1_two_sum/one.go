package twoSum

func twoSum(nums []int, target int) []int {

	var solution []int

	//Loops through the array, used to get the first int.
	for i := 0; i < len(nums); i++ {

		//Also loops through the arry, used to get the second int. Starts 1 index after loop 1.
		for d := i + 1; d < len(nums); d++ {

			//Builds a slice with the index of the first and second int, if the sum matches the target.
			if nums[i]+nums[d] == target {
				solution = append(solution, i, d)
			}
		}
	}
	return solution

	//Problem: uses two for-loops which increases the runtime a lot.
}
