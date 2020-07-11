package solutions

func hammingWeight(num uint32) int {
	result := 0
	for num != 0 {
		result++
		num &= num - 1
	}
	return result
}
