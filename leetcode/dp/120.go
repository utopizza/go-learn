package dp

import (
	"github.com/utopizza/go-learn/leetcode/utils"
	"math"
)

func minimumTotal(triangle [][]int) int {
	for i := 1; i < len(triangle); i++ {
		for j := 0; j < len(triangle[i]); j++ {
			if j == 0 {
				triangle[i][j] += triangle[i-1][j]
			} else if j == len(triangle[i])-1 {
				triangle[i][j] += triangle[i-1][j-1]
			} else {
				triangle[i][j] += utils.Min(triangle[i-1][j-1], triangle[i-1][j])
			}
		}
	}
	minVal := math.MaxInt
	bottomRow := triangle[len(triangle)-1]
	for _, v := range bottomRow {
		minVal = utils.Min(minVal, v)
	}
	return minVal
}
