/*
LeetCode 921: https://leetcode.com/problems/minimum-add-to-make-parentheses-valid/
*/

package leetcode

func minAddToMakeValid(S string) int {
	count1 := 0
	for _, r := range S {
		if r == '(' {
			count1++
		} else if count1 > 0 {
			count1--
		}
	}

	count2 := 0
	for i := len(S) - 1; i >= 0; i-- {
		if S[i] == ')' {
			count2++
		} else if count2 > 0 {
			count2--
		}
	}

	return count1 + count2
}
