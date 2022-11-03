package validPalindrome

func isPalindrome(s string) bool {
   
    cleanString := ""
    var k int
    
    for i := 0; i < len(s); i++ {
        if s[i] > 64 && s[i] < 91 {
            cleanString += string(s[i] + 32)
        } else if s[i] > 47 && s[i] < 58 {
            cleanString += string(s[i])
        } else if s[i] < 97 || s[i] > 122 {
            continue
        } else {
            cleanString += string(s[i])
        }
    }
    k = len(cleanString)-1
    
    for j := 0; j < len(cleanString)/2; j++ {
        if cleanString[j] != cleanString[k] {
            return false
        }
        k--
    }    
    return true
}