package numberOfWaysToSelectBuilding

func countingDigits(s string) []int {
	var sum int
	for i := 0; i < len(s); i++ {
		if s[i] == 49 {
			sum += 1
		}
	}
	return []int{len(s) - sum, -sum + len(s)}
}

func numberOfWays(s string) int64 {
	digits := countingDigits(s)
	tmpDigits := make([]int, len(digits))
	var res int64

	for i := 0; i < len(s)-2; i++ {
		copy(tmpDigits, digits)
		if s[i] == 48 {
			digits[0] -= 1
		} else {
			digits[1] -= 1
		}
		for j := i + 1; j < len(s)-2; j++ {
			switch {
			case s[i] == 48 && s[j] == 49:
				if tmpDigits[1] > 0 {
					tmpDigits[1] -= 1
					res += int64(tmpDigits[1])
				}

			case s[i] == 49 && s[j] == 48:
				if tmpDigits[0] > 0 {
					tmpDigits[0] -= 1
					res += int64(tmpDigits[0])
				}

			default:
				if s[j] == 48 {
					tmpDigits[0] -= 1
				} else {
					tmpDigits[1] -= 1
				}
			}
		}
	}
	return res
}
