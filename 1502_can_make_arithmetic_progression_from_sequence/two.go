package canMakeArithmeticProgressionFromSequence

func abs2(x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func canMakeArithmeticProgression2(arr []int) bool {
    if len(arr) == 2 {
        return true
    }
    
    sort.Ints(arr)
    difference := arr[0] - arr[1]
    
    for i := 0; i < len(arr)-1; i++ {
        if arr[i]-arr[i+1] != difference {
            return false
        }
    }
    return true
}