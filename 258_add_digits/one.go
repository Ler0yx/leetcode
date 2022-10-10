package addDigits

func addDigits(num int) int {

	var sum int

	for {
		for num > 0 {
			tmp := num % 10
			num /= 10
			sum += tmp

		}
		if sum < 10 {
			return sum
		}
		num = sum
		sum = 0
	}
	return sum
}
