/*
LeetCode 709: https://leetcode.com/problems/to-lower-case/
*/

package leetcode

import "strings"

func toLowerCase(str string) string {
	var builder strings.Builder
	for _, ch := range str {
		if ch >= 'A' && ch <= 'Z' {
			ch += 32
		}

		builder.WriteRune(ch)
	}

	return builder.String()
}
