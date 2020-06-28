/*
LeetCode 907: https://leetcode.com/problems/sum-of-subarray-minimums/
*/

package leetcode

func sumSubarrayMins(A []int) int {
	sum := 0
	stack := make([][]int, 0)
	for i, v := range A {
		for len(stack) > 0 && stack[len(stack)-1][0] >= v {
			num := stack[len(stack)-1][0]
			idx := stack[len(stack)-1][1]
			stack = stack[:len(stack)-1]

			prev := -1
			if len(stack) > 0 {
				prev = stack[len(stack)-1][1]
			}

			sum += num * (idx - prev) * (i - idx)
			sum = sum % 1000000007
		}

		stack = append(stack, []int{v, i})
	}

	prev, next := 0, len(A)-1
	for i := range stack {
		num, idx := stack[i][0], stack[i][1]

		sum += num * (idx - prev + 1) * (next - idx + 1)
		sum = sum % 1000000007
		prev = idx + 1
	}

	return sum
}
