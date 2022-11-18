func canConstruct(ransomNote string, magazine string) bool {
    alphaMap := make([]int, 26)
    for i := 0; i < len(magazine); i++ {
        alphaMap[magazine[i] - 'a']++
    }

    for i := 0; i < len(ransomNote); i++ {
        alphaMap[ransomNote[i] - 'a']--
        if alphaMap[ransomNote[i] - 'a'] < 0 {
            return false
        }
    }
    
    return true
}