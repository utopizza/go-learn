package dp

func longestPalindrome(s string) string {
	arr := []rune(s)
	var longestStr string
	for i := 0; i < len(s); i++ {
		// even
		longestStr = getLonger(longestStr, expend(arr, i, i+1))
		// odd
		longestStr = getLonger(longestStr, expend(arr, i-1, i+1))
	}
	return longestStr
}

func expend(arr []rune, left, right int) string {
	for left >= 0 && right < len(arr) {
		if arr[left] == arr[right] {
			left--
			right++
		} else {
			break
		}
	}
	return string(arr[left+1 : right])
}

func getLonger(a, b string) string {
	if len(a) >= len(b) {
		return a
	}
	return b
}
