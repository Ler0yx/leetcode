package keyboardRow

func findWords(words []string) []string {
    
    firstRow := map[string]bool{"q":true, "w":true, "e":true, "r":true, "t":true, "y":true, "u":true, "i":true, "o":true, "p":true, "Q":true, "W":true, "E":true, "R":true, "T":true,                                 "Y":true, "U":true, "I":true, "O":true, "P":true}
    secondRow := map[string]bool{"a":true, "s":true, "d":true, "f":true, "g":true, "h":true, "j":true, "k":true, "l":true, "A":true, "S":true, "D":true, "F":true, "G":true,                                          "H":true, "J":true, "K":true, "L":true}
    thirdRow := map[string]bool{"z":true, "x":true, "c":true, "v":true, "b":true, "n":true, "m":true, "Z":true, "X":true, "C":true, "V":true, "B":true, "N":true, "M":true}
    solution := []string{}
    
    for i := range words {
        switch {
        case firstRow[string(words[i][0])]:
            for j, v := range words[i] {
                if !firstRow[string(v)] {
                    break
                }
                if j == len(words[i])-1 {
                    solution = append(solution, words[i])
                }
            }
            
        case secondRow[string(words[i][0])]:
            for j, v := range words[i] {
                if !secondRow[string(v)] {
                    break
                }
                if j == len(words[i])-1 {
                    solution = append(solution, words[i])
                }
            }
            
        case thirdRow[string(words[i][0])]:
            for j, v := range words[i] {
                if !thirdRow[string(v)] {
                    break
                }
                if j == len(words[i])-1 {
                    solution = append(solution, words[i])
                }
            }
        }
    }
    return solution
}