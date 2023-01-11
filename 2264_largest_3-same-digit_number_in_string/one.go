package largest3-Same-DigitNumberInString

func largestGoodInteger(num string) string {
    var n byte = 47
    for i := 0; i < len(num)-2; i++ {
        if num[i] == num[i+1] && num[i] == num[i+2] {
            if num[i] > n {
                n = num[i]
            }
        i += 2    
        }       
    }
    if n == 47 {
        return ""
    }
    return string(n)+string(n)+string(n)
}