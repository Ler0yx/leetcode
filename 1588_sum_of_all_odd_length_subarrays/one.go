package sumOfAllOddLengthSubarrays

func sumOddLengthSubarrays(arr []int) int {
	var sum int

	for i := 0; i < len(arr); i++ { // 0 1 2 3 4
		for j := i; j < len(arr); j += 2 { // 0 2 4 6
			for k := i; k <= j; k++ {
				sum += arr[k]
			}
		}
	}
	return sum
}
