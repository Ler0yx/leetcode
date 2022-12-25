package mergeSimilarItems

func mergeSimilarItems(items1 [][]int, items2 [][]int) [][]int {
    items := make(map[int]int)
    var ret [][]int

    for i := 0; i < len(items1); i++ {
            items[items1[i][0]] = items1[i][1]
    }
    for k := 0; k < len(items2); k++ {
            items[items2[k][0]] += items2[k][1]
    }
    for m := 0; m <= 1000; m++ {
        if _, ok := items[m]; ok {
            ret = append(ret, []int{m, items[m]})
        }        
    }
    return ret
}