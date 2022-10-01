package generateAStringWithCharactersThatHaveOddCounts

func generateTheString(n int) string {

	var result string

	for i := 0; i < n; i++ {
		if n%2 == 0 {
			result += "a"
			if i == n-2 {
				return result + "b"
			}
		} else {
			result += "a"
			if i == n-3 {
				return result + "bc"
			}
		}
	}
	return result
}
