package linklist

func reverseBetween(head *ListNode, left int, right int) *ListNode {
	tempHead := &ListNode{0, head}

	// 确定左边界
	leftNodePrev, leftNode := tempHead, tempHead
	for i := 1; i <= left-1; i++ {
		leftNodePrev = leftNodePrev.Next
	}
	if leftNodePrev != nil {
		leftNode = leftNodePrev.Next
	}

	// 确定右边界
	rightNodeProc, rightNode := tempHead, tempHead
	for j := 1; j <= right; j++ {
		rightNode = rightNode.Next
	}
	if rightNode != nil {
		rightNodeProc = rightNode.Next
	}

	// 用栈反转吧
	var stack []*ListNode
	for p := leftNode; ; {
		stack = append(stack, p)
		if p == rightNode {
			break
		}
		p = p.Next
	}

	// 连接回去原链表
	for p := leftNodePrev; p != nil && len(stack) > 0; p = p.Next {
		p.Next = stack[len(stack)-1]
		stack = stack[:len(stack)-1]
	}
	leftNode.Next = rightNodeProc

	return tempHead.Next
}
