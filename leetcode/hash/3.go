package hash

import (
	"github.com/utopizza/go-learn/leetcode/utils"
)

func lengthOfLongestSubstring(s string) int {
	arr := []rune(s)
	set := make(map[rune]bool)
	var l, maxLen int
	for r := 0; r < len(arr); r++ {
		c := arr[r]
		for {
			if _, exist := set[c]; exist {
				delete(set, arr[l])
				l++
			} else {
				break
			}
		}
		maxLen = utils.Max(maxLen, r-l+1)
		set[c] = true
	}
	return maxLen
}
