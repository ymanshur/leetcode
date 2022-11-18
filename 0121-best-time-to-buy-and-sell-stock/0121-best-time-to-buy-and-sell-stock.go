func maxProfit(prices []int) int {
    l, r := 0, 1 // left=buy, right=sell
    maxP := 0
    for r < len(prices) {
        // profitable?
        if prices[r] > prices[l] {
            profit := prices[r] - prices[l]
            if profit > maxP {
                maxP = profit
            }
        } else {
            l = r
        }
        r++
    }
    return maxP
}