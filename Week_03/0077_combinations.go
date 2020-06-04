package solutions

import "container/list"

// https://leetcode-cn.com/problems/combinations/

// backtracking approach, time O(k*n!/((n-k)!*k!))
func combine(n int, k int) [][]int {
	result := [][]int{}

	l := list.New()
	var loop func(n, k, start int)
	loop = func(n, k, start int) {
		if l.Len() == k {
			// todo dump to result

			temp := []int{}
			for i := l.Back(); i != nil; i = i.Prev() {
				temp = append(temp, i.Value.(int))
			}
			result = append(result, temp)
			return
		}

		for i := start; i <= n-(k-l.Len())+1; i++ {
			l.PushFront(i)
			loop(n, k, i+1)
			l.Remove(l.Front())
		}
	}

	loop(n, k, 1)
	return result
}
