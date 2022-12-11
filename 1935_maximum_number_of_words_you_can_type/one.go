package maximumNumberOfWordsYouCanType

import (
	"fmt"
	"strings"
)

func canBeTypedWords(text string, brokenLetters string) int {
	letters := make(map[byte]bool)
	var counter int
	var texts []string
	for i := 0; i < len(brokenLetters); i++ {
		letters[brokenLetters[i]] = true
	}
	fmt.Println(letters)
	texts = strings.Split(text, " ")

	for j := 0; j < len(texts); j++ {
		for k := 0; k < len(texts[j]); k++ {
			if letters[texts[j][k]] {
				break
			}
			if k == len(texts[j])-1 {
				counter++
			}
		}
	}
	return counter
}
