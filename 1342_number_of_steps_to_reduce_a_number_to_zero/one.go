package numberOfStepsToReduceANumberToZero

func numberOfSteps(num int) int {
    for i := 0; num > 0; i++ {
        if num%2 == 0 {
            num /= 2
        } else {
            num -= 1
        }
        if num == 0 {
            return i+1
        }
    }
    return 0
}