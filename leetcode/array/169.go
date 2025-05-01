package array

func majorityElement(nums []int) int {
	cnt := 0
	cur := nums[0]
	for _, n := range nums {
		if cnt == 0 {
			cur = n
		}
		if n == cur {
			cnt++
		} else {
			cnt--
		}
	}
	return cur
}
