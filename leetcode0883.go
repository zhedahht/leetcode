/*
LeetCode 883: https://leetcode.com/problems/projection-area-of-3d-shapes/
*/

package leetcode

import "math"

func projectionArea(grid [][]int) int {
	if len(grid) == 0 || len(grid[0]) == 0 {
		return 0
	}

	max := func(m, n int) int {
		return int(math.Max(float64(m), float64(n)))
	}

	result := 0
	xProj, yProj := make([]int, len(grid[0])), make([]int, len(grid))
	for _, row := range grid {
		for j, num := range row {
			xProj[j] = max(xProj[j], num)
			if num > 0 {
				result += 1
			}
		}
	}

	for i := 0; i < len(grid[0]); i++ {
		for j := 0; j < len(grid); j++ {
			yProj[j] = max(yProj[j], grid[j][i])
		}
	}

	for _, p := range xProj {
		result += p
	}

	for _, p := range yProj {
		result += p
	}

	return result
}
