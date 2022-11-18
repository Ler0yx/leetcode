package checkIfTheSentenceIsAPangram

func checkIfPangram3(sentence string) bool {
    letters := make(map[byte]bool)

    if len(sentence) < 26 {
        return false
    }

    for i := 0; i < len(sentence); i++{
        letters[sentence[i]] = true
        if len(letters) == 26 {
            return true
        }
    }
    return false
}