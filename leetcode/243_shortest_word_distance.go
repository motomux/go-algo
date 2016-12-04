package leetcode

// 243. Shortest Word Distance
func shortestDistance(words []string, word1 string, word2 string) int {
	word1Index, word2Index, res := len(words)*-1, len(words)*-1, len(words)
	found := false
	for i, word := range words {
		if word == word1 {
			word1Index = i
			if res > word1Index-word2Index {
				res = word1Index - word2Index
				found = true
			}
		}
		if word == word2 {
			word2Index = i
			if res > word2Index-word1Index {
				res = word2Index - word1Index
				found = true
			}
		}
	}

	if !found {
		return -1
	}

	return res
}
