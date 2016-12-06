package leetcode

import "container/list"

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

// 199. Binary Tree Right Side View by container/list
func rightSideView2(root *TreeNode) []int {
	if root == nil {
		return []int{}
	}
	res := []int{}
	l := list.New()
	l.PushBack(root)
	for l.Len() > 0 {
		front := l.Front().Value.(*TreeNode)
		res = append(res, front.Val)
		nextList := list.New()
		for el := l.Front(); el != nil; el = l.Front() {
			node := el.Value.(*TreeNode)
			if node.Right != nil {
				nextList.PushBack(node.Right)
			}
			if node.Left != nil {
				nextList.PushBack(node.Left)
			}
			l.Remove(el)
		}
		l = nextList
	}

	return res
}
