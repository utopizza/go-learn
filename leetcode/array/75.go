package array

func sortColors(nums []int) {
	lo, hi := 0, len(nums)-1
	for i := 0; i < len(nums); i++ {
		if nums[i] == 0 {
			nums[lo], nums[i] = nums[i], nums[lo]
			lo++
		}
	}
	for j := len(nums) - 1; j >= 0; j-- {
		if nums[j] == 2 {
			nums[hi], nums[j] = nums[j], nums[hi]
			hi--
		}
	}
}
