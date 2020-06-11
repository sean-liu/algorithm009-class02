package solutions

import "sort"

// https://leetcode-cn.com/problems/assign-cookies/

// greedy, sort kids and cookies in ascending order, then use double index, iterate it
func findContentChildren(g []int, s []int) int {
	sort.Ints(g)
	sort.Ints(s)

	gi, si := 0, 0
	for si < len(s) && gi < len(g) {
		if g[gi] <= s[si] {
			gi++
		}
		si++
	}
	return gi
}
