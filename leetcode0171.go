/*
LeetCode 171: https://leetcode.com/problems/excel-sheet-column-number/
*/

package leetcode

func titleToNumber(s string) int {
	result := 0
	for i := range s {
		result = result*26 + int(s[i]-'A') + 1
	}

	return result
}
