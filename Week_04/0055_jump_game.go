package solutions

// https://leetcode-cn.com/problems/jump-game/

func canJump(nums []int) bool {

	for i, rightMost := 0, 0; i < len(nums); i++ {
		if i > rightMost {
			break
		}

		eachNextPosition := i + nums[i]
		if eachNextPosition > rightMost {
			rightMost = eachNextPosition
		}
		if rightMost >= len(nums)-1 {
			return true
		}
	}
	return false
}
