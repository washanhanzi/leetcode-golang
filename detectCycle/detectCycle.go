package detectCycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func detectCycleHashMap(head *ListNode) *ListNode {
	if head == nil {
		return nil
	}
	valMap := make(map[*ListNode]bool)
	for head.Next != nil {
		_, ok := valMap[head.Next]
		if ok {
			return head.Next
		} else {
			valMap[head] = true
		}
		head = head.Next
	}
	return nil
}

func detectCycleTortoiseAndHare(head *ListNode) *ListNode {
	hare := head
	tortoise := head
	//phase 1
	for hare != nil && hare.Next != nil {
		//hare goes 2 steps
		hare = hare.Next.Next
		//tortoise goes 1 step
		tortoise = tortoise.Next
		//phase 2
		if hare == tortoise {
			//use hare as the intersection pointer
			for {
				if head == hare {
					return head
				}
				hare = hare.Next
				head = head.Next
			}
		}
	}
	return nil
}
