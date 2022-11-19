func maxSubArray(nums []int) int {
    maxSum, currSum := nums[0], nums[0]
    
    for i := 1; i < len(nums); i++ {
        if currSum < 0 {
            currSum = nums[i]
        } else {
            currSum += nums[i]
        }

        if currSum > maxSum {
            maxSum = currSum
        }
    }

    return maxSum
}