package dp

func partition(s string) [][]string {
	n := len(s)

	// dp[i][j]表示s[i..j]是否回文串
	dp := make([][]bool, n)
	for i := range dp {
		dp[i] = make([]bool, n)
		for j := range dp[i] {
			dp[i][j] = true
		}
	}
	for i := n - 1; i >= 0; i-- {
		for j := i + 1; j < n; j++ {
			dp[i][j] = (s[i] == s[j]) && dp[i+1][j-1]
		}
	}

	// 回溯搜索解空间
	var solutions [][]string
	var current []string
	partitionBacktrack(s, 0, &current, &solutions, dp)

	return solutions
}

func partitionBacktrack(s string, i int, current *[]string, solutions *[][]string, dp [][]bool) {
	if i == len(s) {
		*solutions = append(*solutions, append([]string{}, *current...))
		return
	}
	for j := i; j < len(s); j++ {
		if dp[i][j] {
			*current = append(*current, s[i:j+1])
			partitionBacktrack(s, j+1, current, solutions, dp)
			*current = (*current)[:len(*current)-1]
		}
	}
}
