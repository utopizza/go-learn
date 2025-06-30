package double_pointer_array

func removeDuplicates(nums []int) int {
	lo, hi := 0, 0
	for hi < len(nums) {
		for hi < len(nums) && nums[hi] == nums[lo] { // 和lo相同的，直接skip
			hi++
		}
		if hi < len(nums) && nums[hi] != nums[lo] { // 确认hi停止时，是相同还是不同，不同才保存
			lo++
			nums[lo] = nums[hi]
		}
	}
	return lo + 1
}
