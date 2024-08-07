package dp

import "github.com/utopizza/go-learn/leetcode/utils"

func rob(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return utils.Max(nums[0], nums[1])
	}
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = utils.Max(nums[0], nums[1])
	for i := 2; i < len(dp); i++ {
		dp[i] = utils.Max(dp[i-2]+nums[i], dp[i-1])
	}
	return dp[len(dp)-1]
}
