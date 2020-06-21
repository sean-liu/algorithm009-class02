package solutions

// https://leetcode-cn.com/problems/walking-robot-simulation/

// simple apply, turn right -1, turn left -2
func robotSim(commands []int, obstacles [][]int) int {

	current := [2]int{0, 0}

	obstacleMap := make(map[[2]int]bool)
	for i := 0; i < len(obstacles); i++ {
		obstacleMap[[2]int{obstacles[i][0], obstacles[i][1]}] = true
	}

	maxDistance := 0
	// north, east, south, west
	delta := [][]int{{0, 1}, {1, 0}, {0, -1}, {-1, 0}}
	for i, direction := 0, 0; i < len(commands); i++ {
		if commands[i] == -1 {
			direction = (direction + 1) % len(delta)
		} else if commands[i] == -2 {
			direction = (direction + 3) % len(delta)
		} else {
			for j := 0; j < commands[i]; j++ {
				nextPosition := [2]int{current[0] + delta[direction][0], current[1] + delta[direction][1]}
				if _, found := obstacleMap[nextPosition]; !found {
					current = nextPosition
					curDistance := current[0]*current[0] + current[1]*current[1]
					if curDistance > maxDistance {
						maxDistance = curDistance
					}
				}
			}
		}
	}

	return maxDistance
}
