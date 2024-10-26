/*
LeetCode 422: https://leetcode.com/problems/valid-word-square/description/
*/

package leetcode

func validWordSquare(words []string) bool {
	for i, word := range words {
		if len(word) > len(words) {
			return false
		}

		for j := 0; j < len(word); j++ {
			if len(words[j]) <= i || word[j] != words[j][i] {
				return false
			}
		}
	}

	return true
}
