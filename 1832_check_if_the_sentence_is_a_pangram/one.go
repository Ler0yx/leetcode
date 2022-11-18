package checkIfTheSentenceIsAPangram

func checkIfPangram(sentence string) bool {
	letters := make(map[byte]bool)

	for i := 0; i < len(sentence); i++ {
		letters[sentence[i]] = true
	}
	for i := 97; i < 123; i++ {
		if _, ok := letters[byte(i)]; !ok {
			return false
		}
	}
	return true
}
