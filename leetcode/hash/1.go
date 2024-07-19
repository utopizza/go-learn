package hash

func twoSum(nums []int, target int) []int {
	m1 := make(map[int]bool)
	m2 := make(map[int]int)
	for i, item := range nums {
		m1[item] = true
		m2[item] = i
	}

	for i, item := range nums {
		val := target - item
		if m1[val] && i != m2[val] {
			return []int{i, m2[val]}
		}
	}

	return []int{}
}
