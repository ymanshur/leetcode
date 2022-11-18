func maxProfit(prices []int) int {
    maxProfit := 0
    var left int // buy pointer
    var right int // sell pointer
    for left, right = 0, 1; right < len(prices); right++ {
        if prices[right] < prices[left] {
            left = right
        }
        // profit = sell - buy
        if profit := prices[right] - prices[left]; profit > maxProfit {
            maxProfit = profit
        }
    }
    return maxProfit
}