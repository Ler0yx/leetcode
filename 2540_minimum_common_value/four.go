package minimumCommonValue

func getCommon(nums1 []int, nums2 []int) int {
    for len(nums1) > 0 {        
        if len(nums2) == 0 {
            return -1
        }
        if nums1[0] < nums2[0] {
            nums1 = nums1[1:]
        } else if nums2[0] < nums1[0] {
            nums2 = nums2[1:]
        } else {
            return nums1[0]
        }
    }
    return -1
}