package checkIfWordEqualsSummationOfTwoWords

import "strconv"

func isSumEqual(firstWord string, secondWord string, targetWord string) bool {
	var one string
	var two string
	var three string
	var oneInt int
	var twoInt int
	var threeInt int

	for i := 0; i < len(firstWord); i++ {
		tmp := firstWord[i] - 49
		one += string(tmp)
	}
	for i := 0; i < len(secondWord); i++ {
		tmp := secondWord[i] - 49
		two += string(tmp)
	}
	for i := 0; i < len(targetWord); i++ {
		tmp := targetWord[i] - 49
		three += string(tmp)
	}

	oneInt, _ = strconv.Atoi(one)
	twoInt, _ = strconv.Atoi(two)
	threeInt, _ = strconv.Atoi(three)

	if oneInt+twoInt == threeInt {
		return true
	}
	return false
}
