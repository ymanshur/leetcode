func countBits(n int) []int {
    ans := []int{}
    for i := 0; i <= n; i++ {
        ans = append(ans, count(i))
    }

    return ans
}

func count(n int) int {
    if n < 2 {
        return n
    }
    
    total := 0
    for n > 0 {
        total += n % 2
        n /= 2
    }

    return total
}