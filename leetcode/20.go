package leetcode

func isValid(s string) bool {
	var stack []rune
	for _, c := range s {
		if c == '(' || c == '{' || c == '[' {
			stack = append(stack, c)
		} else {
			if len(stack) == 0 {
				return false
			}
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if c == ')' && top != '(' {
				return false
			}
			if c == '}' && top != '{' {
				return false
			}
			if c == ']' && top != '[' {
				return false
			}
		}
	}
	return len(stack) == 0
}
