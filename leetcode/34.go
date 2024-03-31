package leetcode

func searchRange(nums []int, target int) []int {
	result := []int{-1, -1}
	if len(nums) == 0 {
		return result
	}

	var lo, hi int

	// 找最左
	for lo, hi = 0, len(nums)-1; lo < hi; {
		mid := (lo + hi) / 2
		if target <= nums[mid] {
			hi = mid
		} else {
			lo = mid + 1
		}
	}
	if nums[hi] == target {
		result[0] = hi
	}

	// 找最右
	for lo, hi = 0, len(nums)-1; lo < hi; {
		mid := (lo + hi + 1) / 2
		if target >= nums[mid] {
			lo = mid
		} else {
			hi = mid - 1
		}
	}
	if nums[lo] == target {
		result[1] = lo
	}

	return result
}
