func removeElement(nums []int, val int) int {
    elems := 0
    n := len(nums)
    for i := 0; i < n; i++ {
        if nums[i] == val {
            elems++
        } else if elems > 0 {
            nums[i - elems] = nums[i]
            nums[i] = val
        }
    }

    return n - elems
}