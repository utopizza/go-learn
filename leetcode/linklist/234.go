package linklist

func isPalindrome(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return true
	}

	// get mid
	slow, fast := head, head
	for fast != nil && fast.Next != nil {
		fast = fast.Next.Next
		slow = slow.Next
	}
	mid := slow

	// reverse
	head2 := reverseList(mid)

	// compare
	for head != nil && head2 != nil {
		if head.Val != head2.Val {
			return false
		}
		head = head.Next
		head2 = head2.Next
	}

	if head == nil && head2 != nil || head != nil && head2 != nil {
		return false
	}

	return true
}
