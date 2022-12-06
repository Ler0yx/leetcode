package brickWall

import "math"

func leastBricks(wall [][]int) int {

	edges := make(map[int]int)
	keys := make([]int, 0)
	var max int = math.MinInt
	var width int

	for i := 0; i < len(wall); i++ {
		sum := 0
		for j := 0; j < len(wall[i]); j++ {
			sum += wall[i][j]
			if _, ok := edges[sum]; !ok {
				keys = append(keys, sum)
			}
			edges[sum] += 1
		}
		if width == 0 {
			width = sum
		}
	}

	if len(edges) == 1 {
		return len(wall)
	}

	for k := 0; k < len(keys); k++ {
		if edges[keys[k]] > max && keys[k] != width {
			max = edges[keys[k]]
		}
	}
	return len(wall) - max
}
