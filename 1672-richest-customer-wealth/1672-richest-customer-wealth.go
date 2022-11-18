func maximumWealth(accounts [][]int) int {
    maxWealth := 0
    for _, account := range accounts {
        var wealth int
        for _, money := range account {
            wealth += money
        }
        if wealth > maxWealth {
            maxWealth = wealth
        }
    }
    return maxWealth
}

/**
* 1. Calculate the wealth of each customer, the complexity is O(m^2)
* 2. If the result larger than maxWealth, assign the result to maxWealth, the complexity is O(m)
*
* The total complexity is O(m^2 + m) ~= O(m^2)
*/