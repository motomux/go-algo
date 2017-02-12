package linkedlist

// Node represents node of singly linked list
type Node struct {
	Value int
	Next  *Node
}

// LastKthNode returns last kth node of given List
func LastKthNode(node *Node, k int) *Node {
	slow, fast := node, node
	for i := 0; i < k; i++ {
		if fast == nil {
			return nil
		}
		fast = fast.Next
	}

	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}

	return slow
}
