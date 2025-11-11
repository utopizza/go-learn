package stack

import (
	"strconv"
	"strings"
)

func decodeString(s string) string {
	// 用栈，一直入栈，直至遇到']'开始弹出，直至'['
	var stack []string

	i := 0
	for i < len(s) {
		// 左括号入栈
		if i < len(s) && s[i] == '[' {
			stack = append(stack, string(s[i]))
			i++
			continue
		}

		// 遇到字母，取完整单词入栈
		if i < len(s) && 'a' <= s[i] && s[i] <= 'z' {
			var str string
			for i < len(s) && 'a' <= s[i] && s[i] <= 'z' {
				str += string(s[i])
				i++
			}
			stack = append(stack, str)
			continue
		}

		// 遇到数字，取完整数字入栈
		if i < len(s) && '0' <= s[i] && s[i] <= '9' {
			var digit string
			for '0' <= s[i] && s[i] <= '9' {
				digit += string(s[i])
				i++
			}
			stack = append(stack, digit)
			continue
		}

		// 遇到右括号']'，弹出处理repeat后重新入栈
		if i < len(s) && s[i] == ']' {
			// 取括号中的字符串
			var str string
			for len(stack) > 0 && stack[len(stack)-1] != "[" {
				str = stack[len(stack)-1] + str
				stack = stack[:len(stack)-1]
			}
			// 弹出左括号
			stack = stack[:len(stack)-1]
			// 弹出数字
			digit := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			// 计算repeat
			repTime, _ := strconv.Atoi(digit)
			repStr := strings.Repeat(str, repTime)
			// 重新入栈
			stack = append(stack, repStr)
			i++
			continue
		}
	}

	// 栈元素就是最终答案
	var res string
	for _, str := range stack {
		res += str
	}
	return res
}
