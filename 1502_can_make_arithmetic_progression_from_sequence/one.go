package canMakeArithmeticProgressionFromSequence

import (
	"fmt"
	"sort"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func canMakeArithmeticProgression(arr []int) bool {

	var dst int

	if len(arr) == 2 {
		return true
	}

	sort.Ints(arr)

	if arr[0] < 0 && arr[1] > 0 {
		dst = abs(abs(arr[0]) + abs(arr[1]))
	} else {
		dst = abs(abs(arr[0]) - abs(arr[1]))
	}

	for i := 1; i < len(arr)-1; i++ {
		if abs(abs(arr[i])-abs(arr[i+1])) != dst {

			if arr[i] < 0 && arr[i+1] > 0 {
				if abs(abs(arr[i])+abs(arr[i+1])) == dst {
					fmt.Println("hier")
					continue
				}
			}
			return false
		}
	}
	return true
}
