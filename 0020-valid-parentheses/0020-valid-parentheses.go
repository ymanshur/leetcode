func isOpenBracket(c string) bool {
    for _, openBracket := range []string{"(", "{", "["} {
        if c == openBracket {
            return true
        }
    }
    return false
}

func isValid(s string) bool {
    var stack string
    for _, char := range s {
        // if the char is open bracket, push into stack
        if isOpenBracket(string(char)) {
            stack += string(char)
        } else {
            // if the char is close bracket and the stack is empty, return false
            if len(stack) == 0 {
                return false
            }
            // assume the stack is not empty,
            // if the char is not close bracket of top the stack, return false, otherwise, pop from stack 
            topStack := string(stack[len(stack)-1])
            switch string(char) {
            case ")":
                if  topStack != "(" {
                    return false
                }
            case "}":
                if  topStack != "{" {
                    return false
                }
            case "]":
                if  topStack != "[" {
                    return false
                }
            }
            stack = stack[:len(stack)-1]            
        }
    }
    if len(stack) > 0 {
        return false
    }
    return true
}