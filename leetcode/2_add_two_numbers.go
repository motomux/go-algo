package leetcode

// 2. Add Two Numbers
/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		return l2
	}
	if l2 == nil {
		return l1
	}

	var res, node *ListNode
	num := 0
	for l1 != nil || l2 != nil {
		x := 0
		y := 0
		if l1 != nil {
			x = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			y = l2.Val
			l2 = l2.Next
		}
		z := num + x + y
		num = z / 10
		val := z % 10
		newNode := &ListNode{Val: val, Next: nil}
		if node == nil {
			node = newNode
			res = node
		} else {
			node.Next = newNode
			node = newNode
		}
	}

	if num != 0 {
		newNode := &ListNode{Val: num, Next: nil}
		node.Next = newNode
	}

	return res
}
