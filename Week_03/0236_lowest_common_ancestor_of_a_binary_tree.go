package solutions

// https://leetcode-cn.com/problems/lowest-common-ancestor-of-a-binary-tree/

// recursive approach time O(n), space O(n)
func lowestCommonAncestor(root, p, q *TreeNode) *TreeNode {

	var loop func(root *TreeNode) *TreeNode
	loop = func(root *TreeNode) *TreeNode {
		if root == nil {
			return nil
		}

		if root.Val == p.Val || root.Val == q.Val {
			return root
		}

		leftFound := loop(root.Left)
		rightFound := loop(root.Right)

		if leftFound != nil && rightFound != nil {
			return root
		}

		if leftFound != nil {
			return leftFound
		}
		return rightFound
	}

	return loop(root)
}
