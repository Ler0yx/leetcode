package minimumCommonValue

func getCommon(nums1 []int, nums2 []int) int {
	nums1Map := make(map[int]bool)
	for i := 0; i < len(nums1); i++ {
		nums1Map[nums1[i]] = true
	}
	for j := 0; j < len(nums2); j++ {
		if nums1Map[nums2[j]] {
			return nums2[j]
		}
	}
	return -1
}
