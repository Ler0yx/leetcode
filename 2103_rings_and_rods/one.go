package ringsAndRods

import "strconv"

func countPoints(rings string) int {
	ringsSlice := make([][]int, 10)
	for i := range ringsSlice {
		ringsSlice[i] = make([]int, 3)
	}
	var ctr int
	for i := 0; i < len(rings)-1; i++ {
		tmp, _ := strconv.Atoi(string(rings[i+1]))
		if rings[i] == 82 {
			ringsSlice[tmp][0] += 1
		} else if rings[i] == 71 {
			ringsSlice[tmp][1] += 1
		} else if rings[i] == 66 {
			ringsSlice[tmp][2] += 1
		}
	}
	for j := 0; j <= 9; j++ {
		if ringsSlice[j][0] > 0 && ringsSlice[j][1] > 0 && ringsSlice[j][2] > 0 {
			ctr++
		}
	}
	return ctr
}
