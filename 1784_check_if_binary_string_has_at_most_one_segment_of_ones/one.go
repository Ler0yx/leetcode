package checkIfBinaryStringHasAtMostOneSegmentOfOnes

func checkOnesSegment(s string) bool {
    
    var swt bool

    for i := 0; i < len(s); i++ {
        if string(s[i]) == "0" {
            swt = true
        }
        if swt && string(s[i]) == "1" {
            return false
        }
    }
    return true
}