package service

// Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func AddTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	dummyHead := &ListNode{}
	current := dummyHead
	carry := 0

	// Traverse both linked lists
	for l1 != nil || l2 != nil || carry != 0 {
		// Get the values of the current nodes, or 0 if the list is shorter
		val1 := 0
		if l1 != nil {
			val1 = l1.Val
			l1 = l1.Next
		}

		val2 := 0
		if l2 != nil {
			val2 = l2.Val
			l2 = l2.Next
		}

		// Calculate sum of the two digits and the carry
		sum := val1 + val2 + carry
		carry = sum / 10                        // Calculate new carry (if sum >= 10)
		current.Next = &ListNode{Val: sum % 10} // Create new node with the ones place
		current = current.Next                  // Move current pointer forward
	}

	return dummyHead.Next // Return the result starting from the first valid node
}
