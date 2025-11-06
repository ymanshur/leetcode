func majorityElement(nums []int) int {
    major := nums[0]
    cnt := 1
    for i := 1; i < len(nums); i++ {
        if cnt == 0 {
            major = nums[i]
            cnt = 1
            continue
        }

        if nums[i] == major {
            cnt++
        } else {
            cnt--
        }
    }

    return major
}