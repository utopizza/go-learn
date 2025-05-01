package array

func missingNumber(nums []int) int {
	n := len(nums)
	res := 0
	for i := 0; i <= n; i++ {
		if i < n {
			res ^= nums[i]
		}
		res ^= i
	}
	return res
}
