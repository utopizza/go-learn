package array

func singleNumberV2(nums []int) int {
	freq := make(map[int]int)
	for _, x := range nums {
		freq[x]++
	}
	for x, cnt := range freq {
		if cnt == 1 {
			return x
		}
	}
	return -1
}
