package bestTimeToBuyAndSellStock

func maxProfit(prices []int) int {
    var min int = math.MaxInt
    var max int = math.MinInt
    var minIdx int
    var maxIdx int
    var diff int = math.MaxInt

    for i := 0; i < len(prices); i++ {
        if prices[i] <= min {
            min = prices[i]
            minIdx = i
        }
        if prices[i] > max {
            max = prices[i]
            maxIdx = i
        }
    }

    fmt.Println()

    if minIdx > maxIdx {
        for i := 0; i < len(prices); i++ {
            if maxIdx == i {
                max = math.MinInt
                for j := i+1; j < len(prices); j++ {
                    if prices[j] >= max {
                        max = prices[j]
                        maxIdx = j
                    }
                }
            }
            if prices[i] - prices[maxIdx] < diff {                
                diff = prices[i] - prices[maxIdx]
            }
        }
    } else {
        fmt.Println("hier")
        return -(prices[minIdx] - prices[maxIdx])
    }
    return -diff
}