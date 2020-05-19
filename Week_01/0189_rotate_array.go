package solutions

// // 暴力破解
// func rotate(nums []int, k int) {
// 	for ; k > 0; k-- {
// 		extra := nums[len(nums)-1]
// 		for i := len(nums) - 1; i > 0; i-- {
// 			nums[i] = nums[i-1]
// 		}
// 		nums[0] = extra
// 	}
// }

// // 旋转3次
// func rotate(nums []int, k int) {
// 	reverse := func(source []int) {
// 		for i := 0; i < len(source)/2; i++ {
// 			source[i], source[len(source)-1-i] = source[len(source)-1-i], source[i]
// 		}
// 	}

// 	reverse(nums)
// 	reverse(nums[:k%len(nums)])
// 	reverse(nums[k%len(nums):])
// }

// 原地换位
func rotate(nums []int, k int) {
	length := len(nums)
	k = k % length
	requireToSwapCount := 0
	for startIndex := 0; requireToSwapCount < length; startIndex++ {
		currentIndex := startIndex
		previousElement := nums[currentIndex]
		for {
			nextIndex := (currentIndex + k) % length
			extraElement := nums[nextIndex]
			nums[nextIndex] = previousElement
			previousElement = extraElement
			currentIndex = nextIndex
			requireToSwapCount++
			if currentIndex == startIndex {
				break
			}
		}
	}
}

/*
解法
	1.暴力破解 时间 O(n^3)
	2.旋转3次 时间 O(n)
	3.原地换位 时间 O(n)
*/
