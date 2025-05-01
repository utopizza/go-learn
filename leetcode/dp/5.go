package dp

func longestPalindrome(s string) string {
	var res string
	for i := range s {
		longestPalindromeExtend(s, i, i, &res)
		if i+1 < len(s) {
			longestPalindromeExtend(s, i, i+1, &res)
		}
	}
	return res
}

func longestPalindromeExtend(s string, lo, hi int, res *string) {
	for i, j := lo, hi; i >= 0 && j < len(s) && s[i] == s[j]; {
		if j-i+1 > len(*res) {
			*res = s[i : j+1]
		}
		i--
		j++
	}
}
