package double_pointer_linklist

func partition(head *ListNode, x int) *ListNode {
	tempHead := &ListNode{-1, nil}
	tempTail := tempHead

	// 找到小于x的节点，从原链移到新链
	preHead := &ListNode{-1, head}
	p := preHead
	for p != nil && p.Next != nil {
		if p.Next.Val < x {
			q := p.Next
			p.Next = q.Next
			q.Next = nil
			tempTail.Next = q
			tempTail = q
		} else { // 要确保下一个大于等于x才移动，否则会丢失一些节点
			p = p.Next
		}

	}

	// 新链拼接在老链头部
	tempTail.Next = preHead.Next

	return tempHead.Next
}
