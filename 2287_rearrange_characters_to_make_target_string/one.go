package rearrangeCharactersToMakeTargetString

func rearrangeCharacters(s string, target string) int {

	if len(target) > len(s) {
		return 0
	}

	targetMap := make(map[rune]int)
	min := 101

	for _, l := range s {
		targetMap[l] += 1
	}

	for _, l := range target {
		count, ok := targetMap[l]
		if ok {
			counterTarget := 0
			for _, l2 := range target {
				if l == l2 {
					counterTarget += 1
				}
			}
			if min > count/counterTarget {
				min = count / counterTarget
			}
		} else {
			return 0
		}
	}
	return min
}
