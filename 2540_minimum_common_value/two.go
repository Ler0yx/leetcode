package minimumCommonValue

func getCommon(nums1 []int, nums2 []int) int {
    for i, j := 0, 0; i < len(nums1); {        
        if nums1[i] == nums2[j] {
            return nums1[i]
        } else if nums1[i] < nums2[j] {
            i++
        } else {
            if j == len(nums2)-1 {
            return -1
            }
            j++
        }
        
    }
    return -1
}