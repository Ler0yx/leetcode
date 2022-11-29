package findFirstPalindromicStringInTheArray

func firstPalindrome(words []string) string {

    for i:= 0; i < len(words); i++ {
        if len(words[i]) == 1 {
        return words[i]
        }
        k := len(words[i])-1
        for j := 0; j < len(words[i])/2; j++ {
            if words[i][j] != words[i][k] {
                break
            }
            if j == len(words[i])/2-1 {
                return words[i]
            }
            k--
        }
    }
    return ""
}