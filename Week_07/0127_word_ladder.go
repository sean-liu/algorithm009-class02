package solutions

// https://leetcode-cn.com/problems/word-ladder/

// bilateral BFS, expand the smaller one first, O(n*k), n is the number of word, k is avg length of word
func ladderLength(beginWord string, endWord string, wordList []string) int {
	wordMap := make(map[string]interface{})
	charMap := make([]map[byte]interface{}, len(beginWord))

	for i := 0; i < len(wordList); i++ {
		wordMap[wordList[i]] = nil
		for j := 0; j < len(charMap); j++ {
			if charMap[j] == nil {
				charMap[j] = make(map[byte]interface{})
			}
			charMap[j][wordList[i][j]] = nil
		}
	}

	if _, found := wordMap[endWord]; !found {
		return 0
	}

	q1, q2 := map[string]interface{}{beginWord: nil}, map[string]interface{}{endWord: nil}
	for steps := 0; len(q1) > 0 && len(q2) > 0; {
		steps++
		if len(q2) < len(q1) {
			q1, q2 = q2, q1
		}

		q := make(map[string]interface{})
		for key := range q1 {
			for i := 0; i < len(charMap); i++ {
				for possibleChar := range charMap[i] {
					newWordByte := []byte(key)
					newWordByte[i] = possibleChar
					newWord := string(newWordByte)

					if _, found := q2[newWord]; found {
						return steps + 1
					}

					if _, found := wordMap[newWord]; found {
						q[newWord] = nil
						delete(wordMap, newWord)
					}
				}
			}
		}
		q1 = q
	}

	return 0
}
