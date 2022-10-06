package uniqueMorseCodeWords

func uniqueMorseRepresentations2(words []string) int {
    totalDifferences := make(map[string]bool)
    morse := []string{
        ".-","-...","-.-.","-..",".","..-.","--.","....","..",".---","-.-",".-..","--",
        "-.","---",".--.","--.-",".-.","...","-","..-","...-",".--","-..-","-.--","--..",
    }    
    for _, word := range words {
        tmp := ""
        for _, letter := range word {
            tmp += morse[letter-'a']
        }
        totalDifferences[tmp] = true
    }
    return len(totalDifferences)
}