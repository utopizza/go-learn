package dp

func isInterleave(s1 string, s2 string, s3 string) bool {
	m, n := len(s1), len(s2)
	if m+n != len(s3) {
		return false
	}

	dp := make([][]bool, m+1)
	for i := 0; i <= m; i++ {
		dp[i] = make([]bool, n+1)
	}
	for i := 0; i <= m; i++ {
		dp[i][0] = s1[:i] == s3[:i]
	}
	for j := 0; j <= n; j++ {
		dp[0][j] = s2[:j] == s3[:j]
	}

	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if dp[i-1][j] && s1[i-1] == s3[i+j-1] {
				dp[i][j] = true
			} else if dp[i][j-1] && s2[j-1] == s3[i+j-1] {
				dp[i][j] = true
			} else {
				dp[i][j] = false
			}
		}
	}

	return dp[m][n]
}
