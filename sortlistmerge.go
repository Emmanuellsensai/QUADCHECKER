package piscine

func SortedListMerge(n1 *NodeI, n2 *NodeI) *NodeI {
	dummy := &NodeI{}
	current := dummy
	p1 := n1
	p2 := n2

	for p1 != nil && p2 != nil {
		if p1.Data <= p2.Data {
			current.Next = p1
			p1 = p1.Next
		} else {
			current.Next = p2
			p2 = p2.Next
		}
		current = current.Next
	}

	if p1 != nil {
		current.Next = p1
	}
	if p2 != nil {
		current.Next = p2
	}

	return dummy.Next
}
