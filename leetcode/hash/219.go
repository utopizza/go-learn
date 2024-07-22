package hash

func containsNearbyDuplicate(nums []int, k int) bool {
	if k > len(nums) {
		k = k % len(nums)
	}
	m := make(map[int]bool)
	for l, r := 0, 0; r < len(nums); {
		if r <= k { // 构造k宽的窗口
			if m[nums[r]] {
				return true
			}
			m[nums[r]] = true
			r++
		} else { // 滑动窗口
			delete(m, nums[l])
			l++
			if m[nums[r]] {
				return true
			}
			m[nums[r]] = true
			r++
		}
	}
	return false
}
