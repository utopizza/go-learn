package leetcode

func reverseKGroup(head *ListNode, k int) *ListNode {
	preHead := &ListNode{0, head}
	p := preHead
	for {
		// 分组:[p.Next, ..., q]为一组
		q := p
		n := k
		for ; n > 0 && q != nil && q.Next != nil; n-- {
			q = q.Next
		}
		if n > 0 { // q走不完k步了，说明最后一组不够k个无需反转，可以结束
			break
		}

		// 组内反转，头尾接回原链表
		temp := q.Next
		groupHead, groupTail := reverseOneGroup(p.Next, q)
		p.Next = groupHead
		groupTail.Next = temp

		// 下一组head的前节点
		p = groupTail
	}

	return preHead.Next
}

func reverseOneGroup(head, tail *ListNode) (*ListNode, *ListNode) {
	preHead := &ListNode{}

	var stack []*ListNode
	for p := head; ; p = p.Next {
		stack = append(stack, p)
		if p == tail {
			break
		}
	}

	p := preHead
	for len(stack) > 0 {
		top := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		p.Next = top
		p = p.Next
	}

	return tail, head
}
