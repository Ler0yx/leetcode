package aNumberAfterADoubleReversal

import "strconv"

func isSameAfterReversals(num int) bool {
	numString := strconv.Itoa(num)
	if len(numString) == 1 {
		return true
	} else if byte(numString[len(numString)-1]) == 48 {
		return false
	}
	return true
}
