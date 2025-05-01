package array

func removeElement(nums []int, val int) int {
	i := -1
	for j := 0; j <= len(nums)-1; j++ {
		if nums[j] == val {
			continue
		}
		i++
		nums[i] = nums[j]
	}
	return i + 1
}
