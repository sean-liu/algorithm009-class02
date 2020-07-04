package solutions

// https://leetcode-cn.com/problems/climbing-stairs/

// recursive approach
func climbStairs(n int) int {

	second, first := 2, 1

	if n == 1 {
		return first
	}

	for i := 3; i <= n; i++ {
		third := first + second
		first = second
		second = third
	}

	return second
}
