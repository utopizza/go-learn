package leetcode

func searchInsert(nums []int, target int) int {
	lo, hi := 0, len(nums)
	for lo < hi {
		mid := (lo + hi) / 2
		if target == nums[mid] {
			return mid
		} else if target > nums[mid] {
			lo = mid + 1
		} else {
			hi = mid
		}
	}
	return hi
}
