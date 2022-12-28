package shortestDistanceToTargetStringInACircularArray

func closetTarget(words []string, target string, startIndex int) int {
    left, right := 0, 0
    ctr := false
    n := len(words)
    for i, j := startIndex, startIndex;; {
        switch {
            case words[i] == target:
            return right

            case words[j] == target:
            return left

            case (ctr && i == startIndex) || (n == 1 && words[0] != target):
            return -1
            
            case i == len(words)-1:
            i = -1

            case j == 0:
            j = n
        }
        ctr = true
        left++
        right++
        i = (i+1)%n
        j = (j-1)%n
    }
    return 0
}