package leetcode

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
func sumNumbers(root *TreeNode) int {
	if root == nil {
		return 0
	} else if root.Left == nil && root.Right == nil {
		return root.Val
	} else if root.Left == nil {
		return sum(root.Right, root.Val)
	} else if root.Right == nil {
		return sum(root.Left, root.Val)
	}
	return sum(root.Left, root.Val) + sum(root.Right, root.Val)
}

func sum(node *TreeNode, val int) int {
	newVal := val*10 + node.Val
	if node.Left == nil && node.Right == nil {
		return newVal
	} else if node.Left == nil {
		return sum(node.Right, newVal)
	} else if node.Right == nil {
		return sum(node.Left, newVal)
	}

	return sum(node.Left, newVal) + sum(node.Right, newVal)
}
