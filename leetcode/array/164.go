package array

import (
	"math"
	"slices"
)

func maximumGap(nums []int) int {
	minV := slices.Min(nums)
	maxV := slices.Max(nums)
	if gap := maxV - minV; gap <= 1 {
		return gap
	}

	n := len(nums)
	d := (maxV - minV + n - 2) / (n - 1) // d是桶大小

	// 桶结构定义，每个桶内记录最大最小值
	type pair struct {
		min int
		max int
	}
	buckets := make([]pair, (maxV-minV)/d+1)
	for i := range buckets {
		buckets[i] = pair{math.MaxInt, math.MinInt}
	}

	// 对元素分桶，并更新每个桶的最大最小值
	for _, x := range nums {
		b := &buckets[(x-minV)/d]
		b.min = min(b.min, x)
		b.max = max(b.max, x)
	}

	// 遍历桶间gap
	preMax := math.MaxInt
	gap := math.MinInt64
	for _, b := range buckets {
		if b.min != math.MaxInt {
			gap = max(gap, b.min-preMax)
			preMax = b.max
		}
	}
	return gap
}
