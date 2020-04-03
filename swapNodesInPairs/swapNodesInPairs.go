package swapNodesInPairs

type ListNode struct {
	Val  int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode {
	//base case
	if head == nil || head.Next == nil {
		return head
	}
	secondNode := head.Next
	head.Next = swapPairs(secondNode.Next)
	secondNode.Next = head
	return secondNode
}
