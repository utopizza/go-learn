package linklist

func detectCycle(head *ListNode) *ListNode {
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast {
			p := head
			for p != slow {
				slow = slow.Next
				p = p.Next
			}
			return p
		}
	}
	return nil
}
