package dp

func maxProduct(nums []int) int {
	n := len(nums)
	// 表示以第 i 个元素结尾的乘积最大子数组
	dpMax := make([]int, n)
	dpMin := make([]int, n)
	for i := 0; i < n; i++ {
		dpMax[i] = nums[i]
		dpMin[i] = nums[i]
	}
	for i := 1; i < n; i++ {
		dpMax[i] = max(max(dpMax[i-1]*nums[i], dpMin[i-1]*nums[i]), nums[i])
		dpMin[i] = min(min(dpMax[i-1]*nums[i], dpMin[i-1]*nums[i]), nums[i])
	}
	res := dpMax[0]
	for i := 1; i < n; i++ {
		res = max(res, dpMax[i])
	}
	return res
}
