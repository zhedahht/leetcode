/*
LeetCode 161: https://leetcode.com/problems/one-edit-distance/description/
*/

package leetcode

func isOneEditDistance(s string, t string) bool {
	i, j := 0, 0
	dist := 0
	for i < len(s) || j < len(t) {
		if i == len(s) {
			dist++
			j++
		} else if j == len(t) {
			dist++
			i++
		} else if s[i] != t[j] {
			dist++
			if len(s) == len(t) {
				i, j = i+1, j+1
			} else if len(s) < len(t) {
				j++
			} else {
				i++
			}
		} else {
			i, j = i+1, j+1
		}

		if dist > 1 {
			break
		}
	}

	return dist == 1
}
