package solutions

import "math"

// https://leetcode-cn.com/problems/minimum-path-sum/

// dp approach, grid will be the minimun value from top left to this point
func minPathSum(grid [][]int) int {
	m := len(grid)
	if m == 0 {
		return 0
	}
	n := len(grid[0])

	for i := 1; i < m; i++ {
		grid[i][0] += grid[i-1][0]
	}

	for j := 1; j < n; j++ {
		grid[0][j] += grid[0][j-1]
	}

	for i := 1; i < m; i++ {
		for j := 1; j < n; j++ {
			grid[i][j] += int(math.Min(float64(grid[i-1][j]), float64(grid[i][j-1])))
		}
	}

	return grid[m-1][n-1]
}
