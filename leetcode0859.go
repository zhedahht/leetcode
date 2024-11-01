/*
LeetCode 859: https://leetcode.com/problems/buddy-strings/description/
*/

package leetcode

func buddyStrings(s string, goal string) bool {
	if len(s) != len(goal) {
		return false
	}

	diffs := make([]int, 0)
	for i := range s {
		if s[i] != goal[i] {
			diffs = append(diffs, i)
		}
	}

	if len(diffs) == 0 {
		chCounts := make([]int, 26)
		for _, ch := range s {
			chCounts[ch-'a']++
			if chCounts[ch-'a'] >= 2 {
				return true
			}
		}

		return false
	}

	if len(diffs) == 2 {
		return s[diffs[0]] == goal[diffs[1]] && s[diffs[1]] == goal[diffs[0]]
	}

	return false
}
