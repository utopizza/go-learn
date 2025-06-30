package double_pointer_linklist

func detectCycle(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		slow = slow.Next
		fast = fast.Next.Next
		if slow == fast { // 完成第一次2倍速追及slow后，从head重新开始第二次追及slow（但这次同速），第二次追及点就是入环点
			p := head
			for p != slow {
				p = p.Next
				slow = slow.Next
			}
			return p
		}
	}
	return nil
}
