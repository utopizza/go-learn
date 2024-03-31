package leetcode

import (
	"strconv"
	"strings"
)

func countAndSay(n int) string {
	prev := "1"
	for i := 2; i <= n; i++ {
		cur := &strings.Builder{}
		for j, start := 0, 0; j < len(prev); start = j {
			for j < len(prev) && prev[j] == prev[start] {
				j++
			}
			cur.WriteString(strconv.Itoa(j - start)) // 写个数
			cur.WriteByte(prev[start])               // 写字符
		}
		prev = cur.String()
	}
	return prev
}
