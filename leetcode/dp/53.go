package dp

import (
	"github.com/utopizza/go-learn/leetcode/utils"
)

func maxSubArray(nums []int) int {
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	maxSub := dp[0]
	for i := 1; i < len(dp); i++ {
		dp[i] = utils.Max(dp[i-1]+nums[i], nums[i])
		maxSub = utils.Max(maxSub, dp[i])
	}
	return maxSub
}
