package hash

func longestConsecutive(nums []int) int {
	numSet := map[int]bool{}
	for _, num := range nums {
		numSet[num] = true
	}

	longest := 0
	for num := range numSet {
		if !numSet[num-1] { // 判断当前数字是否序列开头，是则计算序列长度，不是则跳过
			currentNum := num
			currentLength := 1
			for numSet[currentNum+1] {
				currentNum++
				currentLength++
			}
			if longest < currentLength {
				longest = currentLength
			}
		}
	}

	return longest
}
