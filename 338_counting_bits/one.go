package countingBits

import "strconv"

func countBits(n int) []int {
	var res []int
	for i := 0; i <= n; i++ {
		tmp := strconv.FormatInt(int64(i), 2)
		tmpSum := 0
		for j := 0; j < len(tmp); j++ {
			if tmp[j] == 49 {
				tmpSum++
			}
		}
		res = append(res, tmpSum)
	}
	return res
}
