package solutions

// https://leetcode-cn.com/problems/binary-tree-inorder-traversal/

// recursive method
func inorderTraversal(root *TreeNode) []int {
	result := make([]int, 0)

	var inorderTraverse func(root *TreeNode)
	inorderTraverse = func(root *TreeNode) {
		if root == nil {
			return
		}

		inorderTraverse(root.Left)
		result = append(result, root.Val)
		inorderTraverse(root.Right)
	}

	inorderTraverse(root)
	return result
}
