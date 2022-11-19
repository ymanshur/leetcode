func maxSubArray(nums []int) int {
    maxSum := nums[0]
    currSum := 0
    
    for _, num := range nums {
        if currSum < 0 {
            currSum = 0
        }
        currSum += num
        if currSum > maxSum {
            maxSum = currSum
        }
    }
    
    return maxSum
}