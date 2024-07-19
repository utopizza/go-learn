package linklist

func reorderList(head *ListNode) {
	// get mid
	tempHead := &ListNode{0, head}
	slow, fast := tempHead, tempHead
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	tailHead := slow.Next
	slow.Next = nil // 这里要断开，否则会有换

	// reverse
	var stack []*ListNode
	for p := tailHead; p != nil; p = p.Next {
		stack = append(stack, p)
	}

	// insert
	for p := head; p != nil; {
		if len(stack) > 0 {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			node := p.Next
			p.Next = top
			top.Next = node
			p = node
		} else {
			p = p.Next
		}
	}
}
