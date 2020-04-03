package swapNodesInPairs

import "testing"

func TestSwapPairs(t *testing.T) {
	node4 := ListNode{Val: 4}
	node3 := ListNode{Val: 3, Next: &node4}
	node2 := ListNode{Val: 2, Next: &node3}
	node1 := ListNode{Val: 1, Next: &node2}

	got := swapPairs(&node1)

	if got.Val != 2 {
		t.Errorf("get head node value: %v, want 2", got.Val)
	} else if got.Next.Val != 1 {
		t.Errorf("got second node value: %v, want 1", got.Next.Val)
	} else if got.Next.Next.Val != 4 {
		t.Errorf("got third node value: %v, want 4", got.Next.Next.Val)
	} else if got.Next.Next.Next.Val != 3 {
		t.Errorf("got fourth node value: %v, want 3", got.Next.Next.Next.Val)
	}
}
