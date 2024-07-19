package tree

func connect(root *Node) *Node {
	if root == nil {
		return root
	}
	queue := []*Node{root}
	for len(queue) > 0 {
		levelCnt := len(queue)
		for i := 0; i < levelCnt; i++ {
			if i+1 < levelCnt {
				queue[i].Next = queue[i+1]
			}
		}
		for i := 0; i < levelCnt; i++ {
			node := queue[0]
			if node.Left != nil {
				queue = append(queue, node.Left)
			}
			if node.Right != nil {
				queue = append(queue, node.Right)
			}
			queue = queue[1:]
		}
	}
	return root
}
