package maximumEnemyFortsThatCanBeCaptured

func captureForts(forts []int) int {
	var ctrL int
	var ctrR int
	var maxL int
	var maxR int

	for i := 0; i < len(forts); i++ {
		if forts[i] == 1 {
			for j := i + 1; j < len(forts); j++ {
				if forts[j] == 0 {
					ctrL++
				} else if forts[j] == 1 {
					ctrL = 0
					i = j - 1
					break
				} else {
					if ctrL > maxL {
						maxL = ctrL
					}
					ctrL = 0
					i = j - 1
					break
				}
			}
		}
		if forts[i] == -1 {
			for k := i + 1; k < len(forts); k++ {
				if forts[k] == 0 {
					ctrR++
				} else if forts[k] == -1 {
					ctrR = 0
					i = k - 1
					break
				} else {
					if ctrR > maxR {
						maxR = ctrR
					}
					ctrR = 0
					i = k - 1
					break
				}
			}
		}
	}
	if maxL > maxR {
		return maxL
	}
	return maxR
}
