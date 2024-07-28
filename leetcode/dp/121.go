package dp

import "github.com/utopizza/go-learn/leetcode/utils"

func maxProfit(prices []int) int {
	profit := 0
	minP := prices[0]
	for _, p := range prices {
		if p < minP {
			minP = p
		} else {
			profit = utils.Max(profit, p-minP)
		}
	}
	return profit
}
