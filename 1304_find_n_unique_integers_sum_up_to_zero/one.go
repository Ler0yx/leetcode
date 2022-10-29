package findNUniqueIntegersSumUpToZero

func sumZero(n int) []int {
    sum := 0
    result := make([]int, n)
    
    if n == 1 {
        result[0] = 0
        return result
    }
    
    for i := 1; i < len(result); i++ {
        result[i] = i
        sum += i
    }
    
    result[0] = -sum
    
    return result
}