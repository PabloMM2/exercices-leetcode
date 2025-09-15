func isValid(s string) bool {
	var stack []rune

	for _, value := range s {
		switch value {
		case '[', '{', '(':
			stack = append(stack, value)
		case ']':
			if len(stack) > 0 && stack[len(stack)-1] == '[' {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, value)
			}
		case '}':
			if len(stack) > 0 && stack[len(stack)-1] == '{' {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, value)
			}
		case ')':
			if len(stack) > 0 && stack[len(stack)-1] == '(' {
				stack = stack[:len(stack)-1]
			} else {
				stack = append(stack, value)
			}
		}
	}
	return len(stack) == 0
}