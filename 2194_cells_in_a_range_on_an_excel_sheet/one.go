package cellsInARangeOnAnExcelSheet

func cellsInRange(s string) []string {

	letter := s[0]
	digit := s[1]
	solution := []string{s[:2]}

	for i := 1; letter < s[3] || digit < s[4]; i++ {

		for digit < s[4] {
			digit += 1
			solution = append(solution, string(letter)+string(digit))
		}

		if letter < s[3] {
			letter += 1
			digit = s[1]
			solution = append(solution, string(letter)+string(digit))
		}
	}
	return solution
}
