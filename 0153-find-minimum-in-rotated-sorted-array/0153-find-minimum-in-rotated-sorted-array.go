func findMin(nums []int) int {
    min := nums[0]
    left, right := 0, len(nums) - 1

    for left <= right {
        if nums[left] < nums[right] {
            if nums[left] < min {
                min = nums[left]
                break
            }
        }

        mid := (left + right) / 2
        if nums[mid] < min {
            min = nums[mid]
        }
        
        if nums[left] <= nums[mid] {
            left = mid + 1
        } else {
            right = mid - 1
        }
        
    }
    
    return min
}