package double_pointer_array

func longestPalindrome(s string) string {
	res := ""
	for i := 0; i < len(s); i++ {
		res = getLongest(res, longestPalindromeExtend(s, i, i))
		if i+1 < len(s) && s[i] == s[i+1] {
			res = getLongest(res, longestPalindromeExtend(s, i, i+1))
		}
	}
	return res
}

func longestPalindromeExtend(s string, lo, hi int) string {
	for lo > 0 && hi < len(s)-1 && s[lo-1] == s[hi+1] {
		lo--
		hi++
	}
	return s[lo:hi]
}

func getLongest(s1, s2 string) string {
	if len(s1) > len(s2) {
		return s1
	}
	return s2
}
