package linklist

func getIntersectionNode(headA, headB *ListNode) *ListNode {
	lenA := 0
	for p := headA; p != nil; p = p.Next {
		lenA++
	}

	lenB := 0
	for p := headB; p != nil; p = p.Next {
		lenB++
	}

	// 链表长的先走diff步
	pA, pB := headA, headB
	if lenA > lenB {
		for i := lenA - lenB; i > 0; i-- {
			pA = pA.Next
		}
	} else {
		for i := lenB - lenA; i > 0; i-- {
			pB = pB.Next
		}
	}

	// 一起走
	for pA != nil && pB != nil {
		if pA == pB {
			return pA
		}
		pA = pA.Next
		pB = pB.Next
	}

	return nil
}
