package dp

func minDistance(word1 string, word2 string) int {
	m, n := len(word1), len(word2)
	if m == 0 || n == 0 {
		return n + m
	}

	// dp[i][j]表示word1前i个字母和word2前j个字母的最小编辑距离
	dp := make([][]int, m+1)
	for i := 0; i < len(dp); i++ {
		dp[i] = make([]int, n+1)
	}
	for i := 0; i < len(dp); i++ {
		dp[i][0] = i
	}
	for j := 0; j < len(dp[0]); j++ {
		dp[0][j] = j
	}

	// 对两个单词编辑，全部有6种编辑方法，但实际不同的只有3种：A插入一个字符、B插入一个字符、A替换一个字符
	// 我们可以总是在单词A和B的末尾插入或修改字符，在其他地方操作也可以，但操作顺序不影响最终结果，因此选择在末尾操作便于计算
	// 因此，基于两个单词末尾是否已经相同，可以拆解不同的子问题：最后一个字符是否需要消耗一次操作，即dp[i-1][j-1]是否需要+1
	for i := 1; i < len(dp); i++ {
		for j := 1; j < len(dp[0]); j++ {
			if word1[i-1] == word2[j-1] {
				dp[i][j] = min(
					dp[i][j-1]+1,
					dp[i-1][j]+1,
					dp[i-1][j-1]) // 最后一个字符相同，不需要消耗操作
			} else {
				dp[i][j] = min(
					dp[i][j-1]+1,
					dp[i-1][j]+1,
					dp[i-1][j-1]+1) // 最后一个字符不同，消耗一次操作
			}
		}
	}

	return dp[m][n]
}
