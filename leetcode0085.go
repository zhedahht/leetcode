/*
LeetCode 85: https://leetcode.com/problems/maximal-rectangle/
*/

package leetcode

func maximalRectangle(matrix [][]byte) int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return 0
	}

	heights := make([]int, len(matrix[0]))
	maxArea := 0
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[0]); j++ {
			if matrix[i][j] == '0' {
				heights[j] = 0
			} else {
				heights[j]++
			}
		}

		// NOTE: largestRectangleArea is in leetcode84.go
		area := largestRectangleArea(heights)
		if maxArea < area {
			maxArea = area
		}
	}

	return maxArea
}
