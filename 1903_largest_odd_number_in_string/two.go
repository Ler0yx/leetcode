package largestOddNumberInString

func largestOddNumber2(num string) string {
    for i := len(num)-1; i >= 0; i-- {
        if int(num[i]-48)%2 != 0 {
            return num[0:i+1]
        }
    }
    return ""
}