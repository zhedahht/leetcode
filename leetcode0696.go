/*
LeetCode 696: https://leetcode.com/problems/count-binary-substrings/description/
*/

package leetcode

func countBinarySubstrings(s string) int {
	digitCounts := [2]int{}
	changes := 0
	result := 0
	for i, r := range s {
		if i == 0 {
			digitCounts[int(r-'0')]++
		} else {
			prev := s[i-1]
			if prev != byte(r) {
				changes++
				if changes >= 2 {
					result += min(digitCounts[0], digitCounts[1])
				}

				digitCounts[int(r-'0')] = 1
			} else {
				digitCounts[int(r-'0')]++
			}
		}
	}

	return result + min(digitCounts[0], digitCounts[1])
}
