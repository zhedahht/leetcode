/*
LeetCode 503: https://leetcode.com/problems/next-greater-element-ii/description/
*/

package leetcode

func nextGreaterElements(nums []int) []int {
	stack := make([]int, 0)
	i := 0
	result := make([]int, len(nums))
	done := make([]bool, len(nums))
	looped := false
	for true {
		for len(stack) > 0 && nums[stack[len(stack)-1]] < nums[i] {
			index := stack[len(stack)-1]
			result[index] = nums[i]
			done[index] = true
			stack = stack[:len(stack)-1]
		}

		if !done[i] {
			stack = append(stack, i)
		}

		if looped {
			break
		}

		i++
		if i >= len(nums) {
			i = 0
		}

		if i == stack[0] {
			looped = true
		}
	}

	for _, index := range stack {
		result[index] = -1
	}

	return result
}
