package checkIfNumberHasEqualDigitCountAndDigitValue

func digitCount(num string) bool {
	digits := make(map[int]int)
	for i := 0; i < len(num); i++ {
		digits[int(num[i]-48)] += 1
	}

	for j := 0; j < len(num); j++ {
		if int(num[j])-48 != digits[j] {
			return false
		}
	}
	return true
}
