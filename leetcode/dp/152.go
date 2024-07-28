package dp

import "github.com/utopizza/go-learn/leetcode/utils"

func maxProduct(nums []int) int {
	n := len(nums)
	dpMax := make([]int, n)
	dpMin := make([]int, n)
	dpMax[0], dpMin[0] = nums[0], nums[0]
	for i := 1; i < n; i++ {
		num := nums[i]
		dpMax[i] = utils.Max(num, utils.Max(dpMax[i-1]*num, dpMin[i-1]*num))
		dpMin[i] = utils.Min(num, utils.Min(dpMax[i-1]*num, dpMin[i-1]*num))
	}
	ans := nums[0]
	for i := range dpMax {
		ans = utils.Max(ans, dpMax[i])
	}
	return ans
}
