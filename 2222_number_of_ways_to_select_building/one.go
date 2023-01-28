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
	var res int
	var ctr0, ctr1 int

	for i := 0; i < len(s)-2; i++ {
		if s[i] == 48 {
			ctr0 = 1
		} else {
			ctr1 = 1
		}
		for j := i + 1; j < len(s)-1; j++ {
			switch {
			case s[i] == 48 && s[j] == 49:
				if digits[0] > 0 {
					res += digits[0] - ctr1
					ctr1++
				}

			case s[i] == 49 && s[j] == 48:
				if digits[1] > 0 {
					res += digits[1] - ctr0
					ctr0++
				}

			default:
				if s[j] == 48 {
					ctr0++
				} else {
					ctr1++
				}
			}
		}
		if s[i] == 48 {
			digits[0] -= 1
		} else {
			digits[1] -= 1
		}
	}
	return int64(res)
}
