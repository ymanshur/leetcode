func longestPalindrome(s string) int {
    freq := map[rune]int{}
    for _, c := range s {
        freq[c]++
    }

    ans := 0
    odd := 0
    for _, total := range freq {
        if total % 2 == 0 {
            ans += total
        } else {
            ans += (total - 1)
            odd++
        }
    }

    if odd > 0 {
        ans++
    }

    return ans
}