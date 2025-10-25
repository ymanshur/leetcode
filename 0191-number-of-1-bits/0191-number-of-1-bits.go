func hammingWeight(n int) int {
    cnt := 0
    for n > 0 {
        reminder := n % 2
        if reminder == 1 {
            cnt++
        }
        n /= 2
    }
    return cnt
}