package shortestDistanceToTargetStringInACircularArray

func closetTarget(words []string, target string, startIndex int) int {
    left, right, counter := 0, 0, 0

    for i := startIndex; i < len(words); i++ {
        if words[i] == target {
            break
        }        
        if (counter > 0 && i == startIndex) || len(words) == 1 && words[0] != target {
            return -1
        }
        if i == len(words)-1 {
            i = -1
        }
        right++
        counter++
    }
    for j := startIndex; j >= 0; j-- {
        if words[j] == target {
            break
        }
        if j == 0 {
            j = len(words)
        }
        left++
    }
    if left <= right {
        return left
    }
    return right
}