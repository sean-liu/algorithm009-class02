package solutions

// https://leetcode-cn.com/problems/number-of-islands/

// use iteration and dfs to solve it
func numIslands(grid [][]byte) int {
	count := 0

	if len(grid) == 0 || len(grid[0]) == 0 {
		return count
	}

	maxj := len(grid[0])

	var dfsWeep func(i, j int)
	dfsWeep = func(i, j int) {
		if i < 0 || j < 0 || i >= len(grid) || j >= maxj || grid[i][j] == '0' {
			return
		}

		grid[i][j] = '0'
		dfsWeep(i+1, j)
		dfsWeep(i, j+1)
		dfsWeep(i-1, j)
		dfsWeep(i, j-1)
	}

	for i := 0; i < len(grid); i++ {
		for j := 0; j < maxj; j++ {
			if grid[i][j] == '1' {
				count++
				dfsWeep(i, j)
			}
		}
	}

	return count
}

// can also use unionFound struct to solve this
