func fizzBuzz(n int) []string {
    result := make([]string, n+1)
    fizz := 0
    buzz := 0
    for i := 1; i <= n; i++ {
        fizz++
        buzz++
        if (fizz == 3) && (buzz == 5) {
            result[i] = "FizzBuzz"
            fizz = 0
            buzz = 0
        } else if fizz == 3 {
            result[i] = "Fizz"
            fizz = 0
        } else if buzz == 5 {
            result[i] = "Buzz"
            buzz = 0
        } else {
            result[i] = strconv.Itoa(i)
        }
    }
    return result[1:]
}

/**
* Greedy solution: traversal
*/