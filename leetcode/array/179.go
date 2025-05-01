package array

import (
	"sort"
	"strconv"
)

func largestNumber(nums []int) string {
	sort.Slice(nums, func(i, j int) bool {
		x, y := strconv.Itoa(nums[i]), strconv.Itoa(nums[j])
		return x+y > y+x
	})
	if nums[0] == 0 {
		return "0"
	}
	var res []byte
	for _, x := range nums {
		res = append(res, strconv.Itoa(x)...)
	}
	return string(res)
}
