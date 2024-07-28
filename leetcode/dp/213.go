package dp

import "github.com/utopizza/go-learn/leetcode/utils"

func robV2(nums []int) int {
	if len(nums) == 1 {
		return nums[0]
	}
	if len(nums) == 2 {
		return utils.Max(nums[0], nums[1])
	}

	n := len(nums)

	// 从第一间偷起，不偷最后一间，范围：nums[0...n-2]
	dp := make([]int, len(nums))
	dp[0] = nums[0]
	dp[1] = utils.Max(nums[0], nums[1])
	for i := 2; i <= n-2; i++ {
		dp[i] = utils.Max(dp[i-2]+nums[i], dp[i-1])
	}

	// 从第二间偷起，偷最后一间，范围：nums[1...n-1]
	dp2 := make([]int, len(nums))
	dp2[1] = nums[1]
	dp2[2] = utils.Max(nums[1], nums[2])
	for i := 3; i <= n-1; i++ {
		dp2[i] = utils.Max(dp2[i-2]+nums[i], dp2[i-1])
	}

	return utils.Max(dp[n-2], dp2[n-1])
}
