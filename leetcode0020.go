/*
LeetCode 20: https://leetcode.com/problems/valid-parentheses/
*/

package leetcode

func isValid(s string) bool {
	stack := make([]byte, 0)
	for _, ch := range []byte(s) {
		if ch == '(' || ch == '[' || ch == '{' {
			stack = append(stack, ch)
		} else if ch == ')' {
			if len(stack) == 0 || stack[len(stack)-1] != '(' {
				return false
			}

			stack = stack[:len(stack)-1]
		} else if ch == ']' {
			if len(stack) == 0 || stack[len(stack)-1] != '[' {
				return false
			}

			stack = stack[:len(stack)-1]
		} else if ch == '}' {
			if len(stack) == 0 || stack[len(stack)-1] != '{' {
				return false
			}

			stack = stack[:len(stack)-1]
		}
	}

	return len(stack) == 0
}
