func removeDuplicates(nums []int) int {
    j := 0
    for i := 0; i < len(nums); i++ {
        if nums[i] != nums[j] {
            j++;
            nums[i], nums[j] = nums[j], nums[i]
        }
    }

    return j + 1
}