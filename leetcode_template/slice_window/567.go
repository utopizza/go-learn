package slice_window

func checkInclusion(p string, s string) bool {
	lenP := len(p)
	lenS := len(s)
	if lenS < lenP {
		return false
	}

	// p的字符集
	set := make(map[byte]int)
	for i := 0; i < len(p); i++ {
		set[p[i]]++
	}

	// s初始窗口字符集
	window := make(map[byte]int)
	for i := 0; i < len(p)-1; i++ {
		window[s[i]]++
	}

	// 滑动窗口，收集答案
	for i := 0; i <= len(s)-len(p); i++ {
		window[s[i+len(p)-1]]++ // 右边界加入窗口
		if isSameSet(window, set) {
			return true
		}
		window[s[i]]-- // 左边界弹出窗口
	}

	return false
}
