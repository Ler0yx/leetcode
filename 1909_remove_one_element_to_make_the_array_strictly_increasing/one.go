package removeOneElementToMakeTheArrayStrictlyIncreasing

func canBeIncreasing(nums []int) bool {
    for i := 0; i < len(nums)-1; i++ {
        if nums[i] >= nums[i+1] {
            if i != 0 && nums[i+1] <= nums[i-1] {
                tmp := nums[i+2:]
                nums = nums[:i+1]
                nums = append(nums, tmp...)
                break
            }
            tmp := nums[i+1:]
            nums = nums[:i]
            nums = append(nums, tmp...)
                break
        }
    }
    for j := 0; j < len(nums)-1; j++ {
        if nums[j] >= nums[j+1] {
            return false
        }
    }
    return true
}