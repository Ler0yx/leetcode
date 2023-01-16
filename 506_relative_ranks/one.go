package relativeRanks

import (
	"sort"
	"strconv"
)

func findRelativeRanks(score []int) []string {
	idx := make(map[int]int)
	res := make([]string, len(score))
	var ctr int
	for i := 0; i < len(score); i++ {
		idx[score[i]] = i
	}
	sort.Ints(score)
	for j := len(score) - 1; j >= 0; j-- {
		ctr++
		switch {
		case j == len(score)-1:
			res[idx[score[j]]] = "Gold Medal"

		case j == len(score)-2:
			res[idx[score[j]]] = "Silver Medal"

		case j == len(score)-3:
			res[idx[score[j]]] = "Bronze Medal"

		default:
			res[idx[score[j]]] = strconv.Itoa(ctr)
		}
	}
	return res
}
