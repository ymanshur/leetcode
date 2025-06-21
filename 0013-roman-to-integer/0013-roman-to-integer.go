func romanToInt(s string) int {
    romanHash := map[rune]int{
        'I': 1,
        'V': 5,
        'X': 10,
        'L': 50,
        'C': 100,
        'D': 500,
        'M': 1000,
    }

    ans := 0
    prev := 0
    for i := len(s) - 1; i >= 0; i-- {
        num := romanHash[rune(s[i])]
        if prev > num {
            ans -= num
        } else {
            ans += num
        }
        prev = num
    }

    return ans
}