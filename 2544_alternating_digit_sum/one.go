package alternatingDigitSum

import (
	"strconv"
	"strings"
)

func alternateDigitSum(n int) int {
	sSlice := strings.Split(strconv.Itoa(n), "")
	var res int
	for i := 0; i < len(sSlice); i++ {
		tmp, _ := strconv.Atoi(sSlice[i])
		if i%2 == 0 {
			res += tmp
		} else {
			res -= tmp
		}
	}
	return res
}
