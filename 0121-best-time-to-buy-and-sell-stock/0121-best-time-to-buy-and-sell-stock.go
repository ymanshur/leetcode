func maxProfit(prices []int) int {
    maxProfit := 0
    for left, right := 0, 1; right < len(prices); right++ {
        if prices[right] < prices[left] {
            left = right
        }
        profit := prices[right] - prices[left] // profit = sell - buy
        if profit > maxProfit {
            maxProfit = profit
        }
    }
    return maxProfit
}