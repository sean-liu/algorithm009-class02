package solutions

// https://leetcode-cn.com/problems/binary-tree-preorder-traversal/

// recursive approach
func preorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)

	var preorderTraverse func(root *TreeNode)
	preorderTraverse = func(root *TreeNode) {
		if root == nil {
			return
		}

		result = append(result, root.Val)
		preorderTraverse(root.Left)
		preorderTraverse(root.Right)
	}

	preorderTraverse(root)
	return result
}
