package solutions

// https://leetcode-cn.com/problems/jump-game-ii/

// reversed jump from end to start, time O(n^2)
// func jump(nums []int) int {
// 	position := len(nums) - 1
// 	steps := 0
// 	for position > 0 {
// 		for i := 0; i < len(nums); i++ {
// 			if i+nums[i] >= position {
// 				position = i
// 				steps++
// 				break
// 			}
// 		}
// 	}
// 	return steps
// }

// jump from start to end, before jump to next position ,try to find max first, time O(n)
func jump(nums []int) int {
	end := 0
	maxPosition := 0
	steps := 0

	for i := 0; i < len(nums)-1; i++ {
		nextJump := i + nums[i]
		if nextJump > maxPosition {
			maxPosition = nextJump
		}

		if i == end {
			end = maxPosition
			steps++
		}
	}

	return steps
}
