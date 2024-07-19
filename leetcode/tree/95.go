package tree

func generateTrees(n int) []*TreeNode {
	if n == 0 {
		return nil
	}
	return generateTreesRecursive(1, n)
}

func generateTreesRecursive(start, end int) []*TreeNode {
	if start > end {
		return []*TreeNode{nil}
	}
	var allTrees []*TreeNode
	// 枚举可行根节点
	for i := start; i <= end; i++ {
		// 获得所有可行的左子树集合
		leftTrees := generateTreesRecursive(start, i-1)
		// 获得所有可行的右子树集合
		rightTrees := generateTreesRecursive(i+1, end)
		// 从左子树集合中选出一棵左子树，从右子树集合中选出一棵右子树，拼接到根节点上
		for _, left := range leftTrees {
			for _, right := range rightTrees {
				currTree := &TreeNode{i, left, right}
				allTrees = append(allTrees, currTree)
			}
		}
	}
	return allTrees
}
