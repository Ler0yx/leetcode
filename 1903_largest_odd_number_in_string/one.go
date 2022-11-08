package largestOddNumberInString

func largestOddNumber(num string) string {
    var substr string
    for i := len(num)-1; i >= 0; i-- {
        if int(num[i]-48)%2 != 0 {
            for j := 0; j <= i; j++ {
                substr += string(num[j])
            }
            break
        }
    }
    return substr
}