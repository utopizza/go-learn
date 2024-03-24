package leetcode

func removeDuplicates(nums []int) int {
	i, j := 0, 0
	for i < len(nums)-1 && j < len(nums) {
		if nums[i] == nums[j] {
			j++
		} else {
			i++
			nums[i] = nums[j]
		}
	}
	return i + 1
}
