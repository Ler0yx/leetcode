package distributeCandies

import "sort"

func distributeCandies(candyType []int) int {

	sort.Ints(candyType)
	counter := 1

	for i := 0; i < len(candyType)-1; i++ {
		if candyType[i] != candyType[i+1] {
			counter += 1
		}
	}

	if len(candyType)/2 > counter {
		return counter
	} else {
		return len(candyType) / 2
	}
}

func main() {

}
