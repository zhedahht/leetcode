/*
LeetCode 42: https://leetcode.com/problems/trapping-rain-water/
*/

package leetcode

import (
	"math"
)

func trap(height []int) int {
	stack := make([]int, 0)
	result := 0
	for i, h := range height {
		for len(stack) > 0 && height[stack[len(stack)-1]] <= h {
			lower := height[stack[len(stack)-1]]
			stack = stack[:len(stack)-1]
			if len(stack) > 0 {
				left := stack[len(stack)-1]
				heigher := int(math.Min(float64(height[left]), float64(h)))
				result += (i - left - 1) * (heigher - lower)
			}
		}

		stack = append(stack, i)
	}

	return result
}
