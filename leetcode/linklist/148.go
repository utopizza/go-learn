package linklist

func sortList(head *ListNode) *ListNode {
	return sort(head, nil)
}

func sort(head, tail *ListNode) *ListNode {
	if head == nil {
		return head
	}

	if head.Next == tail {
		head.Next = nil // 这里为啥
		return head
	}

	slow, fast := head, head
	for fast != tail {
		slow = slow.Next
		fast = fast.Next
		if fast != tail {
			fast = fast.Next
		}
	}
	mid := slow

	sortedLeft := sort(head, mid)
	sortedRight := sort(mid, tail)

	return merge(sortedLeft, sortedRight)
}

func merge(head1, head2 *ListNode) *ListNode {
	tempHead := &ListNode{}
	cur := tempHead
	for head1 != nil && head2 != nil {
		if head1.Val <= head2.Val {
			cur.Next = head1
			head1 = head1.Next
		} else {
			cur.Next = head2
			head2 = head2.Next
		}
		cur = cur.Next
	}
	if head1 != nil {
		cur.Next = head1
	} else if head2 != nil {
		cur.Next = head2
	}
	return tempHead.Next
}
