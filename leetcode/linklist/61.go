package linklist

func rotateRight(head *ListNode, k int) *ListNode {
	if head == nil || k == 0 {
		return head
	}

	k = k % getLinkListLength(head)
	if k == 0 {
		return head
	}

	slow, fast := head, head
	for i := 0; i < k; i++ {
		fast = fast.Next
	}
	for fast.Next != nil {
		slow = slow.Next
		fast = fast.Next
	}

	fast.Next = head
	newHead := slow.Next
	slow.Next = nil

	return newHead
}

func getLinkListLength(head *ListNode) int {
	n := 0
	for p := head; p != nil; p = p.Next {
		n++
	}
	return n
}
