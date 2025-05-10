package dp

func lengthOfLIS(nums []int) int {
	n := len(nums)
	dp := make([]int, n)
	dp[0] = 1
	maxSub := dp[0]
	for i := 1; i < n; i++ {
		m := 0
		for j := 0; j < i; j++ {
			if nums[i] > nums[j] {
				m = max(m, dp[j])
			}
		}
		dp[i] = m + 1
		maxSub = max(maxSub, dp[i])
	}
	return maxSub
}
