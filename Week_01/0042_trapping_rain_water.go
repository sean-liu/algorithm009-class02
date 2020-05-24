package solutions

import "container/list"

// https://leetcode-cn.com/problems/trapping-rain-water/

// // process column by column, check left max and right max to know how much water will be cumulated in this column. time complexity O(n^2)
// func trap(height []int) int {

// 	result := 0

// 	for i := 1; i < len(height)-1; i++ {
// 		leftMax := 0
// 		for j := i - 1; j >= 0; j-- {
// 			if height[j] > leftMax {
// 				leftMax = height[j]
// 			}
// 		}

// 		rightMax := 0
// 		for j := i + 1; j <= len(height)-1; j++ {
// 			if height[j] > rightMax {
// 				rightMax = height[j]
// 			}
// 		}

// 		minMax := leftMax
// 		if rightMax < minMax {
// 			minMax = rightMax
// 		}

// 		if minMax > height[i] {
// 			result += minMax - height[i]
// 		}
// 	}

// 	return result
// }

// // dp to compute left max and right max, time complexity O(n), space complexity O(n)
// func trap(height []int) int {

// 	result := 0
// 	leftMaxResult := make([]int, len(height))
// 	for i := 1; i < len(height)-1; i++ {
// 		leftMaxResult[i] = leftMaxResult[i-1]
// 		if height[i-1] > leftMaxResult[i] {
// 			leftMaxResult[i] = height[i-1]
// 		}
// 	}

// 	rightMaxResult := make([]int, len(height))
// 	for i := len(height) - 2; i > 0; i-- {
// 		rightMaxResult[i] = rightMaxResult[i+1]
// 		if height[i+1] > rightMaxResult[i] {
// 			rightMaxResult[i] = height[i+1]
// 		}
// 	}

// 	for i := 1; i < len(height)-1; i++ {

// 		minMax := leftMaxResult[i]
// 		if rightMaxResult[i] < minMax {
// 			minMax = rightMaxResult[i]
// 		}

// 		if minMax > height[i] {
// 			result += minMax - height[i]
// 		}
// 	}

// 	return result
// }

// // double cursors, time complexity O(n)
// func trap(height []int) int {
// 	result := 0
// 	leftMax, rightMax := 0, 0

// 	for left, right := 0, len(height)-1; left < right; {
// 		if height[left] < height[right] {
// 			if height[left] > leftMax {
// 				leftMax = height[left]
// 			} else {
// 				result += leftMax - height[left]
// 			}

// 			left++
// 		} else {
// 			if height[right] > rightMax {
// 				rightMax = height[right]
// 			} else {
// 				result += rightMax - height[right]
// 			}
// 			right--
// 		}
// 	}

// 	return result
// }

// use stack, time complexity O(n), space complexity O(n)
func trap(height []int) int {
	result := 0
	list := list.New()

	for i := 0; i < len(height); i++ {

		for list.Len() > 0 && height[i] > height[list.Front().Value.(int)] {
			topElement := list.Front()
			topIndex := topElement.Value.(int)

			list.Remove(topElement)
			if list.Len() == 0 {
				break
			}
			nextTopIndex := list.Front().Value.(int)
			distance := i - nextTopIndex - 1
			minBound := height[i]
			if height[nextTopIndex] < minBound {
				minBound = height[nextTopIndex]
			}
			boundedHeight := minBound - height[topIndex]
			result += distance * boundedHeight
		}

		list.PushFront(i)
	}

	return result
}
