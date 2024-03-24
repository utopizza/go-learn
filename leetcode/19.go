package leetcode

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dummy := &ListNode{0, head}
	slow, fast := dummy, head

	// 快针先走n步
	for i := 0; i < n; i++ {
		fast = fast.Next
	}

	// 两针保持同速度一起走，快针到链尾
	for ; fast != nil; fast = fast.Next {
		slow = slow.Next
	}

	// 慢针的next即为待删节点
	slow.Next = slow.Next.Next

	return dummy.Next
}
