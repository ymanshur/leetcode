func majorityElement(nums []int) int {
    freq := map[int]int{}
    n := len(nums)
    for _, num := range nums {
        freq[num]++
        if freq[num] > n / 2 {
            return num
        }
    }
    
    return 0
}