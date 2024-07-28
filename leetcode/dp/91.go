package dp

func numDecodings(s string) int {
	length := len(s)
	dp := make([]int, length+1)
	dp[0] = 1
	for i := 1; i <= length; i++ {
		if s[i-1] != '0' {
			dp[i] += dp[i-1]
		}
		if i > 1 && s[i-2] != '0' && numLegal(s[i-2:i]) {
			dp[i] += dp[i-2]
		}
	}
	return dp[length]
}

func numLegal(s string) bool {
	return (s[0]-'0')*10+(s[1]-'0') <= 26
}
