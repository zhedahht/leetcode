/*
LeetCode 58: https://leetcode.com/problems/length-of-last-word/
*/

package leetcode

import "strings"

func lengthOfLastWord(s string) int {
	segs := strings.Split(s, " ")
	for i := len(segs) - 1; i >= 0; i-- {
		seg := segs[i]
		if len(seg) > 0 {
			return len(seg)
		}
	}

	return 0
}
