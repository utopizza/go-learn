package dp

func isInterleave(s1 string, s2 string, s3 string) bool {
	m, n, t := len(s1), len(s2), len(s3)
	if (m + n) != t {
		return false
	}

	// dp[i][j]表示s1[:i]与s2[:j]是否能交错组成s3[:i+j]
	dp := make([][]bool, m+1)
	for i := 0; i < m+1; i++ {
		dp[i] = make([]bool, n+1)
	}

	dp[0][0] = true
	for i := 0; i <= m; i++ {
		for j := 0; j <= n; j++ {
			if i > 0 {
				dp[i][j] = dp[i][j] || (dp[i-1][j] && s1[i-1] == s3[i+j-1])
			}
			if j > 0 {
				dp[i][j] = dp[i][j] || (dp[i][j-1] && s2[j-1] == s3[i+j-1])
			}
		}
	}

	return dp[m][n]
}
