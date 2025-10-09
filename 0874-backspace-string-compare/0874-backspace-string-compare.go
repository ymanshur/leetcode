func backspaceCompare(s string, t string) bool { 
    i, j := len(s) - 1, len(t) - 1

    for {
        bs := 0

        for i >= 0 && (s[i] == '#' || bs > 0) {
            if s[i] == '#' {
                bs++
                i--
            } else {
                i--
                bs--
            }
        }

        bs = 0

        for j >= 0 && (t[j] == '#' || bs > 0) {
            if t[j] == '#' {
                bs++
                j--
            } else {
                j--
                bs--
            }
        }

        if i < 0 || j < 0 { 
            break
        }

        if s[i] != t[j] {
            break
        }

        i--
        j--
    }

    return i < 0 && j < 0
}