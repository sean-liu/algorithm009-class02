package solutions

// https://leetcode-cn.com/problems/majority-element/

// divide & conquer, time O(nlogn)
func majorityElement(nums []int) int {
	var dfs func(low, high int) int
	dfs = func(low, high int) int {
		if low == high {
			return nums[low]
		}

		mid := (high-low)/2 + low
		left := dfs(low, mid) // mind this, mid at extreme case will be the same as low, cant be mid-1
		right := dfs(mid+1, high)

		countFrequency := func(low, high, target int) int {
			count := 0
			for i := low; i <= high; i++ {
				if nums[i] == target {
					count++
				}
			}
			return count
		}

		if left == right {
			return left
		}

		leftFrequency := countFrequency(low, mid, left)
		rightFrequency := countFrequency(mid+1, high, right)

		if leftFrequency > rightFrequency {
			return left
		}
		return right
	}

	return dfs(0, len(nums)-1)
}

//// Boyer-Moore voting method, time O(n)
// func majorityElement(nums []int) int {
// 	candidate := 0

// 	for i, count := 0, 0; i < len(nums); i++ {
// 		if count == 0 {
// 			candidate = nums[i]
// 		}

// 		if candidate == nums[i] {
// 			count++
// 		} else {
// 			count--
// 		}
// 	}

// 	return candidate
// }
