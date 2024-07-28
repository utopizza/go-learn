package dp

import (
	"github.com/utopizza/go-learn/leetcode/utils"
	"math"
)

func maxProfitV4(k int, prices []int) int {
	n := len(prices)
	if n == 0 {
		return 0
	}

	k = utils.Min(k, n/2)
	buy := make([][]int, n)  // buy[i][j]：前i天j次交易后，手上有股票，最大利润
	sell := make([][]int, n) // sell[i][j]：前i天j次交易后，手上没有股票，最大利润
	for i := range buy {
		buy[i] = make([]int, k+1)
		sell[i] = make([]int, k+1)
	}

	buy[0][0] = -prices[0]
	for i := 1; i <= k; i++ {
		buy[0][i] = math.MinInt64 / 2
		sell[0][i] = math.MinInt64 / 2
	}

	for i := 1; i < n; i++ {
		buy[i][0] = utils.Max(buy[i-1][0], sell[i-1][0]-prices[i])
		for j := 1; j <= k; j++ {
			buy[i][j] = utils.Max(buy[i-1][j], sell[i-1][j]-prices[i])
			sell[i][j] = utils.Max(sell[i-1][j], buy[i-1][j-1]+prices[i])
		}
	}

	res := math.MinInt
	for _, v := range sell[n-1] {
		res = utils.Max(res, v)
	}

	return res
}
