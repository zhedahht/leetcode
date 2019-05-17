/*
LeetCode 84: https://leetcode.com/problems/largest-rectangle-in-histogram/
*/

package leetcode

import (
	"math"
)

func largestRectangleArea(heights []int) int {
	if len(heights) == 0 {
		return 0
	}

	stack := make([]int, 0)
	maxArea := 0
	for i, height := range heights {
		for len(stack) > 0 && heights[stack[len(stack)-1]] > height {
			top := stack[len(stack)-1]
			stack = stack[:len(stack)-1]
			left := -1
			if len(stack) > 0 {
				left = stack[len(stack)-1]
			}

			area := (i - left - 1) * heights[top]
			maxArea = int(math.Max(float64(maxArea), float64(area)))
		}

		stack = append(stack, i)
	}

	prev, last := -1, stack[len(stack)-1]
	for _, index := range stack {
		area := heights[index] * (last - prev)
		prev = index
		maxArea = int(math.Max(float64(maxArea), float64(area)))
	}

	return maxArea
}
