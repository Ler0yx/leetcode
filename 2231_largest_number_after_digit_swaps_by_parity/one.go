package largestNumberAferDigitSwapsByParity

import (
	"sort"
	"strconv"
	"strings"
)

func largestInteger(num int) int {
	var odd []int
	var even []int
	var sol int
	numSlice := strings.Split(strconv.Itoa(num), "")
	for i := 0; i < len(numSlice); i++ {
		tmp, _ := strconv.Atoi(numSlice[i])
		if tmp%2 == 0 {
			even = append(even, tmp)
		} else {
			odd = append(odd, tmp)
		}
	}
	sort.Ints(even)
	sort.Ints(odd)
	for j, fctr := len(numSlice)-1, 1; j >= 0; j-- {
		tmp, _ := strconv.Atoi(numSlice[j])
		if tmp%2 == 0 {
			sol += even[0] * fctr
			even = even[1:]
		} else {
			if len(odd) == 0 {
				break
			}
			sol += odd[0] * fctr
			odd = odd[1:]
		}
		fctr *= 10
	}

	return sol
}
