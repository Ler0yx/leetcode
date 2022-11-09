package bestTimeToBuyAndSellStock

func maxProfit(prices []int) int {
	var min int = prices[0]
    var profit int

    for _, price := range prices {
        if price-min > profit {
            profit = price-min
        }
        if price < min {
            min = price
        }
    }
    return profit
}