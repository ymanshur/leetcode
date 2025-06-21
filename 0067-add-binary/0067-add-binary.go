import (
	"strconv"
)

func addBinary(a string, b string) string {
    if len(a) < len(b) {
        a, b = b, a
    }

    ans := ""
    carr := 0
    j := len(b) - 1
    for i := len(a) - 1; i >= 0; i-- {
        bit1 := int(a[i] - '0')
        bit2 := 0
        if j >= 0 {
            bit2 = int(b[j] - '0')
            j--
        }
        sum := bit1 + bit2 + carr
        carr = sum / 2
        ans = strconv.Itoa(sum % 2) + ans
    }

    if carr > 0 {
        ans = strconv.Itoa(carr) + ans
    }

    return ans
}