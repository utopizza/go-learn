package array

func twoSum(nums []int, target int) []int {
	m := make(map[int]int)
	for i, n := range nums {
		m[n] = i
	}
	for i, n := range nums {
		if m[target-n] > i {
			return []int{i, m[target-n]}
		}
	}
	return []int{-1, -1}
}
