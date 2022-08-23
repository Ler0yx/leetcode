package main

import (
	"fmt"
	"strings"
)

//Creates the hint for the bulls and cows game.
func getHint(secret string, guess string) string {
	var solution string
	var possibleCow []byte
	var solutionSlice []string
	var counterB int
	var counterC int
	cowsMap := make(map[byte]int)

	for i := 0; i < len(secret); i++ {

		//If the strings are equal, the Bull-Counter gets incremented by +1.
		if secret[i] == guess[i] {
			counterB += 1

		//If the strings are unequal, the string gets added to the map (if not already present) and incremented by +1.
		//The string is also added to a slice to iterate over later to find all the Cows.
		} else {
			cowsMap[secret[i]] += 1
			possibleCow = append(possibleCow, byte(guess[i]))
		}
	}

	//Iterates over the Map and compares it to the slice of possible cows.
	for i := 0; i < len(possibleCow); i++ {
		if cowsMap[possibleCow[i]] > 0 {
			cowsMap[possibleCow[i]] -= 1
			counterC += 1
		}
	}
	solutionSlice = []string{fmt.Sprint(counterB), "A", fmt.Sprint(counterC), "B"}
	solution = strings.Join(solutionSlice, "")
	return solution
}

func main() {
	fmt.Println(getHint("7823456872345", "7834926587346"))
}
