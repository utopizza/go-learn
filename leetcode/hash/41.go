package hash

import "github.com/utopizza/go-learn/leetcode/utils"

func firstMissingPositive(nums []int) int {
	n := len(nums)

	// 非正数，改为N+1
	for i := 0; i < n; i++ {
		if nums[i] <= 0 {
			nums[i] = n + 1
		}
	}

	// 绝对值小于等于n的元素，对应位置改为负数，表示该整数已存在
	// 原数组num此时已经变成一个排序好的哈希表，负数的元素已存在
	for i := 0; i < n; i++ {
		num := utils.Abs(nums[i], 0)
		if num <= n {
			nums[num-1] = -utils.Abs(nums[num-1], 0)
		}
	}

	// 返回第一个非负元素的+1值
	for i := 0; i < n; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}

	return n + 1
}
