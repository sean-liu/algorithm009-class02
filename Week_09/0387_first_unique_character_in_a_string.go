package solutions

// https://leetcode-cn.com/problems/first-unique-character-in-a-string/

// loop for twice, first time record, second time find, O(n)
func firstUniqChar(s string) int {
	appeared := [26]int{}

	for i := 0; i < len(s); i++ {
		appeared[s[i]-'a']++
	}

	for i := 0; i < len(s); i++ {
		if appeared[s[i]-'a'] == 1 {
			return i
		}
	}

	return -1
}
