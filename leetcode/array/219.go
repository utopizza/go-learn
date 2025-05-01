package array

func containsNearbyDuplicate(nums []int, k int) bool {
	set := make(map[int]bool)
	k = min(k, len(nums))
	for i := 0; i < k; i++ {
		if set[nums[i]] {
			return true
		}
		set[nums[i]] = true
	}
	for i := k; i < len(nums); i++ {
		if set[nums[i]] {
			return true
		}
		set[nums[i]] = true
		delete(set, nums[i-k])
	}
	return false
}
