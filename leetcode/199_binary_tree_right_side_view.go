package leetcode

// 199. Binary Tree Right Side View
func rightSideView(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	list := []*TreeNode{root}
	for len(list) > 0 {
		res = append(res, list[len(list)-1].Val)
		nextList := make([]*TreeNode, 0)
		for _, node := range list {
			if node.Left != nil {
				nextList = append(nextList, node.Left)
			}
			if node.Right != nil {
				nextList = append(nextList, node.Right)
			}
		}
		list = nextList
	}

	return res
}
