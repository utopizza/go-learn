package array

import "sort"

func mergeIntervals(intervals [][]int) [][]int {
	sort.Slice(intervals, func(i, j int) bool {
		return intervals[i][0] <= intervals[j][0]
	})
	result := [][]int{intervals[0]}
	for i := 1; i < len(intervals); i++ {
		if intervals[i][0] <= result[len(result)-1][1] {
			result[len(result)-1][1] = max(result[len(result)-1][1], intervals[i][1])
		} else {
			result = append(result, intervals[i])
		}
	}
	return result
}
