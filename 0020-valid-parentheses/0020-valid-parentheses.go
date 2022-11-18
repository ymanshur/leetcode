func isOpenBracket(char rune) bool {
    for _, openBracket := range []rune{'(', '{', '['} {
        if char == openBracket {
            return true
        }
    }
    return false
}

func isValid(s string) bool {
    var stack []rune
    var closeBracket = map[rune]rune{
        '(' : ')',
        '{' : '}',
        '[' : ']',
    }
    for _, char := range s {
        // if the char is open bracket, push into stack
        if isOpenBracket(char) {
            stack = append(stack, char)
        } else {
            // if the char is close bracket and the stack is empty, return false
            if len(stack) == 0 {
                return false
            }
            // assume the stack is not empty,
            // if the char is not close bracket of top the stack, return false, otherwise, pop from stack 
            lastIndex := len(stack)-1
            topStack := stack[lastIndex]
            if  char != closeBracket[topStack] {
                    return false
            } else {                
                stack = stack[:len(stack)-1]            
            }
        }
    }
    if len(stack) > 0 {
        return false
    }
    return true
}