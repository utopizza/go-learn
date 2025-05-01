package array

import "math"

func maxProfit(prices []int) int {
	profit := 0
	minVal := math.MaxInt
	for i := 0; i < len(prices); i++ {
		minVal = min(minVal, prices[i])
		profit = max(profit, prices[i]-minVal)
	}
	return profit
}
