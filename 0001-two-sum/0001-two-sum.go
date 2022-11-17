func twoSum(nums []int, target int) []int {
    numsLength := len(nums)
    numMap := make(map[int]int)
    for i := 0; i < numsLength; i++ {
        complement := target - nums[i]
        if j, ok := numMap[complement]; ok && j != i {
            return []int{i, j}
        }
        numMap[nums[i]] = i
    }
    return nil
}