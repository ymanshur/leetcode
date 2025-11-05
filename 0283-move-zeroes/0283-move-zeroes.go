func moveZeroes(nums []int)  {
    l, r := -1, 0
    for r < len(nums) {
        if l == -1 && nums[r] == 0 {
            l = r
        }

        if l != -1 && nums[r] != 0 {
            nums[l], nums[r] = nums[r], nums[l]
            l++
        }

        r++
    }
}