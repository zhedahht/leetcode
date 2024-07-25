/*
LeetCode 434: https://leetcode.com/problems/number-of-segments-in-a-string/description/
*/

package leetcode

func countSegments(s string) int {
	count := 0
	inWord := false
	for _, ch := range s {
		if ch == ' ' {
			inWord = false
		} else if !inWord {
			inWord = true
			count++
		}
	}

	return count
}
