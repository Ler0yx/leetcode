package truncateSentence

func truncateSentence(s string, k int) string {
    
    var counter int
    var letterCount int
    for i := 0; i < len(s); i++ {
        if s[i] == 32 {
            counter++
            if counter == k {
                letterCount = i
                break
            }
        }
        if i == len(s)-1 {
            return s
        }
    }
    solution := s[:letterCount]
    return solution
}