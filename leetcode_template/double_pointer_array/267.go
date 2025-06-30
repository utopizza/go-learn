package double_pointer_array

func twoSumV2(numbers []int, target int) []int {
	lo, hi := 0, len(numbers)-1
	for lo <= hi {
		sum := numbers[lo] + numbers[hi]
		if sum == target {
			return []int{lo + 1, hi + 1}
		} else if sum < target {
			lo++
		} else {
			hi--
		}
	}
	return []int{-1, -1}
}
