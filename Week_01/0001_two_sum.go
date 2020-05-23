package solutions

// https://leetcode-cn.com/problems/two-sum/

// 一遍hash法 时间O(n), 空间O(n)
func twoSum(nums []int, target int) []int {
	if len(nums) < 2 {
		return nil
	}

	unmatched := make(map[int]int)
	for i := 0; i < len(nums); i++ {
		index, found := unmatched[nums[i]]
		if found {
			return []int{index, i}
		}
		unmatched[target-nums[i]] = i
	}

	return nil
}

// 暴力枚举法 O(n^2)
// 两遍hash法
