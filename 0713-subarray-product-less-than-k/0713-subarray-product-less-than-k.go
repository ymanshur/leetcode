func numSubarrayProductLessThanK(nums []int, k int) int {
    l, r := 0, 0
    curr := 1
    ans := 0

    for r < len(nums) {
        if k <= 1 {
            return 0
        }

        curr *= nums[r]

        // maintain window
        for curr >= k {
            curr /= nums[l]
            l++
        }

        // 1 for subarray contains only r
        // r - l for subarray contains r and the rest nums within window
        ans += (r - l) + 1
        r++
    }

    return ans
}