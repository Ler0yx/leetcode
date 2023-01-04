package findLuckyIntegerInAnArray

func findLucky(arr []int) int {
    res := make(map[int]int)
    for i := 0; i < len(arr); i++ {
        res[arr[i]] += 1 
    }
    for j := 500; j > 0; j-- {
        if j == res[j] {
            return j
        }
    }
    return -1
}