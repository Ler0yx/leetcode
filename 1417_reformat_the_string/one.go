package reformatTheString

func reformat(s string) string {
    letters := make([]byte, 0)
    digits := make([]byte, 0)
    var res string
    for i := 0; i < len(s); i++ {
        if s[i] >= 97 && s[i] <= 122 {
            letters = append(letters, s[i])
        }
        if s[i] >= 48 && s[i] <= 57 {
            digits = append(digits, s[i])
        }
    }
    if len(letters) - len(digits) == 0 || len(letters) - len(digits) == 1 {
        for j := 0; j < len(letters); j++ {
        res += string(letters[j])
            if j < len(digits) {
                res += string(digits[j])
            }
        }
    } else if len(digits) - len(letters) == 0 || len(digits) - len(letters) == 1 {
        for j := 0; j < len(digits); j++ {
        res += string(digits[j])
            if j < len(letters) {
                res += string(letters[j])
            }
        }
    } else {
        return ""
    }
    return res
}