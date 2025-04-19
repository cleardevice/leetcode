func isValid(s string) bool {
    stack := []rune{}
    for _, c := range s {
        if c == '(' || c == '{' || c == '[' {
            stack = append(stack, c)
        } else if len(stack) == 0 || !match(stack[len(stack)-1], c) {
            return false
        } else {
            stack = stack[:len(stack)-1]
        }
    }
    return len(stack) == 0
}

func match(l, r rune) bool {
    return (l == '(' && r == ')') || (l == '{' && r == '}') || (l == '[' && r == ']');
}
