package checkIfStringIsAPrefixOfArray

func isPrefixString(s string, words []string) bool {
    var text string

    for i := 0; i < len(words); i++ {
        text += words[i]
        if len(text) == len(s) {
            break
        }
    }
    if s == text {
        return true
    }
    return false
}