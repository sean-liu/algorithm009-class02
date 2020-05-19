package solutions

func plusOne(digits []int) []int {

	requireToCarry := digits[len(digits)-1] == 9
	digits[len(digits)-1] = (digits[len(digits)-1] + 1) % 10
	for i := len(digits) - 2; i >= 0; i-- {
		if requireToCarry {
			requireToCarry = digits[i] == 9
			digits[i] = (digits[i] + 1) % 10
		}
	}

	if requireToCarry {
		result := []int{1}
		result = append(result, digits...)
		return result
	}

	return digits
}

/*
第一次，花了点时间，优化了一下代码，评论区的解题思路和我想的差不多，这题就算先过了。
其中，主要有2个优化过程
	1.把最后一位的加法给拉到循环外来处理
	2.把最大一位是否要进位从循环拉出来
*/
