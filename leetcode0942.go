/*
LeetCode 942: https://leetcode.com/problems/di-string-match/
*/

package leetcode

func diStringMatch(S string) []int {
	d, i := len(S), 0
	result := make([]int, 0)
	for _, c := range S {
		if c == 'D' {
			result = append(result, d)
			d--
		} else {
			result = append(result, i)
			i++
		}
	}

	result = append(result, d)
	return result
}
