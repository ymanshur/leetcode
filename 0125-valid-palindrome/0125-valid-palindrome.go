import "regexp"

func isPalindrome(s string) bool {
    // make the input lowercase
    s = strings.ToLower(s)
    // make the input clean from non-alphanum
    re := regexp.MustCompile(`[^a-z0-9]+`)
    s = re.ReplaceAllString(s, "")
    // iterate the cleaned input with left and right pointer
    for i, j := 0, len(s)-1; i < j; i, j = i+1, j-1 {
        if s[i] != s[j] {
            return false
        }
    }
    return true
}

/**
* 1. Modify the input into lowercase and clean from non-alphanum
* 2. Loop from start and end index, then match both the char
* 3. Return false if found not macth char
*/