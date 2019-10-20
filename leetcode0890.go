/*
LeetCode 890: https://leetcode.com/problems/find-and-replace-pattern/
*/

package leetcode

func findAndReplacePattern(words []string, pattern string) []string {
	match := func(str1, str2 string) bool {
		charMap := make(map[byte]byte)
		for i, b1 := range []byte(str1) {
			if b2, exists := charMap[b1]; exists {
				if b2 != byte(str2[i]) {
					return false
				}
			} else {
				charMap[b1] = byte(str2[i])
			}
		}

		return true
	}

	var result []string
	for _, word := range words {
		if len(word) == len(pattern) && match(word, pattern) && match(pattern, word) {
			result = append(result, word)
		}
	}

	return result
}
