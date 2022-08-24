package main

import "fmt"

func getHint2(secret string, guess string) string {
	records := make([]int, 10)
	for _, ch := range guess {
		records[ch-'0']++
	}

	var bulls, cows int
	for i, ch := range secret {
		if ch == rune(guess[i]) {
			bulls++
		}

		if records[ch-'0'] > 0 {
			records[ch-'0']--
			cows++
		}
	}

	return fmt.Sprintf("%dA%dB", bulls, cows-bulls)
}
