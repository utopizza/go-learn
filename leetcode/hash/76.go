package hash

import "math"

func minWindow(s string, t string) string {
	ori, cnt := map[byte]int{}, map[byte]int{}
	for i := 0; i < len(t); i++ {
		ori[t[i]]++
	}

	sLen := len(s)
	minLen := math.MaxInt32
	ansL, ansR := -1, -1

	for l, r := 0, 0; r < sLen; r++ {

		// 窗口右边一直往右扩展，同时对目标字符计数
		c := s[r]
		if r < sLen && ori[c] > 0 {
			cnt[c]++
		}

		// 窗口左边一直往右缩进，同时更新最小覆盖子串长度
		for minWindowCheck(ori, cnt) && l <= r {
			// 更新最小长度
			if r-l+1 < minLen {
				minLen = r - l + 1
				ansL, ansR = l, l+minLen
			}
			// 扣除计数
			c = s[l]
			if _, ok := ori[c]; ok {
				cnt[c]--
			}
			// 左边界右移
			l++
		}
	}

	if ansL == -1 {
		return ""
	}

	return s[ansL:ansR]
}

func minWindowCheck(ori, cnt map[byte]int) bool {
	for k, v := range ori {
		if cnt[k] < v {
			return false
		}
	}
	return true
}
