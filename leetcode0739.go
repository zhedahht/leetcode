/*
LeetCode 739: https://leetcode.com/problems/daily-temperatures/
*/

package leetcode

func dailyTemperatures(T []int) []int {
	stack := make([]int, 0)
	result := make([]int, len(T))
	for i, t := range T {
		for len(stack) > 0 && T[stack[len(stack)-1]] < t {
			result[stack[len(stack)-1]] = i - stack[len(stack)-1]
			stack = stack[:len(stack)-1]
		}

		stack = append(stack, i)
	}

	for _, i := range stack {
		result[i] = 0
	}

	return result
}
