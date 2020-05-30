package solutions

// https://leetcode-cn.com/problems/ugly-number-ii/

// use preprocess and dynamic programming
func nthUglyNumber(n int) int {
	return foundNumbers[n-1]
}

var foundNumbers = findNumber()

func findNumber() []int {

	getMin := func(a, b, c int) int {
		if a <= b && a <= c {
			return a
		} else if b <= a && b <= c {
			return b
		} else {
			return c
		}
	}

	result := make([]int, 1690)
	result[0] = 1

	for i, i2, i3, i5 := 1, 0, 0, 0; i < 1690; i++ {
		result[i] = getMin(result[i2]*2, result[i3]*3, result[i5]*5)

		if result[i] == result[i2]*2 {
			i2++
		}
		if result[i] == result[i3]*3 {
			i3++
		}
		if result[i] == result[i5]*5 {
			i5++
		}
	}

	return result
}
