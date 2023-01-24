package finding3-DigitEvenNumbers

func findEvenNumbers(digits []int) []int {
    var res []int
    nums := make(map[int]bool)

    for i := 0; i < len(digits); i++ {
        for j := 0; j < len(digits); j++ {
            if j == i {
                continue
            }
            for k := 0; k < len(digits); k++ {
                if digits[i] == 0 || k == j || k == i {
                    continue
                }
                tmp := digits[i]*100+digits[j]*10+digits[k]
                if tmp%2 == 0 && !nums[tmp] {
                    res = append(res, tmp)
                    nums[tmp] = true
                }
            }
        }
    }
    sort.Ints(res)
    return res
}