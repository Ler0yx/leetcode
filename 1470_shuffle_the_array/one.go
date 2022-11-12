package shuffleTheArray

func shuffle(nums []int, n int) []int {
    part1 := nums[:n]
    part2 := nums[n:]
    nums = make([]int, len(part1)*2)

    for i, j := 0, 0; i < len(part1); i++ {
        nums[j] = part1[i]
        nums[j+1] = part2[i]
        j += 2
    }
    return nums
}