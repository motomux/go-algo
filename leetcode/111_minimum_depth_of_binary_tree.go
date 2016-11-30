package leetcode

// 111. Minimum Depth of Binary Tree
func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	} else if root.Left == nil && root.Right == nil {
		return 1
	} else if root.Left == nil {
		return minDepth(root.Right) + 1
	} else if root.Right == nil {
		return minDepth(root.Left) + 1
	}

	leftMin := minDepth(root.Left)
	rightMin := minDepth(root.Right)
	if leftMin < rightMin {
		return leftMin + 1
	}
	return rightMin + 1
}
