package maximumRepeatingSubstring

func maxRepeating(sequence string, word string) int {

	var counter int
	var wordSeqOne = word

	for i := 0; i < len(sequence); {
		if len(sequence) >= i+len(word) && sequence[i:i+len(word)] == word {
			counter += 1
			word += wordSeqOne
			i = 0
		} else {
			i += 1
		}
	}
	return counter
}
