package leetcode

func removeElement(nums []int, val int) int {
	i, j := -1, 0
	for i < len(nums) && j < len(nums) {
		if nums[j] != val {
			i++
			nums[i] = nums[j]
		}
		j++
	}
	return i + 1
}
