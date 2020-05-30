package solutions

// https://leetcode-cn.com/problems/n-ary-tree-level-order-traversal/

// use list to iterate elements
func levelOrder(root *Node) [][]int {
	result := make([][]int, 0)

	currentLayerNodes := []*Node{root}
	for len(currentLayerNodes) > 0 {
		levelResult := make([]int, 0)
		nextLayerNodes := make([]*Node, 0)

		for i := 0; i < len(currentLayerNodes); i++ {
			if currentLayerNodes[i] == nil {
				continue
			}
			levelResult = append(levelResult, currentLayerNodes[i].Val)
			nextLayerNodes = append(nextLayerNodes, currentLayerNodes[i].Children...)
		}

		currentLayerNodes = nextLayerNodes
		if len(levelResult) > 0 {
			result = append(result, levelResult)
		}
	}

	return result
}
