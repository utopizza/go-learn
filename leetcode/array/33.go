package array

func search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}
	lo, hi := 0, len(nums)-1
	for lo <= hi {
		mid := lo + (hi-lo)/2
		if nums[mid] == target {
			return mid
		}
		if nums[0] <= nums[mid] { // 左半部有序
			if nums[0] <= target && target < nums[mid] {
				hi = mid - 1
			} else {
				lo = mid + 1
			}
		} else { // 右半部有序
			if nums[mid] < target && target <= nums[len(nums)-1] {
				lo = mid + 1
			} else {
				hi = mid - 1
			}
		}
	}
	return -1
}
