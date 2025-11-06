func search(nums []int, target int) int {
    l := 0
    r := len(nums) - 1
    for l <= r {
        m := l + (r - l) >> 1

        if target == nums[m] {
            return m
        }
        
        // sorted side
        if nums[l] <= nums[m] {
            if target >= nums[l] && target <= nums[m] {
                r = m - 1
            } else {
                l = m + 1
            }

            continue
        }
        
        if target >= nums[m] && target <= nums[r] {
            l = m + 1
        } else {
            r = m - 1                
        }
    }

    return -1
}