package dp

import "github.com/utopizza/go-learn/leetcode/utils"

func maxProfitV3(prices []int) int {
	buy1, sell1 := -prices[0], 0
	buy2, sell2 := -prices[0], 0
	for i := 1; i < len(prices); i++ {
		buy1 = utils.Max(buy1, -prices[i])
		sell1 = utils.Max(sell1, buy1+prices[i])
		buy2 = utils.Max(buy2, sell1-prices[i])
		sell2 = utils.Max(sell2, buy2+prices[i])
	}
	return sell2
}
