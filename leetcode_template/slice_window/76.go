package slice_window

import "math"

func minWindow(s string, t string) string {
	if len(s) < len(t) {
		return ""
	}

	var res string
	minLen := math.MaxInt32

	// t字符存入set
	set := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		set[t[i]]++
	}

	// s初始window，与t长度相同
	window := make(map[byte]int)
	for i := 0; i < len(t); i++ {
		window[s[i]]++
	}

	// 滑动窗口检索s
	left, right := 0, len(t)-1
	for right < len(s) {
		if !isContainsSet(window, set) { // 未包含，继续右扩
			right++
			if right < len(s) {
				window[s[right]]++
			}
			continue
		}
		for isContainsSet(window, set) { // 已包含，尝试左缩，找最短
			if right-left+1 < minLen {
				minLen = right - left + 1
				res = s[left : right+1]
			}
			window[s[left]]--
			left++
		}
	}

	return res
}
