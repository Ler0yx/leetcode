package maximumNumberOfBallsInABox

func countBalls(lowLimit int, highLimit int) int {
	boxes := make(map[int]int)
	keys := make([]int, 0)
	var max int
	for i := lowLimit; i <= highLimit; i++ {
		tmp := 0
		for j := i; j > 0; {
			tmp += j % 10
			j /= 10
		}
		boxes[tmp] += 1
		keys = append(keys, tmp)
	}
	for k := 0; k < len(keys); k++ {
		if boxes[keys[k]] > max {
			max = boxes[keys[k]]
		}
	}
	return max
}
