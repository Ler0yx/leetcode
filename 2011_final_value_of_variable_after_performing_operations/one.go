package finalValueOfVariableAfterPerformingOperations

func finalValueAfterOperations(operations []string) int {
    
    var x int
    for _, operation := range operations {
        if operation == "--X" || operation == "X--" {
            x--
        } else {
            x++
        }
    }
    return x
}