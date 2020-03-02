/*
LeetCode 844: https://leetcode.com/problems/backspace-string-compare/
*/

package leetcode

func backspaceCompare(S string, T string) bool {
	handleBackspace := func(s string) string {
		stack := make([]rune, 0)
		for _, ch := range s {
			if ch == '#' && len(stack) > 0 {
				stack = stack[:len(stack)-1]
			} else if ch != '#' {
				stack = append(stack, ch)
			}
		}

		return string(stack)
	}

	return handleBackspace(S) == handleBackspace(T)
}
