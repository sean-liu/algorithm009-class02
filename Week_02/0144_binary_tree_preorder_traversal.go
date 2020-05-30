package solutions

import "container/list"

// https://leetcode-cn.com/problems/binary-tree-preorder-traversal/

// stack approach
func preorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)

	l := list.New()

	l.PushFront(root)
	for l.Len() > 0 {
		current := l.Front().Value.(*TreeNode)
		l.Remove(l.Front())

		if current == nil {
			continue
		}

		result = append(result, current.Val)
		l.PushFront(current.Right)
		l.PushFront(current.Left)
	}

	return result
}

// recursive approach, time O(n), space O(n)
// func preorderTraversal(root *TreeNode) []int {
// 	result := make([]int, 0)

// 	var preorderTraverse func(root *TreeNode)
// 	preorderTraverse = func(root *TreeNode) {
// 		if root == nil {
// 			return
// 		}

// 		result = append(result, root.Val)
// 		preorderTraverse(root.Left)
// 		preorderTraverse(root.Right)
// 	}

// 	preorderTraverse(root)
// 	return result
// }
