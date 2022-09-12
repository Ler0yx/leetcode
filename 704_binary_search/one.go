package binarySearch

func search(nums []int, target int) int {

	numsMap := make(map[int]int, len(nums))

	for i, v := range nums {
		numsMap[v] = i
	}

	idx, ok := numsMap[target]
	if ok {
		return idx
	} else {
		return -1
	}
}
