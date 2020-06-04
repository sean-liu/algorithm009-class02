package solutions

import (
	"container/list"
	"sort"
)

// https://leetcode-cn.com/problems/permutations-ii/

// backtracking approach, sort and trim, time and space O(n*n!)
func permuteUnique(nums []int) [][]int {
	result := [][]int{}
	length := len(nums)

	sort.SliceStable(nums, func(i, j int) bool {
		return nums[i] < nums[j]
	})

	used := make([]bool, length)
	l := list.New()

	var dfs func(depth int)
	dfs = func(depth int) {
		if depth == length {
			temp := []int{}
			for i := l.Front(); i != nil; i = i.Next() {
				temp = append(temp, i.Value.(int))
			}
			result = append(result, temp)
			return
		}

		for i := 0; i < length; i++ {
			if used[i] || i > 0 && nums[i] == nums[i-1] && !used[i-1] {
				// !used[i-1] this condition makes possible for some case to reach depth 3,
				// ortherwise, no one will pass due to every case has 1 and 1'
				continue
			}

			l.PushBack(nums[i])
			used[i] = true
			dfs(depth + 1)
			used[i] = false
			l.Remove(l.Back())
		}
	}

	dfs(0)

	return result
}
