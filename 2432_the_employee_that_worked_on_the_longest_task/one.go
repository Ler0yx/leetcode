package theEmployeeThatWorkedOnTheLongestTask

import "sort"

func hardestWorker(n int, logs [][]int) int {
	var max int = logs[0][1]
	var ties []int
	ties = append(ties, logs[0][0])
	for i := 1; i < len(logs); i++ {
		tmp := logs[i][1] - logs[i-1][1]
		if tmp > max {
			max = tmp
			ties = []int{}
			ties = append(ties, logs[i][0])
		} else if tmp == max {
			ties = append(ties, logs[i][0])
		}
	}
	sort.Ints(ties)
	return ties[0]
}
