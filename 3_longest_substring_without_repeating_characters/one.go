package longestSubstringWithoutRepeatingCharacters

func lengthOfLongestSubstring(s string) int {

	usedLetters := make(map[byte]bool)
	var length int
	var max int

	for i := 0; i < len(s); i++ {
		for j := i; j < len(s); j++ {
			if ok := usedLetters[s[j]]; ok {
				usedLetters = map[byte]bool{}
				break
			}
			usedLetters[s[j]] = true
			length += 1
		}
		if length > max {
			max = length
		}
		length = 0
	}
	return max
}
