package array

import "strconv"

func evalRPN(tokens []string) int {
	var res int
	var stack []string
	for _, token := range tokens {
		if token == "+" || token == "-" || token == "*" || token == "/" {
			t1 := stack[len(stack)-2]
			t2 := stack[len(stack)-1]
			v1, _ := strconv.Atoi(t1)
			v2, _ := strconv.Atoi(t2)
			switch token {
			case "+":
				res = v1 + v2
			case "-":
				res = v1 - v2
			case "*":
				res = v1 * v2
			case "/":
				res = v1 / v2
			}
			stack = stack[:len(stack)-2]
			stack = append(stack, strconv.Itoa(res))
		} else {
			stack = append(stack, token)
		}
	}
	res, _ = strconv.Atoi(stack[0])
	return res
}
