func isAnagram(s string, t string) bool {
    if len(t) != len(s) {return false
        return false
    }
    alphaMap := make([]int, 26)
    for _, char := range t {
        alphaMap[char-'a']++
    }
    for _, char := range s {
        alphaMap[char-'a']--
        if alphaMap[char-'a'] < 0 {
            return false
        }
    }
    return true
}