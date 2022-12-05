package sortThePeople

import "sort"

func sortPeople(names []string, heights []int) []string {

	order := make(map[int]int)
	solution := make([]string, len(names))
	var k int

	for i := 0; i < len(heights); i++ {
		order[heights[i]] = i
	}

	sort.Ints(heights)
	for j := len(names) - 1; j >= 0; j-- {
		solution[k] = names[order[heights[j]]]
		k++
	}
	return solution
}
