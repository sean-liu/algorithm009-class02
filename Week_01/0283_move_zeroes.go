package solutions

// https://leetcode-cn.com/problems/move-zeroes/submissions/

// use nonZeroIndex which is a very common method in array related operation
func moveZeroes(nums []int) {
	nonZeroIndex := 0
	for i := 0; i < len(nums); i++ {
		if nums[i] != 0 {
			nums[nonZeroIndex], nums[i] = nums[i], nums[nonZeroIndex]
			nonZeroIndex++
		}
	}
}

// // snowball solution, sounds interesting
// func moveZeroes(nums []int) {
// 	snowballSize := 0
// 	for i := 0; i < len(nums); i++ {
// 		if nums[i] == 0 {
// 			snowballSize++
// 		} else if snowballSize > 0 {
// 			nums[i-snowballSize], nums[i] = nums[i], nums[i-snowballSize]
// 		}
// 	}
// }
