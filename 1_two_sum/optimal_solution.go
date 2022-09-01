package twoSum

func twoSum2(nums []int, target int) []int {

	result := map[int]int{}

	//Loops through the array, saving the index and value.
	for i, v := range nums {

		//Detracts the value from the target
		remainder := target - v

		//Creates a bool that turns true if the remainder is a key of result,
		//meaning that the value at this index + the value at the index remainder = target
		if _, ok := result[remainder]; ok {
			return []int{result[remainder], i}

			//If the bool is false, we create an entry with the value as the key and the index as the value
		} else {
			result[v] = i
		}
	}

	//Just a placeholder return to satisy the compiler
	return []int{}
}

func main() {

}
