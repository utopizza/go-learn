package double_pointer_array

func moveZeroes(nums []int) {
	slow := -1
	for fast := 0; fast < len(nums); fast++ {
		if nums[fast] != 0 {
			slow++
			nums[slow], nums[fast] = nums[fast], nums[slow]
		}
	}
}
