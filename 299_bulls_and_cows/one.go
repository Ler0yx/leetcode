package main

import (
	"fmt"
	"strings"
)

func getHint(secret string, guess string) string {
	var solution string
	var possibleCow []byte
	var solutionSlice []string
	var counterB int
	var counterC int
	cowsMap := make(map[byte]int)

	for i := 0; i < len(secret); i++ {

		if secret[i] == guess[i] {
			counterB += 1
		} else {
			cowsMap[secret[i]] += 1
			possibleCow = append(possibleCow, byte(guess[i]))
		}
	}
	fmt.Println("pC: ", possibleCow)
	fmt.Println("cM: ", cowsMap)
	for i := 0; i < len(possibleCow); i++ {
		fmt.Println(possibleCow[i])
		if cowsMap[possibleCow[i]] > 0 {
			fmt.Println("guess:", possibleCow[i])
			cowsMap[possibleCow[i]] -= 1
			fmt.Println("cM:", cowsMap)
			counterC += 1
			fmt.Println("counterC:", counterC)
		}
	}
	solutionSlice = []string{fmt.Sprint(counterB), "A", fmt.Sprint(counterC), "B"}
	solution = strings.Join(solutionSlice, "")
	return solution
}

func main() {
	fmt.Println(getHint("7823456872345", "7834926587346"))
	//3:2 6:2 2:1 4:1 5:1 7:1
	//8:3 4:1 3:1 5:1 6:1 9:1
	//3:1 6:1 2:0 4:1 5:1 7:0
}
