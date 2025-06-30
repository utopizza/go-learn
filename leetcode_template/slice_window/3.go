package slice_window

func lengthOfLongestSubstring(s string) int {
	m := make(map[byte]bool)
	res := 0
	lo, hi := 0, 0
	for hi < len(s) {
		c := s[hi]
		if !m[c] { // 无重复，继续右扩，更新窗口内容和长度最大值
			m[c] = true
			res = max(res, hi-lo+1)
			hi++
		} else { // 遇到重复，左边缩进直至无重复
			for m[c] {
				delete(m, s[lo])
				lo++
			}
		}
	}
	return res
}
