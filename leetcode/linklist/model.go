package linklist

type ListNode struct {
	Val  int
	Next *ListNode
}

type Node struct {
	Val    int
	Next   *Node
	Random *Node
}

type DoubleListNode struct {
	Key  int
	Val  int
	Pre  *DoubleListNode
	Next *DoubleListNode
}
