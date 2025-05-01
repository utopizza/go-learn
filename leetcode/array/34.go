package array

func searchRange(nums []int, target int) []int {
	l, r := -1, -1
	lo, hi := 0, len(nums)-1
	// 找左边界
	for lo, hi = 0, len(nums)-1; lo < hi; {
		mid := (lo + hi) / 2
		if target <= nums[mid] {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	if 0 <= hi && hi <= len(nums)-1 && nums[hi] == target {
		l = hi
	}
	// 找右边界
	for lo, hi = 0, len(nums)-1; lo < hi; {
		mid := (lo + hi + 1) / 2
		if target >= nums[mid] {
			lo = mid
		} else {
			hi = mid - 1
		}
	}
	if 0 <= lo && lo <= len(nums)-1 && nums[lo] == target {
		r = lo
	}
	return []int{l, r}
}
