package checkIfTheSentenceIsAPangram

func checkIfPangram2(sentence string) bool {
	letters := make(map[byte]bool)

	for i := 0; i < len(sentence); i++ {
		letters[sentence[i]] = true
		if len(letters) == 26 {
			return true
		}
	}
	return false
}
