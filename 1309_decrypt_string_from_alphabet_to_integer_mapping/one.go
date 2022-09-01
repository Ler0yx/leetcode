package decryptStringFromAlphabetToIntegerMapping

func freqAlphabets(s string) string {

	var solution []byte

	for i := 0; i < len(s); i++ {
		if i < len(s)-2 && s[i+2] == '#' {
			if s[i] == '1' {
				solution = append(solution, s[i+1]+58)
				i += 2
			} else {
				solution = append(solution, s[i+1]+68)
				i += 2
			}
		} else {
			solution = append(solution, s[i]+48)
		}
	}
	return string(solution)
}