package solutions

// https://leetcode-cn.com/problems/construct-binary-tree-from-preorder-and-inorder-traversal/

// recursive approach, use the first element of preorder to determine the root, then use root to split inorder and preorder
//, time O(n), space O(n)
func buildTree(preorder []int, inorder []int) *TreeNode {
	if len(preorder) == 0 || len(inorder) == 0 || len(preorder) != len(inorder) {
		return nil
	}

	resultNode := &TreeNode{Val: preorder[0]}

	findInInorder := func(target int) int {
		for i := 0; i < len(inorder); i++ {
			if inorder[i] == target {
				return i
			}
		}
		return -1
	}

	index := findInInorder(resultNode.Val)

	resultNode.Left = buildTree(preorder[1:len(inorder[:index])+1], inorder[:index])
	resultNode.Right = buildTree(preorder[len(inorder[:index])+1:], inorder[index+1:])
	return resultNode
}
