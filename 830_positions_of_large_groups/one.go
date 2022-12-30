package positionsOfLargeGroups

func largeGroupPositions(s string) [][]int {
	var res [][]int
	var ctr int
	for i := 1; i < len(s); i++ {
		if s[i-1] == s[i] {
			ctr += 1
			if i == len(s)-1 && ctr >= 2 {
				res = append(res, []int{i - ctr, i})
			}
		} else {
			if ctr > 1 {
				res = append(res, []int{i - 1 - ctr, i - 1})
			}
			ctr = 0
		}
	}
	return res
}
