package removeLetterToEqualizeFrequency

func equalFrequency(word string) bool {
    amounts := make(map[byte]int)
    amts := make([]int, 0)
    swt := false
    for i := 0; i < len(word); i++ {
        amounts[word[i]] += 1
    }
    for j := 96; j < 123; j++ {
        amts = append(amts, amounts[j])
    }
    sort.Ints
    for k := 0; k < len(amts); k++ {
        if swt = false && amts[0]-amts[1] == -1 {
            continue
        } else {
            break
        }
    }
}