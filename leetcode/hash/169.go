package hash

func majorityElement(nums []int) int {
	num := nums[0]
	count := 1
	for i := 1; i < len(nums); i++ {
		if nums[i] == num {
			count++
		} else {
			count--
			if count == 0 {
				num = nums[i+1]
			}
		}
	}
	return num
}
