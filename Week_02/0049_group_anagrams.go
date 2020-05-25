package solutions

// https://leetcode-cn.com/problems/group-anagrams/

// // sum frequency of each string then use hashmap to check, O(NK), N is length of strs, K is max length of each string
// func groupAnagrams(strs []string) [][]string {

// 	getKey := func(raw string) [26]int {
// 		sumInfo := [26]int{}
// 		for _, v := range []byte(raw) {
// 			sumInfo[v-'a']++
// 		}

// 		return sumInfo
// 	}

// 	result := make([][]string, 0)
// 	resultKeyMap := make(map[[26]int]int)
// 	for i := 0; i < len(strs); i++ {
// 		key := getKey(strs[i])
// 		index, found := resultKeyMap[key]
// 		if !found {
// 			index = len(result)
// 			resultKeyMap[key] = index
// 			result = append(result, []string{})
// 		}
// 		result[index] = append(result[index], strs[i])
// 	}

// 	return result
// }

// sum frequency of each string then use prime number to check, O(NK), N is length of strs, K is max length of each string
// **sometimes we can use prime numbers to store some collective data
func groupAnagrams(strs []string) [][]string {

	getKey := func(raw string) int {

		prime := []int{2, 3, 5, 7, 11, 13, 17, 19, 23, 29, 31, 41, 43, 47, 53, 59, 61, 67, 71, 73, 79, 83, 89, 97, 101, 103}

		key := 1
		for _, v := range []byte(raw) {
			key *= prime[v-'a']
		}

		return key
	}

	result := make([][]string, 0)
	resultKeyMap := make(map[int]int)
	for i := 0; i < len(strs); i++ {
		key := getKey(strs[i])
		index, found := resultKeyMap[key]
		if !found {
			index = len(result)
			resultKeyMap[key] = index
			result = append(result, []string{})
		}
		result[index] = append(result[index], strs[i])
	}

	return result
}
