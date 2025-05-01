package array

func containsDuplicate(nums []int) bool {
	set := make(map[int]bool)
	for _, n := range nums {
		if set[n] {
			return true
		}
		set[n] = true
	}
	return false
}
