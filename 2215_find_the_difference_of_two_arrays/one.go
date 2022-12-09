package findTheDifferenceOfTwoArrays

func findDifference(nums1 []int, nums2 []int) [][]int {
    
    ints1 := make(map[int]bool)
    copied1 := make(map[int]bool)
    ints2 := make(map[int]bool)
    copied2 := make(map[int]bool)
    solution := make([][]int, 2)

    for i := 0; i < len(nums1); i++ {
        ints1[nums1[i]] = true
    }
    for j := 0; j < len(nums2); j++ {
        ints2[nums2[j]] = true
    }
    for k := 0; k < len(nums1); k++ {
        if !ints2[nums1[k]] && !copied1[nums1[k]] {
            solution[0] = append(solution[0], nums1[k])
            copied1[nums1[k]] = true
        }
    }
    for l := 0; l < len(nums2); l++ {
        if !ints1[nums2[l]] && !copied2[nums2[l]] {
            solution[1] = append(solution[1], nums2[l])
            copied2[nums2[l]] = true
        }
    }
    return solution
}