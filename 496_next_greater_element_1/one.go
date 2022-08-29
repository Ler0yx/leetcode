package main

import "fmt"

func nextGreaterElement(nums1 []int, nums2 []int) []int {

	var solution []int

	for i := 0; i < len(nums1); i++ {
		for j := 0; j < len(nums2); j++ {
			if nums1[i] == nums2[j] {
				for k := j + 1; ; k++ {
					if k >= len(nums2) {
						solution = append(solution, -1)
						break
					} else if nums2[k] > nums1[i] {
						solution = append(solution, nums2[k])
						break
					}
				}
			}
		}
	}
	return solution
}

func main() {

	nums1 := []int{4, 1, 2}
	nums2 := []int{1, 3, 4, 2}

	fmt.Println(nextGreaterElement(nums1, nums2))

}
