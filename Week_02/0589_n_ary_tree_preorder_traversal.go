package solutions

/**
 * Definition for a Node.
 * type Node struct {
 *     Val int
 *     Children []*Node
 * }
 */

func preorder(root *Node) []int {
	result := make([]int, 0)

	var preorderTravel func(*Node)
	preorderTravel = func(root *Node) {
		if root == nil {
			return
		}

		result = append(result, root.Val)
		for i := 0; i < len(root.Children); i++ {
			preorderTravel(root.Children[i])
		}
		return
	}

	preorderTravel(root)
	return result
}
