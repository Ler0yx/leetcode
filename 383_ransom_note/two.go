package ransomNote

func canConstruct(ransomNote string, magazine string) bool {
	digits := make(map[byte]int)
	if len(ransomNote) > len(magazine) {
		return false
	}
	for i := 0; i < len(magazine); i++ {
		digits[magazine[i]] += 1
	}
	for j := 0; j < len(ransomNote); j++ {
		digits[ransomNote[j]] -= 1
		if digits[ransomNote[j]] < 0 {
			return false
		}
	}
	return true
}
