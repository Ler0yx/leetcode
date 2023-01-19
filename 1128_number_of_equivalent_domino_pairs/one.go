package numberOfEquivalentDominoPairs

func numEquivDominoPairs(dominoes [][]int) int {
	var ctr int
	for i := 0; i < len(dominoes)-1; i++ {
		for j := i + 1; j < len(dominoes); j++ {
			if dominoes[i][0] == dominoes[j][0] && dominoes[i][1] == dominoes[j][1] {
				ctr++
			} else if dominoes[i][0] == dominoes[j][1] && dominoes[i][1] == dominoes[j][0] {
				ctr++
			}
		}
	}
	return ctr
}
