package array

func removeDuplicates(nums []int) int {
	slow := 0
	for fast := slow + 1; fast <= len(nums)-1; fast++ {
		if nums[slow] == nums[fast] {
			continue
		}
		slow++
		nums[slow] = nums[fast]
	}
	return slow + 1
}
