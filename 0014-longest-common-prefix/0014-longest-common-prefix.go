func longestCommonPrefix(strs []string) string {
    if len(strs) == 0 {
        return ""
    }

    lcp := strs[0]
    for i := 1; i < len(strs); i++ {
        fmt.Println(lcp)
        lcp = longestCommonPrefixTwo(lcp, strs[i])
        if lcp == "" {
            break
        }
    }

    if lcp == "" {
        return ""
    }

    return lcp
}

func longestCommonPrefixTwo(s1, s2 string) string {
    n := len(s1)
    if len(s2) < n {
        n = len(s2)
    }

    cnt := 0
    for i := range n {
        if s1[i] != s2[i] {
            break
        }

        cnt++
    }

    if cnt == 0 {
        return ""
    }

    fmt.Println(cnt)

    return s1[:cnt]
}