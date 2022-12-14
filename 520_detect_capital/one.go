package detectCapital

func detectCapitalUse(word string) bool {
    
    switch {
        case len(word) == 1:
            return true

        case word[0] < 97 && word[1] < 97:
        for i := 2; i < len(word); i++ {
            if word[i] >= 96 {
                return false
            }
        }

        default:
        for j := 1; j < len(word); j++ {
            if word[j] < 97 {
                return false
            }
        }
    }
    return true
}