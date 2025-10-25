func containsDuplicate(nums []int) bool {
    unique := make(map[int]int, len(nums))
    for _, num := range nums {
        if _, ok := unique[num]; ok {
            return true
        }
        unique[num] = 1
    }
    return false
}