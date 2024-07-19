package tree

func recoverTree(root *TreeNode) {
	var inorderArr []int
	inorderMap := make(map[int]*TreeNode)

	// inorder traversal
	var stack []*TreeNode
	for len(stack) > 0 || root != nil {
		// left child
		for root != nil {
			stack = append(stack, root)
			root = root.Left
		}
		// pop and visit
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		inorderArr = append(inorderArr, top.Val)
		inorderMap[top.Val] = top
		// right child
		root = top.Right
	}

	// find incorrect nodes
	var minNode, maxNode *TreeNode
	for i := 0; i < len(inorderArr)-1; i++ {
		if inorderArr[i] >= inorderArr[i+1] {
			maxNode = inorderMap[inorderArr[i]]
			break
		}
	}
	for j := len(inorderArr) - 1; j > 0; j-- {
		if inorderArr[j] <= inorderArr[j-1] {
			minNode = inorderMap[inorderArr[j]]
			break
		}
	}

	// fix
	minNode.Val, maxNode.Val = maxNode.Val, minNode.Val
}
