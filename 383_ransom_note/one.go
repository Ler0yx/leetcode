package ransomNote

import "fmt"

func canConstruct(ransomNote string, magazine string) bool {

	noteMap := make(map[byte]int)
	magMap := make(map[byte]int)

	if len(ransomNote) > len(magazine) {
		return false
	}

	for i := 0; i < len(ransomNote); i++ {
		noteMap[ransomNote[i]] += 1
	}
	for i := 0; i < len(magazine); i++ {
		magMap[magazine[i]] += 1
	}

	if len(noteMap) > len(magMap) {
		return false
	}

	for j := 0; j < 123; j++ {
		fmt.Println(noteMap)
		fmt.Println(magMap)
		if noteMap[byte(j)] > magMap[byte(j)] {
			fmt.Println("hier")
			return false
		}
	}
	return true
}
