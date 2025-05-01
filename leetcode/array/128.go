package array

func longestConsecutive(nums []int) int {
	m := make(map[int]bool)
	for _, x := range nums {
		m[x] = true
	}
	longest := 0
	for n := range m {
		if !m[n-1] {
			currentN := n
			currentL := 1
			for m[currentN+1] {
				currentN++
				currentL++
			}
			longest = max(longest, currentL)
		}
	}
	return longest
}
