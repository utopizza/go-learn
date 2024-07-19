package linklist

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
		return head
	}
	var newHead *ListNode
	for cur := head; cur != nil; {
		next := cur.Next
		cur.Next = newHead
		newHead = cur
		cur = next
	}
	return newHead
}

func reverseList2(head *ListNode) *ListNode {
	newHead, _ := reverseListRecur(head)
	return newHead
}

func reverseListRecur(node *ListNode) (*ListNode, *ListNode) {
	if node == nil || node.Next == nil {
		return node, node
	}
	head, tail := reverseListRecur(node.Next)
	if tail != nil {
		tail.Next = node
		tail = tail.Next
		tail.Next = nil
	}
	return head, tail
}
