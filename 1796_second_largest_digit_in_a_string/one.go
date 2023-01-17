package secondLargestDigitInAString

import (
	"sort"
	"strconv"
	"strings"
)

func secondHighest(s string) int {
	var res string
	for i := 0; i < len(s); i++ {
		if s[i] < 97 {
			res += string(s[i])
		}
	}
	sInt := strings.Split(res, "")
	sort.Strings(sInt)
	for j := len(sInt) - 2; j >= 0; j-- {
		if sInt[j] < sInt[len(sInt)-1] {
			tmp, _ := strconv.Atoi(sInt[j])
			return tmp
		}
	}
	return -1
}
