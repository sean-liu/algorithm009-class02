package solutions

// https://leetcode-cn.com/problems/valid-anagram/

// use map to sum the occurence of chars, time complexity O(n)
func isAnagram(s string, t string) bool {

	if len(s) != len(t) {
		return false
	}

	comparisonResult := make([]int, 26)
	for i := 0; i < len(s); i++ {
		comparisonResult[s[i]-'a']++
		comparisonResult[t[i]-'a']--
	}

	for i := 0; i < len(comparisonResult); i++ {
		if comparisonResult[i] != 0 {
			return false
		}
	}

	return true
}
