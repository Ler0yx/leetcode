package findWordsThatCanBeFormedByCharacters

func countCharacters(words []string, chars string) int {
  
    wordMap := make(map[byte]int)
    charsMap := make(map[byte]int)
    count := 0
    
    for i := 0; i < len(chars); i++ {
        charsMap[chars[i]] += 1
    }
    
    for i:=0; i < len(words); i++ {
        for j := 0; j < len(words[i]); j++ {
            wordMap[words[i][j]] += 1
        }
        
        for k := 0; k < len(words[i]); k++ {
            if wordMap[words[i][k]] > charsMap[words[i][k]] {
                wordMap = map[byte]int{}
                break
            }
            if k == len(words[i])-1 {
                count += len(words[i])
                wordMap = map[byte]int{}
            }          
        }
    }
    return count
}