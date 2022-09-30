package countLargestGroup

func countLargestGroup(n int) int {
    
    var sum int
    var sums []int
    var max int
    var count int    
    groups := make(map[int]int)
    
    for i := 1; i <= n; i++ {
        number := i
        sum = 0
        for number > 0 {            
            tmp := number%10
            number /= 10
            sum += tmp
        }
        
        if _, ok := groups[sum]; !ok {
            sums = append(sums, sum)
        }        
        groups[sum] += 1
    }
    
    for i := 0; i < len(sums); i++ {
        if groups[sums[i]] > max {
            max = groups[sums[i]]
        }
    }

    for i := 0; i < len(sums); i++ {               
        if groups[sums[i]] == max {
            count += 1
        }
    }
    return count
}