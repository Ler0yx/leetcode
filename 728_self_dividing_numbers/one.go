package selfDividingNumbers

func selfDividingNumbers(left int, right int) []int {
    
    var solution []int

    for i := left; i <= right; i++ {
        tmp := -1
        num := i
        for tmp != 0 {
            tmp = num%10
            num /= 10
            if tmp == 0 || i%tmp != 0 {
                break
            }
            if num == 0 {
                solution = append(solution, i)
            }
        }
    }
    return solution
}