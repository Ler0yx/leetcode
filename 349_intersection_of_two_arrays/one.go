package intersectionOfTwoArrays

func intersection(nums1 []int, nums2 []int) []int {

	intersectionMap := map[int]bool{}
	intersections := []int{}

	if len(nums1) < len(nums2) {
		for i := 0; i < len(nums1); i++ {
			if !intersectionMap[nums1[i]] {
				for j := 0; j < len(nums2); j++ {
					if nums1[i] == nums2[j] {
						intersections = append(intersections, nums1[i])
						intersectionMap[nums1[i]] = true
						break
					}
				}
			}
		}
	} else {
		for i := 0; i < len(nums2); i++ {
			if !intersectionMap[nums2[i]] {
				for j := 0; j < len(nums1); j++ {
					if nums2[i] == nums1[j] {
						intersections = append(intersections, nums2[i])
						intersectionMap[nums2[i]] = true
						break
					}
				}
			}
		}
	}
	return intersections
}
