package array

import "sort"

func hIndex(citations []int) int {
	sort.Ints(citations)
	n := len(citations)
	var res int
	for i := 0; i < n; i++ {
		res = max(min(n-i, citations[i]), res)
	}
	return res
}
