package tree

func levelOrder(root *TreeNode) [][]int {
	var result [][]int
	var queue []*TreeNode
	if root != nil {
		queue = append(queue, root)
	}
	for len(queue) > 0 {
		levelLength := len(queue)
		var levelNodes []int
		for i := 0; i < levelLength; i++ {
			node := queue[0]
			queue = queue[1:]
			levelNodes = append(levelNodes, node.Val)
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
		}
		result = append(result, levelNodes)
	}
	return result
}
