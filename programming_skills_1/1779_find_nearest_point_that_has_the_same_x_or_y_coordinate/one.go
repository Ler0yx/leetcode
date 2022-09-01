package findNearestPointThatHasTheSameXOrYCoordinate

import (
	"math"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func nearestValidPoint(x int, y int, points [][]int) int {

	smallestIdx := 0
	manhattanDst := math.MaxInt64

	for i := 0; i < len(points); i++ {
		tmp := 0
		if points[i][0] == x || points[i][1] == y {
			tmp = abs(x-points[i][0]) + abs(y-points[i][1])
			if tmp < manhattanDst {
				manhattanDst = tmp
				smallestIdx = i
			}
		}
	}
	if manhattanDst != math.MaxInt64 {
		return smallestIdx
	} else {
		return -1
	}
}

func main() {

}
