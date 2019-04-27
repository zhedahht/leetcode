/*
LeetCode 22: https://leetcode.com/problems/generate-parentheses/
*/

package leetcode

func generateParenthesis(n int) []string {
	result := make([]string, 0)
	return helper22(n, n, "", result)
}

func helper22(left, right int, str string, result []string) []string {
	if left == 0 && right == 0 {
		result = append(result, str)
	} else {
		if left > 0 {
			result = helper22(left-1, right, str+"(", result)
		}

		if right > left {
			result = helper22(left, right-1, str+")", result)
		}
	}

	return result
}
