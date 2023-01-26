package ransomNote

func canConstruct(ransomNote string, magazine string) bool {
	digits := make(map[byte]int)
	if len(ransomNote) > len(magazine) {
		return false
	}
	for i := 0; i < len(ransomNote); i++ {
		digits[ransomNote[i]] += 1
	}
	for j := 0; j < len(magazine); j++ {
		if _, ok := digits[magazine[j]]; !ok {
			continue
		}
		digits[magazine[j]] -= 1
		if digits[magazine[j]] == 0 {
			delete(digits, magazine[j])
		}
		if len(digits) == 0 {
			return true
		}
	}
	return false
}
