/*
LeetCode 962: https://leetcode.com/problems/maximum-width-ramp/
*/

package leetcode

func maxWidthRamp(A []int) int {
	stack := make([]int, 0)
	for i, num := range A {
		if len(stack) == 0 || A[stack[len(stack)-1]] > num {
			stack = append(stack, i)
		}
	}

	max := 0
	for i := len(A) - 1; i >= 0; i-- {
		for len(stack) > 0 && A[stack[len(stack)-1]] <= A[i] {
			index := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			if max < i-index {
				max = i - index
			}
		}
	}

	return max
}
