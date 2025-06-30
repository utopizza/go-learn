package double_pointer_array

func removeElement(nums []int, val int) int {
	lo := -1
	for hi := 0; hi < len(nums); hi++ {
		if nums[hi] == val {
			continue
		}
		lo++
		nums[lo] = nums[hi]
	}
	return lo + 1
}
