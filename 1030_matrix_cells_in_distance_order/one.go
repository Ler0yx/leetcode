package matrixCellsInDistanceOrder

func abs (x int) int {
    if x < 0 {
        return -x
    }
    return x
}

func allCellsDistOrder(rows int, cols int, rCenter int, cCenter int) [][]int {
    
    distances := make(map[int][][]int)
    solution := [][]int{}
    
    for i := 0; i < rows; i++ {
        
        for j := 0; j < cols; j++ {
            distance := abs(rCenter-i)+abs(cCenter-j)
            tmp := []int{i, j}
            distances[distance] = append(distances[distance], tmp)
        }         
    }
    for k := 0; k < len(distances); k++ {
        for l := 0; l < len(distances[k]); l++ {
            solution = append(solution, distances[k][l])
        }         
    }
    return solution
}