package tree

func zigzagLevelOrder(root *TreeNode) [][]int {
	result := levelOrder(root)
	for i := 0; i < len(result); i++ {
		if i%2 == 1 {
			for l, r := 0, len(result[i])-1; l < r; {
				result[i][l], result[i][r] = result[i][r], result[i][l]
				l++
				r--
			}
		}
	}
	return result
}
