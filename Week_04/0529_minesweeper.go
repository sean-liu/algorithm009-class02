package solutions

// https://leetcode-cn.com/problems/minesweeper/

// use iteration & dfs
func updateBoard(board [][]byte, click []int) [][]byte {
	xBound := len(board)
	if xBound == 0 {
		return board
	}
	yBound := len(board[0])

	isInBoundary := func(x, y int) bool {
		return x >= 0 && x < xBound && y >= 0 && y < yBound
	}

	deltaX := []int{-1, 1, 1, -1, -1, 1, 0, 0}
	deltaY := []int{-1, 1, -1, 1, 0, 0, -1, 1}

	checkMineCount := func(x, y int) int {
		mineCount := 0
		for i := 0; i < 8; i++ {
			if newX, newY := x+deltaX[i], y+deltaY[i]; isInBoundary(newX, newY) && board[newX][newY] == 'M' {
				mineCount++
			}
		}
		return mineCount
	}

	var expand func(x, y int)
	expand = func(x, y int) {

		if !isInBoundary(x, y) {
			return
		}

		if board[x][y] == 'E' {
			board[x][y] = 'B'
			mineCount := checkMineCount(x, y)
			if mineCount > 0 {
				board[x][y] = byte('0' + mineCount)
			} else {
				for i := 0; i < 8; i++ {
					expand(x+deltaX[i], y+deltaY[i])
				}
			}
		} else if board[x][y] == 'M' {
			board[x][y] = 'X'
		}
	}

	expand(click[0], click[1])
	return board
}
