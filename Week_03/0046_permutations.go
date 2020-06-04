package solutions

import "container/list"

// https://leetcode-cn.com/problems/permutations/

// backtracking approach, time & space O(n*n!)
func permute(nums []int) [][]int {
	result := [][]int{}
	length := len(nums)

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
			if used[i] {
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
