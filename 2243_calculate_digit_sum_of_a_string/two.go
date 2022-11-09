package calculateDigitSumOfAString

import (
	"fmt"
	"strconv"
)

func digitSum(s string, k int) string {

	if len(s) <= k {
		return s
	}

	j := k
	i := 0
	sum := 0
	sumStr := ""

	for {
		for ; i < j; i++ {
			sum += int(s[i] - 48)
		}
		sumStr += strconv.Itoa(sum)
		if j == len(s) {
			fmt.Println("S:", s)
			s = sumStr
			sumStr = ""
			i = 0
			j = k
			if len(s) <= k {
				break
			}
		} else if j+k > len(s) {
			j += len(s) - j
		} else {
			j += k
		}
		sum = 0
	}
	return s
}
