package array

func insert(intervals [][]int, newInterval []int) [][]int {
	lo, hi := newInterval[0], newInterval[1]
	var ans [][]int
	merged := false // 记录第一个与目标区间无交集的后继区间
	for _, interval := range intervals {
		if interval[0] > hi { // 当前区间在目标区间的右侧且无交集
			if !merged {
				ans = append(ans, []int{lo, hi})
				merged = true
			}
			ans = append(ans, interval)
		} else if interval[1] < lo { // 当前区间在目标区间的左侧且无交集
			ans = append(ans, interval)
		} else { // 当前区间与目标区间有交集，计算并集
			lo = min(lo, interval[0])
			hi = max(hi, interval[1])
		}
	}
	if !merged {
		ans = append(ans, []int{lo, hi})
	}
	return ans
}
