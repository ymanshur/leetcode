import (
    "strconv"
    "fmt"
)

func compress(chars []byte) int {
    j := 0
    cnt := 0
    n := len(chars)
    for i := 0; i < n; i++ {
        if chars[i] == chars[j] {
            cnt++
            continue
        }

        j++
        if cnt > 1 {
            digits := strconv.Itoa(cnt)
            for _, d := range digits {
                chars[j] = byte(d)
                j++
            }
        }
        
        chars[j] = chars[i]
        cnt = 1
    }

    j++

    if cnt > 1 {
        digits := strconv.Itoa(cnt)
        for _, d := range digits {
            chars[j] = byte(d)
            j++
        }
    }

    return j
}