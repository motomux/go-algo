package main

// Node represents node of singly linked list
type Node struct {
	Value int
	Next  *Node
}

// RemoveDup removes duplicate nodes in the given list
func RemoveDup(node *Node) {
	if node == nil {
		return
	}

	table := make(map[int]bool, 0)
	table[node.Value] = true
	for node.Next != nil {
		next := node.Next
		if _, ok := table[next.Value]; ok {
			node.Next = next.Next
		} else {
			table[next.Value] = true
			node = next
		}
	}
}
